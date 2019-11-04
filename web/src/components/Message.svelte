<script>
  export let details = {
    sender: "Joe Smith",
    message: "Default Message",
    sent: false
  };

  let messageStyle = details.sent
    ? "bg-blue-400 rounded-br-none"
    : "bg-green-400 rounded-bl-none";
  let lineStyle = details.message.length > 30 ? "rounded-lg" : "rounded-full";
</script>

<li class="{details.sender ? 'self-end' : 'self-start'} mb-1">
  <div class="text-base text-gray-600 text-right">{details.sender}</div>
  <div class="text-white break-word px-2 py-1 {lineStyle} {messageStyle}">
    {details.message}
  </div>
</li>
