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
				<link rel="stylesheet" href="https://devspace.cloud/quickstart.css">
			</head>
			<body>
				<img src="https://devspace.cloud/img/congrats.gif" />
				<h1>You deployed this project with DevSpace!</h1>
				<div>
					<h2>Now it's time to code:</h2>
					<ol>
						<li>Edit this text in <code>main.go</code> and save the file</li>
						<li>Check the logs to see how DevSpace restarts your container</li>
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
