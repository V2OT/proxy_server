FROM alpine

ADD proxy_server /

EXPOSE 8080

ENTRYPOINT /proxy_server
