<script>
  import Clappr from "clappr";
  import { onMount } from "svelte";
  import { user } from "../store/state";
  import { createSocket } from "../store/socket";

  let player;
  let socket = createSocket($user.ip);

  onMount(() => {
    player = new Clappr.Player({
      source: `http://${$user.ip}/media/index.m3u8`,
      width: "100%",
      height: "100%",
      poster:
        "https://upload.wikimedia.org/wikipedia/commons/4/4d/Rembrandt_-_The_Anatomy_Lesson_of_Dr_Nicolaes_Tulp.jpg",
      parentId: "#player",
      resizable: true,
      chromeless: $user.type === "viewer"
    });

    // We listen to events on the streamers player to send to other viewers
    if ($user.type === "streamer") {
      player.listenTo(player, "pause", () => {
        socket.send("pause", {});
      });

      player.listenTo(player, "play", () => {
        socket.send("play", {});
      });

      player.listenTo(player, "seek", time => {
        console.log(`seek ${time}`);
        socket.send("seek", { time: time });
      });

      player.listenTo(player, "timeupdate", progress => {
        console.log(progress);
      });
    }

    // Only viewers should react to the streamers player events
    if ($user.type === "viewer") {
      socket.bind("play", () => {
        player.play();
      });

      socket.bind("pause", () => {
        player.pause();
      });

      socket.bind("seek", data => {
        player.seek(Math.round(data.time));
      });
    }
  });
</script>

<div id="player" class="w-full h-64 md:h-full md:p-4" />
