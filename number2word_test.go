package number2word

import (
    "testing"
    "regexp"
)

// TestConvertText calls number2word.Convert with a number, checking
// for a valid return value.
func TestConvertText(t *testing.T) {
    text := "five hundred fourteen"
    want := regexp.MustCompile(`\b`+text+`\b`)
    msg, err := Convert(514)
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Convert("five hundred fourteen") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestConvertAndText calls greetings.ConvertAnd with a number, checking
// for a valid return value.
func TestConvertAndText(t *testing.T) {
    text := "five hundred and fourteen"
    want := regexp.MustCompile(`\b`+text+`\b`)
    msg, err := ConvertAnd(514)
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Convert("five hundred and fourteen") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}