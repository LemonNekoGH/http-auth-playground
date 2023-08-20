package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// Try to implement HTTP Basic Auth which specified in RFC 7617
type Handler struct{}

func (h Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	authStr := req.Header.Get("Authorization")
	// check header
	if authStr == "" || !strings.HasPrefix(authStr, "Basic ") {

		return
	}

	// check username and password
	authStr = strings.TrimPrefix(authStr, "Basic ")
	decoded, err := base64.StdEncoding.DecodeString(authStr)
	if err != nil {
		resp.Header().Add("WWW-Authenticate", "Basic realm=\"protected\"")
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	usernameAndPassword := strings.Split(string(decoded), ":")
	// incorrect
	if usernameAndPassword[0] != "Foo" || usernameAndPassword[1] != "Bar" {
		resp.Header().Add("WWW-Authenticate", "Basic realm:\"protected\"")
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	// correct
	resp.Header().Add("Content-Type", "text/html")
	resp.Write([]byte("<code>Authorized!</code>"))
}

func main() {
	http.ListenAndServe("0.0.0.0:5444", Handler{})
}
