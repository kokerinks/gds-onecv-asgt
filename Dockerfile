FROM golang:1.22.1

WORKDIR .

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 3000

ENTRYPOINT ["./main"]