FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o tbot app/myteambotslack/main.go

EXPOSE 8080

CMD ["./tbot"]