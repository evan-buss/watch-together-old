<script>
  import page from "page";
  import { onMount } from "svelte";

  // import { createSocket } from "./store/socket";
  import NavBar from "./components/NavBar.svelte";
  import PlayerPage from "./views/PlayerPage.svelte";
  import LoginPage from "./views/LoginPage.svelte";

  let route;
  let routeParams;

  function setRoute(r) {
    return function({ params }) {
      route = r;
      routeParams = params;
    };
  }

  page("/", setRoute(LoginPage));
  page("/movie", setRoute(PlayerPage));
  page({ hashbang: true });
</script>

<NavBar />
<svelte:component this={route} bind:params={routeParams} />
