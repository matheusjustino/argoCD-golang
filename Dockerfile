FROM golang:1.17 as build

WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go


FROM alpine:3.12

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=build /app/server .
COPY --from=build /app/.env .

CMD ["./server"]