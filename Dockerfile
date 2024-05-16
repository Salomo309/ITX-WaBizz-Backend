FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN mkdir -p /app/storage
RUN go build -o main ./server
EXPOSE 8080
CMD ["./main"]
