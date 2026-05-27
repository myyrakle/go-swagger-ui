package swagger

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Scheme string

const (
	SchemeHTTP  Scheme = "http"
	SchemeHTTPS Scheme = "https"
)

func PreprocessSchemeInJSONSpec(specSource []byte, preferredScheme Scheme) ([]byte, error) {
	var spec map[string]any
	if err := json.Unmarshal([]byte(specSource), &spec); err != nil {
		return nil, fmt.Errorf("swagger doc parsing failed: %w", err)
	}

	preprocessScheme(spec, preferredScheme)

	out, err := json.Marshal(spec)
	if err != nil {
		return nil, fmt.Errorf("swagger doc serialization failed: %w", err)
	}

	return out, nil
}

func PreprocessSchemeInYAMLSpec(specSource []byte, preferredScheme Scheme) ([]byte, error) {
	var spec map[string]any
	if err := yaml.Unmarshal(specSource, &spec); err != nil {
		return nil, fmt.Errorf("swagger doc parsing failed: %w", err)
	}

	preprocessScheme(spec, preferredScheme)

	out, err := yaml.Marshal(spec)
	if err != nil {
		return nil, fmt.Errorf("swagger doc serialization failed: %w", err)
	}

	return out, nil
}

func preprocessScheme(spec map[string]any, preferredScheme Scheme) {
	if schemes, ok := spec["schemes"].([]any); ok && len(schemes) > 1 {
		for i, v := range schemes {
			if s, ok := v.(string); ok && s == string(preferredScheme) && i > 0 {
				schemes[0], schemes[i] = schemes[i], schemes[0]
				spec["schemes"] = schemes
				break
			}
		}
	}
}
