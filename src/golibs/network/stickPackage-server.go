package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

func main() {
	netListen, err := net.Listen("tcp", ":5000")
	CheckError(err)

	defer netListen.Close()

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	allbuf := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		readLen, err := conn.Read(buffer)
		//fmt.Println("readLen: ", readLen, len(allbuf))
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read error")
			return
		}

		if len(allbuf) != 0 {
			allbuf = append(allbuf, buffer...)
		} else {
			allbuf = buffer[:]
		}
		var readP int = 0
		for {
			//fmt.Println("allbuf content:", string(allbuf))

			//buffer长度小于7
			if readLen-readP < 7 {
				allbuf = buffer[readP:]
				break
			}

			msgLen, _ := strconv.Atoi(string(allbuf[readP+4 : readP+7]))
			logLen := 7 + msgLen
			//fmt.Println(readP, readP+logLen)
			//buffer剩余长度>将处理的数据长度
			if len(allbuf[readP:]) >= logLen {
				//fmt.Println(string(allbuf[4:7]))
				fmt.Println(string(allbuf[readP : readP+logLen]))
				readP += logLen
				//fmt.Println(readP, readLen)
				if readP == readLen {
					allbuf = nil
					break
				}
			} else {
				allbuf = buffer[readP:]
				break
			}
		}
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
