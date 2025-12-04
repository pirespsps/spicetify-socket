main();

async function main() {

  getCurrentPlaying()

  setTimeout(() => {
    Spicetify.Player.play()
  }, 5000)

  const ws = new WebSocket("ws://localhost:8080/ws");

  ws.onopen = async () => {
    console.log("spicetify connected")

    ws.send(JSON.stringify({
      sender: "spicetify",
      message: ""
    }))

    setTimeout(async () => {
      ws.send(JSON.stringify(await currentJson()))
    }, 3000)

  };

  Spicetify.Player.addEventListener("songchange", (event) => {
    ws.send(JSON.stringify(currentJson()))
  });

  ws.onmessage = async (event) => {
    const msg = JSON.parse(event.data);

    if (msg.message === "play") {

      if (Spicetify.Player.isPlaying()) {
        Spicetify.Player.pause();
      } else {
        Spicetify.Player.play();
      }

      ws.send(
        JSON.stringify(
          {
            "sender": "spicetify",
            "message": "ok"
          }
        )
      )
    }

    if (msg.message === "next") {
      Spicetify.Player.next()
      ws.send(JSON.stringify(currentJson()))
    }
    if (msg.message === "previous") {
      Spicetify.Player.back()
      ws.send(JSON.stringify(currentJson()))
    };
    if (msg.message == "current") {
      ws.send(JSON.stringify(currentJson()))
    }
  };

}


function getCurrentPlaying() {
  const url = "https://accounts.spotify.com/api/token"

  let token = "";

  const data = {
    grant_type: "client_credentials",
    client_id: "b498ca8c85ff4c53a454834caf7b5edd",
    client_secret: "b6da51fec068455da6531798897c7431",
  }

  fetch(url, {
    method: "POST",
    body: data,
    headers: ["Content-Type","application/x-www-form-urlencoded"],
  })
    .then((response) => {
      console.log(response)
      token = response.data.access_token
    })
    .catch((error) => {
      console.log(error)
    });

    return token;
}


async function currentJson() {

  let data = await Spicetify.CosmosAsync.get("https://api.spotify.com/v1/me/player/currently-playing")

  let track;
  let album;
  let artist;
  let duration;

  if (!data || !data.item) {

    data = await SPicetify.CosmosAsync.get("https://api.spotify.com/v1/me/player/recently-played?limit=1")
    track = data.items[0].name
    album = data.items[0].album.images[0].url
    artist = data.items[0].artists[0].name
    duration = data.items[0].duration_ms

  } else {

    track = data.item.name
    album = data.item.album.images[0].url
    artist = data.item.artists[0].name
    duration = data.item.duration_ms

  }

  const obj = {
    track: track,
    album: album,
    artist: artist,
    duration: duration
  }

  return {
    sender: "spicetify",
    message: JSON.stringify(obj)
  }
}