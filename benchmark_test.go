package vtlgen

import "testing"

func BenchmarkDatasourceGenerateMappingTemplates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateMappingTemplates("testdata/mapping-templates")
	}
}
