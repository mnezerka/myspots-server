var app;

var host = '';

if (window.location.href.includes("file:///")) {
    host = "http://localhost:8081"
    console.log("Setting host to local mock server:", host)
}

let empty_function = function() {}


//////////////////////////////////////////////// Application

function App() {
    console.log("App:constructor:enter")

    this.errors =  store({}, 'errors');

    document.addEventListener('wizards', function (event) {
        app.innerHTML = template(event.detail);
    }); 

    this.identity = new Identity();
    this.identity.registerObserver(this.onUpdateIdentity.bind(this));
    
    this.activeSpotId = null;

    this.spotInfo = new SpotInfo();

    this.toolbar = new Toolbar({
        identity: this.identity,
        onLogin: this.onLogin.bind(this),
        onLogout: this.onLogout.bind(this),
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

    this.map = new Map({
        onSpotClick: this.onSpotClick.bind(this)
    });

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
    
    this.map.fit();

    console.log("App:onUpdateSlots:leave")
}

App.prototype.setActiveSpot = function(spot) {
    console.log("App:setActiveSpot:enter", spot);

    this.activeSpotId = spot.id;

    console.log("App:setActiveSpot:leave");
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

App.prototype.onSpotClick = function(spot) {
    console.log("App:onSpotClick:enter");

    this.setActiveSpot(spot);

    this.spotInfo.show(spot);

    console.log("App:onSpotClick:leave");
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

