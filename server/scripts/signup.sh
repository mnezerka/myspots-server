 #!/usr/bin/env bash

 curl -X POST -H 'content-type: application/json' --data '{"name": "mn", "email": "mn@example.com", "password": "mn"}' http://localhost:8080/signup