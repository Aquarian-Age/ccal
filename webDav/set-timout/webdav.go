package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/net/webdav"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

var (
	dir      = flag.String("dir", ".", "share directory path")
	addr     = flag.String("addr", "", "IP Address")
	port     = flag.Int("port", 1744, "port num")
	ssl      = flag.Bool("ssl", false, "https: true http: false")
	key      = flag.String("ssl-key", "key.pem", "https key file")
	cert     = flag.String("ssl-cert", "cert.pem", "https cert file")
	user     = flag.String("user", "", "user name")
	pass     = flag.String("pass", "", "password")
	readOnly = flag.Bool("R", false, "read only (defalut: false)")
	logFile  = flag.String("F", "webdav.log", "log file name and path ")
	logStat  = flag.Bool("log", false, "log file status (defalut: false)")
)

// go run main.go  -dir "/path/" -R
func main() {

	flag.Parse()

	var argCount = len(os.Args)
	if argCount == 1 && *dir == "." {
		flag.Usage()
		os.Exit(0)
	}

	if argCount == 2 {
		arg1 := os.Args[1]
		if reflect.TypeOf(arg1).String() == "string" {
			*dir = arg1
		} else {
			flag.Usage()
			os.Exit(0)
		}
	}

	p, err := PathExists(*dir)

	if !p {
		fmt.Printf("%s\n", err)
		os.Exit(2)
	}

	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	if *logStat {
		writer3, err := os.OpenFile(*logFile, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			log.Printf("Create file %s failed: %v\n", *logFile, err)
		}
		fmt.Fprintf(os.Stdout, "%v\n", io.MultiWriter(writer1, writer2, writer3))
	} else {
		fmt.Fprintf(os.Stdout, "%v\n", io.MultiWriter(writer1, writer2))
	}

	if *addr == "" {
		*addr = ":" + strconv.Itoa(*port)
	} else {
		*addr = *addr + ":" + strconv.Itoa(*port)
	}
	var sslStr string
	if *ssl {
		sslStr = "https://"
	} else {
		sslStr = "http://"
	}
	fmt.Printf("%s%s\n", sslStr, *addr)
	fmt.Printf("dir: %s\n", *dir)

	if *readOnly {
		fmt.Println("Server Read Only Mode!")
	}

	if *user != "" && *pass != "" {
		fmt.Printf("UserName:%s,PassWord:%s\n", *user, *pass)
	}

	fs := &webdav.Handler{
		FileSystem: webdav.Dir(*dir),
		LockSystem: webdav.NewMemLS(),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if *user != "" && *pass != "" {
			username, password, ok := req.BasicAuth()
			if !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="webDav"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if username != *user || password != *pass {
				http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
				return
			}
		}
		if *readOnly {
			switch req.Method {
			case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
				http.Error(w, "WebDAV: Read Only", http.StatusForbidden)
				return
			}
		}
		ld := webdav.LockDetails{}
		ld.Duration = -time.Second //超时设置
		fmt.Fprintf(os.Stdout, "Method: %v Url: %v RemoteAddr:%v\n", req.Method, req.URL, req.RemoteAddr)
		fs.ServeHTTP(w, req)
	})

	if *ssl {
		err = http.ListenAndServeTLS(*addr, *cert, *key, nil)
	} else {
		err = http.ListenAndServe(*addr, nil)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
