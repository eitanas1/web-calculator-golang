FROM ubuntu:latest
MAINTAINER eitanas1 <eitanas1@gmail.com>
WORKDIR /app

EXPOSE 8080

RUN apk update && \
    apk install -y sudo && \
    addgroup --gid 400 web-calculator && \
    adduser --uid 200 --gid 400 --disabled-password --gecos "" web-calculator && \
    echo 'web-calculator ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers && \
    apk install -y curl

USER web-calculator

ENV LOG_LEVEL=INFO \
    USERNAME=username \
    USER_PASSWORD=user_password

COPY --chown=web-calculator:web-calculator web-calculator .

CMD ["/app/web-calculator"]
