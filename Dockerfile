FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/server /app/server

CMD ["/app/server"]