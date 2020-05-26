package controllers

import (
	"fmt"
	"net/http"
)

const info = `<h1>Goexe Web Server</h1>
<body>
This web server recieves compilable code snippets and reponse with the correct output.
This route to send a request containing a code snippet is /submit. The request must be a POST request.
</body>
`

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, info)
}