package analyse

import (
	"math"
	"testing"
)

var (
	testContents = []string{
		"這是一個伸手不見五指的黑夜。我叫孫悟空，我愛北京，我愛Python和C++。",
		"我不喜歡日本和服。",
		"雷猴回歸人間。",
		"工信處女幹事每月經過下屬科室都要親口交代24口交換機等技術性器件的安裝工作",
		"我需要廉租房",
		"永和服裝飾品有限公司",
		"我愛北京天安門",
		"abc",
		"隱馬爾可夫",
		"雷猴是個好網站",
		"「Microsoft」一詞由「MICROcomputer（微型計算機）」和「SOFTware（軟件）」兩部分組成",
		"草泥馬和欺實馬是今年的流行詞彙",
		"伊藤洋華堂總府店",
		"中國科學院計算技術研究所",
		"羅密歐與朱麗葉",
		"我購買了道具和服裝",
		"PS: 我覺得開源有一個好處，就是能夠敦促自己不斷改進，避免敞帚自珍",
		"湖北省石首市",
		"湖北省十堰市",
		"總經理完成了這件事情",
		"電腦修好了",
		"做好了這件事情就一了百了了",
		"人們審美的觀點是不同的",
		"我們買了一個美的空調",
		"線程初始化時我們要注意",
		"一個分子是由好多原子組織成的",
		"祝你馬到功成",
		"他掉進了無底洞里",
		"中國的首都是北京",
		"孫君意",
		"外交部發言人馬朝旭",
		"領導人會議和第四屆東亞峰會",
		"在過去的這五年",
		"還需要很長的路要走",
		"60週年首都閱兵",
		"你好人們審美的觀點是不同的",
		"買水果然後來世博園",
		"買水果然後去世博園",
		"但是後來我才知道你是對的",
		"存在即合理",
		"的的的的的在的的的的就以和和和",
		"I love你，不以為恥，反以為rong",
		"因",
		"",
		"hello你好人們審美的觀點是不同的",
		"很好但主要是基於網頁形式",
		"hello你好人們審美的觀點是不同的",
		"為什麼我不能擁有想要的生活",
		"後來我才",
		"此次來中國是為了",
		"使用了它就可以解決一些問題",
		",使用了它就可以解決一些問題",
		"其實使用了它就可以解決一些問題",
		"好人使用了它就可以解決一些問題",
		"是因為和國家",
		"老年搜索還支持",
		"乾脆就把那部蒙人的閒法給廢了拉倒！RT @laoshipukong : 27日，全國人大常委會第三次審議侵權責任法草案，刪除了有關醫療損害責任「舉證倒置」的規定。在醫患糾紛中本已處於弱勢地位的消費者由此將陷入萬劫不復的境地。 ",
		"大",
		"",
		"他說的確實在理",
		"長春市長春節講話",
		"結婚的和尚未結婚的",
		"結合成分子時",
		"旅遊和服務是最好的",
		"這件事情的確是我的錯",
		"供大家參考指正",
		"哈爾濱政府公佈塌橋原因",
		"我在機場入口處",
		"邢永臣攝影報道",
		"BP神經網絡如何訓練才能在分類時增加區分度？",
		"南京市長江大橋",
		"應一些使用者的建議，也為了便於利用NiuTrans用於SMT研究",
		"長春市長春藥店",
		"鄧穎超生前最喜歡的衣服",
		"胡錦濤是熱愛世界和平的政治局常委",
		"程序員祝海林和朱會震是在孫健的左面和右面, 範凱在最右面.再往左是李松洪",
		"一次性交多少錢",
		"兩塊五一套，三塊八一斤，四塊七一本，五塊六一條",
		"小和尚留了一個像大和尚一樣的和尚頭",
		"我是中華人民共和國公民;我爸爸是共和黨黨員; 地鐵和平門站",
		"張曉梅去人民醫院做了個B超然後去買了件T恤",
		"AT&T是一件不錯的公司，給你發offer了嗎？",
		"C++和c#是什麼關係？11+122=133，是嗎？PI=3.14159",
		"你認識那個和主席握手的的哥嗎？他開一輛黑色的士。",
		"槍桿子中出政權"}

	Tags = [][]string{
		[]string{"Python", "C++", "伸手不見五指", "孫悟空", "黑夜", "北京", "這是", "一個"},
		[]string{"和服", "喜歡", "日本"},
		[]string{"雷猴", "人間", "回歸"},
		[]string{"工信處", "女幹事", "24", "交換機", "科室", "親口", "器件", "技術性", "下屬", "交代", "每月", "安裝", "經過", "工作"},
		[]string{"廉租房", "需要"},
		[]string{"飾品", "永和", "服裝", "有限公司"},
		[]string{"天安門", "北京"},
		[]string{"abc"},
		[]string{"馬爾可夫"},
		[]string{"雷猴", "網站"},
		[]string{"SOFTware", "Microsoft", "MICROcomputer", "微型", "一詞", "軟件", "計算機", "組成", "部分"},
		[]string{"草泥馬", "欺實", "詞彙", "流行", "今年"},
		[]string{"洋華堂", "總府", "伊藤"},
		[]string{"中國科學院計算技術研究所"},
		[]string{"朱麗葉", "羅密歐"},
		[]string{"道具", "服裝", "購買"},
		[]string{"自珍", "敞帚", "PS", "開源", "不斷改進", "敦促", "好處", "避免", "能夠", "覺得", "就是", "自己", "一個"},
		[]string{"石首市", "湖北省"},
		[]string{"十堰市", "湖北省"},
		[]string{"總經理", "這件", "完成", "事情"},
		[]string{"修好", "電腦"},
		[]string{"一了百了", "做好", "這件", "事情"},
		[]string{"審美", "觀點", "人們", "不同"},
		[]string{"美的", "空調", "我們", "一個"},
		[]string{"線程", "初始化", "注意", "我們"},
		[]string{"好多", "原子", "分子", "組織", "一個"},
		[]string{"馬到功成"},
		[]string{"無底洞"},
		[]string{"首都", "北京", "中國"},
		[]string{"孫君意"},
		[]string{"馬朝旭", "外交部", "發言人"},
		[]string{"第四屆", "東亞", "峰會", "領導人", "會議"},
		[]string{"五年", "過去"},
		[]string{"很長", "需要"},
		[]string{"60", "閱兵", "週年", "首都"},
		[]string{"審美", "你好", "觀點", "人們", "不同"},
		[]string{"世博園", "水果", "然後"},
		[]string{"世博園", "水果", "然後"},
		[]string{"後來", "但是", "知道"},
		[]string{"合理", "存在"},
		[]string{},
		[]string{"rong", "love", "不以為恥", "以為"},
		[]string{},
		[]string{},
		[]string{"hello", "審美", "你好", "觀點", "人們", "不同"},
		[]string{"網頁", "基於", "形式", "主要"},
		[]string{"hello", "審美", "你好", "觀點", "人們", "不同"},
		[]string{"想要", "擁有", "為什麼", "生活", "不能"},
		[]string{"後來"},
		[]string{"此次", "為了", "中國"},
		[]string{"解決", "使用", "一些", "問題", "可以"},
		[]string{"解決", "使用", "一些", "問題", "可以"},
		[]string{"解決", "其實", "使用", "一些", "問題", "可以"},
		[]string{"好人", "解決", "使用", "一些", "問題", "可以"},
		[]string{"是因為", "國家"},
		[]string{"老年", "搜索", "支持"},
		[]string{"閒法", "中本", "laoshipukong", "RT", "27", "責任法", "蒙人", "萬劫不復", "舉證", "倒置", "醫患", "那部", "拉倒", "侵權", "全國人大常委會", "草案", "境地", "糾紛", "刪除", "弱勢"},
		[]string{},
		[]string{},
		[]string{"在理", "確實"},
		[]string{"長春", "春節", "講話", "市長"},
		[]string{"結婚", "尚未"},
		[]string{"分子", "結合"},
		[]string{"旅遊", "最好", "服務"},
		[]string{"的確", "這件", "事情"},
		[]string{"指正", "參考", "大家"},
		[]string{"塌橋", "哈爾濱", "公佈", "原因", "政府"},
		[]string{"入口處", "機場"},
		[]string{"邢永臣", "攝影", "報道"},
		[]string{"區分度", "BP", "神經網絡", "訓練", "分類", "才能", "如何", "增加"},
		[]string{"長江大橋", "南京市"},
		[]string{"SMT", "NiuTrans", "使用者", "便於", "用於", "建議", "利用", "為了", "研究", "一些"},
		[]string{"長春市", "藥店", "長春"},
		[]string{"鄧穎超", "生前", "衣服", "喜歡"},
		[]string{"政治局", "熱愛", "常委", "胡錦濤", "和平", "世界"},
		[]string{"右面", "孫健", "範凱", "李松洪", "朱會震", "海林", "左面", "程序員", "再往"},
		[]string{"一次性", "多少"},
		[]string{"四塊", "五塊", "三塊", "一斤", "兩塊", "一本", "一套", "一條"},
		[]string{"和尚", "和尚頭", "一樣", "一個"},
		[]string{"和平門", "共和黨", "地鐵", "黨員", "公民", "爸爸", "中華人民共和國"},
		[]string{"張曉梅", "T恤", "B超", "醫院", "人民", "然後"},
		[]string{"offer", "AT&T", "不錯", "一件", "公司"},
		[]string{"c#", "PI", "C++", "3.14159", "133", "122", "11", "關係", "什麼"},
		[]string{"的士", "的哥", "他開", "握手", "一輛", "黑色", "主席", "認識", "那個"},
		[]string{"槍桿子", "政權"},
	}

	Lyric = `
我沒有心
我沒有真實的自我
我只有消瘦的臉孔
所謂軟弱
所謂的順從一向是我
的座右銘

而我
沒有那海洋的寬闊
我只要熱情的撫摸
所謂空洞
所謂不安全感是我
的墓誌銘

而你
是否和我一般怯懦
是否和我一般矯作
和我一般囉唆

而你
是否和我一般退縮
是否和我一般肌迫
一般地困惑

我沒有力
我沒有滿腔的熱火
我只有滿肚的如果
所謂勇氣
所謂的認同感是我
隨便說說

而你
是否和我一般怯懦
是否和我一般矯作
是否對你來說
只是一場遊戲
雖然沒有把握

而你
是否和我一般退縮
是否和我一般肌迫
是否對你來說
只是逼不得已
雖然沒有藉口
`
	LyciWeight = Segments{
		Segment{text: "所謂", weight: 1.010262},
		Segment{text: "是否", weight: 0.738650},
		Segment{text: "一般", weight: 0.607600},
		Segment{text: "雖然", weight: 0.336754},
		Segment{text: "退縮", weight: 0.336754},
		Segment{text: "肌迫", weight: 0.336754},
		Segment{text: "矯作", weight: 0.336754},
		Segment{text: "沒有", weight: 0.336754},
		Segment{text: "怯懦", weight: 0.271099},
		Segment{text: "隨便", weight: 0.168377},
	}

	LyciWeight2 = Segments{
		Segment{text: "所謂", weight: 1.215739},
		Segment{text: "一般", weight: 0.731179},
		Segment{text: "雖然", weight: 0.405246},
		Segment{text: "退縮", weight: 0.405246},
		Segment{text: "肌迫", weight: 0.405246},
		Segment{text: "矯作", weight: 0.405246},
		Segment{text: "怯懦", weight: 0.326238},
		Segment{text: "逼不得已", weight: 0.202623},
		Segment{text: "右銘", weight: 0.202623},
		Segment{text: "寬闊", weight: 0.202623},
	}
)

func TestExtractTags(t *testing.T) {
	var te TagExtracter
	te.LoadDictionary("../dict.txt")
	te.LoadIdf("idf.txt")

	for index, sentence := range testContents {
		result := te.ExtractTags(sentence, 20)
		if len(result) != len(Tags[index]) {
			t.Fatalf("%s = %v", sentence, result)
		}
		for i, tag := range result {
			if tag.text != Tags[index][i] {
				t.Fatalf("%s != %s", tag, Tags[index][i])
			}
		}
	}
}

func TestExtratTagsWithWeight(t *testing.T) {
	var te TagExtracter
	te.LoadDictionary("../dict.txt")
	te.LoadIdf("idf.txt")
	result := te.ExtractTags(Lyric, 10)
	for index, tag := range result {
		if LyciWeight[index].text != tag.text ||
			math.Abs(LyciWeight[index].weight-tag.weight) > 1e-6 {
			t.Fatalf("%v != %v", tag, LyciWeight[index])
		}
	}
}

func TestExtractTagsWithStopWordsFile(t *testing.T) {
	var te TagExtracter
	te.LoadDictionary("../dict.txt")
	te.LoadIdf("idf.txt")
	te.LoadStopWords("stop_words.txt")
	result := te.ExtractTags(Lyric, 7)
	for index, tag := range result {
		if LyciWeight2[index].text != tag.text ||
			math.Abs(LyciWeight2[index].weight-tag.weight) > 1e-6 {
			t.Fatalf("%v != %v", tag, LyciWeight2[index])
		}
	}
}
