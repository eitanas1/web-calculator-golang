FROM busybox:stable
MAINTAINER eitanas1 <eitanas1@gmail.com>
WORKDIR /app

EXPOSE 8080

ENV HOST=localhost
# 	USERNAME=username \
# 	USER_PASSWORD=user_password \
# 	ROLE=role \
# 	ORG_ID=org_id \
# 	GRAFANA_URL=grafana_url \
# 	ADMIN_USERNAME=admin_username \
# 	ADMIN_PASSWORD=admin_password

COPY web-calculator .

#RUN ./web-calculator

CMD ["/app/web-calculator"]
