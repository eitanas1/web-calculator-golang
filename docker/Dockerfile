FROM curlimages/curl:latest
MAINTAINER eitanas1 <eitanas1@gmail.com>

# ENV USER_FULL_NAME=user_full_name \
# 	EMAIL=email \
# 	USERNAME=username \
# 	USER_PASSWORD=user_password \
# 	ROLE=role \
# 	ORG_ID=org_id \
# 	GRAFANA_URL=grafana_url \
# 	ADMIN_USERNAME=admin_username \
# 	ADMIN_PASSWORD=admin_password

COPY web-calculator /

# RUN apk update \
# 	&& apk add jq

CMD ["/web-calculator"]
