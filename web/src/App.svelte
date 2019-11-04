<script>
  import page from "page";

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
</script>

<svelte:component this={route} bind:params={routeParams} />
