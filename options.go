package gen_elasticsearch_go

type Query struct {
	Data map[string]interface{}
}

type OptionsType uint8

var esQueryBoolMustDefaultTemplate = `{"query":{"bool":{"must":[%s]}}}`
var esQueryBoolDefaultTemplate = `{"query":{"bool":{%s}}}`

var esSentenceQueryTemplate = `{"size":3,"_source":"content","query":{"bool":{"must":[{"range":{"len":{"gte":10,"lte":100}}},%s]}}}`
var esTermTemplate = `{"term":{"%s":{"value":"%s"}}}`
var esTermTemplateInt = `{"term":{"%s":{"value":%d}}}`
var esTermTemplateF = `{"term":{"%s":{"value":%f}}}`
var esTermTemplateInt1 = `{"term":{"%s":{"value":%d}}}`
var esShouldTemplate = `{"bool":{"should":[%s],"minimum_should_match":1}}`
var esMatchTemplate = `{"match":{"%s":"%s"}}`
var esMatchPhraseTemplate = `{"match_phrase":{"%s":"%s"}}`
var esRangeTemplateInt = `{"range":{"%s":{"gte":%d,"lte":%d}}}`
var esRangeTemplateIntMin = `{"range":{"%s":{"gte":%d}}}`
var esRangeTemplateIntMax = `{"range":{"%s":{"lte":%d}}}`
var esRangeTemplateIntComp = `{"range":{"%s":{"%s":%d}}}`
var esMustNotEmptyTemplate = `{"bool":{"must_not":[{"term":{"%s":{"value":""}}}]}}`
var esNestedTemplate = `{"nested":{"path":"%s","query":{"bool":{"must":[%s]}}}}`
var esNestedMustsMustNotsTemplate = `{"nested":{"path":"%s","query":{"bool":{"must":[%s],"must_not":[%s]}}}}`
var esNestedShouldsTemplate = `{"nested":{"path":"%s","query":{"bool":{"should":[%s],"minimum_should_match":1}}}}`
var esNestedMustsMustNotsShouldsTemplate = `{"nested":{"path":"%s","query":{"bool":{"must":[%s],"must_not":[%s],"should":[%s],"minimum_should_match":1}}}}`
var esRangeTemplateTimeStr = `{"range":{"%s":{"gte":"%d","lte":"%d"}}}`
var esRangeTemplateTimeStrMin = `{"range":{"%s":{"gte":"%d"}}}`
var esRangeTemplateTimeStrMax = `{"range":{"%s":{"lte":"%d"}}}`

var esRangeTemplateFloat = `{"range":{"%s":{"gte":%f,"lte":%f}}}`
var esRangeTemplateFloatMin = `{"range":{"%s":{"gte":%f}}}`
var esRangeTemplateFloatMax = `{"range":{"%s":{"lte":%f}}}`

const (
	OptionsTypeTerm OptionsType = iota
	OptionsTypeMatchPrefix
	OptionsTypeRange
	OptionsTypeExists
)

type Options struct {
	VectorKey string
	Vector    []float64
	Musts     []string
	MustNots  []string
	Limit     int
	EsIndex   string
	Others    []string
}

func EsSetQuery[T any](fn func(T) bool) func(key string, otype OptionsType, vals ...T) {
	return func(key string, otype OptionsType, vals ...T) {

		return
	}
}

func EsSetQueryShould[T any](fn func(T) bool) func(key string, otype OptionsType, vals ...T) {

}
