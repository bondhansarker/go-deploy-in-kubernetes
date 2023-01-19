ARG GO_VERSION=1.18
FROM golang:${GO_VERSION}-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 3000
CMD ["/app/main"]