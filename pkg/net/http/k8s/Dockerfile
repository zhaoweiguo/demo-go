FROM alpine

ADD demo6 /
ADD DigiCertSHA2SecureServerCA2.crt /usr/share/ca-certificates/
RUN apk add --no-cache ca-certificates curl openssl