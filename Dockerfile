FROM golang:1.18

COPY . /app
WORKDIR /app
EXPOSE 8081
RUN go mod tidy
RUN CGO_ENABLE=0 GOOS=linux go build -o main
CMD ["./main"]