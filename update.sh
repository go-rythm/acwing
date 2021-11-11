#!/bin/sh
SHELL_PATH=$(cd `dirname ${0}`; pwd);
cd "$SHELL_PATH/util/update"
go run main.go $@