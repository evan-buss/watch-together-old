<script>
  import { user } from "../store/state";

  export let details = {
    sender: "Joe Smith",
    message: "Default Message"
  };

  let myMessage = $user.name === details.sender;

  // Change the color and border of messages
  let messageStyle = myMessage
    ? "bg-blue-400 rounded-br-none"
    : "bg-green-400 rounded-bl-none";
  // Change the border depending on the message length
  let lineStyle = details.message.length > 30 ? "rounded-lg" : "rounded-full";
</script>

<li class="{myMessage ? 'self-end' : 'self-start'} px-1 mb-1">
  <div
    class="{myMessage ? 'text-right' : 'text-left'} px-2 text-base text-gray-600">
    {details.sender}
  </div>
  <div class="text-white break-word px-2 py-1 {lineStyle} {messageStyle}">
    {details.message}
  </div>
</li>
