#!/bin/sh
SHELL_PATH=$(cd `dirname ${0}`; pwd);
cd "$SHELL_PATH/util/submit"
go run main.go -n $@
cd "$SHELL_PATH"
goimports -w submit.go