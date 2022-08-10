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
		  <link rel="stylesheet" href="css/style.css">
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
				  <li>Edit <u>this text</u> within the <code>./main.go</code> file of this project and save the changes</li>
				  <li>Abort <code>go run main.go</code> via <code>CTRL/CMD + C</code></li>
				  <li>Run <code>go run main.go</code> again to restart the application</li>
				  <li>Reload the browser to see the changes</li>
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
