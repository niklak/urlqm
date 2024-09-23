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
			wantP:   nil,
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

func TestParams_Sort(t *testing.T) {
	tests := []struct {
		name  string
		p     Params
		wantP Params
	}{
		{
			name:  "Simple",
			p:     Params{{"b", "2"}, {"a", "2"}, {"a", "1"}, {"c", "3"}},
			wantP: Params{{"a", "2"}, {"a", "1"}, {"b", "2"}, {"c", "3"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Sort()
			if !reflect.DeepEqual(tt.p, tt.wantP) {
				t.Errorf("Params.Sort() = %v, want %v", tt.p, tt.wantP)
			}
		})
	}
}

func TestParams_SetOrder(t *testing.T) {
	type args struct {
		order []string
	}
	tests := []struct {
		name  string
		p     Params
		args  args
		wantP Params
	}{
		{
			name:  "No order",
			args:  args{order: nil},
			p:     Params{{"k3", "v3"}, {"k2", "v2"}, {"k1", "v1"}},
			wantP: Params{{"k3", "v3"}, {"k2", "v2"}, {"k1", "v1"}},
		},
		{
			name:  "With priority param",
			args:  args{order: []string{"q"}},
			p:     Params{{"b", "2"}, {"a", "1"}, {"q", "3"}},
			wantP: Params{{"q", "3"}, {"b", "2"}, {"a", "1"}},
		},
		{
			name:  "With full order",
			args:  args{order: []string{"q", "a", "b"}},
			p:     Params{{"b", "2"}, {"a", "1"}, {"q", "3"}},
			wantP: Params{{"q", "3"}, {"a", "1"}, {"b", "2"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.SetOrder(tt.args.order...)
			if !reflect.DeepEqual(tt.p, tt.wantP) {
				t.Errorf("Params.SetOrder() = %v, want %v", tt.p, tt.wantP)
			}
		})
	}
}

func TestParams_Add(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name  string
		p     Params
		args  args
		wantP Params
	}{
		{
			name:  "Simple",
			p:     Params{{"k1", "v1"}, {"k2", "v2"}},
			args:  args{"k3", "v3"},
			wantP: Params{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Add(tt.args.key, tt.args.value)
			if !reflect.DeepEqual(tt.p, tt.wantP) {
				t.Errorf("Params.Add() = %v, want %v", tt.p, tt.wantP)
			}
		})
	}
}

func TestParams_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		p    Params
		args args
		want string
	}{
		{
			name: "Not found",
			p:    Params{{"k1", "v1"}, {"k2", "v2"}},
			args: args{"k3"},
			want: "",
		},
		{
			name: "Found",
			p:    Params{{"k1", "v1"}, {"k2", "v2"}},
			args: args{"k2"},
			want: "v2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Get(tt.args.key); got != tt.want {
				t.Errorf("Params.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParams_GetAll(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		p    Params
		args args
		want []string
	}{
		{
			name: "Not found",
			p:    Params{{"k1", "v1"}, {"k2", "v2"}},
			args: args{"k3"},
			want: nil,
		},
		{
			name: "Found",
			args: args{"k2"},
			p:    Params{{"k1", "v1"}, {"k2", "v2"}, {"k2", "v3"}},
			want: []string{"v2", "v3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetAll(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Params.GetAll() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestParams_Extract(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		p         Params
		args      args
		wantP     Params
		wantValue string
	}{
		{
			name:      "Not found",
			p:         Params{{"k1", "v1"}, {"k2", "v2"}},
			args:      args{"k3"},
			wantP:     Params{{"k1", "v1"}, {"k2", "v2"}},
			wantValue: "",
		},
		{
			name:      "Found",
			p:         Params{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}},
			args:      args{"k2"},
			wantP:     Params{{"k1", "v1"}, {"k3", "v3"}},
			wantValue: "v2",
		},
		{
			name:      "Found many",
			p:         Params{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}, {"k2", "v4"}},
			args:      args{"k2"},
			wantP:     Params{{"k1", "v1"}, {"k3", "v3"}, {"k2", "v4"}},
			wantValue: "v2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValue := tt.p.Extract(tt.args.key); gotValue != tt.wantValue {
				t.Errorf("Params.Extract() = %v, want %v", gotValue, tt.wantValue)
			}

			if !reflect.DeepEqual(tt.p, tt.wantP) {
				t.Errorf("Params.Extract() = %v, want %v", tt.p, tt.wantP)
			}
		})
	}
}

func TestParams_ExtractAll(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name       string
		p          Params
		args       args
		wantP      Params
		wantValues []string
	}{
		{
			name:       "Not found",
			p:          Params{{"k1", "v1"}, {"k2", "v2"}},
			args:       args{"k3"},
			wantP:      Params{{"k1", "v1"}, {"k2", "v2"}},
			wantValues: nil,
		},
		{
			name:       "Found",
			p:          Params{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}, {"k2", "v4"}},
			args:       args{"k2"},
			wantP:      Params{{"k1", "v1"}, {"k3", "v3"}},
			wantValues: []string{"v2", "v4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValues := tt.p.ExtractAll(tt.args.key); !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("Params.ExtractAll() = %v, want %v", gotValues, tt.wantValues)
			}

			if !reflect.DeepEqual(tt.p, tt.wantP) {
				t.Errorf("Params.ExtractAll() = %v, want %v", tt.p, tt.wantP)
			}
		})
	}
}

func TestParams_Set(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name  string
		p     Params
		wantP Params
		args  args
	}{
		{
			name:  "new param",
			p:     Params{{"k1", "v1"}, {"k2", "v2"}},
			wantP: Params{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}},
			args:  args{"k3", "v3"},
		},
		{
			name:  "replace existing",
			p:     Params{{"k1", "v1"}, {"k2", "v2"}},
			wantP: Params{{"k1", "v1"}, {"k2", "v3"}},
			args:  args{"k2", "v3"},
		},
		{
			name:  "replace existing multiple values",
			p:     Params{{"k1", "v1"}, {"k2", "v2"}, {"k2", "v3"}, {"k3", "v3"}, {"k2", "v4"}},
			wantP: Params{{"k1", "v1"}, {"k2", "v4"}, {"k3", "v3"}},
			args:  args{"k2", "v4"},
		},
		{
			name:  "empty key",
			args:  args{"", "value"},
			p:     Params{{"k1", "v1"}},
			wantP: Params{{"k1", "v1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Set(tt.args.key, tt.args.value)

			if !reflect.DeepEqual(tt.p, tt.wantP) {
				t.Errorf("Params.Set() = %v, want %v", tt.p, tt.wantP)
			}
		})
	}
}

func TestParams_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		p     Params
		args  args
		wantP Params
	}{
		{
			name:  "Not found",
			p:     Params{{"k1", "v1"}, {"k2", "v2"}},
			args:  args{"k3"},
			wantP: Params{{"k1", "v1"}, {"k2", "v2"}},
		},
		{
			name:  "Found",
			p:     Params{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}},
			args:  args{"k2"},
			wantP: Params{{"k1", "v1"}, {"k3", "v3"}},
		},
		{
			name:  "Found many",
			p:     Params{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}, {"k2", "v4"}},
			args:  args{"k2"},
			wantP: Params{{"k1", "v1"}, {"k3", "v3"}, {"k2", "v4"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Delete(tt.args.key)

			if !reflect.DeepEqual(tt.p, tt.wantP) {
				t.Errorf("Params.Delete() = %v, want %v", tt.p, tt.wantP)
			}
		})
	}
}

func TestParams_DeleteAll(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		p     Params
		args  args
		wantP Params
	}{
		{
			name:  "Not found",
			p:     Params{{"k1", "v1"}, {"k2", "v2"}},
			args:  args{"k3"},
			wantP: Params{{"k1", "v1"}, {"k2", "v2"}},
		},
		{
			name:  "Found",
			p:     Params{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}},
			args:  args{"k2"},
			wantP: Params{{"k1", "v1"}, {"k3", "v3"}},
		},
		{
			name:  "Found many",
			p:     Params{{"k1", "v1"}, {"k2", "v2"}, {"k3", "v3"}, {"k2", "v4"}},
			args:  args{"k2"},
			wantP: Params{{"k1", "v1"}, {"k3", "v3"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.DeleteAll(tt.args.key)

			if !reflect.DeepEqual(tt.p, tt.wantP) {
				t.Errorf("Params.DeleteAll() = %v, want %v", tt.p, tt.wantP)
			}
		})
	}
}

func TestParams_Has(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		p    Params
		args args
		want bool
	}{
		{
			name: "Found",
			p:    Params{{"k1", "v1"}, {"k2", "v2"}},
			args: args{"k2"},
			want: true,
		},
		{
			name: "Not found",
			p:    Params{{"k1", "v1"}, {"k2", "v2"}},
			args: args{"k3"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Has(tt.args.key); got != tt.want {
				t.Errorf("Params.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}
