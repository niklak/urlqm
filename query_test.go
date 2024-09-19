package urlp

import (
	"reflect"
	"testing"
)

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
