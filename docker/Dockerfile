FROM busybox:stable
MAINTAINER eitanas1 <eitanas1@gmail.com>
WORKDIR /app

EXPOSE 8080

ENV LOG_LEVEL=INFO \
    USERNAME=username \
    USER_PASSWORD=user_password

COPY web-calculator .

CMD ["/app/web-calculator"]
