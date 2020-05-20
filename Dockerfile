FROM golang:1.14.3-alpine as build-env

RUN mkdir /twitch-chat-logger
WORKDIR /twitch-chat-logger
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/twitch-chat-logger

FROM scratch
COPY --from=build-env /bin/twitch-chat-logger
ENTRYPOINT ["/bin/twitch-chat-logger"]