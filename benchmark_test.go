package vtlgen

import "testing"

func BenchmarkGenerateMappingTemplatesAndFunctions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateMappingTemplatesAndFunctions("testdata/mapping-templates")
	}
}
