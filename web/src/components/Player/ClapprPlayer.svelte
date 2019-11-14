<script>
  import Clappr from "clappr";
  import { onMount } from "svelte";
  import { user } from "../../store/state";
  import { getSocket } from "../../store/socket";
  import PlayPause from "./PlayPause.svelte";
  import VideoProgress from "./VideoProgress.svelte";
  import Volume from "./Volume.svelte";

  let player;
  let socket = getSocket($user.ip);
  let isPaused;
  let volume;

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
        isPaused = true;
      });

      player.listenTo(player, "play", () => {
        socket.send("play", {});
        isPaused = false;
      });

      player.listenTo(player, "seek", time => {
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
        isPaused = false;
      });

      socket.bind("pause", () => {
        player.pause();
        isPaused = true;
      });

      socket.bind("seek", data => {
        player.seek(Math.round(data.time));
      });
    }

    // volume = player.getVolume();
  });

  // function setVolume(event) {
  //   player.setVolume(event.details.volume);
  // }
</script>

<div id="player" class="w-full h-full md:p-4 flex flex-col-reverse">
  <!-- <div class="flex flex-row py-4 px-4 md:px-0">
    <PlayPause
      {isPaused}
      on:toggle={() => (isPaused ? player.play() : player.pause())} />
    <VideoProgress />
    <Volume {volume} on:adjust={setVolume} />
  </div> -->
</div>
