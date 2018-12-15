FROM scratch

LABEL author="Nickson" email="nicksonmjk@gmail.com"

# import certs, user and group files
COPY ./build/ca-certificates.crt /etc/ssl/certs/
COPY ./build/passwd /etc/passwd

# copy the static executable
COPY --chown=1000:1000 ./build/wapi /go/bin/wapi

USER 1000

VOLUME [ "/data" ]

EXPOSE 8080

# run binary
ENTRYPOINT [ "/go/bin/wapi" ]