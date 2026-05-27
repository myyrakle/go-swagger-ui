package swagger

import (
	"strings"
	"testing"
)

func TestPreprocessSchemeInYAMLSpecMovesPreferredSchemeFirst(t *testing.T) {
	spec := []byte(`swagger: "2.0"
schemes:
  - http
  - https
info:
  title: Example
  version: "1.0"
`)

	out, err := PreprocessSchemeInYAMLSpec(spec, SchemeHTTPS)
	if err != nil {
		t.Fatalf("PreprocessSchemeInYAMLSpec returned error: %v", err)
	}

	schemes := yamlSchemes(t, string(out))
	if len(schemes) != 2 || schemes[0] != "https" || schemes[1] != "http" {
		t.Fatalf("schemes = %v, want [https http]", schemes)
	}
}

func yamlSchemes(t *testing.T, source string) []string {
	t.Helper()

	var schemes []string
	inSchemes := false
	for _, line := range strings.Split(source, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "schemes:" {
			inSchemes = true
			continue
		}
		if inSchemes && strings.HasPrefix(trimmed, "- ") {
			schemes = append(schemes, strings.TrimPrefix(trimmed, "- "))
			continue
		}
		if inSchemes && trimmed != "" {
			break
		}
	}
	return schemes
}
