 #!/usr/bin/env bash

# curl -s = --silent = hide progress bar
# jq -r = strip quotes
 curl -s -X POST -H 'content-type: application/json' --data '{"email": "mn@example.com", "password": "mn"}' http://localhost:8080/login | jq -r ".accessToken"