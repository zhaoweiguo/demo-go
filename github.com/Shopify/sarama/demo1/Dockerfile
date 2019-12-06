FROM alpine
MAINTAINER zhaoweiguo<docker@zhaoweiguo.com>
RUN apk update && apk add busybox-extras ca-certificates curl
ADD sarama /
ADD conf/config.yml /conf/config.yml