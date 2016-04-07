package analyse

import (
	"math"
	"testing"
)

var (
	sentence = "此外，公司擬對全資子公司吉林歐亞置業有限公司增資4.3億元，增資後，吉林歐亞置業註冊資本由7000萬元增加到5億元。吉林歐亞置業主要經營範圍為房地產開發及百貨零售等業務。目前在建吉林歐亞城市商業綜合體項目。2013年，實現營業收入0萬元，實現淨利潤-139.13萬元。"

	tagRanks = Segments{
		Segment{text: "吉林", weight: 1.0},
		Segment{text: "歐亞", weight: 0.87807810644},
		Segment{text: "置業", weight: 0.562048250306},
		Segment{text: "實現", weight: 0.520905743929},
		Segment{text: "收入", weight: 0.384283870648},
		Segment{text: "增資", weight: 0.360590945312},
		Segment{text: "子公司", weight: 0.353131980904},
		Segment{text: "城市", weight: 0.307509449283},
		Segment{text: "全資", weight: 0.306324426665},
		Segment{text: "商業", weight: 0.306138241063},
	}
)

func TestTextRank(t *testing.T) {
	var tr TextRanker
	tr.LoadDictionary("../dict.txt")
	results := tr.TextRank(sentence, 10)
	for index, tw := range results {
		if tw.text != tagRanks[index].text || math.Abs(tw.weight-tagRanks[index].weight) > 1e-6 {
			t.Fatalf("%v != %v", tw, tagRanks[index])
		}
	}
}
ss