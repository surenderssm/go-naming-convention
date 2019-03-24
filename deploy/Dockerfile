# A minimal Docker image based on Alpine Linux with a complete package index and only 5 MB in size!
FROM alpine

# expose port
EXPOSE 8083

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Adding executable into image
ADD build/app /

CMD ["/app"]
