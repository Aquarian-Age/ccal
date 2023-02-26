package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

var (
	path = flag.String("path", ".", "")
	port = flag.String("port", ":1744", "")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println("username: openSUSE password: Tumbleweed")
	}
	flag.Parse()

	fs := &webdav.Handler{
		FileSystem: webdav.Dir(*path),
		LockSystem: webdav.NewMemLS(),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		username, password, ok := req.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if username != "openSUSE" || password != "Tumbleweed" {
			http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
			return
		}
		fs.ServeHTTP(w, req)
	})
	fmt.Println("port=", *port, ", path=", *path)
	http.ListenAndServe(*port, nil)
}
