package main

import (
	"flag"
	"fmt"
	"net/http"

	"golang.org/x/net/webdav"
)

var (
	addr = flag.String("addr", ":1744", "")
	path = flag.String("path", ".", "") // 文件路径，默认当前目录
)

func main() {
	flag.Parse()
	fmt.Println("addr=", *addr, ", path=", *path)
	http.ListenAndServe(*addr, &webdav.Handler{
		FileSystem: webdav.Dir(*path),
		LockSystem: webdav.NewMemLS(),
	})
}
