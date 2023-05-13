# MySpots Server

## Test

```bash
$ curl http://localhost:8080/api/v1/spots
[{"id":"xxx","name":"London"}]

$ curl http://localhost:8080/api/v1/spots/xxx
{"id":"xxx","name":"London"}

$ curl -X POST -H 'content-type: application/json' --data '{"name": "new spot"}' http://localhost:8080/api/v1/spots
{"id":"yyy","name":"karen"}
```