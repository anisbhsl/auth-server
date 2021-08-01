FROM golang:1.14.1-buster
WORKDIR /app
COPY . .
RUN CGO_ENABLED=1 GOOS=linux  go build -a -installsuffix cgo -o auth-bin main.go
ENTRYPOINT ["./auth-bin"]