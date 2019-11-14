<script>
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  export let create;

  let username = "";
  let serverAddress = "";
  let buttonClicked = false;

  let noUser = false;
  let noServer = false;

  function dispatcher(event) {
    event.preventDefault();
    buttonClicked = true;
    noUser = username === "";
    noServer = serverAddress === "";

    if (!noUser && !noServer) {
      dispatch(create ? "create" : "join", {
        name: username,
        ip: serverAddress
      });
    }
  }
</script>

<form>
  <div class="mb-4">
    <label class="block text-gray-500 text-sm mb-2" for="username">
      Username
    </label>
    <input
      class="appearance-none rounded w-full py-2 px-3 text-gray-700 bg-gray-400
      leading-tight focus:outline-none focus:shadow-outline"
      id="username"
      type="text"
      bind:value={username}
      placeholder="John Smith" />
    {#if noUser}
      <p class="text-red-500 text-xs italic">Please enter a username.</p>
    {/if}
  </div>

  <div class="mb-6">
    <label class="block text-gray-500 text-sm mb-2" for="room_code">
      Video Server Address
    </label>
    <input
      class="appearance-none rounded w-full py-2 px-3 text-gray-700 bg-gray-400
       leading-tight focus:outline-none focus:shadow-outline"
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
  <div class="flex items-center justify-center mt-3">
    <button
      on:click={dispatcher}
      class="bg-gray-700 hover:bg-gray-900 focus:outline-none
      focus:shadow-outline text-white font-bold py-2 px-4 rounded cursor-pointer">
      {create ? 'Create Room' : 'Join Room'}
    </button>
  </div>
</form>
