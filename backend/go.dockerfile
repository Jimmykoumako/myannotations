# backend/Dockerfile

FROM golang:1.17

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o mas .

CMD ["./mas"]
