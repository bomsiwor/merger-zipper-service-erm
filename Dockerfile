FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

FROM alpine:edge
WORKDIR /app

COPY --from=builder /app/main .

RUN apk --no-cache add ca-certificates tzdata

ENTRYPOINT [ "/app/main" ]