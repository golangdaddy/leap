#!/bin/bash

GONOPROXY=github.com/richardboase go get -u github.com/richardboase/npgpublic

go run . $1 || exit 10
chmod -R 775 build/ || exit 10
(cd build && go build . errors) || exit 10
(cd build/app && npm install && npm run dev) || exit 10