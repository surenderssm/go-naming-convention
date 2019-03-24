package processor

import "testing"

// Test all the formatter

func getTestToken() []string {

	tokens := []string{"golang", "founder", "ken", "thompson"}
	return tokens
}

func TestForValidCaseType(t *testing.T) {

	if ok := IsValidCaseType(string(CamelCase)); !ok {
		t.Error("supported casetype should be valid")
	}
	if ok := IsValidCaseType("randomcase"); ok {
		t.Error("unsupported casetype should be invalid")
	}
}
func TestFormatForCamelCase(t *testing.T) {

	output := Format(getTestToken(), CamelCase)
	expected := "golangFounderKenThompson"
	if output != expected {
		t.Fatal("TestFormatForCamelCase-output :", output, "-Instead of :", expected)
	}
}

func TestFormatForLowerCamelCase(t *testing.T) {

	output := Format(getTestToken(), LowerCamelCase)
	expected := "golangFounderKenThompson"
	if output != expected {
		t.Fatal("TestFormatForLowerCamelCase-output :", output, "-Instead of :", expected)
	}
}

func TestFormatForPascalCase(t *testing.T) {

	output := Format(getTestToken(), PascalCase)
	expected := "GolangFounderKenThompson"
	if output != expected {
		t.Fatal("TestFormatForPascalCase-output :", output, "-Instead of :", expected)
	}
}

func TestFormatForUpperCamelCase(t *testing.T) {

	output := Format(getTestToken(), UpperCamelCase)
	expected := "GolangFounderKenThompson"
	if output != expected {
		t.Fatal("TestFormatForUpperCamelCase-output :", output, "-Instead of :", expected)
	}
}

func TestFormatForSankeCase(t *testing.T) {

	output := Format(getTestToken(), SankeCase)
	expected := "golang_founder_ken_thompson"
	if output != expected {
		t.Fatal("TestFormatForSankeCase-output :", output, "-Instead of :", expected)
	}
}

func TestFormatForDarwinCase(t *testing.T) {

	output := Format(getTestToken(), DarwinCase)
	expected := "Golang_Founder_Ken_Thompson"
	if output != expected {
		t.Fatal("TestFormatForSankeCase-output :", output, "-Instead of :", expected)
	}
}

func TestFormatForTitleCase(t *testing.T) {

	output := Format(getTestToken(), TitleCase)
	expected := "GolangFounderKenThompson"
	if output != expected {
		t.Fatal("TestFormatForTitleCase-output :", output, "-Instead of :", expected)
	}
}

func TestFormatForLowerCase(t *testing.T) {

	output := Format(getTestToken(), LowerCase)
	expected := "golangfounderkenthompson"
	if output != expected {
		t.Fatal("TestFormatForLowerCase-output :", output, "-Instead of :", expected)
	}
}

func TestFormatForUpperCase(t *testing.T) {

	output := Format(getTestToken(), UpperCase)
	expected := "GOLANGFOUNDERKENTHOMPSON"
	if output != expected {
		t.Fatal("TestFormatForUpperCase-output :", output, "-Instead of :", expected)
	}
}
