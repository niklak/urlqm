package test

import (
	"net/url"
	"testing"

	"github.com/niklak/urlqm"
)

const simpleRawQuery = "q=%D1%81%D0%BE%D1%82%D0%BD%D1%96+%D0%BC%D1%96%D0%BB%D1%8C%D0%B9%D0%BE%D0%BD%D1%96%D0%B2&" +
	"key1=value&key2=a+long+value&key3=a+little+bit+longer+value&key1=the+other+value+&uuid=aa83de98-5b1c-40af-8607-390f7b9271c1"

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
	q, err := urlqm.ParseQuery(query)
	if err != nil {
		return "", err
	}
	return q.Get(key), nil
}

func getParamUrlPAll(query string, key string) ([]string, error) {
	q, err := urlqm.ParseQuery(query)
	if err != nil {
		return []string{}, err
	}
	return q.GetAll(key), nil
}

func BenchmarkGetQueryParamOne(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		urlqm.GetQueryParam(query, "uuid")
	}
}

func BenchmarkGetQueryParamOneUrlP(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		getParamUrlP(query, "uuid")
	}
}

func BenchmarkGetQueryParamOneStd(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		getParamStd(query, "uuid")
	}
}

func BenchmarkGetQueryParamAll(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		urlqm.GetQueryParamAll(query, "key1")
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
		urlqm.ParseParams(query)
	}
}

func BenchmarkEncodeParamsStd(b *testing.B) {
	// std query parse and encode breaks original query param order
	q, _ := url.ParseQuery(simpleRawQuery)
	for i := 0; i < b.N; i++ {
		_ = q.Encode()
	}
}

func BenchmarkEncodeParamsUrlP(b *testing.B) {
	q, _ := urlqm.ParseQuery(simpleRawQuery)
	for i := 0; i < b.N; i++ {
		_ = q.Encode()
	}
}

func BenchmarkEncodeParamsSortedUrlP(b *testing.B) {
	q, _ := urlqm.ParseQuery(simpleRawQuery)
	for i := 0; i < b.N; i++ {
		q.Sort()
		_ = q.Encode()
	}
}

func BenchmarkAddParam(b *testing.B) {
	var query string
	for i := 0; i < b.N; i++ {
		query = simpleRawQuery
		urlqm.AddQueryParam(&query, "key6", "value1")
	}
}

func BenchmarkAddParamUrlP(b *testing.B) {
	q, _ := urlqm.ParseQuery(simpleRawQuery)
	var query urlqm.Params
	for i := 0; i < b.N; i++ {
		query = q
		query.Add("key6", "value1")
	}
}

func BenchmarkAddParamStd(b *testing.B) {
	q, _ := url.ParseQuery(simpleRawQuery)
	var query url.Values
	for i := 0; i < b.N; i++ {
		query = q
		query.Add("key6", "value1")
	}
}

func BenchmarkDeleteParam(b *testing.B) {
	var query string
	for i := 0; i < b.N; i++ {
		query = simpleRawQuery
		urlqm.DeleteQueryParam(&query, "key1")
	}
}

func BenchmarkDeleteParamUrlP(b *testing.B) {
	q, _ := urlqm.ParseQuery(simpleRawQuery)
	var query urlqm.Params
	for i := 0; i < b.N; i++ {
		query = q
		query.Delete("key1")
	}
}

func BenchmarkDeleteParamAll(b *testing.B) {
	var query string
	for i := 0; i < b.N; i++ {
		query = simpleRawQuery
		urlqm.DeleteQueryParamAll(&query, "key1")
	}
}

func BenchmarkDeleteParamAllUrlP(b *testing.B) {
	q, _ := urlqm.ParseQuery(simpleRawQuery)
	var query urlqm.Params
	for i := 0; i < b.N; i++ {
		query = q
		query.DeleteAll("key1")
	}
}

func BenchmarkDeleteParamAllStd(b *testing.B) {
	q, _ := url.ParseQuery(simpleRawQuery)
	var query url.Values
	for i := 0; i < b.N; i++ {
		query = q
		query.Del("key1")
	}
}

func BenchmarkExtractParam(b *testing.B) {
	var query string
	for i := 0; i < b.N; i++ {
		query = simpleRawQuery
		urlqm.ExtractQueryParam(&query, "key1")
	}
}

func BenchmarkExtractParamUrlP(b *testing.B) {
	q, _ := urlqm.ParseQuery(simpleRawQuery)
	var query urlqm.Params
	for i := 0; i < b.N; i++ {
		query = q
		query.Extract("key1")
	}
}

func BenchmarkExtractParamAll(b *testing.B) {
	var query string
	for i := 0; i < b.N; i++ {
		query = simpleRawQuery
		urlqm.ExtractQueryParamAll(&query, "key1")
	}
}

func BenchmarkExtractParamAllUrlP(b *testing.B) {
	q, _ := urlqm.ParseQuery(simpleRawQuery)
	var query urlqm.Params
	for i := 0; i < b.N; i++ {
		query = q
		query.ExtractAll("key1")
	}
}
func BenchmarkSetParam(b *testing.B) {
	var query string
	for i := 0; i < b.N; i++ {
		query = simpleRawQuery
		urlqm.SetQueryParam(&query, "key10", "new+value")
	}
}

func BenchmarkSetParamUrlP(b *testing.B) {
	q, _ := urlqm.ParseQuery(simpleRawQuery)
	var query urlqm.Params
	for i := 0; i < b.N; i++ {
		query = q
		query.Set("key10", "new+value")
	}
}

func BenchmarkSetParamStd(b *testing.B) {
	q, _ := url.ParseQuery(simpleRawQuery)
	var query url.Values
	for i := 0; i < b.N; i++ {
		query = q
		query.Set("key10", "new+value")
	}
}

func BenchmarkSetParamExisting(b *testing.B) {
	var query string
	for i := 0; i < b.N; i++ {
		query = simpleRawQuery
		urlqm.SetQueryParam(&query, "key1", "new+value")
	}
}

func BenchmarkSetParamExistingUrlP(b *testing.B) {
	q, _ := urlqm.ParseQuery(simpleRawQuery)
	var query urlqm.Params
	for i := 0; i < b.N; i++ {
		query = q
		query.Set("key1", "new+value")
	}
}

func BenchmarkSetParamExistingStd(b *testing.B) {
	q, _ := url.ParseQuery(simpleRawQuery)
	var query url.Values
	for i := 0; i < b.N; i++ {
		query = q
		query.Set("key1", "new+value")
	}
}

func BenchmarkParseSetParamUrlP(b *testing.B) {
	var query urlqm.Params
	for i := 0; i < b.N; i++ {
		query, _ = urlqm.ParseQuery(simpleRawQuery)
		query.Set("key10", "new+value")
	}
}

func BenchmarkParseSetParamStd(b *testing.B) {
	var query url.Values
	for i := 0; i < b.N; i++ {
		query, _ = url.ParseQuery(simpleRawQuery)
		query.Set("key10", "new+value")
	}
}

func BenchmarkParseSetParamExistingUrlP(b *testing.B) {
	var query urlqm.Params
	for i := 0; i < b.N; i++ {
		query, _ = urlqm.ParseQuery(simpleRawQuery)
		query.Set("key1", "new+value")
	}
}

func BenchmarkParseSetParamExistingStd(b *testing.B) {

	var query url.Values
	for i := 0; i < b.N; i++ {
		query, _ = url.ParseQuery(simpleRawQuery)
		query.Set("key1", "new+value")
	}
}

func BenchmarkHasQueryParam(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		urlqm.HasQueryParam(query, "uuid")
	}
}

func BenchmarkHasQueryParamUrlP(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		q, _ := urlqm.ParseQuery(query)
		q.Has("uuid")
	}
}

func BenchmarkHasQueryParamStd(b *testing.B) {
	query := simpleRawQuery
	for i := 0; i < b.N; i++ {
		q, _ := url.ParseQuery(query)
		q.Has("uuid")
	}
}
