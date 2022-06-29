package karasu

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/takei0107/karasu/cmd/httperror"
	"github.com/takei0107/karasu/cmd/logger"
)

func Run() error {
	// FIXME 設定ファイルを探す
	cfg_file, err := os.Open("./conf/conf.toml")
	if err != nil {
		return err
	}
	cfg, err := load_config(cfg_file)
	if err != nil {
		return err
	}
	docroot := cfg.Document.Root
	// root待受
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		logger.Info.Printf("URI = %s\n", path)
		// FIXME pathのチェック処理を入れる
		if !strings.HasPrefix(path, "/") {
			path += "/"
		}
		file_path := docroot + path
		logger.Info.Printf("FILE_PATH = %s\n", file_path)
		f, err := os.Open(file_path)
		if err != nil {
			handleHttpError(w, err)
			return
		}
		_, err = io.Copy(w, f)
		if err != nil {
			handleHttpError(w, err)
			return
		}
	})
	port := strconv.Itoa(cfg.Server.Port)
	logger.Info.Printf("Server start. Listen port:%s", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}
	return nil
}

func handleHttpError(w http.ResponseWriter, e error) {
	httpError := httperror.New(e)
	http.Error(w, httpError.Message(), httpError.StatusCode())
	logger.Error.Println(httpError)
}
