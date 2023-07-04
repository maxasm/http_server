package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io"
	"path/filepath"
)

	// event logger
var eventLogger = log.New(os.Stdout, "[Event] ", log.Ltime | log.Ldate)	
	// error logger	
var errorLogger = log.New(os.Stdout, "[Error] ", log.Ltime | log.Ldate)

func read_file(fname string) ([]byte,error) {
	f, err_f := os.Open(fname)
	if err_f != nil {
		errorLogger.Printf("Error: %s\n", err_f)
		return []byte{}, err_f
	}
	defer f.Close()
	data, err_data := io.ReadAll(f)
	if err_data != nil {
		errorLogger.Printf("Error: %s\n", err_data)
	}
	return data, nil	
}

func main() {
	
	root_dir := "./dist"
	
	// This is where you handle all routes with a Hashmap -> map[route]handle_func
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request){
		req_path := root_dir + req.URL.Path
		eventLogger.Printf("request on path: %s\n", req_path)
		if req_path == "./dist/" {
			req_path += "index.html"
		}
		
		switch filepath.Ext(req_path) {
			case ".js":
				w.Header().Set("Content-Type","text/javascript")
			case ".html":
				w.Header().Set("Content-Type", "text/html")
			case ".ico":
				w.Header().Set("Content-Type", "image/vnd.microsoft.icon")
			case ".css":
				w.Header().Set("Content-Type", "text/css")
		}
		
		data, err_read_file := read_file(req_path)
		if err_read_file != nil {
			data, _ := read_file("./dist/index.html")
			w.Write(data)
		}
		w.Write(data)
	})


	fmt.Printf("Listening on port :8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		errorLogger.Printf("Error: %s\n", nil)
	}
	
}

