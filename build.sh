#! /bin/bash

set -e

export LOGCAT=true

APP_NAME=gin-base

BIN_APP=bin/${APP_NAME}

#if [[ -f "${BIN_APP}" ]];then
#    echo "delete $BIN_APP ok"
#    rm -f ${BIN_APP}
#fi
#
#go build -v -o bin/${APP_NAME} .
#
#${BIN_APP} -c=config/local/config.yaml start

gowatch -o ./bin/${APP_NAME} -p . -args='-c=config/local/config.yaml,start'