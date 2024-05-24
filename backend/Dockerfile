FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/server ./cmd/server
COPY ./internal ./internal
COPY ./pkg ./pkg

RUN go build -o /my-reading-app ./cmd/server

EXPOSE 8080

CMD [ "/my-reading-app" ]
