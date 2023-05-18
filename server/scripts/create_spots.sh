 #!/usr/bin/env bash

 curl -s -H "authorization: JWT $TOKEN" -X POST -H 'content-type: application/json' --data '{"name": "Ostry", "coordinates": [49.6010961, 18.6571078]}' http://localhost:8080/spots
