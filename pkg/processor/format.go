package processor

import (
	"bytes"
	"strings"
)

// NamingConventionCase available in programming convention
type NamingConventionCase string

const (

	// CamelCase https://en.wikipedia.org/wiki/Camel_case
	CamelCase NamingConventionCase = "camel"

	// LowerCamelCase https://en.wikipedia.org/wiki/Camel_case (~CamelCase)
	LowerCamelCase NamingConventionCase = "lowercamel"

	// PascalCase http://wiki.c2.com/?PascalCase (~UpperCamelCase)
	PascalCase NamingConventionCase = "pascal"

	// UpperCamelCase https://en.wikipedia.org/wiki/Camel_case (~PascalCase)
	UpperCamelCase NamingConventionCase = "uppercamel"

	// SankeCase https://en.m.wikipedia.org/wiki/Snake_case (lowerCase + "_")
	SankeCase NamingConventionCase = "snake"

	// DarwinCase https://en.wikipedia.org/wiki/Camel_case The combination of "TitleCase " and "snake case"
	DarwinCase NamingConventionCase = "darwin"

	// TitleCase .
	TitleCase NamingConventionCase = "title"

	// LowerCase  .as
	LowerCase NamingConventionCase = "lower"

	// UpperCase .
	UpperCase NamingConventionCase = "upper"
)

// Format given tokens in the given NamingConventionCase
func Format(tokens []string, namingConventionCase NamingConventionCase) string {

	if len(tokens) == 0 {
		return ""
	}

	var buffer bytes.Buffer
	length := len(tokens)
	for index, token := range tokens {
		// first token

		if index == 0 {

			switch namingConventionCase {

			case PascalCase, UpperCamelCase, TitleCase, DarwinCase:
				token = strings.Title(token)
			case UpperCase:
				token = strings.ToUpper(token)
			default:
				token = strings.ToLower(token)
			}
		} else {

			switch namingConventionCase {

			case SankeCase, LowerCase:
				token = strings.ToLower(token)
			case UpperCase:
				token = strings.ToUpper(token)
			default:
				token = strings.Title(token)
			}
		}

		buffer.WriteString(token)

		// add a splitter only for snake and darwin
		// don't add splitter for last token
		if (namingConventionCase == SankeCase || namingConventionCase == DarwinCase) && (index < (length - 1)) {
			buffer.WriteString("_")
		}
	}
	output := buffer.String()
	return output
}

// IsValidCaseType check if the given case is supported by formatter
func IsValidCaseType(caseType string) bool {

	caseType = strings.ToLower(caseType)

	namingConventionCase := NamingConventionCase(caseType)
	switch namingConventionCase {
	case CamelCase, LowerCamelCase, PascalCase, UpperCamelCase, SankeCase, DarwinCase, TitleCase, LowerCase, UpperCase:
		return true
	default:
		return false
	}
}