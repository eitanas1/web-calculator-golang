FROM busybox:stable
MAINTAINER eitanas1 <eitanas1@gmail.com>
WORKDIR /app

EXPOSE 8080

ENV HOST=localhost

COPY web-calculator .

CMD ["/app/web-calculator"]
