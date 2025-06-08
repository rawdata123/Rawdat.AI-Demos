package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Set up handler for the root path
	http.HandleFunc("/", handleRoot)

	// Start server
	log.Println("Starting Flying Unicorns server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, htmlContent)
}

const htmlContent = `<!DOCTYPE html>
<html>
<head>
    <title>Flying Unicorns</title>
    <style>
        body {
            margin: 0;
            overflow: hidden;
            background: linear-gradient(to bottom, #87CEEB, #E0F7FA);
            font-family: Arial, sans-serif;
        }
        .container {
            position: relative;
            width: 100vw;
            height: 100vh;
            overflow: hidden;
        }
        .rainbow {
            position: absolute;
            width: 100%;
            height: 300px;
            bottom: 100px;
        }
        .unicorn {
            position: absolute;
            transition: transform 0.5s ease;
        }
        .controls {
            position: absolute;
            bottom: 20px;
            left: 50%;
            transform: translateX(-50%);
            background: rgba(255, 255, 255, 0.7);
            padding: 10px;
            border-radius: 10px;
            display: flex;
            gap: 10px;
        }
        button {
            padding: 8px 16px;
            background: #ff69b4;
            color: white;
            border: none;
            border-radius: 5px;
            cursor: pointer;
        }
        button:hover {
            background: #ff1493;
        }
    </style>
</head>
<body>
    <div class="container">
        <svg class="rainbow" viewBox="0 0 1000 300">
            <defs>
                <linearGradient id="grass" x1="0%" y1="0%" x2="0%" y2="100%">
                    <stop offset="0%" style="stop-color:#7cfc00;stop-opacity:1" />
                    <stop offset="100%" style="stop-color:#228B22;stop-opacity:1" />
                </linearGradient>
            </defs>
            
            <!-- Rainbow Arcs -->
            <path d="M-100,300 C200,50 800,50 1100,300" stroke="#FF0000" stroke-width="20" fill="none" />
            <path d="M-100,300 C200,70 800,70 1100,300" stroke="#FF7F00" stroke-width="20" fill="none" />
            <path d="M-100,300 C200,90 800,90 1100,300" stroke="#FFFF00" stroke-width="20" fill="none" />
            <path d="M-100,300 C200,110 800,110 1100,300" stroke="#00FF00" stroke-width="20" fill="none" />
            <path d="M-100,300 C200,130 800,130 1100,300" stroke="#0000FF" stroke-width="20" fill="none" />
            <path d="M-100,300 C200,150 800,150 1100,300" stroke="#4B0082" stroke-width="20" fill="none" />
            <path d="M-100,300 C200,170 800,170 1100,300" stroke="#9400D3" stroke-width="20" fill="none" />
            
            <!-- Ground/Grass -->
            <rect x="-100" y="300" width="1200" height="100" fill="url(#grass)" />
        </svg>
        
        <div id="unicorn-container"></div>
        
        <div class="controls">
            <button id="add-unicorn">Add Unicorn</button>
            <button id="clear-unicorns">Clear All</button>
        </div>
    </div>

    <script>
        // Define unicorn SVG as a string without template literals
        const unicornSvg = '<svg width="100" height="80" viewBox="0 0 100 80" class="unicorn">' +
            '<ellipse cx="50" cy="50" rx="30" ry="20" fill="#FFFFFF" />' +
            '<rect x="30" y="65" width="5" height="20" fill="#FFFFFF" />' +
            '<rect x="40" y="65" width="5" height="20" fill="#FFFFFF" />' +
            '<rect x="60" y="65" width="5" height="20" fill="#FFFFFF" />' +
            '<rect x="70" y="65" width="5" height="20" fill="#FFFFFF" />' +
            '<ellipse cx="20" cy="45" rx="15" ry="10" fill="#FFFFFF" />' +
            '<path d="M30,35 Q40,20 50,35 Q60,15 70,35" stroke="#FF69B4" stroke-width="5" fill="none" />' +
            '<path d="M80,45 Q95,30 90,55" stroke="#FF69B4" stroke-width="5" fill="none" />' +
            '<path d="M15,35 L5,15" stroke="gold" stroke-width="3" />' +
            '<circle cx="15" cy="42" r="2" fill="#000000" />' +
            '</svg>';

        // Initialize
        document.addEventListener('DOMContentLoaded', function() {
            const container = document.getElementById('unicorn-container');
            const addButton = document.getElementById('add-unicorn');
            const clearButton = document.getElementById('clear-unicorns');
            
            // Random color generator
            function getRandomColor() {
                const letters = '0123456789ABCDEF';
                let color = '#';
                for (let i = 0; i < 6; i++) {
                    color += letters[Math.floor(Math.random() * 16)];
                }
                return color;
            }
            
            // Create unicorn function
            function createUnicorn() {
                const unicornDiv = document.createElement('div');
                unicornDiv.innerHTML = unicornSvg;
                const unicorn = unicornDiv.firstElementChild;
                
                // Random position
                const x = Math.random() * (window.innerWidth - 100);
                const y = Math.random() * (window.innerHeight - 300);
                
                // Random size
                const scale = 0.5 + Math.random() * 1.5;
                
                // Random color tint
                const color = getRandomColor();
                const mane = unicorn.querySelector('path[stroke="#FF69B4"]');
                const tail = unicorn.querySelectorAll('path[stroke="#FF69B4"]')[1];
                if (mane) mane.setAttribute('stroke', color);
                if (tail) tail.setAttribute('stroke', color);
                
                unicorn.style.left = x + 'px';
                unicorn.style.top = y + 'px';
                unicorn.style.transform = 'scale(' + scale + ')';
                
                container.appendChild(unicorn);
                
                // Animate the unicorn
                animateUnicorn(unicorn);
            }
            
            // Animation function
            function animateUnicorn(unicorn) {
                const speed = 1 + Math.random() * 5;
                const direction = Math.random() > 0.5 ? 1 : -1;
                let position = parseFloat(unicorn.style.left) || 0;
                
                function move() {
                    position += speed * direction;
                    
                    // Boundary check and reverse direction
                    if (position > window.innerWidth) {
                        position = -100;
                    } else if (position < -100) {
                        position = window.innerWidth;
                    }
                    
                    unicorn.style.left = position + 'px';
                    
                    // Floating effect
                    const floatY = Math.sin(position / 50) * 10;
                    unicorn.style.marginTop = floatY + 'px';
                    
                    // Continue animation
                    requestAnimationFrame(move);
                }
                
                move();
            }
            
            // Event listeners
            addButton.addEventListener('click', createUnicorn);
            clearButton.addEventListener('click', function() {
                container.innerHTML = '';
            });
            
            // Create initial unicorns
            for (let i = 0; i < 5; i++) {
                createUnicorn();
            }
        });
    </script>
</body>
</html>`
