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
 ```