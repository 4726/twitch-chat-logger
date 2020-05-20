FROM golang:1.14.3-alpine

ENV GO111MODULE=on

WORKDIR /

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/twitch-chat-logger

EXPOSE 14000
ENTRYPOINT ["/bin/twitch-chat-logger", "server"]
