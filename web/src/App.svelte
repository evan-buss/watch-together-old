<script>
  import Router, { wrap, replace } from "svelte-spa-router";
  import { onMount } from "svelte";

  import NavBar from "./components/NavBar.svelte";
  import PlayerPage from "./views/PlayerPage.svelte";
  import LoginPage from "./views/LoginPage.svelte";
  import LibraryPage from "./views/LibraryPage.svelte";

  import { user, sidebarVisible } from "./store/state";

  // When authentication fails for any route send user to home page
  function conditionsFailed() {
    replace("/");
  }

  const routes = {
    "/": wrap(LoginPage, (location, querystring) => {
      return true;
    }),
    "/library": wrap(LibraryPage, (location, querystring) => {
      if ("type" in $user && $user.type === "streamer") {
        return true;
      }
      return false;
    }),
    "/movie": wrap(PlayerPage, (location, querystring) => {
      if ("name" in $user && $user.name != "") {
        return true;
      }
      return false;
    })
  };
</script>

<NavBar full={!$sidebarVisible} />
<Router {routes} on:conditionsFailed={conditionsFailed} />
