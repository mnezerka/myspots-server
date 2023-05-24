
function Map(args) {
    console.log("Map:constructor:enter");

    this.onSpotClick = args.onSpotClick || empty_function;

    this.map = L.map('map').setView([51.505, -0.09], 13);
    
    // active position on map
    this.pos = null;
    this.posMarker = null;

    this.markers = {};
    
    this.groupBasic = L.featureGroup([]).addTo(this.map);
    
    this.icons = {
        parking: L.divIcon({ html: '<i class="fa-solid fa-square-parking fa-2x"></i>', iconSize: [16, 16], className: 'myDivIcon' })
    }

    this.tiles = L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(this.map);

    this.map.on('click', this.onMapClick.bind(this));
    
    // custom control - localize
    L.Control.Localize = L.Control.extend({
        onAdd: function(map) {
            let el = L.DomUtil.create('div');
            el.className = "myspots-control";
            el.innerHTML = '<i class="fa-regular fa-circle-dot"></i>';
            
            L.DomEvent.on(el, "click", function(ev) {
                L.DomEvent.stopPropagation(ev);
                this.localize();
            }.bind(this))

            this.elCtrlLocalize = el;
            return this.elCtrlLocalize;
        }.bind(this),
    
        onRemove: function(map) {
            L.DomEvent.off(this.elCtrlLocalize);
        }
    });
    
    this.ctrlLocalize = new L.Control.Localize({position: 'topright'}).addTo(this.map); 

    // custom control - fit all content

    L.Control.Fit = L.Control.extend({
        onAdd: function(map) {
            let el = L.DomUtil.create('div');
            el.className = "myspots-control";
            el.innerHTML = '<i class="fa-solid fa-down-left-and-up-right-to-center"></i>';
            
            L.DomEvent.on(el, "click", function(ev) {
                L.DomEvent.stopPropagation(ev);
                this.fit();
            }.bind(this))

            this.elCtrlFit = el;
            return this.elCtrlFit;
        }.bind(this),
    
        onRemove: function(map) {
            L.DomEvent.off(this.elCtrlFit);
        }
    });
    
    this.ctrlFit = new L.Control.Fit({position: 'topright'}).addTo(this.map); 
    
       // custom control - center

    L.Control.Center = L.Control.extend({
        onAdd: function(map) {
            let el = L.DomUtil.create('div');
            el.className = "myspots-control";
            el.innerHTML = '<i class="fa-solid fa-arrows-to-dot"></i>';
            
            L.DomEvent.on(el, "click", function(ev) {
                L.DomEvent.stopPropagation(ev);
                this.center();
            }.bind(this))

            this.elCtrlCenter = el;
            return this.elCtrlCenter;
        }.bind(this),

        onRemove: function(map) {
            L.DomEvent.off(this.elCtrlCenter);
        }
    });
    
    this.ctrlCenter = new L.Control.Center({position: 'topright'}).addTo(this.map); 

    console.log("Map:constructor:leave")
}

Map.prototype.onMapClick = function(e) {
    this.setPos(e.latlng);
}

Map.prototype.onMarkerClick = function(e) {
  //alert("hi. you clicked the marker at " + e.latlng);
    this.onSpotClick(e.target.spot)
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

Map.prototype.fit = function() {
    console.log("Map:fit:enter");

    points = [];

    for (var key in this.markers) {
        points.push(this.markers[key].getLatLng());
    }

    console.log("Map:fit:points", points);
    
    if (points.length > 0) {
        this.map.fitBounds(points);
    }

    console.log("Map:fit:leave");
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
        
        const markerOptions = {
            icon: this.icons.parking
        }

        const p = new L.LatLng(s.coordinates[0], s.coordinates[1]);
        var m = L.marker(p, markerOptions)
            .addTo(this.map)
            .on('click', this.onMarkerClick.bind(this));

        m.spot = s;
        this.markers[s.id] = m;
    }
}

