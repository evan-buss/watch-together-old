<script>
  import GenericMessage from "./GenericMessage.svelte";
  import MovieMessage from "./MovieMessage.svelte";

  import { notifications } from "./notification";
</script>

<div class="z-30 fixed top-0 right-0 mt-2 mr-2 w-full max-w-sm flex flex-col">
  {#each $notifications as notification (notification.time)}
    {#if notification.type == 'movie'}
      <MovieMessage
        on:close={() => notifications.remove(notification)}
        {notification} />
    {:else}
      <GenericMessage
        on:close={() => notifications.remove(notification)}
        {notification} />
    {/if}
  {/each}
</div>
