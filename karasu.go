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
		logger.Info.Printf("request path = %s\n", path)
		file_path := docroot + resolveFilePath(path, cfg.Document)
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

func resolveFilePath(requestPath string, documentCfg documentConfig) string {
	if !strings.HasPrefix(requestPath, "/") {
		requestPath = "/" + requestPath
	}
	for _, locationCfg := range documentCfg.Location {
		locationPattern := locationCfg.Pattern
		if !strings.HasPrefix(locationPattern, "/") {
			locationPattern = "/" + locationPattern
		}
		if requestPath == locationPattern {
			locationPath := locationCfg.Path
			if !strings.HasPrefix(locationPath, "/") {
				locationPath = "/" + locationPath
			}
			return locationPath
		}
	}
	return requestPath
}

func handleHttpError(w http.ResponseWriter, e error) {
	httpError := httperror.New(e)
	http.Error(w, httpError.Message(), httpError.StatusCode())
	logger.Error.Println(httpError)
}
