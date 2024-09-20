package urlp

import "testing"

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
