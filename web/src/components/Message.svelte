<script>
  export let details = {
    sender: "Joe Smith",
    message: "Default Message",
    sent: false
  };

  let messageStyle = details.sent
    ? "self-end bg-blue-400"
    : "self-start bg-green-400";
  let lineStyle = details.message.length > 30 ? "rounded-lg" : "rounded-full";
</script>


<div class="px-2 py-1 m-1 text-white {lineStyle} {messageStyle}">
  {details.message}
</div>
