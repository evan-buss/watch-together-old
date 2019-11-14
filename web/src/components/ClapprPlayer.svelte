<script>
  import Clappr from "clappr";
  import { onMount } from "svelte";
  import { user } from "../store/state";
  import { getSocket } from "../store/socket";

  let player;
  let socket = getSocket($user.ip);

  onMount(() => {
    player = new Clappr.Player({
      source: `http://${$user.ip}/media/index.m3u8`,
      width: "100%",
      height: "100%",
      poster:
        "https://cdn.collider.com/wp-content/uploads/2010/06/inception-movie-poster-7.jpg",
      parentId: "#player",
      resizable: true,
      chromeless: $user.type === "viewer",
      playback: {
        externalTracks: [
          {
            lang: "en",
            label: "English",
            src: `/media/index.vtt`,
            kind: "subtitles"
          }
        ]
      }
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
        // console.log(`seek ${time}`);
        socket.send("seek", { time: time });
      });

      player.listenTo(player, "timeupdate", progress => {
        // console.log(progress);
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
