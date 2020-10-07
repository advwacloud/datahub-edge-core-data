FROM alpine
ADD core-data-service /core-data-service
ENTRYPOINT [ "/core-data-service" ]
