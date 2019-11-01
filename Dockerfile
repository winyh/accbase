FROM alpine
ADD test-srv /test-srv
ENTRYPOINT [ "/test-srv" ]
