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
	[[document.location]]
	pattern = '/hoge'
	path = '/fuga/hoge'
	[[document.location]]
	pattern = '/foo'
	path = '/bar'
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

	want2 := "/hoge"
	got2 := conf.Document.Location[0].Pattern
	if want2 != got2 {
		t.Errorf("want = %v, but got = %v", want2, got2)
	}

	want3 := "/fuga/hoge"
	got3 := conf.Document.Location[0].Path
	if want3 != got3 {
		t.Errorf("want = %v, but got = %v", want3, got3)
	}

	want4 := "/foo"
	got4 := conf.Document.Location[1].Pattern
	if want4 != got4 {
		t.Errorf("want = %v, but got = %v", want4, got4)
	}

	want5 := "/bar"
	got5 := conf.Document.Location[1].Path
	if want5 != got5 {
		t.Errorf("want = %v, but got = %v", want5, got5)
	}
}
