FROM golang:latest AS builder

RUN apt-get update && \
    apt-get install -y git automake autoconf libtool bison

RUN mkdir -p /go/src/github.com/maurofran/iam
COPY . /go/src/github.com/maurofran/iam
WORKDIR /go/src/github.com/maurofran/iam
# RUN make setup
RUN make build

FROM alpine:latest
RUN apk add --no-cache ca-certificates
# Create app dir
RUN mkdir -p /var/app/
WORKDIR /var/app
# Copy Binary
COPY --from=builder /go/src/github.com/maurofran/iam/bin/linux_amd64/iamd /var/app
ENV DB_URL mongodb://localhost/iam_db
# Run Binary
CMD ["/var/app/iamd"]