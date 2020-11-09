FROM golang:alpine

WORKDIR /bin

COPY . .

RUN go build -o ./bin/webB ./main.go

ENTRYPOINT ./bin/webB

EXPOSE 8000
