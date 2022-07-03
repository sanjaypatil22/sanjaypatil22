package main

import (
	"html/template"
	"net/http"
	"os"
)

type Context struct {
	Version, Commit, Hostname string
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	var hostname string
	var err error

	hostname, err = os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	context := Context{
		Version:  "release.Version",
		Commit:   "release.Commit",
		Hostname: hostname,
	}

	page := `
		<html>
			<head>
				<title>Simple Go Helloworld</title>
			</head>

			<style>
				.data{
					font-family: Arial;
					margin: 0 auto;
					width: 500px;
				}
				.data h1 {
					padding: 5px;
					background-color: #ff6347;
					color: white;
					text-align: center;
				}
				.data h2 {
					color: #808080;
					text-align: center;
				}
				.release {
					text-align: right;	
				}
			</style>

			<body>
				<div class="data">
					<h1>Simple Go Helloworld</h1>
					<h2>Response from: {{.Hostname}}</h2>
				</div>
			</body>
		</html>
	`

	t := template.New("Simple Go Helloworld")
	t, _ = t.Parse(page)
	t.Execute(w, context)
}

func main() {
	http.HandleFunc("/", helloworld)
	http.HandleFunc("/hello", helloworld)
	http.ListenAndServe(":3000", nil)
}
