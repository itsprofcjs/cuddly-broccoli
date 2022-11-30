FROM alpine

RUN mkdir /app

COPY loggerServiceApp /app

CMD ["/app/loggerServiceApp"]