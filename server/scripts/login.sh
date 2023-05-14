 #!/usr/bin/env bash

 curl -X POST -H 'content-type: application/json' --data '{"email": "mn@example.com", "password": "pwd"}' http://localhost:8080/login