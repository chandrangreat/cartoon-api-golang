FROM golang:alpine

WORKDIR /build

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o main .

ENTRYPOINT ["./main"]