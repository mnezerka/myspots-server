<LibLoader url="https://api.mapy.cz/loader.js" on:loaded="{onLoaded}" />

<script>
    import LibLoader from './LibLoader.svelte';

    ///var stred = SMap.Coords.fromWGS84(14.41, 50.08);
    //var mapa = new SMap(JAK.gel("mapa"), stred, 10);
    /*
    mapa.addDefaultLayer(SMap.DEF_BASE).enable();
    mapa.addDefaultControls();                
    */
    
    function onLoaded() {
        console.log("hello");
        console.log(SMap);
    }
    
    function createMap(container) {

    ///var stred = SMap.Coords.fromWGS84(14.41, 50.08);
    //var mapa = new SMap(JAK.gel("mapa"), stred, 10);
 
        /*
        let m = L.map(container, {preferCanvas: true }).setView(initialView, 5);
    L.tileLayer(
	    'https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png',
	    {
	      attribution: `&copy;<a href="https://www.openstreetmap.org/copyright" target="_blank">OpenStreetMap</a>,
	        &copy;<a href="https://carto.com/attributions" target="_blank">CARTO</a>`,
	      subdomains: 'abcd',
	      maxZoom: 14,
	    }
	  ).addTo(m);

    return m;
    */
  } 
    
    function mapAction(container) {
        /*
        map = createMap(container);

        toolbar.addTo(map);

        markerLayers = L.layerGroup()
                
        return {
            destroy: () => {
                toolbar.remove();
                map.remove();
                map = null;
            }
        };
        */
    }

</script>

<h1>MySpots</h1>

a place for map

<div class="map" style="height:100%;width:100%" use:mapAction />