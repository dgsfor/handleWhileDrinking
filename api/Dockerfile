FROM alpine:v3.7-1

MAINTAINER dgsfor@gmail.com

WORKDIR /www/api-hwd.com
COPY hwd-api hwd-api
ADD conf conf
RUN chmod +x hwd-api

# api
EXPOSE 8081

CMD ["./hwd-api", "--env", "prod"]
