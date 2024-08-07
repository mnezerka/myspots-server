#!/usr/bin/env bash

source common.sh

curl -s -H "$(get_auth_header)" -X POST -H 'content-type: application/json' --data '{"name": "Ostry", "coordinates": [49.6010961, 18.6571078]}' "${URL}spots"
