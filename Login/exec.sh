#!/usr/bin/env bash
docker build -t tv-golang-web-app .
docker run --env-file .env -p 3000:3000 -it tv-golang-web-app
