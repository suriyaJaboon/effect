# Dockerfile
FROM golang:1.15.7-alpine3.13 as builder

COPY . /offset
WORKDIR /offset
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o main

#second stage
FROM scratch
# FROM alpine:latest
WORKDIR /offset
COPY --from=builder /offset/main .
EXPOSE 1323
CMD ["./main"]