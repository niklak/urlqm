package urlp

import (
	"errors"
	"net/url"
	"reflect"
	"testing"
)

func TestEncodeParams(t *testing.T) {
	type args struct {
		params []Param
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "No params", args: args{[]Param{}}, want: ""},
		{name: "No params nil", args: args{nil}, want: ""},
		{name: "Simple", args: args{[]Param{{"a", "1"}}}, want: "a=1"},
		{
			name: "Unordered multiple values",
			args: args{[]Param{{"a", "1"}, {"b", "2"}, {"a", "3"}}},
			want: "a=1&b=2&a=3",
		},
		{
			name: "Encoded chars",
			args: args{[]Param{{"q", `"daily news"`}}},
			want: "q=%22daily+news%22",
		},
		{
			name: "Previously bad encoded",
			args: args{[]Param{{"q", `100%+truth`}}},
			want: "q=100%25%2Btruth",
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
		wantValues []Param
		wantErr    bool
	}{
		{
			name:       "Empty",
			args:       args{},
			wantValues: []Param{},
			wantErr:    false,
		},
		{
			name:       "Ampersand as separator",
			args:       args{"k1=v1&k2=v2"},
			wantValues: []Param{{"k1", "v1"}, {"k2", "v2"}},
			wantErr:    false,
		},
		{
			name:       "Semicolon as separator",
			args:       args{"k1=v1;k2=v2"},
			wantValues: []Param{{"k1", "v1"}, {"k2", "v2"}},
			wantErr:    false,
		},
		{
			name:       "Mixed separators",
			args:       args{"k1=v1;k2=v2&k3=v3"},
			wantValues: []Param{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}},
			wantErr:    false,
		},
		{
			name:       "Encoded chars",
			args:       args{`q=%22daily+news%22`},
			wantValues: []Param{{"q", `"daily news"`}},
			wantErr:    false,
		},
		{
			name:       "Encoded chars err",
			args:       args{`a=1&q=100%+truth&b=2&brightness=90%`},
			wantValues: []Param{{"a", "1"}, {"q", "100%+truth"}, {"b", "2"}, {"brightness", "90%"}},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		var e url.EscapeError
		t.Run(tt.name, func(t *testing.T) {
			gotValues, err := ParseParams(tt.args.query)
			if tt.wantErr {
				if !errors.As(err, &e) {
					t.Errorf("ParseParams() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("ParseParams() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestSortOrderParams(t *testing.T) {
	type args struct {
		params []Param
		order  []string
	}
	tests := []struct {
		name       string
		args       args
		wantValues []Param
	}{
		{
			name:       "No order",
			args:       args{params: []Param{{"k3", "v3"}, {"k2", "v2"}, {"k1", "v1"}}, order: nil},
			wantValues: []Param{{"k3", "v3"}, {"k2", "v2"}, {"k1", "v1"}},
		},
		{
			name:       "With priority param",
			args:       args{params: []Param{{"b", "2"}, {"a", "1"}, {"q", "3"}}, order: []string{"q"}},
			wantValues: []Param{{"q", "3"}, {"b", "2"}, {"a", "1"}},
		},
		{
			name:       "With full order",
			args:       args{params: []Param{{"b", "2"}, {"a", "1"}, {"q", "3"}}, order: []string{"q", "a", "b"}},
			wantValues: []Param{{"q", "3"}, {"a", "1"}, {"b", "2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortOrderParams(&tt.args.params, tt.args.order...)
			gotValues := tt.args.params
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("SortParams() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestSortParams(t *testing.T) {
	type args struct {
		params []Param
	}
	tests := []struct {
		name       string
		args       args
		wantParams []Param
	}{
		{
			name: "simple sort",
			args: args{
				params: []Param{
					{"b", "2"},
					{"a", "2"},
					{"a", "1"},
					{"c", "3"},
				},
			},
			wantParams: []Param{{"a", "2"}, {"a", "1"}, {"b", "2"}, {"c", "3"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortParams(tt.args.params)

			if !reflect.DeepEqual(tt.args.params, tt.wantParams) {
				t.Errorf("SortParams() = %v, want %v", tt.args.params, tt.wantParams)
			}
		})
	}
}
