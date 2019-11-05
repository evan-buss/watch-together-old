<script>
  import { push } from "svelte-spa-router";
  import { user } from "../store/state";

  let isCreator = true;

  let username = "";
  let roomname = "";
  let accessCode = "";

  $: noRoom = roomname === "";
  $: noUser = username === "";
  $: noCode = accessCode === "";

  function createRoom(event) {
    event.preventDefault();
    user.login({
      type: "streamer",
      name: username,
      room: roomname,
      code: "abc123"
    });
    push("/movie");
  }

  function joinRoom(event) {
    event.preventDefault();
    user.login({
      type: "viewer",
      name: username,
      code: accessCode
    });
    // Check with the server, ensure the room code is valid,
    // push("/movie/roomCode");
    // else push('/')
    push("/movie");
  }
</script>

<div class="h-full flex items-center justify-center">
  <div class="w-full max-w-sm rounded-lg m-0">
    <ul
      class="w-full flex justify-center border-b cursor-pointer border-gray-600">
      <li class="{isCreator && '-mb-px'} flex-grow">
        <div
          on:click={() => (isCreator = true)}
          class="{isCreator ? 'border-gray-600 border-l border-t border-r rounded-t text-blue-500 bg-gray-800' : 'text-blue-800'}
          py-2 px-4 font-semibold">
          Create Video Room
        </div>
      </li>
      <li class="{!isCreator && '-mb-px'} flex-grow">
        <div
          on:click={() => (isCreator = false)}
          class="{!isCreator ? 'border-gray-600 border-l border-t border-r rounded-t text-blue-500 bg-gray-800' : 'text-blue-800'}
          py-2 px-4 font-semibold">
          Join Video Room
        </div>
      </li>
    </ul>
    <div
      class="border border-t-0 border-gray-600 bg-gray-800 rounded
      rounded-t-none px-8 pt-6 pb-8 mb-4">
      {#if isCreator}
        <form>
          <div class="mb-4">
            <label
              class="block text-gray-400 text-sm font-bold mb-2"
              for="username">
              Username
            </label>
            <input
              class="shadow appearance-none border-2 rounded w-full py-2 px-3
              text-gray-400 leading-tight focus:outline-none
              focus:shadow-outline bg-gray-700 {noUser && 'border-red-500'}"
              id="username"
              type="text"
              bind:value={username}
              placeholder="John Smith" />
            {#if noUser}
              <p class="text-red-500 text-xs italic">
                Please enter a username.
              </p>
            {/if}
          </div>

          <div class="mb-6">
            <label
              class="block text-gray-400 text-sm font-bold mb-2"
              for="room_name">
              Room Name
            </label>
            <input
              class="shadow appearance-none border-2 rounded w-full py-2 px-3
              text-gray-400 bg-gray-700 mb-3 leading-tight focus:outline-none
              focus:shadow-outline {noRoom && 'border-red-500'}"
              id="room_name"
              type="text"
              bind:value={roomname}
              placeholder="John's Movie Night" />
            {#if noRoom}
              <p class="text-red-500 text-xs italic">
                Please enter a room name.
              </p>
            {/if}

          </div>
          <div class="flex items-center justify-center">
            <input
              on:click={createRoom}
              class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2
              px-4 rounded focus:outline-none focus:shadow-outline"
              type="submit"
              value="Create Room" />
          </div>
        </form>
      {:else}
        <form>
          <div class="mb-4">
            <label
              class="block text-gray-400 text-sm font-bold mb-2"
              for="username">
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
              <p class="text-red-500 text-xs italic">
                Please enter a username.
              </p>
            {/if}
          </div>

          <div class="mb-6">
            <label
              class="block text-gray-400 text-sm font-bold mb-2"
              for="room_code">
              Room Access Code
            </label>
            <input
              class="shadow appearance-none border-2 rounded w-full py-2 px-3
              text-gray-400 bg-gray-700 mb-3 leading-tight focus:outline-none
              focus:shadow-outline {noCode && 'border-red-500'}"
              id="room_code"
              type="text"
              bind:value={accessCode}
              placeholder="abc123" />
            {#if noRoom}
              <p class="text-red-500 text-xs italic">
                Please enter a room access code.
              </p>
            {/if}

          </div>
          <div class="flex items-center justify-center">
            <input
              on:click={joinRoom}
              class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2
              px-4 rounded focus:outline-none focus:shadow-outline"
              type="submit"
              value="Join Room" />
          </div>
        </form>
      {/if}

    </div>
  </div>
</div>
