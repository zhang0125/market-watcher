FROM golang:1.17 as build

WORKDIR /go/release
ADD . .
RUN go build -o market-watcher

FROM alpine as prod
WORKDIR /app

COPY --from=build /go/release/market-watcher /app/

CMD ["./market-watcher"]
