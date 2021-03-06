FROM golang:1.8.0

RUN curl https://glide.sh/get | sh

RUN mkdir -p /go/src/booking-go-rest-api
WORKDIR /go/src/booking-go-rest-api

EXPOSE 3000

COPY ./entrypoint.sh /
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]