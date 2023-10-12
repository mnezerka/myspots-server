
//////////////////////////////////////////////// MapContextMenu

function MapContextMenu(args) {
    console.log("MapContextMenu:constructor:enter")

    this.el = document.getElementById("map-context-menu");
    this.el.style.display = "none";

    // When the user clicks anywhere outside of the modal, close it
    window.onclick = function(event) {
        console.log("here");
        if (event.target == this.el) {
            this.hide();
        }
    }

    console.log("MapContextMenu:constructor:leave")
}

MapContextMenu.prototype.show = function() {
    this.el.style.display = "block";
}

MapContextMenu.prototype.hide = function() {
    this.el.style.display = "none";
}

MapContextMenu.prototype.setPosFromEvent = function(e) {
    console.log("MapContextMenu:setPosFromEvent", e)
    this.el.style.left = e.originalEvent.pageX + "px";
    this.el.style.top = e.originalEvent.pageY + "px";
}
//////////////////////////////////////////////// Map

function Map(args) {
    console.log("Map:constructor:enter");

    this.positionStore = args.positionStore || null;
    this.positionStore.registerObserver(this.onUpdatePosition.bind(this));

    this.spots = args.spots || null;
    this.spots.registerObserver(this.onUpdateSpots.bind(this));

    this.spotStore = args.spotStore || null;

    this.map = L.map('map').setView([51.505, -0.09], 13);

    //this.contextMenu = new MapContextMenu()
    
    // active position on map
    this.spotMarker = null;

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
    
    // this.map.on('contextmenu', this.onMapContextMenu.bind(this));
   
    // custom control - localize
    L.Control.Localize = L.Control.extend({
        onAdd: function(map) {
            let el = L.DomUtil.create('div');
            el.className = "myspots-control";
            el.innerHTML = '<i class="fa-regular fa-circle-dot"></i>';
            el.setAttribute('title', 'Show my location');
            
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
            el.setAttribute('title', 'Fit all content');
            
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
            el.setAttribute('title', 'Center');
            
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
    console.log("click in map", e);
    this.positionStore.setPos(e.latlng);
}

Map.prototype.onMapContextMenu = function(e) {
    console.log("context menu", e)
    //e.preventDefault()
    this.contextMenu.setPosFromEvent(e);
    this.contextMenu.show();
}
 
Map.prototype.onMarkerClick = function(e) {
    this.spotStore.setSpot(e.target.spot) 
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

    this.positionStore.setPos(p);
    this.map.panTo(p);

    console.log("Map:onLocalized:leave")
}

Map.prototype.center = function() {
    console.log("Map:center:enter");

    if (this.positionStore.getPos()) {
        this.map.panTo(this.positionStore.getPos());
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

// called when global actual position was changed
Map.prototype.onUpdatePosition = function(store) {
    console.log("Map:onUpdatePosition:enter", store.pos );
    
    // if we have position
    if (store.pos) {
        // if marker is already created => move it to new pos
        if (this.posMarker) {
            this.posMarker.setLatLng(store.pos)
        } else {
            this.posMarker = L.marker(store.pos).addTo(this.map) 
        }
    } else {
        // remove marker from map if exists
        if (this.posMarker) {
            this.map.removeLayer(this.posMarker);
            this.posMarker = null;
        }
    }

    console.log("Map:onUpdatePosition:leave");
}

Map.prototype.onUpdateSpots = function(spots) {
    console.log("Map:onUpdateSlots:enter");
    
    if (spots.getSpots()) {
        for (let i = 0; i < spots.getSpots().length; i++) {
            const s = spots.getSpots()[i];
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

        this.fit();
    }

    console.log("Map:onUpdateSlots:leave")
}
