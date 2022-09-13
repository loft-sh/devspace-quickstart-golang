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
				<title>DevSpace Demo</title>
				<link rel="stylesheet" href="https://devspace.sh/css/quickstart.css">
			</head>
			<body>
				<section>
					<div class="container">
						<div class="left">
							<h1>You deployed this project with<img class="logo" src="https://static.loft.sh/branding/logos/devspace/horizontal/mono/devspace_horizontal_mono.svg" alt="DevSpace" /></h1>
							<h2>Now it's time to code:</h2>
							<ol>
								<ul class="dots">
									<li class="red"></li>
									<li class="yellow"></li>
									<li class="blue"></li>
								</ul>
								<li>Edit this text in <code>main.go</code> and save the file</li>
								<li>Restart the application with <code>go run main.go</code></li>
								<li>Reload browser to see the changes</li>
							</ol>
						</div>
						<div><img src="https://static.loft.sh/devspace/quickstarts/devspace-astronaut.gif" /></div>
					</div>
				</section>
			</body>
		</html>
		`)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Started server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
