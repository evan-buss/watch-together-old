<script>
  import { link, location } from "svelte-spa-router";
  import { user, sidebarVisible } from "../store/state";
</script>

<nav
  class="w-full fixed top-0 left-0 h-16 flex flex-row flex-no-wrap items-center
  justify-between bg-blue-100">
  <div>
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
  <h1>Watch Together</h1>
  {#if $location === '/movie'}
    <div
      on:click={() => sidebarVisible.update(visible => !visible)}
      class="rounded p-2 mx-2 bg-blue-500 hover:bg-blue-400 text-white">
      <i class="la la-comment-dots" />
    </div>
  {:else}
    <div />
  {/if}
</nav>
