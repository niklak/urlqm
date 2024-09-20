package urlp

import (
	"reflect"
	"testing"
)

func TestParams_Encode(t *testing.T) {
	tests := []struct {
		name       string
		p          Params
		wantEncode string
	}{
		// TODO: Add test cases.
		{name: "No params", p: Params{}, wantEncode: ""},
		{name: "No params nil", p: nil, wantEncode: ""},
		{name: "Simple", p: Params{{"a", "1"}}, wantEncode: "a=1"},
		{
			name:       "Unordered multiple values",
			p:          Params{{"a", "1"}, {"b", "2"}, {"a", "3"}},
			wantEncode: "a=1&b=2&a=3",
		},
		{
			name:       "Encoded chars",
			p:          Params{{"q", `"daily news"`}},
			wantEncode: "q=%22daily+news%22",
		},
		{
			name:       "Previously bad encoded",
			p:          Params{{"q", `100%+truth`}},
			wantEncode: "q=100%25%2Btruth",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotEncode := tt.p.Encode(); gotEncode != tt.wantEncode {
				t.Errorf("Params.Encode() = %v, want %v", gotEncode, tt.wantEncode)
			}
		})
	}
}

func TestQueryParams(t *testing.T) {
	type args struct {
		rawQuery string
	}
	tests := []struct {
		name    string
		args    args
		wantP   Params
		wantErr bool
	}{
		{
			name:    "Empty",
			args:    args{},
			wantP:   Params{},
			wantErr: false,
		},
		{
			name:    "Ampersand as separator",
			args:    args{"k1=v1&k2=v2"},
			wantP:   Params{{"k1", "v1"}, {"k2", "v2"}},
			wantErr: false,
		},
		{
			name:    "Semicolon as separator",
			args:    args{"k1=v1;k2=v2"},
			wantP:   Params{{"k1", "v1"}, {"k2", "v2"}},
			wantErr: false,
		},
		{
			name:    "Mixed separators",
			args:    args{"k1=v1;k2=v2&k3=v3"},
			wantP:   Params{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}},
			wantErr: false,
		},
		{
			name:    "Encoded chars",
			args:    args{`q=%22daily+news%22`},
			wantP:   Params{{"q", `"daily news"`}},
			wantErr: false,
		},
		{
			name:    "Encoded chars err",
			args:    args{`a=1&q=100%+truth&b=2&brightness=90%`},
			wantP:   Params{{"a", "1"}, {"q", "100%+truth"}, {"b", "2"}, {"brightness", "90%"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, err := QueryParams(tt.args.rawQuery)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("QueryParams() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}
