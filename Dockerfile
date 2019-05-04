FROM hhy5861/blockshine-go-base:1.0.1

MAINTAINER Mike Huang <hhy5861@gmail.com>

ENV APP_NAME backend-user-server

ADD ./${APP_NAME}-* /app/${APP_NAME}
COPY ./config /app/config/
