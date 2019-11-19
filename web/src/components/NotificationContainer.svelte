<script>
  import NotificationMessage from "./NotificationMessage.svelte";

  import { notifications } from "../store/notification";
</script>

<div class="z-30 fixed top-0 right-0 mt-2 mr-2 w-full max-w-sm flex flex-col">
  {#each $notifications as notification (notification.time)}
    <NotificationMessage
      on:close={() => notifications.remove(notification)}
      {notification} />
  {/each}
</div>
