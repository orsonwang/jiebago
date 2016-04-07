package analyse_test

import (
	"fmt"

	"jiebago/analyse"
)

func Example_extractTags() {
	var t analyse.TagExtracter
	t.LoadDictionary("../dict.txt")
	t.LoadIdf("idf.txt")

	sentence := "這是一個伸手不見五指的黑夜。我叫孫悟空，我愛北京，我愛Python和C++。"
	segments := t.ExtractTags(sentence, 5)
	fmt.Printf("Top %d tags:", len(segments))
	for _, segment := range segments {
		fmt.Printf(" %s /", segment.Text())
	}
	// Output:
	// Top 5 tags: Python / C++ / 伸手不見五指 / 孫悟空 / 黑夜 /
}

func Example_textRank() {
	var t analyse.TextRanker
	t.LoadDictionary("../dict.txt")
	sentence := "此外，公司擬對全資子公司吉林歐亞置業有限公司增資4.3億元，增資後，吉林歐亞置業註冊資本由7000萬元增加到5億元。吉林歐亞置業主要經營範圍為房地產開發及百貨零售等業務。目前在建吉林歐亞城市商業綜合體項目。2013年，實現營業收入0萬元，實現淨利潤-139.13萬元。"

	result := t.TextRank(sentence, 10)
	for _, segment := range result {
		fmt.Printf("%s %f\n", segment.Text(), segment.Weight())
	}
	// Output:
	// 吉林 1.000000
	// 歐亞 0.878078
	// 置業 0.562048
	// 實現 0.520906
	// 收入 0.384284
	// 增資 0.360591
	// 子公司 0.353132
	// 城市 0.307509
	// 全資 0.306324
	// 商業 0.306138
}
