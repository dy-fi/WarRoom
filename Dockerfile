FROM golang:latest

LABEL maintainer Dylan Finn

WORKDIR /wr

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server .

EXPOSE 8000
CMD ["./server"]