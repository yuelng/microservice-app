package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"io"
	"log"
	"strings"
	"time"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost:5672/", "AMQP URI")
	exchange     = flag.String("exchange", "signal-exchange", "Durable, non-auto-deleted AMQP exchange name")
	exchangeType = flag.String("exchange-type", "direct", "Exchange type - direct|fanout|topic|x-custom")
	queue        = flag.String("queue", "signal-queue", "Ephemeral AMQP queue name")
	bindingKey   = flag.String("key", "signal-key", "AMQP binding key")
	consumerTag  = flag.String("consumer-tag", "signal-consumer", "AMQP consumer")
	lifetime     = flag.Duration("lifetime", 0*time.Second, "lifetime of process before shutdown (0s=infinite)")
)

func init() {
	flag.Parse()
}

type Rabbit struct {
	Publish    chan *SockMsgRabbitProxy // this channel receives socket messages from the message hub
	Deliveries <-chan amqp.Delivery     // this channel listens to rabbitmq deliveries and consumes them
	Done       chan error               // this channel receives a message when there the delivery channel breaks
	Channel    *amqp.Channel            // this is the rabbitmq channel we are using to publish and subscribe to messages
}

var r = Rabbit{
	Publish: make(chan *SockMsgRabbitProxy),
	Done:    make(chan error),
}

func dial(amqpURI []byte) (*amqp.Connection, error) {
	log.Printf("rabbit dialing %q", amqpURI)
	connection, err := amqp.Dial(string(amqpURI))
	if err != nil {
		return nil, fmt.Errorf("Dial: %s", err)
	}
	return connection, nil
}

func createChannel(c *amqp.Connection) *amqp.Channel {
	log.Printf("got Connection, getting Channel")
	channel, err := c.Channel()
	if err != nil {
		panic(fmt.Errorf("Channel: %s", err))
	}

	log.Printf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange)
	if err := channel.ExchangeDeclare(
		*exchange,     // name
		*exchangeType, // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // noWait
		nil,           // arguments
	); err != nil {
		panic(fmt.Errorf("Exchange Declare: %s", err))
	}
	return channel
}

func createQueue(channel *amqp.Channel, queueName string) (*amqp.Queue, error) {
	log.Printf("declared Exchange, declaring Queue %q", queueName)
	queue, err := channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Declare: %s", err)
	}

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, bindingKey)

	if err = channel.QueueBind(
		queue.Name,  // name of the queue
		*bindingKey, // bindingKey
		*exchange,   // sourceExchange
		false,       // noWait
		nil,         // arguments
	); err != nil {
		return nil, fmt.Errorf("Queue Bind: %s", err)
	}
	return &queue, nil
}

func getDeliveriesChannel(channel *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	return channel.Consume(
		queueName,    // name
		*consumerTag, // consumerTag,
		false,        // noAck
		false,        // exclusive
		false,        // noLocal
		false,        // noWait
		nil,          // arguments
	)
}

// 1. init the rabbitmq conneciton
// 2. expect messages from the message hub on the Publish channel
// 3. if the connection is closed, try to restart it

func (r *Rabbit) run() {
	initRabbitConn()
	for {
		select {
		case sm := <-r.Publish:
			fmt.Printf("publish message: %s\n", sm)
			publish(r.Channel, sm)
		case err := <-r.Done:
			log.Println(err)
			initRabbitConn()
		}
	}
}

// try to start a new connection, channel and deliveries channel. if failed, try again in 5 sec.
func initRabbitConn() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if conn, err := dial([]byte(*uri)); err != nil {
					log.Println(err)
					log.Println("node will only be able to respond to local connections")
					log.Println("trying to reconnect in 5 seconds...")
				} else {
					close(quit)
					r.Channel = createChannel(conn)
					createQueue(r.Channel, *queue)
					r.Deliveries, _ = getDeliveriesChannel(r.Channel, *queue)
					go handleConsume(r.Deliveries, r.Done)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func publish(c *amqp.Channel, sm *SockMsgRabbitProxy) error {
	body, _ := json.Marshal(sm)
	log.Printf("publishing %dB body (%s)", len(body), body)
	if c == nil {
		return fmt.Errorf("connection to rabbitmq might not be ready yet")
	}
	if err := c.Publish(
		*exchange,   // publish to an exchange
		*bindingKey, // routing to 0 or more queues
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "application/json",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}
	return nil
}

func handleConsume(deliveries <-chan amqp.Delivery, done chan error) {
	fmt.Println("waiting for rabbitmq deliveries...")
	for d := range deliveries {
		d.Ack(false)
		s := string(d.Body)
		log.Printf(
			"got %dB delivery: [%v] %s",
			len(s),
			d.DeliveryTag,
			s,
		)
		dec := json.NewDecoder(strings.NewReader(s))
		for {
			var m SockMsgRabbitProxy
			if err := dec.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("from: %s, to: %s, broadcast: %t, payload: %s\n", m.From, m.To, m.Broadcast, m.Payload)
			h.send <- &m
		}
	}
	done <- fmt.Errorf("error: deliveries channel closed")
}
