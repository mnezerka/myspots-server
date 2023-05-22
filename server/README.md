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

## Deploy

```bash
docker build . -t myspots

docker save myspots -o myspots.tar

# copy to remote system
scp myspots.tar yourserver:

# login to remote system and import image
ssh to youur server and
cat myspots.tar | docker load
```
