FROM alpine:3.18.4

RUN apk update
RUN apk add ca-certificates

# COPY ui/dist /web/dist
COPY public /web/public
COPY app /web

WORKDIR /web
RUN chmod +x app

ENV PORT=8080

EXPOSE ${PORT}

CMD ["./app"]