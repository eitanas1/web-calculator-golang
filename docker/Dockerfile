FROM curlimages/curl:latest
MAINTAINER eitanas1 <eitanas1@gmail.com>
WORKDIR /app

EXPOSE 8080

# Update the package list, install sudo, create a non-root user, and grant password-less sudo permissions
RUN apk update && \
    apk install -y sudo && \
    addgroup --gid 400 nonroot && \
    adduser --uid 200 --gid 400 --disabled-password --gecos "" nonroot && \
    echo 'nonroot ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers

# Set the non-root user as the default user
USER nonroot

ENV LOG_LEVEL=INFO \
    USERNAME=username \
    USER_PASSWORD=user_password

COPY web-calculator .

CMD ["/app/web-calculator"]
