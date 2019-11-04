<script>
  import page from "page";

  let isCreator = true;

  let username = "";
  let roomName = "";
  let accessCode = "";

  $: noRoom = roomName === "";
  $: noUser = username === "";
  $: noCode = accessCode === "";
</script>

<!-- <div class="h-screen flex items-center justify-center">
  <div
    class="w-full max-w-sm rounded-lg bg-gray-100 shadow flex flex-col
    justify-center items: center">
    <input type="text" class="mb-4" />
    <div>
      <button
        on:click={() => page('/movie')}
        class="rounded text-black p-2 bg-gray-300 text-gray-600
        hover:text-green-600">
        Create
      </button>
    </div>
  </div>
</div> -->

<div class="h-screen flex items-center justify-center">
  <div class="w-full max-w-sm rounded-lg m-0">
    <ul class="w-full flex justify-center border-b cursor-pointer">
      <li class="{isCreator ? '-mb-px' : ''} flex-grow">
        <div
          on:click={() => (isCreator = true)}
          class="{isCreator ? 'border-l border-t border-r rounded-t text-blue-700' : 'text-blue-500'}
          py-2 px-4 font-semibold bg-white ">
          Create Video Room
        </div>
      </li>
      <li class="{!isCreator ? '-mb-px' : ''} flex-grow">
        <div
          on:click={() => (isCreator = false)}
          class="{!isCreator ? 'border-l border-t border-r rounded-t text-blue-700' : 'text-blue-500'}
          bg-white py-2 px-4 hover:text-blue-800 font-semibold">
          Join Video Room
        </div>
      </li>
    </ul>
    <div class="border border-t-0 rounded rounded-t-none px-8 pt-6 pb-8 mb-4">
      {#if isCreator}
        <form>
          <div class="mb-4">
            <label
              class="block text-gray-700 text-sm font-bold mb-2"
              for="username">
              Username
            </label>
            <input
              class="shadow appearance-none border rounded w-full py-2 px-3
              text-gray-700 leading-tight focus:outline-none
              focus:shadow-outline {noUser ? 'border-red-500' : ''}"
              id="username"
              type="text"
              bind:value={username}
              placeholder="John Smith" />
            {#if noUser}
              <p class="text-red-500 text-xs italic">
                Please enter a room name.
              </p>
            {/if}
          </div>

          <div class="mb-6">
            <label
              class="block text-gray-700 text-sm font-bold mb-2"
              for="room_name">
              Room Name
            </label>
            <input
              class="shadow appearance-none border rounded w-full py-2 px-3
              text-gray-700 mb-3 leading-tight focus:outline-none
              focus:shadow-outline {noRoom ? 'border-red-500' : ''}"
              id="room_name"
              type="text"
              bind:value={roomName}
              placeholder="John's Movie Night" />
            {#if noRoom}
              <p class="text-red-500 text-xs italic">
                Please enter a room name.
              </p>
            {/if}

          </div>
          <div class="flex items-center justify-center">
            <button
              on:click={() => page('/movie')}
              class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2
              px-4 rounded focus:outline-none focus:shadow-outline"
              type="button">
              Create Room
            </button>
          </div>
        </form>
      {:else}
        <form>
          <div class="mb-4">
            <label
              class="block text-gray-700 text-sm font-bold mb-2"
              for="username">
              Username
            </label>
            <input
              class="shadow appearance-none border rounded w-full py-2 px-3
              text-gray-700 leading-tight focus:outline-none
              focus:shadow-outline {noUser ? 'border-red-500' : ''}"
              id="username"
              type="text"
              bind:value={username}
              placeholder="John Smith" />
            {#if noUser}
              <p class="text-red-500 text-xs italic">
                Please enter a room name.
              </p>
            {/if}
          </div>

          <div class="mb-6">
            <label
              class="block text-gray-700 text-sm font-bold mb-2"
              for="room_code">
              Room Access Code
            </label>
            <input
              class="shadow appearance-none border rounded w-full py-2 px-3
              text-gray-700 mb-3 leading-tight focus:outline-none
              focus:shadow-outline {noCode ? 'border-red-500' : ''}"
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
            <button
              on:click={() => page('/movie')}
              class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2
              px-4 rounded focus:outline-none focus:shadow-outline"
              type="button">
              Join Room
            </button>
          </div>
        </form>
      {/if}

    </div>
  </div>
</div>
