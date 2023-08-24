FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/main cmd/main.go

FROM alpine:3.18

RUN adduser -D -u 1001 appuser

COPY --from=build /app/main /main

RUN chown appuser /main

USER appuser

ENTRYPOINT ["/main"]