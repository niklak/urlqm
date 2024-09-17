package urlp

import "testing"

func TestEncodeParams(t *testing.T) {
	type args struct {
		params []QueryParam
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"No params", args{[]QueryParam{}}, ""},
		{"No params nil", args{nil}, ""},
		{"Simple", args{[]QueryParam{{"a", "1"}}}, "a=1"},
		{"Unordered multiple values",
			args{[]QueryParam{{"a", "1"}, {"b", "2"}, {"a", "3"}}},
			"a=1&b=2&a=3",
		},
		{"Encoded chars", args{[]QueryParam{{"q", `"daily news"`}}},
			"q=%22daily+news%22"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeParams(tt.args.params); got != tt.want {
				t.Errorf("EncodeParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
