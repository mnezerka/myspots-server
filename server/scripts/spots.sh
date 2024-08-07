#!/usr/bin/env bash

source common.sh

# -s --silent = hide progress bar
curl -s -H "$(get_auth_header)" "${URL}spots"
