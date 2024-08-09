# MySpots

Keep your favorite map spots in cloud application.

## Deployment

Enter certbot container:
```bash
docker-compose exec certbot sh
```

Create certificates:
```bash
certbot certonly -d yonderland.org,myspots.yonderland.org
# webroot dir is /var/www/certbot 
```

Generate Diffie-Hellman group to be refered from nginx conf:
```bash
 openssl dhparam -out /etc/ssl/certs/dhparam.pem 2048
 ```

### Deploy server

```bash
docker build . -t myspots

docker save myspots -o myspots.tar

# copy to remote system
scp myspots.tar yourserver:

# login to remote system and import image
ssh to youur server and
cat myspots.tar | docker load
```

 ## Development

 ```bash
 cd server
 docker-compose up -d
``` 

 or on MacOS:

 ```bash
 cd server
 lima nerctl compose up -d
 ```
 
 build and start server:
 ```bash
 cd server
 export DB_HOST=localhost
 go build -o server && ./server
 ```# MySpots Server

## Test

```bash
$ curl http://localhost:8080/api/v1/spots
[{"id":"xxx","name":"London"}]

$ curl http://localhost:8080/api/v1/spots/xxx
{"id":"xxx","name":"London"}

$ curl -X POST -H 'content-type: application/json' --data '{"name": "new spot"}' http://localhost:8080/api/v1/spots
{"id":"yyy","name":"karen"}
```


