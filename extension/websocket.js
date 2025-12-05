var token = ""
getAccessToken()
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

async function getAccessToken(){
  const refresh_token = "AQADe7SihJ0PrFqyckmgCJBuak-R0RchrsLk7jxL3es5OQU_jJnSi4olIJ0LS0mfZ06s4eeX-wHgL9LHpcP9fEGwBbyBfx6T4sGq-jF3fQ6c06ucV5u0UUEprpMCK6fJJKM"

    const u = "https://accounts.spotify.com/api/token"

    const body = new URLSearchParams();
    body.append("grant_type","refresh_token")
    body.append("refresh_token",refresh_token)

    let response = await fetch(u,{
        method:"POST",
        headers: {
            "Content-Type" : "application/x-www-form-urlencoded",
            "Authorization": "Basic " + btoa("b498ca8c85ff4c53a454834caf7b5edd" + ":" + "b6da51fec068455da6531798897c7431")
        },
        body
    })

    let data = await response.json();

    token = data.access_token
}


async function getCurrentPlaying() {
  if (!token) {
        console.log("error in access_token:", data);
        return;
    }

    const url = "https://api.spotify.com/v1/me/player/currently-playing";

    let response = await fetch(url, {
        method: "GET",
        headers: { "Authorization": `Bearer ${token}` }
    });

    let json = {}

    try{
        let track = await response.json();

        json = {
        name: track.item.name,
        band: track.item.artists[0].name,
        image: track.item.album.images[0].url,
        duration: track.item.duration_ms
    }

    }catch(error){

        json = {
            name: "Not playing anything!",
            band: "Spotify",
            image: "n/a",
            duration: 111111
        }
    }

    console.log(json)

    return {
      sender: "spicetify",
      message: JSON.stringify(json)
    }
}