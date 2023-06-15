FROM golang:1.20

RUN mdkir -p /usr/src/app
WORKDIR /usr/src/app
COPY . /usr/src/app
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin/app ./...
EXPOSE 8080
CMD ["go run", "main.go"]