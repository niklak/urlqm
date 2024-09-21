package test

import (
	"net/url"
	"testing"

	"github.com/niklak/urlp"
)

const simpleRawQuery = "key1=value&key2=a+long+value&key3=a+little+bit+longer+value&key1=the+other+value+&uuid=aa83de98-5b1c-40af-8607-390f7b9271c1"

func getParamStd(query string, key string) (string, error) {
	// Because of if you need to get even a one key, you have to parse the whole query once or more times.
	q, err := url.ParseQuery(query)
	if err != nil {
		return "", err
	}
	return q.Get(key), nil
}

func getParamStdAll(query, key string) ([]string, error) {
	q, err := url.ParseQuery(query)
	if err != nil {
		return []string{}, err
	}
	return q[key], nil
}

func getParamUrlP(query string, key string) (string, error) {
	q, err := urlp.QueryParams(query)
	if err != nil {
		return "", err
	}
	return q.Get(key), nil
}

func getParamUrlPAll(query string, key string) ([]string, error) {
	q, err := urlp.QueryParams(query)
	if err != nil {
		return []string{}, err
	}
	return q.GetAll(key), nil
}

func BenchmarkGetQueryParamOne(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		urlp.GetQueryParam(query, "key3")
	}
}

func BenchmarkGetQueryParamOneUrlP(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		getParamUrlP(query, "key3")
	}
}

func BenchmarkGetQueryParamOneStd(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		getParamStd(query, "key3")
	}
}

func BenchmarkGetQueryParamAll(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		urlp.GetQueryParamAll(query, "key1")
	}
}

func BenchmarkGetQueryParamAllUrlP(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		getParamUrlPAll(query, "key1")
	}
}

func BenchmarkGetQueryParamAllStd(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		getParamStdAll(query, "key1")
	}
}

func BenchmarkParseParamsStd(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		url.ParseQuery(query)
	}
}

func BenchmarkParseParamsUrlP(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		urlp.ParseParams(query)
	}
}
