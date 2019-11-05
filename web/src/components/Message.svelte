<script>
  import { user } from "../store/state";

  export let details = {
    sender: "Joe Smith",
    message: "Default Message"
  };

  let myMessage = $user.name === details.sender;

  // Change the color and border of messages
  let messageStyle = myMessage ? "bg-blue-400 ml-4" : "bg-green-400 mr-4";
</script>

<li class="{myMessage ? 'self-end' : 'self-start'} px-1 my-1">
  {#if !myMessage}
  <div
    class="{myMessage ? 'text-right' : 'text-left'} px-2 text-xs text-gray-200">
    {details.sender}
  </div>
  {/if}
  <div class="rounded-lg text-white break-word px-2 py-1 {messageStyle}">
    {details.message}
  </div>
</li>
