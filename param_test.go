package urlp

import (
	"reflect"
	"testing"
)

func TestEncodeParams(t *testing.T) {
	type args struct {
		params []QueryParam
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "No params", args: args{[]QueryParam{}}, want: ""},
		{name: "No params nil", args: args{nil}, want: ""},
		{name: "Simple", args: args{[]QueryParam{{"a", "1"}}}, want: "a=1"},
		{
			name: "Unordered multiple values",
			args: args{[]QueryParam{{"a", "1"}, {"b", "2"}, {"a", "3"}}},
			want: "a=1&b=2&a=3",
		},
		{
			name: "Encoded chars",
			args: args{[]QueryParam{{"q", `"daily news"`}}},
			want: "q=%22daily+news%22",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeParams(tt.args.params); got != tt.want {
				t.Errorf("EncodeParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseParams(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name       string
		args       args
		wantValues []QueryParam
		wantErr    bool
	}{
		{
			name:       "Empty",
			args:       args{},
			wantValues: []QueryParam{},
			wantErr:    false,
		},
		{
			name:       "Ampersand as separator",
			args:       args{"k1=v1&k2=v2"},
			wantValues: []QueryParam{{"k1", "v1"}, {"k2", "v2"}},
			wantErr:    false,
		},
		{
			name:       "Semicolon as separator",
			args:       args{"k1=v1;k2=v2"},
			wantValues: []QueryParam{{"k1", "v1"}, {"k2", "v2"}},
			wantErr:    false,
		},
		{
			name:       "Mixed separators",
			args:       args{"k1=v1;k2=v2&k3=v3"},
			wantValues: []QueryParam{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}},
			wantErr:    false,
		},
		{
			name:       "Encoded chars",
			args:       args{`q=%22daily+news%22`},
			wantValues: []QueryParam{{"q", `"daily news"`}},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValues, err := ParseParams(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("ParseParams() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}
