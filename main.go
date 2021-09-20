package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
		<html>
			<head>
				<link rel="stylesheet" href="https://devspace.sh/css/quickstart.css">
			</head>
			<body>
				<img src="https://devspace.sh/images/congrats.gif" />
				<h1>You deployed this project with DevSpace!</h1>
				<div>
					<h2>Now it's time to code:</h2>
					<ol>
						<li>Edit this text in <code>main.go</code> and save the file</li>
						<li>Restart the application with <code>go run main.go</code></li>
						<li>Reload browser to see the changes</li>
					</ol>
				</div>
			</body>
		</html>
		`)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
