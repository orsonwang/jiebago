package util

import (
	"regexp"
	"testing"
)

func TestRegexpSplit(t *testing.T) {
	result := RegexpSplit(regexp.MustCompile(`\p{Han}+`),
		"BP神經網絡如何訓練才能在分類時增加區分度？", -1)
	if len(result) != 2 {
		t.Fatal(result)
	}
	result = RegexpSplit(regexp.MustCompile(`(\p{Han})+`),
		"BP神經網絡如何訓練才能在分類時增加區分度？", -1)
	if len(result) != 3 {
		t.Fatal(result)
	}
	result = RegexpSplit(regexp.MustCompile(`([\p{Han}#]+)`),
		",BP神經網絡如何訓練才能在分類時#增加區分度？", -1)
	if len(result) != 3 {
		t.Fatal(result)
	}
}
