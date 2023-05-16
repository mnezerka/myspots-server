
var elBtnCenter;
var elBtnLogout;
var elBtnLogin;

var elLoginForm;
var elEmail;
var elPassword;

var elMap;

var map;

var userProfile;
var toolbar;

function UserProfile() {
    
    this.name = '';
    this.email = ''; 
    this.token = localStorage.getItem('token');
    this.loggedIn = this.token ? true : false;

    this.observers = []; 
}

UserProfile.prototype.registerObserver = function(observer) {
    this.observers.push(observer);
}

UserProfile.prototype.notifyAll = function() {
    this.observers.forEach(function(observer) {
        observer.update();
    })
}

UserProfile.prototype.logout = function() {
    console.log("UserProfile:logout:enter");
    localStorage.removeItem('token');
    this.isLogged = false;
    this.token = null;
    this.notifyAll();
}

UserProfile.prototype.fetchProfile = async function() {

    console.log("UserProfile:fetchProfile:enter");

    try {
        const response = await fetch("/profile", {
            method: 'GET',
            headers: {
                'Authorization': `JWT ${this.token}`
            },
            referrer: 'no-referrer'
        });
        
        console.log("fetchProfile:response", response);

        if (response.ok) {
            const data = await response.json();
            
            console.log("fetchProfile:data", data);
            this.name = data.name;
            this.email = data.email;
            this.notifyAll();
        } else {
            console.warn('Unexpected response code: ', response.status);
        }
    } catch (err) {
        console.warn('Something went wrong.', err);
    }
}

function Toolbar() {

}

Toolbar.prototype.update = function() {
    console.log("Toolbar:update:enter")

    elUserInfo = document.getElementById("user-info");
    if (userProfile.isLogged) {
        elUserInfo.textContent = `${userProfile.name} <${userProfile.email}>`;
    } else {
        elUserInfo.textContent =  'anonyous';
    }
}

function onLoaded() {
    console.log("onLoaded:enter");
    
    toolbar = new Toolbar();
    userProfile.registerObserver(toolbar); 

    elBtnCenter = document.getElementById("button-center");
    elBtnLogin = document.getElementById("button-login");
    elBtnLogout = document.getElementById("button-logout");

    elLoginForm = document.getElementById("loginForm")
    elEmail = document.getElementById("email")
    elPassword = document.getElementById("password")

    elMap =  document.getElementById("map")

    elBtnCenter.addEventListener("click", function(e) {
        centerMap();
    })
    
    elBtnLogin.addEventListener("click", onLogin);
    elBtnLogout.addEventListener("click", onLogout);
    elLoginForm.addEventListener("submit", onLoginSubmit);

    initMap();

    updateUI();
    
    if (userProfile.isLogged)  {
        userProfile.fetchProfile();
    }
    
}

function initMap() {

    console.log("initMap:enter");

    map = L.map('map').setView([51.505, -0.09], 13);

    const tiles = L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);

    /*
    function onMapClick(e) {
        popup
            .setLatLng(e.latlng)
            .setContent(`You clicked the map at ${e.latlng.toString()}`)
            .openOn(map);
    }

    map.on('click', onMapClick);
    */
 
}

function centerMap() {
    console.log("centerMap:enter");

    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(onCenterMap);
    } else {
        console.warn("Geolocation is not supported by this browser.");
    }
}

function onLogin(e) {
    console.log("login");
}

function onLogout(e) {
    console.log("logout");
    userProfile.logout();
    updateUI();
}

function onCenterMap(position) {
    console.log("onCenterMap", position)
    if (map) {
        map.panTo(new L.LatLng(position.coords.latitude, position.coords.longitude));
    }
}

function updateUI() {
    console.log("updateUI:token", localStorage.getItem('token'));
    
    //updateToolbar();

    if(localStorage.getItem('token') == null) {
        elLoginForm.style.display = "block"
        elMap.style.display = "none"
    } else {
        elLoginForm.style.display = "none"
        elMap.style.display = "block"

        //centerMap()
        map.invalidateSize();
    }
}

function updateToolbar() {
    console.log("updateToolbar:enter");
    let token = localStorage.getItem('token');
    if (token == null) {
        console.log("updateToolbar:logged out");
        elBtnLogout.style.display = "none";
        elBtnLogin.style.display = "inherited";
    } else {
        console.log("updateToolbar:logged in");
        elBtnLogout.style.display = "inherited";
        elBtnLogin.style.display = "none";

        fetch("/profile", {
            method: 'GET',
            headers: {
                'Authorization': `JWT ${token}`
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
            console.log(data)
            localStorage.setItem('name', data.name)
            localStorage.setItem('email', data.email)
        }).catch(function(err) {
            console.warn('Something went wrong.', err);
        });
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
            localStorage.setItem('token', data.accessToken)
            updateUI();
        }).catch(function(err) {
            console.warn('Something went wrong.', err);
        });
    }
}

userProfile = new UserProfile();