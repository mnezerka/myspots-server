#!/usr/bin/env bash

: ${URL:?"Need to set URL non-empty"}

curl -X POST -H 'content-type: application/json' --data '{"name": "mn", "email": "mn@example.com", "password": "mn"}' "${URL}signup"
