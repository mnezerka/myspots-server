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
