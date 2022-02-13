FROM alpine:3.15.0

ADD server /bin/

EXPOSE 8080

CMD ["/bin/server"]