FROM golang:1.25-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./out/main .

FROM alpine:3.19

RUN apk --no-cache add curl

RUN apk add --no-cache sqlite-libs

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
#USER appuser

COPY --from=builder /app/out/main .

EXPOSE 8000
CMD ["./main"]