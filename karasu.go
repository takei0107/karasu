package karasu

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Run() error {
	// FIXME 設定ファイルを探す
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
		log_info(fmt.Sprintf("URI = %s", path))
		// FIXME pathのチェック処理を入れる
		if !strings.HasPrefix(path, "/") {
			path += "/"
		}
		file_path := docroot + path
		log_info(fmt.Sprintf("FILE_PATH = %s", file_path))
		f, err := os.Open(file_path)
		if err != nil {
			http.Error(w, "Page Not Found.", 404)
			log_error("details", err)
			return
		}
		_, err = io.Copy(w, f)
		if err != nil {
			http.Error(w, "Internal Server Error.", 500)
			log_error("details", err)
			return
		}
	})
	port := strconv.Itoa(cfg.Server.Port)
	log_info(fmt.Sprintf("Server start. Listen port:%s", port))
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
	return nil
}
