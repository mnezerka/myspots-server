
var app;

var host = ""

let empty_function = function() {}

//////////////////////////////////////////////// UserProfile 

function Identity() {

    this.name = '';
    this.email = ''; 
    this.token = localStorage.getItem('token');
    this.isLogged = false;

    this.observers = []; 
}

Identity.prototype.registerObserver = function(observer) {
    this.observers.push(observer);
}

Identity.prototype.notifyAll = function() {
    this.observers.forEach(function(observer) { observer(this)}.bind(this))
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
        const response = await fetch(host + "/profile", {
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
        const response = await fetch(host + "/login", {
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

            this.fetchProfile();
        } else {
            console.warn('Unexpected response code: ', response.status);
        }

    } catch (err) {
        console.warn('Something went wrong.', err);
    }
}

//////////////////////////////////////////////// Toolbar

function Spots(args) {
    this.identity = args.identity || null;
    this.spots = [];

    this.observers = []; 
}

Spots.prototype.registerObserver = function(observer) {
    this.observers.push(observer);
}

Spots.prototype.notifyAll = function() {
    this.observers.forEach(function(observer) { observer(this)}.bind(this))
}

Spots.prototype.fetch = async function() {

    console.log("Spots:fetch:enter");

    if (!this.identity || !this.identity.isLogged) {
        return;
    }

    try {
        const response = await fetch(host + "/spots", {
            method: 'GET',
            headers: {
                'Authorization': `JWT ${this.identity.token}`
            },
            referrer: 'no-referrer'
        });

        if (response.ok) {
            const data = await response.json();
            
            console.log("Spots:fetch:data", data);
            this.spots = data;
            this.notifyAll();
        } else {
            console.warn('Unexpected response code: ', response.status);
        }
    } catch (err) {
        console.warn('Something went wrong.', err);
    }
}

Spots.prototype.create = async function(spot) {

    console.log("Spots:create:enter", spot);

    if (!this.identity || !this.identity.isLogged) {
        return;
    }

    try {
        const response = await fetch(host + "/spots", {
            method: 'POST',
            body: `{"name": "${spot.name}", "coordinates": [${spot.position.lat}, ${spot.position.lng}]}`,
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `JWT ${this.identity.token}`
            },
            referrer: 'no-referrer'
        });
        
        if (response.ok) {
            const data = await response.json();
            console.log("Spots:create:data", data);
            this.notifyAll();

        } else {
            console.warn('Unexpected response code: ', response.status);
        }
    } catch (err) {
        console.warn('Something went wrong.', err);
    }
}


//////////////////////////////////////////////// Toolbar

function Toolbar(args) {

    this.identity = args.identity;
    this.onLogin = args.onLogin || empty_function;
    this.onLogout = args.onLogout || empty_function;
    this.onLocalize = args.onLocalize || empty_function;
    this.onCenter = args.onCenter || empty_function;
    this.onAdd = args.onAdd || empty_function;

    this.identity.registerObserver(this.update.bind(this)); 

    this.elBtnLogin = document.getElementById("button-login");
    this.elBtnLogout = document.getElementById("button-logout");
    this.elBtnLocalize = document.getElementById("button-localize");
    this.elBtnCenter = document.getElementById("button-center");
    this.elBtnAdd = document.getElementById("button-add");
    this.elUserInfo = document.getElementById("user-info");

    this.elBtnLogin.addEventListener("click", this.onLogin);
    this.elBtnLogout.addEventListener("click", this.onLogout);
    this.elBtnLocalize.addEventListener("click", this.onLocalize);
    this.elBtnCenter.addEventListener("click", this.onCenter);
    this.elBtnAdd.addEventListener("click", this.onAdd);
}

Toolbar.prototype.update = function(publisher) {
    console.log("Toolbar:update:enter", publisher);

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

function LoginModal(args) {
    console.log("LoginModal:constructor:enter")

    this.onClose = args.onClose || empty_function;
    this.onSubmit = args.onSubmit || empty_function;

    this.el = document.getElementById("login-modal");
    this.el.style.display = "none";

    this.elClose = this.el.getElementsByClassName("close")[0];
    this.elSubmit = this.el.getElementsByClassName("submit")[0];

    this.elEmail = this.el.getElementsByClassName("input email")[0];
    this.elPassword = this.el.getElementsByClassName("input password")[0];

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

//////////////////////////////////////////////// AddSpotModal

function AddSpotModal(args) {
    console.log("AddSpotModal:constructor:enter")

    this.onClose = args.onClose || empty_function;
    this.onSubmit = args.onSubmit || empty_function;

    this.el = document.getElementById("add-spot-modal");
    this.el.style.display = "none";

    this.elClose = this.el.getElementsByClassName("close")[0];
    this.elSubmit = this.el.getElementsByClassName("submit")[0];

    this.elName = this.el.getElementsByClassName("input name")[0];

    this.elSubmit.addEventListener("click", this.onClickSubmit.bind(this));
    this.elClose.addEventListener("click", this.onClose);

    // When the user clicks anywhere outside of the modal, close it
    window.onclick = function(event) {
        if (event.target == this.ell) {
            this.onClose();
        }
    }

    console.log("AddSpotModal:constructor:leave")
}

AddSpotModal.prototype.onClickSubmit = function(e) {
    e.preventDefault();

    let name = this.elName.value;

    this.onSubmit(name);
}

AddSpotModal.prototype.show = function() {
    this.el.style.display = "block";
}

AddSpotModal.prototype.hide = function() {
    this.el.style.display = "none";
}

//////////////////////////////////////////////// Map

function Map() {
    console.log("Map:constructor:enter")

    this.map = L.map('map').setView([51.505, -0.09], 13);
    
    // active position on map
    this.pos = null;
    this.posMarker = null;

    this.markers = {};

    this.tiles = L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(this.map);

    this.map.on('click', this.onMapClick.bind(this));

    console.log("Map:constructor:leave")
}

Map.prototype.onMapClick = function(e) {
    this.setPos(e.latlng);
}

Map.prototype.localize = function() {
    console.log("Map:localize:enter");

    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(this.onLocalized.bind(this));
    } else {
        console.warn("Geolocation is not supported by this browser.");
    }

    console.log("Map:localize:leave");
}

Map.prototype.onLocalized = function(position) {
    console.log("Map:onLocalized:enter", position)

    const p = new L.LatLng(position.coords.latitude, position.coords.longitude);

    this.setPos(p);
    this.map.panTo(p);

    console.log("Map:onLocalized:leave")
}

Map.prototype.center = function() {
    console.log("Map:center:enter");

    if (this.pos) {
        this.map.panTo(this.pos);
    }

    console.log("Map:center:leave");
}


Map.prototype.setPos = function(pos) {
    this.pos = pos;
    if (this.posMarker) {
        this.posMarker.setLatLng(this.pos)
    } else {
        this.posMarker = L.marker(this.pos).addTo(this.map) 
    }
}

Map.prototype.getPos = function() {
    return this.pos;
}

Map.prototype.updateSpots = function(spots) {

    for (let i = 0; i < spots.spots.length; i++) {
        const s = spots.spots[i];
        if (s.id in this.markers) {
            continue;
        }

        const p = new L.LatLng(s.coordinates[0], s.coordinates[1]);
        this.markers[s.id] = L.marker(p);
        this.markers[s.id].addTo(this.map);
        this.markers[s.id].bindPopup(s.name);
    }
}


//////////////////////////////////////////////// Application

function App() {
    console.log("App:constructor:enter")

    this.identity = new Identity();
    this.identity.registerObserver(this.onUpdateIdentity.bind(this));

    this.toolbar = new Toolbar({
        identity: this.identity,
        onLogin: this.onLogin.bind(this),
        onLogout: this.onLogout.bind(this),
        onLocalize: this.onLocalize.bind(this),
        onCenter: this.onCenter.bind(this),
        onAdd: this.onAdd.bind(this)
    });
    this.toolbar.update();

    this.loginModal = new LoginModal({
        onClose: this.onLoginModalClose.bind(this),
        onSubmit: this.onLoginModalSubmit.bind(this)
    });
 
    this.addSpotModal = new AddSpotModal({
        onClose: this.onAddSpotModalClose.bind(this),
        onSubmit: this.onAddSpotModalSubmit.bind(this)
    });

    this.spots = new Spots({
        identity: this.identity
    });
    this.spots.registerObserver(this.onUpdateSpots.bind(this));

    this.map = new Map();

    this.identity.fetchProfile();

    console.log("App:constructor:leave")
}

App.prototype.onUpdateIdentity = function() {
    console.log("App:onUpdateIdentity:enter");

    if (this.identity.isLogged && this.spots.spots.length == 0) {
        this.spots.fetch();
    }

    console.log("App:onUpdateIdentity:leave")
}

App.prototype.onUpdateSpots = function() {
    console.log("App:onUpdateSlots:enter");

    this.map.updateSpots(this.spots);

    console.log("App:onUpdateSlots:leave")
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

App.prototype.onLocalize = function(e) {
    console.log("App:onLocalize:enter", this);
    this.map.localize();
    console.log("App:onLocalize:leave");
}

App.prototype.onCenter = function(e) {
    console.log("App:onCenter:enter");
    this.map.center();
    console.log("App:onCenter:leave");
}

App.prototype.onAdd = function(e) {
    console.log("App:onAdd:enter");
    //this.loginModal.show();

    let pos = this.map.getPos();
    console.log(pos);
    if (pos) {
        this.addSpotModal.show();
    }

    console.log("App:onAdd:leave");
}

App.prototype.onAddSpotModalClose = function(e) {
    console.log("App:onAddSpotModalClose:enter");
    this.addSpotModal.hide();
    console.log("App:onAddSpotModalClose:leave");
}

App.prototype.onAddSpotModalSubmit = async function(name) {

    console.log("App:onAddSpotModalSubmit:enter", name);

    pos = this.map.getPos();

    if (name == "") {
        // do nothing, keep form open
        return;
    }

    if (pos) {
        console.log(pos);
        this.spots.create({
            "name": name,
            "position": pos
        })
    }

    this.addSpotModal.hide();

    console.log("App:onAddSpotModalSubmit:leave");
}

//////////////////////////////////////////////// main

function onLoaded() {
    console.log("onLoaded:enter");

    // initialize application when the page is completely loaded
    app = new App();


    console.log("onLoaded:leave");
}

//;map.invalidateSize();