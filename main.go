package main

import ("net/http"
		"io")

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index (w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
		<!DOCTYPE html>
		<html lang=en>
		<head>
			<meta charset="UTF-8">
			<title>Title</title>
		</head>
		<style>
			html, body, h1 {
				padding: 0;
				border: 0;
				margin: 0;
				box-sizing: border-box;
			}
			main {
				display: flex;
				justify-content: center;
				align-items: center;
				background-image: linear-gradient(green, white, red);
				height: 100vh;
			}
			h1 {
				font-size: 8rem;
				color: white;
			}
		</style>
		<body>
		<main>
			<ul>
				<h1>Hello Venus</h1>
			</ul>
		</main>
		<ul>
			<li>001</li>
			<li>002</li>
			<li>003</li>
			<li>004</li>
			<li>005</li>
			<li>006</li>
			<li>007</li>
			<li>008</li>
			<li>009</li>
			<li>010</li>
			<li>011</li>
			<li>012</li>
			<li>013</li>
			<li>014</li>
			<li>015</li>
			<li>016</li>
			<li>017</li>
			<li>018</li>
			<li>019</li>
			<li>020</li>
		</ul>


		</body>

		`)
}