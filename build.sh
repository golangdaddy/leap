#!/bin/bash

go run . || exit 10
chmod -R 775 build/ || exit 10
(cd build/app && rm -r node_modules && npm install) || exit 10
