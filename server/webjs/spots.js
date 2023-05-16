
var app;

//////////////////////////////////////////////// UserProfile 

function Identity() {

    this.name = '';
    this.email = ''; 
    this.token = localStorage.getItem('token');
    this.loggedIn = false;

    this.observers = []; 
}

Identity.prototype.registerObserver = function(observer) {
    this.observers.push(observer);
}

Identity.prototype.notifyAll = function() {
    this.observers.forEach(function(observer) {
        console.log(observer);
        observer.update();
    })
}

Identity.prototype.logout = function() {
    console.log("UserProfile:logout:enter");
    localStorage.removeItem('token');
    this.isLogged = false;
    this.token = null;
    this.notifyAll();
}

Identity.prototype.fetchProfile = async function() {

    console.log("UserProfile:fetchProfile:enter");

    if (!this.token) {
        return;
    }

    try {
        const response = await fetch("/profile", {
            method: 'GET',
            headers: {
                'Authorization': `JWT ${this.token}`
            },
            referrer: 'no-referrer'
        });

        if (response.ok) {
            const data = await response.json();
            
            console.log("fetchProfile:data", data);
            this.name = data.name;
            this.email = data.email;
            this.isLogged = true;
            this.notifyAll();
        } else {
            console.warn('Unexpected response code: ', response.status);
        }
    } catch (err) {
        console.warn('Something went wrong.', err);
    }
}

Identity.prototype.login = async function(email, password) {

    console.log("UserProfile:login:enter", email, password);

    try {
        const response = await fetch("/login", {
            method: 'POST',
            body: '{"email": "' + email + '", "password": "' + password + '"}',
            headers: {
                'Content-Type': 'application/json'
            },
            referrer: 'no-referrer'
        });
        
        if (response.ok) {
            const data = await response.json();
            console.log("UserProfile:login:data", data);
            localStorage.setItem('token', data.accessToken)
            this.token= data.accessToken;
            this.isLogged = true;
            this.notifyAll();
        } else {
            console.warn('Unexpected response code: ', response.status);
        }

    } catch (err) {
        console.warn('Something went wrong.', err);
    }
}

//////////////////////////////////////////////// Toolbar

function Toolbar(identity, onLogin, onLogout, onCenter) {
    this.identity = identity;
    this.onLogin = onLogin;
    this.onLogout = onLogout;
    this.onCenter = onCenter;

    this.identity.registerObserver(this); 

    this.elBtnLogin = document.getElementById("button-login");
    this.elBtnLogout = document.getElementById("button-logout");
    this.elBtnCenter = document.getElementById("button-center");
    this.elUserInfo = document.getElementById("user-info");

    this.elBtnLogin.addEventListener("click", this.onLogin);
    this.elBtnLogout.addEventListener("click", this.onLogout);
    this.elBtnCenter.addEventListener("click", this.onCenter);
}

Toolbar.prototype.update = function() {
    console.log("Toolbar:update:enter");

    if (this.identity.isLogged) {
        this.elUserInfo.textContent = `${this.identity.name} <${this.identity.email}>`;
        this.elBtnLogin.style.display = "none"
        this.elBtnLogout.style.display = "block"
    } else {
        this.elUserInfo.textContent =  'anonyous';
        this.elBtnLogin.style.display = "block"
        this.elBtnLogout.style.display = "none"
    }
    console.log("Toolbar:update:leave")
}

//////////////////////////////////////////////// LoginModal
// https://www.w3schools.com/howto/howto_css_modals.asp 

function LoginModal(onClose, onSubmit) {
    console.log("LoginModal:constructor:enter")

    this.onClose = onClose;
    this.onSubmit = onSubmit;

    this.el = document.getElementById("login-modal");
    this.el.style.display = "none";

    this.elClose = this.el.getElementsByClassName("close")[0];
    this.elSubmit = this.el.getElementsByClassName("submit")[0];

    this.elEmail = this.el.getElementsByClassName("input email")[0];
    this.elPassword = this.el.getElementsByClassName("input password")[0];

    console.log(this.elEmail, this.elPassword);

    this.elSubmit.addEventListener("click", this.onClickSubmit.bind(this));
    this.elClose.addEventListener("click", this.onClose);

    // When the user clicks anywhere outside of the modal, close it
    window.onclick = function(event) {
        if (event.target == this.ell) {
            this.onClose();
        }
    }

    console.log("LoginModal:constructor:leave")
}

LoginModal.prototype.onClickSubmit = function(e) {
    e.preventDefault();

    let email = this.elEmail.value;
    let password = this.elPassword.value;

    this.onSubmit(email, password);
}

LoginModal.prototype.show = function() {
    this.el.style.display = "block";
}

LoginModal.prototype.hide = function() {
    this.el.style.display = "none";
}

//////////////////////////////////////////////// Map

function Map() {
    console.log("Map:constructor:enter")

    this.map = L.map('map').setView([51.505, -0.09], 13);

    const tiles = L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(this.map);

    /*
    function onMapClick(e) {
        popup
            .setLatLng(e.latlng)
            .setContent(`You clicked the map at ${e.latlng.toString()}`)
            .openOn(map);
    }

    map.on('click', onMapClick);
    */

    console.log("Map:constructor:leave")
}

Map.prototype.centerMap = function() {
    console.log("Map:centerMap:enter");

    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(this.onCenterMap.bind(this));
    } else {
        console.warn("Geolocation is not supported by this browser.");
    }

    console.log("Map:centerMap:leave");
}

Map.prototype.onCenterMap = function(position) {
    console.log("Map:onCenterMap:enter", position)

    this.map.panTo(new L.LatLng(position.coords.latitude, position.coords.longitude));

    console.log("Map:onCenterMap:leave")
}

//////////////////////////////////////////////// Application

function App() {
    console.log("App:constructor:enter")

    this.identity = new Identity();

    this.toolbar = new Toolbar(this.identity, this.onLogin.bind(this), this.onLogout.bind(this), this.onCenter.bind(this));
    this.toolbar.update();

    this.loginModal = new LoginModal(this.onLoginModalClose.bind(this), this.onLoginModalSubmit.bind(this), );

    this.map = new Map();

    this.identity.fetchProfile();

    console.log("App:constructor:leave")
}

App.prototype.onLogin = function(e) {
    console.log("App:onLogin:enter");
    this.loginModal.show();
    console.log("App:onLogin:leave");
}

App.prototype.onLoginModalClose = function(e) {
    console.log("App:onLoginModalClose:enter");
    this.loginModal.hide();
    console.log("App:onLoginModalClose:leave");
}

App.prototype.onLoginModalSubmit = async function(email, password) {

    console.log("App:onLoginModalSubmit:enter", email, password);

    if (email == "" || password == "") {
        // do nothing, keep dialog open
        return;
    }

    this.identity.login(email, password);
    this.loginModal.hide();
}

App.prototype.onLogout = function(e) {
    console.log("App:onLogout:enter");
    this.identity.logout();
    console.log("App:onLogout:leave");
}

App.prototype.onCenter = function(e) {
    console.log("App:onCenter:enter", this);
    this.map.centerMap();
    console.log("App:onCenter:leave");
}

//////////////////////////////////////////////// main

function onLoaded() {
    console.log("onLoaded:enter");

    // initialize application when the page is completely loaded
    app = new App();

    console.log("onLoaded:leave");
}

//;map.invalidateSize();