FROM golang:1.18.3-alpine3.16 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o server

FROM alpine:3.16.0
COPY --from=builder /app/server ./
ENTRYPOINT ["./server"]