package karasu

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Run() error {
	cfg_file, err := os.Open("./conf/conf.toml")
	if err != nil {
		panic(err)
	}
	cfg, err := load_config(cfg_file)
	if err != nil {
		panic(err)
	}
	docroot := cfg.Document.Root
	// root待受
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		log.Printf("URI = %s\n", path)
		if !strings.HasPrefix(path, "/") {
			path += "/"
		}
		full_path := docroot + path
		log.Printf("FULL_PATH = %s\n", full_path)
		f, err := os.Open(full_path)
		if err != nil {
			http.Error(w, "Page Not Found.", 404)
			log.Printf("[ERROR] details = %v\n", err)
			return
		}
		_, err = io.Copy(w, f)
		if err != nil {
			http.Error(w, "Internal Server Error.", 500)
			log.Printf("[ERROR] details = %v\n", err)
			return
		}
	})
	port := strconv.Itoa(cfg.Server.Port)
	err = http.ListenAndServe(":"+port, nil)
	log.Printf("[INFO] Server start. Listen port:%s", port)
	if err != nil {
		panic(err)
	}
	return nil
}
