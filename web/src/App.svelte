<script>
  import Router, { wrap, replace } from "svelte-spa-router";
  import { onMount } from "svelte";

  import NavBar from "./components/NavBar.svelte";
  import PlayerPage from "./views/PlayerPage.svelte";
  import LoginPage from "./views/LoginPage.svelte";
  import LibraryPage from "./views/LibraryPage.svelte";

  import NotificationContainer from "./components/Notification/NotificationContainer.svelte";

  import { user, sidebarVisible } from "./store/state";

  // When authentication fails for any route send user to home page
  function conditionsFailed() {
    if ("type" in $user) {
      replace("/movie");
    } else {
      replace("/");
    }
  }

  const routes = {
    "/": wrap(LoginPage, (location, querystring) => {
      if ("type" in $user) {
        return false;
      }
      return true;
    }),
    "/library": wrap(LibraryPage, (location, querystring) => {
      // Only allow the streamer to view the library and change videos
      if ("type" in $user && $user.type === "streamer") {
        return true;
      }
      return false;
    }),
    "/movie": wrap(PlayerPage, (location, querystring) => {
      // Users must sign in before accessing the stream
      if ("name" in $user && $user.name != "") {
        return true;
      }
      return false;
    })
  };
</script>

<style>
  .page-content {
    height: calc(100vh - 4em);
  }
</style>

<NavBar full={!$sidebarVisible} />
<NotificationContainer />
<div class="page-content bg-gray-300">
  <Router {routes} on:conditionsFailed={conditionsFailed} />
</div>
