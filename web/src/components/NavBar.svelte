<script>
  import { link, location, push } from "svelte-spa-router";
  import { user, sidebarVisible } from "../store/state";
  import { notifications } from "./Notification/notification";

  function logout() {
    user.logout();
    push("/");
  }

  let count = 0;
</script>

<nav
  class="flex flex-row flex-no-wrap items-center justify-between bg-gray-800
  h-16">
  <div class="flex-shrink md:flex-1 float-left">
    <a
      class="rounded px-4 py-2 mx-px md:mx-2 bg-blue-500 hover:bg-blue-400
      text-white"
      href="/"
      use:link>
      Home
    </a>
    {#if $user.type}
      <a
        class="rounded px-4 py-2 mx-px md:mx-2 bg-blue-500 hover:bg-blue-400
        text-white"
        href="/movie"
        use:link>
        Movie
      </a>
    {/if}
    <!-- Only show library page when on movie page -->
    {#if $user.type === 'streamer'}
      <a
        class="rounded px-4 py-2 mx-px md:mx-2 bg-blue-500 hover:bg-blue-400
        text-white"
        href="/library"
        use:link>
        Library
      </a>
    {/if}
    <button
      on:click={() => notifications.addTimed({ title: count++, message: 'timed notification' }, 4000)}>
      CREATE
    </button>
    <button
      on:click={() => notifications.addPersistant({
          title: count++,
          message: 'Persistant notification'
        })}>
      CREATE
    </button>
  </div>
  <div class="flex-shrink ">
    {#if $user.type}
      <a
        on:click={logout}
        class="rounded px-4 py-2 mx-px md:mx-2 bg-blue-500 hover:bg-blue-400
        text-white"
        href="/">
        Logout
      </a>
    {/if}
    <!-- Only show library page when on movie page -->
    {#if $location === '/movie'}
      <div
        on:click={() => sidebarVisible.update(visible => !visible)}
        class="hidden sm:inline rounded py-2 px-3 mx-2 bg-blue-500
        hover:bg-blue-400 text-white cursor-pointer">
        <i class="la la-comment-dots" />
      </div>
    {/if}
  </div>
</nav>
