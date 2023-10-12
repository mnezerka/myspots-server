
function Toolbar(args) {

    this.app = args.app;
    this.spotStore = args.spotStore || null;
    this.identity = args.identity;
    this.onLogin = args.onLogin || empty_function;
    this.onLogout = args.onLogout || empty_function;
    this.onAdd = args.onAdd || empty_function;
    this.onCancel = args.onCancel || empty_function;

    this.app.registerObserver(this.updateApp.bind(this)); 

    this.identity.registerObserver(this.updateIdentity.bind(this)); 

    this.elBtnLogin = document.getElementById("button-login");
    this.elBtnLogout = document.getElementById("button-logout");
    this.elBtnAdd = document.getElementById("button-add");
    this.elBtnCancel = document.getElementById("button-cancel");
    this.elUserInfo = document.getElementById("user-info");

    this.elBtnLogin.style.display = "none"
    this.elBtnLogout.style.display = "none"

    this.elBtnLogin.addEventListener("click", this.onLogin);
    this.elBtnLogout.addEventListener("click", this.onLogout);
    this.elBtnAdd.addEventListener("click", this.onAdd);
    this.elBtnCancel.addEventListener("click", this.onCancel);

    this.spotInfo = new SpotInfo();

    this.spotForm = new SpotForm({});

    this.spotStore.registerObserver(this.onSpotChange.bind(this)); 
}

Toolbar.prototype.updateIdentity = function(identity) {
    console.log("Toolbar:updateIdentity:enter", identity);

    if (identity.isLogged) {
        this.elUserInfo.textContent = `${identity.name} <${identity.email}>`;
        this.elBtnLogin.style.display = "none"
        this.elBtnLogout.style.display = ''
    } else {
        this.elUserInfo.textContent =  'anonyous';
        this.elBtnLogin.style.display = ''
        this.elBtnLogout.style.display = "none"
    }
    console.log("Toolbar:updateIdentity:leave")
}

Toolbar.prototype.updateApp = function(publisher) {
    console.log("Toolbar:updateApp:enter", publisher);
    
    if (this.app.state == APP_STATE_IDLE) {
        console.log("Toolbar: app is idle state")
        this.elBtnAdd.style.display = "none"
        this.elBtnCancel.style.display = "none"
        this.spotForm.hide();
    } else if (this.app.state == APP_STATE_POSITION) {
        console.log("Toolbar: app is in position state")

        if (this.identity.isLogged) {
            this.elBtnAdd.style.display = '';
            this.elBtnCancel.style.display = '';
            this.spotForm.show();
        } 
    } else if (this.app.state == APP_STATE_SPOT_VIEW) {
        console.log("Toolbar: app is in spot view state")
        this.elBtnAdd.style.display = 'none';
        this.elBtnCancel.style.display = '';
        this.spotForm.hide();

    } else {
        console.log("Toolbar: app is in unknown state")
    }

    console.log("Toolbar:updateApp:leave")
}

Toolbar.prototype.onSpotChange = function(spot) {
    console.log("Toolbar:onSpotChange:enter");
    
    if (spot.getSpot()) {
        this.spotInfo.show(spot.getSpot());
    } else {
        this.spotInfo.hide();
    }

    console.log("Toolbar:onSpotChange:leave");
}