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
      <LoginForm on:create={createRoom} on:join={joinRoom} create={isCreator} />
    </div>
  </div>
</div>
