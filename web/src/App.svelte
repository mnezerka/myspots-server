
<script>
    import {onMount} from "svelte";

    let mapyCzReady = false;
    let mapyCzLoaded = false;
    let mounted = false;
    let elMap;

    function loadedMapyCz() {
        console.log("loadedMapyCz");
        mapyCzLoaded = true;
        Loader.async = true;
        Loader.load(null, null, readyMapyCz);
    }

    function readyMapyCz() {
        mapyCzReady = true;
        console.log("ready mapy cz");
        if (mounted) {
            RenderMap();
        }
    }

    onMount(() => {
        console.log("onMount")
        mounted = true;
        let x = 1
        if (mapyCzReady) {
            RenderMap();
        }
    });
    
    function RenderMap() {
        console.log("Render Map", elMap);
        var stred = SMap.Coords.fromWGS84(14.41, 50.08);
        var mapa = new SMap(JAK.gel("map"), stred, 10);
        mapa.addDefaultLayer(SMap.DEF_BASE).enable();
        mapa.addDefaultControls();
    }

</script>

<svelte:head>
    <script src="https://api.mapy.cz/loader.js" on:load={loadedMapyCz}></script>
</svelte:head>

<div bind:this={elMap} id="map"/>

<style>
	:global(html, body, #app, #map) {
		width: 100%;
        height: 100%;
        margin: 0;
	}
</style>