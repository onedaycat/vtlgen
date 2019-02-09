package vtlgen

import "testing"

func BenchmarkGenerateMappingTemplates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateMappingTemplates("./mapping-templates")
	}
}
