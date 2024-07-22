FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum main.go ./

RUN go mod download && go mod verify

RUN go build -o /albums

EXPOSE 8080

CMD ["/albums"]