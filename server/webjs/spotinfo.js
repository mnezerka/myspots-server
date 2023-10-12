function SpotInfo(args) {
    console.log("SpotInfo:constructor:enter")

    this.el = document.getElementById("spot-info");
    this.el.style.display = "none";

    console.log("SpotInfo:constructor:leave")
}

SpotInfo.prototype.show = function(spot) {
    this.el.innerHTML = `
        <div class="title">${spot.name}</div>
        <div class="description">${spot.description}</div>
    `
    this.el.style.display = "block";
}

SpotInfo.prototype.hide = function() {
    this.el.style.display = "none";
}

