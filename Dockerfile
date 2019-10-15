FROM alpine
EXPOSE 8080
ADD demo-go /bin
ENTRYPOINT [ "/bin/demo-go" ]
