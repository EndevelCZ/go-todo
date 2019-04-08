#!/usr/bin/env sh
GO_WORK_DIR=${GO_WORK_DIR:-$GOPATH/cmd}
ls
echo ${GO_WORK_DIR}
echo $GO_WORK_DIR
cd ${GO_WORK_DIR}

exec "$@"