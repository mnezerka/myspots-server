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
    "/spots": []
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
