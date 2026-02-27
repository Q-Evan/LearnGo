FROM golang:1.26.0-alpine3.23 AS builder

WORKDIR /app
COPY . .
RUN go build -o server .

FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/server .

EXPOSE 8080
CMD [ "./server" ]