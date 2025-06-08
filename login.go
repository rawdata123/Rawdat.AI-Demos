package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Login</title>
	<!-- Font Awesome CDN for icons -->
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/css/all.min.css" integrity="sha512-..." crossorigin="anonymous" 
referrerpolicy="no-referrer" />
	<style>
		body {
			font-family: Arial, sans-serif;
			background: url('https://media.giphy.com/media/VbnUQpnihPSIgIXuZv/giphy.gif') no-repeat center center fixed;
			background-size: cover;
			display: flex;
			justify-content: center;
			align-items: center;
			height: 100vh;
			margin: 0;
		}
		.login-container {
			background: rgba(255, 255, 255, 0.95);
			padding: 30px 40px;
			box-shadow: 0px 0px 15px rgba(0,0,0,0.3);
			border-radius: 10px;
			width: 320px;
			backdrop-filter: blur(5px);
		}
		h2 {
			text-align: center;
			margin-bottom: 20px;
			color: #333;
		}
		.input-group {
			position: relative;
			margin-top: 15px;
		}
		.input-group i {
			position: absolute;
			top: 12px;
			left: 10px;
			color: #aaa;
		}
		.input-group input {
			width: 100%;
			padding: 10px 10px 10px 35px;
			border: 1px solid #ccc;
			border-radius: 4px;
			box-sizing: border-box;
		}
		input[type="submit"] {
			width: 100%;
			background-color: #4CAF50;
			color: white;
			padding: 10px;
			margin-top: 20px;
			border: none;
			border-radius: 4px;
			cursor: pointer;
			font-size: 16px;
		}
		input[type="submit"]:hover {
			background-color: #45a049;
		}
	</style>
</head>
<body>
	<div class="login-container">
		<h2><i class="fas fa-sign-in-alt"></i> Login</h2>
		<form action="/login" method="POST">
			<div class="input-group">
				<i class="fas fa-user"></i>
				<input type="text" name="username" placeholder="Username" required>
			</div>
			<div class="input-group">
				<i class="fas fa-lock"></i>
				<input type="password" name="password" placeholder="Password" required>
			</div>
			<input type="submit" value="Login">
		</form>
	</div>
</body>
</html>
`
func main() {
	http.HandleFunc("/", serveLoginForm)
	http.HandleFunc("/login", handleLogin)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func serveLoginForm(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("login").Parse(tpl))
	t.Execute(w, nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Hardcoded user validation
	if username == "admin" && password == "1234" {
		http.Redirect(w, r, "https://juicebox.publicvm.com", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
