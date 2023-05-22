//////////////////////////////////////////////// Toolbar

function Toolbar(args) {

    this.identity = args.identity;
    this.onLogin = args.onLogin || empty_function;
    this.onLogout = args.onLogout || empty_function;
    this.onLocalize = args.onLocalize || empty_function;
    this.onCenter = args.onCenter || empty_function;
    this.onAdd = args.onAdd || empty_function;
    this.onFit = args.onFit || empty_function;

    this.identity.registerObserver(this.update.bind(this)); 

    this.elBtnLogin = document.getElementById("button-login");
    this.elBtnLogout = document.getElementById("button-logout");
    this.elBtnFit = document.getElementById("button-fit");
    this.elBtnLocalize = document.getElementById("button-localize");
    this.elBtnCenter = document.getElementById("button-center");
    this.elBtnAdd = document.getElementById("button-add");
    this.elUserInfo = document.getElementById("user-info");

    this.elBtnLogin.addEventListener("click", this.onLogin);
    this.elBtnLogout.addEventListener("click", this.onLogout);
    this.elBtnLocalize.addEventListener("click", this.onLocalize);
    this.elBtnCenter.addEventListener("click", this.onCenter);
    this.elBtnAdd.addEventListener("click", this.onAdd);
    this.elBtnFit.addEventListener("click", this.onFit);
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

