FROM cgswong/aws:aws
ADD s3-sync /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/s3-sync
