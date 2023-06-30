FROM alpine

RUN apk update
RUN apk add ca-certificates

COPY ui/dist /web/dist
COPY public /web/public
COPY app /web

WORKDIR /web
RUN chmod +x app

ENV port=8080

EXPOSE ${port}

CMD ["./app"]