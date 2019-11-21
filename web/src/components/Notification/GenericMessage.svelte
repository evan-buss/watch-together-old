<script>
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();
  export let notification;
</script>

<div class="relative rounded shadow w-full max-w-md bg-white p-2 mb-2">
  <div class="flex flex-row w-full justify-between">
    <h1 class="font-bold text-blue-600 truncate">{notification.title}</h1>
    <i
      on:click={() => dispatch('close')}
      class="absolute right-0 top-0 mt-1 mr-1 la la-times-circle text-2xl
    text-red-500 hover:text-red-300 cursor-pointer" />
  </div>
  <span class="text-gray-600 break-words">{notification.message}</span>
</div>
