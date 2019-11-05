<script>
  import { link, location, push } from "svelte-spa-router";
  import { user, sidebarVisible } from "../store/state";

  function logout() {
    user.logout();
    push("/");
  }
</script>

<nav
  class="flex flex-row flex-no-wrap items-center justify-between bg-gray-800
  h-16">
  <div class="flex-1">
    <a
      class="rounded px-4 py-2 mx-2 bg-blue-500 hover:bg-blue-400 text-white"
      href="/">
      Home
    </a>
    <!-- Only show library page when on movie page -->
    {#if $user.type === 'streamer' && $location === '/movie'}
      <a
        class="rounded px-4 py-2 mx-2 bg-blue-500 hover:bg-blue-400 text-white"
        href="/library"
        use:link>
        Library
      </a>
      <!-- Only show movie page when on library page -->
    {:else if $user.type === 'streamer' && $location === '/library'}
      <a
        class="rounded px-4 py-2 mx-2 bg-blue-500 hover:bg-blue-400 text-white"
        href="/movie"
        use:link>
        Movie
      </a>
    {/if}
  </div>
  <h1
    class="flex-1 text-xl text-center font-semibold font-mono hidden sm:inline
    text-white">
    Watch Together
  </h1>
  <div class="flex-1 flex justify-end">
    <div
      on:click={logout}
      class="rounded px-4 py-2 mx-2 bg-blue-500 hover:bg-blue-400 text-white
      cursor-pointer">
      Logout
    </div>
    {#if $location === '/movie'}
      <div
        on:click={() => sidebarVisible.update(visible => !visible)}
        class="hidden sm:inline rounded py-2 px-3 mx-2 bg-blue-500 hover:bg-blue-400
        text-white cursor-pointer">
        <i class="la la-comment-dots" />
      </div>
    {/if}
  </div>
</nav>
