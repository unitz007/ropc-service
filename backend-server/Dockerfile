FROM golang:1.19
LABEL authors="charles"

RUN mkdir /app

WORKDIR /app

COPY ./ropc-service .
COPY .env .
RUN mkdir assets
COPY ./assets assets

RUN export GIN_MODE=release # run GIN in production mode

CMD ["./ropc-service"]