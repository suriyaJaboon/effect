# Dockerfile
FROM golang:1.15.7-alpine3.13 as builder

COPY . /topics
WORKDIR /topics
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o main

#second stage
FROM scratch
# FROM alpine:latest
WORKDIR /topics
COPY --from=builder /topics/main .
EXPOSE 3000
CMD ["./main"]