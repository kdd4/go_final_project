FROM golang:1.25.1

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

ENV TODO_PORT=7540

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server

EXPOSE $TODO_PORT

CMD ["/server"]