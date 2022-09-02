FROM golang:1.18.5-alpine3.16 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o server

FROM alpine:3.16.2
COPY --from=builder /app/server ./
ENTRYPOINT ["./server"]