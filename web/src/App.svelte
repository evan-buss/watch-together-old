<script>
  import page from "page";
  import { onMount } from "svelte";

  import { createSocket } from "./store/socket";
  import PlayerPage from "./views/PlayerPage.svelte";

  let route;
  let routeParams;

  function setRoute(r) {
    return function({ params }) {
      route = r;
      routeParams = params;
    };
  }

  page("/", setRoute(PlayerPage));
  page({ hashbang: true });

  onMount(() => {
    createSocket();
  });
</script>

<svelte:component this={route} bind:params={routeParams} />
