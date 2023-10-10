FROM golang:1.19
LABEL authors="charles"

RUN mkdir /app

WORKDIR /app

COPY ./backend-server .
COPY .env .
RUN mkdir assets
COPY ./assets assets

RUN export GIN_MODE=release # run GIN in production mode

CMD ["./backend-server"]