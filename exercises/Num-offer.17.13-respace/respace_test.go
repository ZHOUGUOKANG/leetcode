package Num_offer_17_13_respace

import "testing"

func TestRespace(t *testing.T) {
	input := []string{"looked", "just", "like", "her", "brother"}
	sentence := "jesslookedjustliketimherbrother"
	expected := 7
	output := respace(input, sentence)
	if expected != output {
		t.Error(input, sentence, expected, output)
	}
}
func TestRespace1(t *testing.T) {
	input := []string{"potimzz"}
	sentence := "potimzzpotimzz"
	expected := 0
	output := respace(input, sentence)
	if expected != output {
		t.Error(input, sentence, expected, output)
	}
}
func TestRespace2(t *testing.T) {
	input := []string{"sssjjs", "hschjf", "hhh", "fhjchfcfshhfjhs", "sfh", "jsf", "cjschjfscscscsfjcjfcfcfh", "hccccjjfchcffjjshccsjscsc", "chcfjcsshjj", "jh", "h", "f", "s", "jcshs", "jfjssjhsscfc"}
	sentence := "sssjjssfshscfjjshsjjsjchffffs"
	expected := 7
	output := respace(input, sentence)
	if expected != output {
		t.Error(input, sentence, expected, output)
	}
}
