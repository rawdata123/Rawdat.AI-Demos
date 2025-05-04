package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type TravoltaQuote struct {
	Movie string `json:"movie"`
	Quote string `json:"quote"`
	Style string `json:"style,omitempty"`
}

var quotes = []TravoltaQuote{
	{"Pulp Fiction", "You know what they call a Quarter Pounder with Cheese in Paris?", ""},
	{"Pulp Fiction", "I got my technique down and everything, I don't be tickling or nothin'.", "cool"},
	{"Grease", "Tell me about it, stud.", "leather jacket"},
	{"Saturday Night Fever", "You paint your hair. I work on my hair a long time and you hit it.", "disco"},
	{"Face/Off", "I want to take his face... off.", "unhinged"},
	{"Look Who's Talking", "You're the only one who knows what it's like to be me.", ""},
	{"Hairspray", "Our world is changing, boys and girls!", "drag queen energy"},
	{"Battlefield Earth", "While you were still learning how to SPELL YOUR NAME, I was conquering galaxies!", "so bad it’s good"},
}

func getRandomQuote() TravoltaQuote {
	rand.Seed(time.Now().UnixNano())
	return quotes[rand.Intn(len(quotes))]
}

// Serve HTML directly from Go
func uiHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Travolta Quote Machine</title>
  <style>
    body {
  font-family: 'Comic Sans MS', cursive;
  margin: 0;
  padding: 0;
  height: 100%;
  overflow: hidden;
  text-align: center;
  color: white;
  background: linear-gradient(45deg, #ff00cc, #3333ff);
  background-size: 400% 400%;
  animation: flashlights 6s ease infinite;
}

@keyframes flashlights {
  0% { background-position: 0% 50%; }
  25% { background-position: 50% 50%; background-color: #1e90ff; }
  50% { background-position: 100% 50%; background-color: #ffcc00; }
  75% { background-position: 50% 50%; background-color: #ff1493; }
  100% { background-position: 0% 50%; background-color: #00ffff; }
}
    html, body {
      margin: 0;
      padding: 0;
      height: 100%;
      overflow: hidden;
      font-family: 'Comic Sans MS', cursive;
      background: black;
      color: white;
    }

    .floor {
      position: absolute;
      top: 0;
      left: 0;
      height: 100%;
      width: 100%;
      display: grid;
      grid-template-columns: repeat(10, 1fr);
      grid-template-rows: repeat(6, 1fr);
      z-index: -1;
    }

    .tile {
      animation: flash 3s infinite;
    }

    @keyframes flash {
      0%, 100% { background: #ff1493; }
      25% { background: #1e90ff; }
      50% { background: #32cd32; }
      75% { background: #ffa500; }
    }

    #quoteBox {
  font-size: 1.5em;
  margin: 30px auto;
  padding: 20px;
  border: 2px dashed #ffcc00;
  width: 60%;
  background: rgba(255, 255, 255, 0.15);
  animation: pulse 3s infinite;
}

@keyframes pulse {
  0%, 100% { box-shadow: 0 0 20px #ff1493; }
  50% { box-shadow: 0 0 40px #1e90ff; }
}


    .tile:nth-child(even) {
      animation-delay: 1s;
    }

    .tile:nth-child(odd) {
      animation-delay: 2s;
    }

    .container {
      position: relative;
      text-align: center;
      padding: 50px;
      z-index: 1;
    }

    h1 {
      font-size: 3em;
      color: #ffcc00;
      text-shadow: 2px 2px #800080;
    }

    #quoteBox {
      font-size: 1.5em;
      margin: 30px auto;
      padding: 20px;
      border: 2px dashed #ffcc00;
      width: 60%;
      background: rgba(255, 255, 255, 0.1);
    }

    button {
      padding: 15px 30px;
      font-size: 1em;
      margin: 10px;
      background: #800080;
      color: white;
      border: none;
      border-radius: 5px;
      cursor: pointer;
    }

    button:hover {
      background: #a000a0;
    }

    #discoBall {
      position: absolute;
      top: 20px;
      right: 20px;
      width: 80px;
      height: 80px;
      z-index: 2;
      animation: spin 6s linear infinite;
    }

    @keyframes spin {
      0% { transform: rotate(0deg); }
      100% { transform: rotate(360deg); }
    }
  </style>
</head>
<body>
  <img src="https://media.giphy.com/media/3o6Zt481isNVuQI1l6/giphy.gif" alt="Disco Ball" id="discoBall" />
  
  <div class="floor" id="danceFloor"></div>

  <div class="container">
    <h1>🕺 Travolta Quote Machine</h1>
    <button onclick="getQuote()">Hit Me With a Quote</button>
    <button onclick="playMusic()">🎧 Play Some Funk</button>
    <button onclick="toggleDance()">💃 DANCE MODE</button>
    <div id="quoteBox">Click the button to get a Travolta classic!</div>
    <div id="gifBox">
  <img id="travoltaGif" src="" alt="Travolta GIF" style="max-width: 400px; margin-top: 20px; display: none;">
</div>
  </div>

  <audio id="discoAudio" loop>
    <source src="https://www.soundhelix.com/examples/mp3/SoundHelix-Song-8.mp3" type="audio/mpeg">
    Your browser does not support the audio element.
  </audio>

  <script>
  window.onload = function () {
    const floor = document.getElementById('danceFloor');
    const audio = document.getElementById('discoAudio');

    async function getQuote() {
  const res = await fetch('/quote');
  const data = await res.json();
  document.getElementById('quoteBox').innerHTML =
    '<p><strong>"' + data.quote + '"</strong></p><p>– <em>' + data.movie + '</em></p>';

  const gifs = [
    "https://media.giphy.com/media/3o6ZtpxSZbQRRnwCKQ/giphy.gif",   // classic confused Travolta
    "https://media.giphy.com/media/l0MYt5jPR6QX5pnqM/giphy.gif",     // Travolta dancing
    "https://media.giphy.com/media/10cU0MYvS6YkU4/giphy.gif",        // white suit dance
    "https://media.giphy.com/media/3oEduN5sfUqV0jDFyQ/giphy.gif",    // face off moment
    "https://media.giphy.com/media/xT0xezQGU5xCDJuCPe/giphy.gif",    // Grease dance
    "https://media.giphy.com/media/l0ExncehJzexFpRHq/giphy.gif",     // dramatic pose
    "https://media.giphy.com/media/3o7TKtnuHOHHUjR38Y/giphy.gif"     // finger pointing disco king
  ];

  const gif = gifs[Math.floor(Math.random() * gifs.length)];
  const gifEl = document.getElementById('travoltaGif');
  gifEl.src = gif;
  gifEl.style.display = "block";
}


    function playMusic() {
      if (audio.paused) {
        audio.play();
      } else {
        audio.pause();
      }
    }

    function toggleDance() {
      if (floor.childElementCount === 0) {
        for (let i = 0; i < 60; i++) {
          const tile = document.createElement('div');
          tile.className = 'tile';
          floor.appendChild(tile);
        }
      } else {
        floor.innerHTML = '';
      }
    }

    // Expose functions to global scope
    window.getQuote = getQuote;
    window.playMusic = playMusic;
    window.toggleDance = toggleDance;
  };
</script>
</body>
</html>
`
	t := template.Must(template.New("travolta").Parse(tmpl))
	t.Execute(w, nil)
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	quote := getRandomQuote()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func main() {
	http.HandleFunc("/", uiHandler)
	http.HandleFunc("/quote", quoteHandler)

	port := ":1978"
	log.Printf("Travoltaserver is strutting on http://localhost%s 🕺", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
