<script>
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  export let create;

  let username = "";
  let serverAddress = "";

  $: noServer = serverAddress === "";
  $: noUser = username === "";

  function dispatcher() {
    dispatch(create ? "create" : "join", {
      name: username,
      ip: serverAddress
    });
  }
</script>

<form>
  <div class="mb-4">
    <label class="block text-gray-400 text-sm font-bold mb-2" for="username">
      Username
    </label>
    <input
      class="shadow appearance-none border-2 rounded w-full py-2 px-3
      text-gray-400 bg-gray-700 leading-tight focus:outline-none
      focus:shadow-outline {noUser && 'border-red-500'}"
      id="username"
      type="text"
      bind:value={username}
      placeholder="John Smith" />
    {#if noUser}
      <p class="text-red-500 text-xs italic">Please enter a username.</p>
    {/if}
  </div>

  <div class="mb-6">
    <label class="block text-gray-400 text-sm font-bold mb-2" for="room_code">
      Video Server Address
    </label>
    <input
      class="shadow appearance-none border-2 rounded w-full py-2 px-3
      text-gray-400 bg-gray-700 mb-3 leading-tight focus:outline-none
      focus:shadow-outline {noServer && 'border-red-500'}"
      id="room_code"
      type="text"
      bind:value={serverAddress}
      placeholder="192.168.1.1:8080" />
    {#if noServer}
      <p class="text-red-500 text-xs italic">
        {#if create}
          Enter the IP address and port number of your server.
        {:else}
          Enter the IP address and port number of your friend's server.
        {/if}
      </p>
    {/if}

  </div>
  <div class="flex items-center justify-center">
    <input
      on:click={dispatcher}
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4
      rounded focus:outline-none focus:shadow-outline"
      type="submit"
      value={create ? 'Create Room' : 'Join Room'} />
  </div>
</form>
