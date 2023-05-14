
var elBtnCenter;

var elLoginForm;
var elEmail;
var elPassword;

var elMap;

function onLoaded() {
    console.log("onLoaded");

    elBtnCenter = document.getElementById("button-center")

    elLoginForm = document.getElementById("loginForm")
    elEmail = document.getElementById("email")
    elPassword = document.getElementById("password")

    elMap =  document.getElementById("map")

    elBtnCenter.addEventListener("click", function(e) {
        console.log("center");
        centerMap();
    })

    elLoginForm.addEventListener("submit", onLoginSubmit)

    updateUI();
}

function centerMap() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(onCenterMap);
    } else { 
        console.warn("Geolocation is not supported by this browser.");
    }
}

function onCenterMap(position) {
    console.log("onCenterMap", position)
    if (map) {
        map.panTo(new L.LatLng(position.coords.latitude, position.coords.longitude)); 
    }
}

function updateUI() {
    console.log("token", localStorage.getItem('token'));

    if(localStorage.getItem('token') == null) {
        elLoginForm.style.display = "block"
        elMap.style.display = "none"
    } else {
        elLoginForm.style.display = "none"
        elMap.style.display = "block"
        
        centerMap()
    }
}

function onLoginSubmit(e) {
    e.preventDefault();
    let email = elEmail.value.trim();
    let password = elPassword.value.trim();

    if (email != "" && password != "") {
        
        fetch("/login", {
            method: 'POST',
            body: '{"email": "' + email + '", "password": "' + password + '"}',
            headers: {
                'Content-Type': 'application/json'
            },
            referrer: 'no-referrer'
        }).then(function(response) {
            // The API call was successful!
            if (response.ok) {
                return response.json();
            } else {
                return Promise.reject(response);
            }
        }).then(function (data) {
            // This is the JSON from our response
            console.log(data);
            localStorage.setItem('token', data.accessToken)
            updateUI();
        }).catch(function(err) {
            console.warn('Something went wrong.', err); 
        });
    }
}