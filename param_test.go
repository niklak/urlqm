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

func TestSortOrderParams(t *testing.T) {
	type args struct {
		params []QueryParam
		order  []string
	}
	tests := []struct {
		name       string
		args       args
		wantValues []QueryParam
	}{
		{
			name:       "No order",
			args:       args{params: []QueryParam{{"k3", "v3"}, {"k2", "v2"}, {"k1", "v1"}}, order: nil},
			wantValues: []QueryParam{{"k3", "v3"}, {"k2", "v2"}, {"k1", "v1"}},
		},
		{
			name:       "With priority param",
			args:       args{params: []QueryParam{{"b", "2"}, {"a", "1"}, {"q", "3"}}, order: []string{"q"}},
			wantValues: []QueryParam{{"q", "3"}, {"b", "2"}, {"a", "1"}},
		},
		{
			name:       "With full order",
			args:       args{params: []QueryParam{{"b", "2"}, {"a", "1"}, {"q", "3"}}, order: []string{"q", "a", "b"}},
			wantValues: []QueryParam{{"q", "3"}, {"a", "1"}, {"b", "2"}},
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

func TestGetParam(t *testing.T) {
	type args struct {
		query string
		key   string
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
		wantErr   bool
	}{
		{
			name:      "Not found",
			args:      args{query: "a=1&b=2&c=3", key: "d"},
			wantValue: "",
		},
		{
			name:      "No value",
			args:      args{query: "a=1&b=&c=3", key: "b"},
			wantValue: "",
		},
		{
			name:      "Found",
			args:      args{query: "a=1&b=2&c=3", key: "b"},
			wantValue: "2",
		},
		{
			name:      "Found with deprecated separator",
			args:      args{query: "a=1;b=2;c=3", key: "b"},
			wantValue: "2",
		},
		{
			name:      "encoded",
			args:      args{query: `q=%22daily+news%22&theme=dark`, key: "q"},
			wantValue: `"daily news"`,
		},
		{
			name:      "bad encoding",
			args:      args{query: `q=%-daily+news%22&theme=dark`, key: "q"},
			wantValue: ``,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := GetParam(tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("GetParam() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestGetParamValues(t *testing.T) {
	type args struct {
		query string
		key   string
	}
	tests := []struct {
		name       string
		args       args
		wantValues []string
		wantErr    bool
	}{
		{
			name:       "Not found",
			args:       args{query: "a=1&b=2&c=3", key: "d"},
			wantValues: []string{},
		},
		{
			name:       "Found empty",
			args:       args{query: "a=1&b=&c=3", key: "b"},
			wantValues: []string{""},
		},
		{
			name:       "Found multiple values",
			args:       args{query: "a=1&b=2&c=3&d=4&b=5&e=6", key: "b"},
			wantValues: []string{"2", "5"},
		},
		{
			name:       "Found multiple values with deprecated separator",
			args:       args{query: "a=1;b=2;c=3;d=4;b=5;e=6", key: "b"},
			wantValues: []string{"2", "5"},
		},
		{
			name:       "encoded",
			args:       args{query: `q=%22daily+news%22&theme=dark`, key: "q"},
			wantValues: []string{`"daily news"`},
		},
		{
			name:       "bad encoding",
			args:       args{query: `q=%-daily+news%22&theme=dark`, key: "q"},
			wantValues: []string{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValues, err := GetParamValues(tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetParamValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("GetParamValues() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestPopParam(t *testing.T) {
	type args struct {
		query string
		key   string
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
		wantErr   bool
		wantQuery string
	}{
		{
			name:      "Not found",
			args:      args{query: "a=1&b=2&c=3", key: "d"},
			wantValue: "",
			wantQuery: "a=1&b=2&c=3",
		},
		{
			name:      "Found",
			args:      args{query: "a=1&b=2&c=3&d=4", key: "b"},
			wantValue: "2",
			wantQuery: "a=1&c=3&d=4",
		},
		{
			name:      "Found empty",
			args:      args{query: "a=1&b=&c=3&d=4", key: "b"},
			wantValue: "",
			wantQuery: "a=1&c=3&d=4",
		},
		{
			name:      "Single param",
			args:      args{query: "a=1", key: "a"},
			wantValue: "1",
			wantQuery: "",
		},
		{
			name:      "Found with deprecated separator",
			args:      args{query: "a=1;b=2;c=3", key: "b"},
			wantValue: "2",
			wantQuery: "a=1;c=3",
		},
		{
			name:      "Found with mixed separators",
			args:      args{query: "a=1;b=2&c=3&d=4", key: "b"},
			wantValue: "2",
			wantQuery: "a=1;c=3&d=4",
		},
		{
			name:      "encoded",
			args:      args{query: `q=%22daily+news%22&theme=dark`, key: "q"},
			wantValue: `"daily news"`,
			wantQuery: `theme=dark`,
		},
		{
			name:      "bad encoding",
			args:      args{query: `q=%-daily+news%22&theme=dark`, key: "q"},
			wantValue: ``,
			wantQuery: `q=%-daily+news%22&theme=dark`,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := PopParam(&tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetParamValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("PopParam() = %v, want %v", gotValue, tt.wantValue)
			}
			if tt.args.query != tt.wantQuery {
				t.Errorf("PopParam() query = %v, want %v", tt.args.query, tt.wantQuery)
			}
		})
	}
}

func TestPopParamValues(t *testing.T) {
	type args struct {
		query string
		key   string
	}
	tests := []struct {
		name       string
		args       args
		wantValues []string
		wantErr    bool
		wantQuery  string
	}{
		{
			name:       "Not found",
			args:       args{query: "a=1&b=2&c=3", key: "d"},
			wantValues: []string{},
			wantQuery:  "a=1&b=2&c=3",
		},
		{
			name:       "Found",
			args:       args{query: "a=1&b=2&c=3&d=4", key: "b"},
			wantValues: []string{"2"},
			wantQuery:  "a=1&c=3&d=4",
		},
		{
			name:       "Found",
			args:       args{query: "a=1&b=2&c=3&b=4&d=5&b=6", key: "b"},
			wantValues: []string{"2", "4", "6"},
			wantQuery:  "a=1&c=3&d=4",
		},
		{
			name:       "encoded",
			args:       args{query: `q=%22daily+news%22&theme=dark`, key: "q"},
			wantValues: []string{`"daily news"`},
			wantQuery:  `theme=dark`,
		},
		{
			name:       "bad encoding",
			args:       args{query: `q=%-daily+news%22&theme=dark`, key: "q"},
			wantValues: []string{},
			wantQuery:  `q=%-daily+news%22&theme=dark`,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValues, err := PopParamValues(&tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopParamValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("PopParamValues() = %v, want %v", gotValues, tt.wantValues)
			}

			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("PopParamValues() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

// TODO: add tests cases with errors
