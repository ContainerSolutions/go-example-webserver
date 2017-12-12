FROM golang:1.8

COPY *.go /go/
RUN go build -o /server
COPY test.sh /
RUN chmod +x /test.sh

CMD /server
EXPOSE 8080
