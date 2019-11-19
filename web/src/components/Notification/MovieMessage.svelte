<script>
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();
  export let notification;
</script>

<div
  class="relative mb-2 rounded shadow-md w-full max-w-md bg-white flex flex-row">
  <!-- Notification Movie Poster -->
  {#if notification.poster}
    <img src={notification.poster} alt="movie poster" class="h-32 rounded-l" />
  {:else}
    <div
      class="h-32 w-24 bg-blue-700 rounded-l flex items-center justify-center">
      <i class="la la-video text-4xl text-white" />
    </div>
  {/if}
  <div class="flex-grow p-2">
    <!-- Notification Message Title Header -->
    <h1 class="pr-8 font-bold text-blue-600 break-words">
      {notification.title}
    </h1>
    <!-- Notification Message Body -->
    <span class="text-gray-600 break-words">{notification.message}</span>
  </div>
  <i
    on:click={() => dispatch('close')}
    class="absolute right-0 top-0 mt-1 mr-1 la la-times-circle text-2xl
    text-red-500 hover:text-red-300 cursor-pointer" />
</div>
