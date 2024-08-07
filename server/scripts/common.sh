#!/usr/bin/env bash

TOKEN_FILE=".token"

function get_token () {
    cat $TOKEN_FILE
}

function get_auth_header () {
    TOKEN=$(get_token)
    echo "authorization: Bearer $TOKEN"
}
