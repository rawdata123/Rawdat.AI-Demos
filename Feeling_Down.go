package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", energyHandler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func energyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Energy Sunshine</title>
	<style>
		body {
			background: linear-gradient(to top, #ffe259, #ffa751 90%);
			font-family: 'Segoe UI', sans-serif;
			color: #333;
			text-align: center;
			padding: 0;
			margin: 0;
			min-height: 100vh;
		}
		.header {
			margin-top: 40px;
			font-size: 2.5rem;
			font-weight: bold;
			color: #ff9800;
			text-shadow: 1px 1px 8px #fff3e0;
		}
		.sunshine {
			font-size: 5rem;
		}
		.balloons {
			margin: 30px 0;
		}
		.message {
			font-size: 1.3rem;
			margin: 30px auto 25px auto;
			max-width: 500px;
			background: rgba(255,255,255,0.7);
			border-radius: 12px;
			padding: 18px 12px;
			box-shadow: 0 2px 12px #ffecb3;
		}
		iframe {
			border-radius: 12px;
			box-shadow: 0 2px 12px #ffecb3;
		}
	</style>
</head>
<body>
	<div class="header">
		<span class="sunshine">☀️</span> Welcome to Energy Sunshine! <span 
class="sunshine">☀️</span>
	</div>
	<div class="balloons">
		<span style="font-size:3rem;">🎈🎈🎈</span>
	</div>
	<div class="message">
		Feeling a bit down? Take a deep breath, smile, and let the sunshine in!<br>
		Imagine yourself floating high above the clouds in a hot air balloon, the sun warming your 
face and the world below glowing with possibility.<br>
		You're doing better than you think. Keep going!
	</div>
	<!-- Uplifting music with sunshine and hot air balloons -->
	<div>
		<iframe width="360" height="215" 
src="https://www.youtube.com/embed/BdowBfeAIyc?autoplay=1&rel=0" title="Peaceful Calming Hot Air Balloon 
Music" allow="autoplay; encrypted-media" allowfullscreen></iframe>
	</div>
	<div style="margin-top:30px;">
		<span style="font-size:2rem;">🌞</span>
		<span style="font-size:2rem;">🎈</span>
		<span style="font-size:2rem;">🌤️</span>
	</div>
</body>
</html>
`)
}

