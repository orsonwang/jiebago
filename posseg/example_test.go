package posseg_test

import (
	"fmt"

	"jiebago/posseg"
)

func Example() {
	var seg posseg.Segmenter
	seg.LoadDictionary("../dict.txt")

	for segment := range seg.Cut("我愛北京天安門", true) {
		fmt.Printf("%s %s\n", segment.Text(), segment.Pos())
	}
	// Output:
	// 我 r
	// 愛 v
	// 北京 ns
	// 天安門 ns
}
