# MySpots

Keep your favorite map spots in cloud application.

## Usage 

Install dependencies:
```bash
npm install
```
Start a development server:
```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## Building

To create a production version of your app:

```bash
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://kit.svelte.dev/docs/adapters) for your target environment.


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