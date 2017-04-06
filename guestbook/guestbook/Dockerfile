FROM busybox:ubuntu-14.04

ADD ./bin/guestbook /app/guestbook
ADD ./guestbook/public/index.html /app/public/index.html
ADD ./guestbook/public/script.js /app/public/script.js
ADD ./guestbook/public/style.css /app/public/style.css

WORKDIR /app
CMD ["./guestbook"]
EXPOSE 3000
