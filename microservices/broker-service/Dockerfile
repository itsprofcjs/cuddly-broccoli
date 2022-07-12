# base go image

FROM golang as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./api

RUN chmod +x /app/brokerApp


# build a tiny docker image

FROM alpine

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]