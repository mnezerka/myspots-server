 #!/usr/bin/env bash

# curl -s = --silent = hide progress bar
#curl -s -H "authorization: JWT $TOKEN" http://localhost:8080/spots
curl -s -H "authorization: JWT $TOKEN" https://myspots.yonderland.org/spots
