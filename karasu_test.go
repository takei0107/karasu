package karasu

import "testing"

func TestResolvePath(t *testing.T) {
	type args struct {
		path string
		cfg  []locationConfig
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ファイル名のみ",
			args: args{
				path: "a.txt",
				cfg: []locationConfig{
					locationConfig{
						Pattern: "a.txt",
						Path:    "b.txt",
					},
				},
			},
			want: "/b.txt",
		},
		{
			name: "ディレクトリ+ファイル名",
			args: args{
				path: "/foo/bar.txt",
				cfg: []locationConfig{
					locationConfig{
						Pattern: "/foo/bar.txt",
						Path:    "/fuga/hoge/bar.txt",
					},
				},
			},
			want: "/fuga/hoge/bar.txt",
		},
		{
			name: "ディレクトリを書き換える",
			args: args{
				path: "/hoge/fuga/foo.txt",
				cfg: []locationConfig{
					locationConfig{
						Pattern: "/hoge/fuga/",
						Path:    "/fuga/hoge/",
					},
				},
			},
			want: "/fuga/hoge/foo.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveFilePath(tt.args.path, tt.args.cfg)
			if got != tt.want {
				t.Errorf("got = %v, want = %v\n", got, tt.want)
			}
		})
	}
}
