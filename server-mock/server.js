const http = require("http");

const standardHeaders = {
    "Content-Type": "application/json",
    'Access-Control-Allow-Origin': '*',
    'Access-Control-Allow-Methods': 'OPTIONS, POST, GET',
    'Access-Control-Allow-Headers': '*',
    'Access-Control-Request-Headers': '*',
    'Access-Control-Max-Age': 2592000, // 30 days
  };

let routes = {
    "/login": {
        "accessToken": "access-token-value",
        "refreshToken": "refresh-token-value"
    },
    "/profile": {
        "name": "mn",
        "email": "michal.nezerka@gmail.com",
    },
    "/spots": [
        {
            "id": "646b5d2167e11af57a59ed41",
            "name": "SAP",
            "description": "This is SAP location in Brno",
            "coordinates": [
                49.18041,
                16.605665
            ]
        },
        {
            "id": "746b5d2167e11af57a59ed41",
            "name": "Hala Rondo",
            "description": "A place where people play ice hockey",
            "coordinates": [
                49.1854922,
                16.6020447
            ]
        }
    ]
}

console.log("----- Routes -----");
console.log(routes);
console.log("----- Routes -----");

const requestListener = function(req, res) {
    
    console.log(req.method + ' ' + req.url + ' HTTP/' + req.httpVersion, req.headers.host);
    
    let headers = standardHeaders;

    // console.log(req.headers)

    if (req.url in routes) {
        
        if (req.method === 'OPTIONS') {
            res.writeHead(204, headers);
            res.end();
            console.log(" -> 204");
            return;
        }
        
        res.writeHead(200, headers);
        res.end(JSON.stringify(routes[req.url]));
        console.log(" -> ", routes[req.url]);
    } else {
        res.writeHead(404, standardHeaders);
        res.end("The requested resource was not found")
        console.log(" -> 404 Not Found ");
    }
};

// pass on the request listener
const server = http.createServer(requestListener);

// set port number as per choice
server.listen(8081);
