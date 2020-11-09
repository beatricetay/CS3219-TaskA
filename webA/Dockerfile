FROM golang:alpine

WORKDIR /bin

COPY . .

RUN go build -o ./bin/webA ./main.go

ENTRYPOINT ./bin/webA

EXPOSE 5000
