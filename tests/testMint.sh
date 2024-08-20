#!/bin/bash

curl -X POST https://lyla-five.vercel.app/api/handcash/mint \
    -H "Content-Type: application/json" \
    -d '{"key1":"value1","key2":"value2"}'
