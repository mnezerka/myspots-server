
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

<main>

    <h1>MySpots</h1>

    <p>a place for map is here</p>

    <div bind:this={elMap} id="map" style="height: 500px; width: 500px" />

</main>