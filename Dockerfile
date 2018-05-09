FROM alpine
ADD json2file /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/json2file
