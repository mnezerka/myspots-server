#!/usr/bin/env bash

: ${URL:?"Need to set URL non-empty"}
USER="mn@example.com"
PASSWORD="mn"
TOKEN_FILE=".token"

echo "logging as $USER/$PASSWORD"

# curl -s = --silent = hide progress bar
# jq -r = strip quotes
curl -s -X POST -H "content-type: application/json" --data "{\"email\": \"$USER\", \"password\": \"$PASSWORD\"}" "${URL}login" | jq -r ".accessToken" > $TOKEN_FILE

echo "storing token to \"$TOKEN_FILE\" file"

cat $TOKEN_FILE
