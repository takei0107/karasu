package karasu

import (
	"strings"
	"testing"
)

func Test_init_config(t *testing.T) {
	doc := `
	[server]
	port = 80
	[document]
	root = "/var/www/docroot"
	`

	conf, err := load_config(strings.NewReader(doc))
	if err != nil {
		t.Errorf("init_config() has error.")
		t.Errorf(err.Error())
	}

	want1 := 80
	got1 := conf.Server.Port
	if want1 != got1 {
		t.Errorf("want = %v, but got = %v", want1, got1)
	}

	want := "/var/www/docroot"
	got := conf.Document.Root
	if want != got {
		t.Errorf("want = %v, but got = %v", want, got)
	}
}
