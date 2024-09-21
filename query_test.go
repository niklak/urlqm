package urlp

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

func TestGetQueryParamValues(t *testing.T) {
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
			gotValues, err := GetQueryParamValues(tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQueryParamValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("GetQueryParamValues() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestPopQueryParam(t *testing.T) {
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
			gotValue, err := PopQueryParam(&tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetQueryParamValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("PopQueryParam() = %v, want %v", gotValue, tt.wantValue)
			}
			if tt.args.query != tt.wantQuery {
				t.Errorf("PopQueryParam() query = %v, want %v", tt.args.query, tt.wantQuery)
			}
		})
	}
}

func TestPopQueryParamValues(t *testing.T) {
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
			gotValues, err := PopQueryParamValues(&tt.args.query, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PopQueryParamValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("PopQueryParamValues() = %v, want %v", gotValues, tt.wantValues)
			}

			if !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("PopQueryParamValues() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestAddQueryParam(t *testing.T) {
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
			args:      args{query: "a=1&b=2", key: "c", value: "3"},
			wantQuery: "a=1&b=2&c=3",
		},
		{
			name:      "Encode key and value",
			args:      args{query: "a=1=b=2", key: "слово", value: "опис"},
			wantQuery: "a=1=b=2&%D1%81%D0%BB%D0%BE%D0%B2%D0%BE=%D0%BE%D0%BF%D0%B8%D1%81",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddQueryParam(&tt.args.query, tt.args.key, tt.args.value)

			if tt.args.query != tt.wantQuery {
				t.Errorf("TestAddQueryParam() query = %v, want %v", tt.args.query, tt.wantQuery)
			}
		})
	}
}
