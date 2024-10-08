package urlqm

import (
	"reflect"
	"testing"
)

func TestGetQueryParam(t *testing.T) {
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
		{
			name: "long query",
			args: args{
				query: `key1=value&key2=a+long+value&key3=a+little+bit+longer+value&key1=the+other+value+&uuid=aa83de98-5b1c-40af-8607-390f7b9271c1`,
				key:   "uuid",
			},
			wantValue: "aa83de98-5b1c-40af-8607-390f7b9271c1",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := GetQueryParam(tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("GetQueryParam() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestGetQueryParamAll(t *testing.T) {
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
			wantValues: nil,
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
			gotValues, err := GetQueryParamAll(tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQueryParamAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("GetQueryParamAll() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestExtractQueryParam(t *testing.T) {
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
		{
			name:      "Found last",
			args:      args{query: "a=1&b=2&c=3&d=4", key: "d"},
			wantValue: "4",
			wantQuery: "a=1&b=2&c=3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := ExtractQueryParam(&tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQueryParamAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("ExtractQueryParam() = %v, want %v", gotValue, tt.wantValue)
			}
			if tt.args.query != tt.wantQuery {
				t.Errorf("ExtractQueryParam() query = %v, want %v", tt.args.query, tt.wantQuery)
			}
		})
	}
}

func TestExtractQueryParamAll(t *testing.T) {
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
			wantValues: nil,
			wantQuery:  "a=1&b=2&c=3",
		},
		{
			name:       "Found",
			args:       args{query: "a=1&b=2&c=3&d=4", key: "b"},
			wantValues: []string{"2"},
			wantQuery:  "a=1&c=3&d=4",
		},
		{
			name:       "Found many",
			args:       args{query: "a=1&b=2&c=3&b=4&d=5&b=6&b=8", key: "b"},
			wantValues: []string{"2", "4", "6", "8"},
			wantQuery:  "a=1&c=3&d=5",
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
			wantValues: nil,
			wantQuery:  `q=%-daily+news%22&theme=dark`,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValues, err := ExtractQueryParamAll(&tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractQueryParamAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("ExtractQueryParamAll() = %v, want %v", gotValues, tt.wantValues)
			}

			if !reflect.DeepEqual(tt.args.query, tt.wantQuery) {
				t.Errorf("ExtractQueryParamAll() = %v, want %v", tt.args.query, tt.wantQuery)
			}
		})
	}
}

func TestAddQueryParam(t *testing.T) {
	type args struct {
		query  string
		key    string
		values []string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery string
	}{
		{
			name:      "Empty query no key no value",
			args:      args{},
			wantQuery: "",
		},
		{
			name:      "Empty query no value",
			args:      args{query: "", key: "a"},
			wantQuery: "a=",
		},
		{
			name:      "Simple addition",
			args:      args{query: "a=1&b=2", key: "c", values: []string{"3"}},
			wantQuery: "a=1&b=2&c=3",
		},
		{
			name:      "Encode key and value",
			args:      args{query: "a=1=b=2", key: "слово", values: []string{"опис"}},
			wantQuery: "a=1=b=2&%D1%81%D0%BB%D0%BE%D0%B2%D0%BE=%D0%BE%D0%BF%D0%B8%D1%81",
		},
		{
			name:      "Add value to existing key",
			args:      args{query: "a=1&b=2", key: "a", values: []string{"3", "6"}},
			wantQuery: "a=1&b=2&a=3&a=6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddQueryParam(&tt.args.query, tt.args.key, tt.args.values...)

			if tt.args.query != tt.wantQuery {
				t.Errorf("TestAddQueryParam() query = %v, want %v", tt.args.query, tt.wantQuery)
			}
		})
	}
}

func TestSetQueryParam(t *testing.T) {
	type args struct {
		query string
		key   string
		value string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery string
	}{
		{
			name:      "Empty query no key",
			args:      args{},
			wantQuery: "",
		},
		{
			name:      "Empty query",
			args:      args{query: "", key: "a", value: "1"},
			wantQuery: "a=1",
		},
		{
			name:      "new key",
			args:      args{query: "a=1&b=2&c=3", key: "d", value: "4"},
			wantQuery: "a=1&b=2&c=3&d=4",
		},
		{
			name:      "existing key",
			args:      args{query: "a=1&b=2&c=3", key: "b", value: "5"},
			wantQuery: "a=1&b=5&c=3",
		},
		{
			name:      "existing multiple keys",
			args:      args{query: "a=1&b=2&c=3&b=4&e=5", key: "b", value: "6"},
			wantQuery: "a=1&b=6&c=3&e=5",
		},
		{
			name:      "no value",
			args:      args{query: "a=1&b=2&c=3", key: "b"},
			wantQuery: "a=1&b=&c=3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetQueryParam(&tt.args.query, tt.args.key, tt.args.value)
			if tt.args.query != tt.wantQuery {
				t.Errorf("TestSetQueryParam() query = %v, want %v", tt.args.query, tt.wantQuery)
			}
		})
	}
}

func TestDeleteQueryParam(t *testing.T) {
	type args struct {
		query string
		key   string
	}
	tests := []struct {
		name      string
		wantQuery string
		args      args
	}{
		{
			name:      "Not found",
			args:      args{query: "a=1&b=2&c=3", key: "d"},
			wantQuery: "a=1&b=2&c=3",
		},
		{
			name:      "Found",
			args:      args{query: "a=1&b=2&c=3&d=4", key: "b"},
			wantQuery: "a=1&c=3&d=4",
		},
		{
			name:      "Found empty",
			args:      args{query: "a=1&b=&c=3&d=4", key: "b"},
			wantQuery: "a=1&c=3&d=4",
		},
		{
			name:      "Single param",
			args:      args{query: "a=1", key: "a"},
			wantQuery: "",
		},
		{
			name:      "Found with deprecated separator",
			args:      args{query: "a=1;b=2;c=3", key: "b"},
			wantQuery: "a=1;c=3",
		},
		{
			name:      "Found with mixed separators",
			args:      args{query: "a=1;b=2&c=3&d=4", key: "b"},
			wantQuery: "a=1;c=3&d=4",
		},
		{
			name:      "encoded",
			args:      args{query: `q=%22daily+news%22&theme=dark`, key: "q"},
			wantQuery: `theme=dark`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteQueryParam(&tt.args.query, tt.args.key)

			if tt.args.query != tt.wantQuery {
				t.Errorf("DeleteQueryParam() query = %v, want %v", tt.args.query, tt.wantQuery)
			}
		})
	}
}

func TestDeleteQueryParamAll(t *testing.T) {
	type args struct {
		query string
		key   string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery string
	}{
		{
			name:      "Not found",
			args:      args{query: "a=1&b=2&c=3", key: "d"},
			wantQuery: "a=1&b=2&c=3",
		},
		{
			name:      "Found",
			args:      args{query: "a=1&b=2&c=3&d=4", key: "b"},
			wantQuery: "a=1&c=3&d=4",
		},
		{
			name:      "Found many",
			args:      args{query: "a=1&b=2&c=3&b=4&d=5&b=6&b=8", key: "b"},
			wantQuery: "a=1&c=3&d=5",
		},
		{
			name:      "encoded",
			args:      args{query: `q=%22daily+news%22&theme=dark`, key: "q"},
			wantQuery: `theme=dark`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteQueryParamAll(&tt.args.query, tt.args.key)

			if !reflect.DeepEqual(tt.args.query, tt.wantQuery) {
				t.Errorf("DeleteQueryParamAll() = %v, want %v", tt.args.query, tt.wantQuery)
			}
		})
	}
}

func TestHasQueryParam(t *testing.T) {
	type args struct {
		query string
		key   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Not found",
			args: args{query: "a=1&b=2&c=3", key: "d"},
			want: false,
		},
		{
			name: "Found",
			args: args{query: "a=1&b=2&c=3", key: "b"},
			want: true,
		},
		{
			name: "encoded key",
			args: args{query: "%D1%81%D0%BB%D0%BE%D0%B2%D0%BE=%D0%BA%D1%96%D1%82", key: "слово"},
			want: false,
		},
		{
			name: "Found first",
			args: args{query: "a=1&b=2&c=3", key: "a"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasQueryParam(tt.args.query, tt.args.key); got != tt.want {
				t.Errorf("HasQueryParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
