<script>
  import { push } from "svelte-spa-router";
  import { user } from "../store/state";
  import LoginForm from "../components/LoginForm.svelte";

  let isCreator = true;

  function createRoom(event) {
    event.preventDefault();
    user.login({
      type: "streamer",
      ...event.detail
    });
    push("/movie");
  }

  function joinRoom(event) {
    event.preventDefault();
    user.login({
      type: "viewer",
      ...event.detail
    });
    push("/movie");
  }
</script>

<div class="h-full flex items-center justify-center">

  <div class="w-full max-w-md rounded-lg m-0">
    <!-- Tab Bar -->
    <ul
      class="w-full flex justify-center cursor-pointer relative z-0">
      <li class="{isCreator && 'shadow-lg'} flex-grow rounded-lg rounded-b-none">
        <div
          on:click={() => (isCreator = true)}
          class="{isCreator ? 'rounded-t-lg text-gray-900 bg-gray-100' : 'text-gray-500'}
          p-4 font-semibold text-center">
          Create Video Room
        </div>
      </li>
      <li class="{!isCreator && 'shadow-lg'} flex-grow rounded-lg rounded-b-none">
        <div
          on:click={() => (isCreator = false)}
          class="{!isCreator ? 'rounded-t-lg text-gray-900 bg-gray-100' : 'text-gray-500'}
          p-4 font-semibold text-center">
          Join Video Room
        </div>
      </li>
    </ul>
    <!-- Form Content Container -->
    <div
      class="px-8 py-10 {isCreator ?'rounded-tl-none':'rounded-tr-none'} bg-gray-100 rounded-lg
      z-10 relative shadow-lg">
      <LoginForm on:create={createRoom} on:join={joinRoom} create={isCreator} />
    </div>
  </div>
</div>
