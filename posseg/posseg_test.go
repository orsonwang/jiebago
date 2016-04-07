package posseg

import (
	"testing"
)

var (
	seg          Segmenter
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

	defaultCutResult = [][]Segment{[]Segment{Segment{"這", "r"}, Segment{"是", "v"}, Segment{"一個", "m"}, Segment{"伸手不見五指", "i"}, Segment{"的", "uj"}, Segment{"黑夜", "n"}, Segment{"。", "x"}, Segment{"我", "r"}, Segment{"叫", "v"}, Segment{"孫悟空", "nr"}, Segment{"，", "x"}, Segment{"我", "r"}, Segment{"愛", "v"}, Segment{"北京", "ns"}, Segment{"，", "x"}, Segment{"我", "r"}, Segment{"愛", "v"}, Segment{"Python", "eng"}, Segment{"和", "c"}, Segment{"C++", "nz"}, Segment{"。", "x"}},
		[]Segment{Segment{"我", "r"}, Segment{"不", "d"}, Segment{"喜歡", "v"}, Segment{"日本", "ns"}, Segment{"和服", "nz"}, Segment{"。", "x"}},
		[]Segment{Segment{"雷猴", "n"}, Segment{"回歸", "v"}, Segment{"人間", "n"}, Segment{"。", "x"}},
		[]Segment{Segment{"工信處", "n"}, Segment{"女幹事", "n"}, Segment{"每月", "r"}, Segment{"經過", "p"}, Segment{"下屬", "v"}, Segment{"科室", "n"}, Segment{"都", "d"}, Segment{"要", "v"}, Segment{"親口", "n"}, Segment{"交代", "n"}, Segment{"24", "m"}, Segment{"口", "n"}, Segment{"交換機", "n"}, Segment{"等", "u"}, Segment{"技術性", "n"}, Segment{"器件", "n"}, Segment{"的", "uj"}, Segment{"安裝", "v"}, Segment{"工作", "vn"}},
		[]Segment{Segment{"我", "r"}, Segment{"需要", "v"}, Segment{"廉租房", "n"}},
		[]Segment{Segment{"永和", "nz"}, Segment{"服裝", "vn"}, Segment{"飾品", "n"}, Segment{"有限公司", "n"}},
		[]Segment{Segment{"我", "r"}, Segment{"愛", "v"}, Segment{"北京", "ns"}, Segment{"天安門", "ns"}},
		[]Segment{Segment{"abc", "eng"}},
		[]Segment{Segment{"隱", "n"}, Segment{"馬爾可夫", "nr"}},
		[]Segment{Segment{"雷猴", "n"}, Segment{"是", "v"}, Segment{"個", "q"}, Segment{"好", "a"}, Segment{"網站", "n"}},
		[]Segment{Segment{"「", "x"}, Segment{"Microsoft", "eng"}, Segment{"」", "x"}, Segment{"一", "m"}, Segment{"詞", "n"}, Segment{"由", "p"}, Segment{"「", "x"}, Segment{"MICROcomputer", "eng"}, Segment{"（", "x"}, Segment{"微型", "b"}, Segment{"計算機", "n"}, Segment{"）", "x"}, Segment{"」", "x"}, Segment{"和", "c"}, Segment{"「", "x"}, Segment{"SOFTware", "eng"}, Segment{"（", "x"}, Segment{"軟件", "n"}, Segment{"）", "x"}, Segment{"」", "x"}, Segment{"兩", "m"}, Segment{"部分", "n"}, Segment{"組成", "v"}},
		[]Segment{Segment{"草泥馬", "n"}, Segment{"和", "c"}, Segment{"欺實", "v"}, Segment{"馬", "n"}, Segment{"是", "v"}, Segment{"今年", "t"}, Segment{"的", "uj"}, Segment{"流行", "v"}, Segment{"詞彙", "n"}},
		[]Segment{Segment{"伊藤", "nr"}, Segment{"洋華堂", "n"}, Segment{"總府", "n"}, Segment{"店", "n"}},
		[]Segment{Segment{"中國科學院計算技術研究所", "nt"}},
		[]Segment{Segment{"羅密歐", "nr"}, Segment{"與", "p"}, Segment{"朱麗葉", "nr"}},
		[]Segment{Segment{"我", "r"}, Segment{"購買", "v"}, Segment{"了", "ul"}, Segment{"道具", "n"}, Segment{"和", "c"}, Segment{"服裝", "vn"}},
		[]Segment{Segment{"PS", "eng"}, Segment{":", "x"}, Segment{" ", "x"}, Segment{"我", "r"}, Segment{"覺得", "v"}, Segment{"開源", "n"}, Segment{"有", "v"}, Segment{"一個", "m"}, Segment{"好處", "d"}, Segment{"，", "x"}, Segment{"就是", "d"}, Segment{"能夠", "v"}, Segment{"敦促", "v"}, Segment{"自己", "r"}, Segment{"不斷改進", "l"}, Segment{"，", "x"}, Segment{"避免", "v"}, Segment{"敞", "v"}, Segment{"帚", "ng"}, Segment{"自珍", "b"}},
		[]Segment{Segment{"湖北省", "ns"}, Segment{"石首市", "ns"}},
		[]Segment{Segment{"湖北省", "ns"}, Segment{"十堰市", "ns"}},
		[]Segment{Segment{"總經理", "n"}, Segment{"完成", "v"}, Segment{"了", "ul"}, Segment{"這件", "mq"}, Segment{"事情", "n"}},
		[]Segment{Segment{"電腦", "n"}, Segment{"修好", "v"}, Segment{"了", "ul"}},
		[]Segment{Segment{"做好", "v"}, Segment{"了", "ul"}, Segment{"這件", "mq"}, Segment{"事情", "n"}, Segment{"就", "d"}, Segment{"一了百了", "l"}, Segment{"了", "ul"}},
		[]Segment{Segment{"人們", "n"}, Segment{"審美", "vn"}, Segment{"的", "uj"}, Segment{"觀點", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"我們", "r"}, Segment{"買", "v"}, Segment{"了", "ul"}, Segment{"一個", "m"}, Segment{"美的", "nr"}, Segment{"空調", "n"}},
		[]Segment{Segment{"線程", "n"}, Segment{"初始化", "l"}, Segment{"時", "n"}, Segment{"我們", "r"}, Segment{"要", "v"}, Segment{"注意", "v"}},
		[]Segment{Segment{"一個", "m"}, Segment{"分子", "n"}, Segment{"是", "v"}, Segment{"由", "p"}, Segment{"好多", "m"}, Segment{"原子", "n"}, Segment{"組織", "v"}, Segment{"成", "v"}, Segment{"的", "uj"}},
		[]Segment{Segment{"祝", "v"}, Segment{"你", "r"}, Segment{"馬到功成", "i"}},
		[]Segment{Segment{"他", "r"}, Segment{"掉", "v"}, Segment{"進", "v"}, Segment{"了", "ul"}, Segment{"無底洞", "ns"}, Segment{"里", "f"}},
		[]Segment{Segment{"中國", "ns"}, Segment{"的", "uj"}, Segment{"首都", "d"}, Segment{"是", "v"}, Segment{"北京", "ns"}},
		[]Segment{Segment{"孫君意", "nr"}},
		[]Segment{Segment{"外交部", "nt"}, Segment{"發言人", "l"}, Segment{"馬朝旭", "nr"}},
		[]Segment{Segment{"領導人", "n"}, Segment{"會議", "n"}, Segment{"和", "c"}, Segment{"第四屆", "m"}, Segment{"東亞", "ns"}, Segment{"峰會", "n"}},
		[]Segment{Segment{"在", "p"}, Segment{"過去", "t"}, Segment{"的", "uj"}, Segment{"這", "r"}, Segment{"五年", "t"}},
		[]Segment{Segment{"還", "d"}, Segment{"需要", "v"}, Segment{"很", "d"}, Segment{"長", "a"}, Segment{"的", "uj"}, Segment{"路", "n"}, Segment{"要", "v"}, Segment{"走", "v"}},
		[]Segment{Segment{"60", "m"}, Segment{"週年", "t"}, Segment{"首都", "d"}, Segment{"閱兵", "v"}},
		[]Segment{Segment{"你好", "l"}, Segment{"人們", "n"}, Segment{"審美", "vn"}, Segment{"的", "uj"}, Segment{"觀點", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"買", "v"}, Segment{"水果", "n"}, Segment{"然後", "c"}, Segment{"來", "v"}, Segment{"世博園", "nr"}},
		[]Segment{Segment{"買", "v"}, Segment{"水果", "n"}, Segment{"然後", "c"}, Segment{"去", "v"}, Segment{"世博園", "nr"}},
		[]Segment{Segment{"但是", "c"}, Segment{"後來", "t"}, Segment{"我", "r"}, Segment{"才", "d"}, Segment{"知道", "v"}, Segment{"你", "r"}, Segment{"是", "v"}, Segment{"對", "p"}, Segment{"的", "uj"}},
		[]Segment{Segment{"存在", "v"}, Segment{"即", "v"}, Segment{"合理", "vn"}},
		[]Segment{Segment{"的的", "u"}, Segment{"的的", "u"}, Segment{"的", "uj"}, Segment{"在的", "u"}, Segment{"的的", "u"}, Segment{"的", "uj"}, Segment{"就", "d"}, Segment{"以", "p"}, Segment{"和和", "nz"}, Segment{"和", "c"}},
		[]Segment{Segment{"I", "x"}, Segment{" ", "x"}, Segment{"love", "eng"}, Segment{"你", "r"}, Segment{"，", "x"}, Segment{"不以為恥", "i"}, Segment{"，", "x"}, Segment{"反", "zg"}, Segment{"以為", "c"}, Segment{"rong", "eng"}},
		[]Segment{Segment{"因", "p"}},
		[]Segment{},
		[]Segment{Segment{"hello", "eng"}, Segment{"你好", "l"}, Segment{"人們", "n"}, Segment{"審美", "vn"}, Segment{"的", "uj"}, Segment{"觀點", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"很好", "a"}, Segment{"但", "c"}, Segment{"主要", "b"}, Segment{"是", "v"}, Segment{"基於", "p"}, Segment{"網頁", "n"}, Segment{"形式", "n"}},
		[]Segment{Segment{"hello", "eng"}, Segment{"你好", "l"}, Segment{"人們", "n"}, Segment{"審美", "vn"}, Segment{"的", "uj"}, Segment{"觀點", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"為什麼", "r"}, Segment{"我", "r"}, Segment{"不能", "v"}, Segment{"擁有", "v"}, Segment{"想要", "v"}, Segment{"的", "uj"}, Segment{"生活", "vn"}},
		[]Segment{Segment{"後來", "t"}, Segment{"我", "r"}, Segment{"才", "d"}},
		[]Segment{Segment{"此次", "r"}, Segment{"來", "v"}, Segment{"中國", "ns"}, Segment{"是", "v"}, Segment{"為了", "p"}},
		[]Segment{Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解決", "v"}, Segment{"一些", "m"}, Segment{"問題", "n"}},
		[]Segment{Segment{",", "x"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解決", "v"}, Segment{"一些", "m"}, Segment{"問題", "n"}},
		[]Segment{Segment{"其實", "d"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解決", "v"}, Segment{"一些", "m"}, Segment{"問題", "n"}},
		[]Segment{Segment{"好人", "n"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解決", "v"}, Segment{"一些", "m"}, Segment{"問題", "n"}},
		[]Segment{Segment{"是因為", "c"}, Segment{"和", "c"}, Segment{"國家", "n"}},
		[]Segment{Segment{"老年", "t"}, Segment{"搜索", "v"}, Segment{"還", "d"}, Segment{"支持", "v"}},
		[]Segment{Segment{"乾脆", "d"}, Segment{"就", "d"}, Segment{"把", "p"}, Segment{"那部", "r"}, Segment{"蒙人", "n"}, Segment{"的", "uj"}, Segment{"閒法", "n"}, Segment{"給", "p"}, Segment{"廢", "v"}, Segment{"了", "ul"}, Segment{"拉倒", "v"}, Segment{"！", "x"}, Segment{"RT", "eng"}, Segment{" ", "x"}, Segment{"@", "x"}, Segment{"laoshipukong", "eng"}, Segment{" ", "x"}, Segment{":", "x"}, Segment{" ", "x"}, Segment{"27", "m"}, Segment{"日", "m"}, Segment{"，", "x"}, Segment{"全國人大常委會", "nt"}, Segment{"第三次", "m"}, Segment{"審議", "v"}, Segment{"侵權", "v"}, Segment{"責任法", "n"}, Segment{"草案", "n"}, Segment{"，", "x"}, Segment{"刪除", "v"}, Segment{"了", "ul"}, Segment{"有關", "vn"}, Segment{"醫療", "n"}, Segment{"損害", "v"}, Segment{"責任", "n"}, Segment{"「", "x"}, Segment{"舉證", "v"}, Segment{"倒置", "v"}, Segment{"」", "x"}, Segment{"的", "uj"}, Segment{"規定", "n"}, Segment{"。", "x"}, Segment{"在", "p"}, Segment{"醫患", "n"}, Segment{"糾紛", "n"}, Segment{"中本", "ns"}, Segment{"已", "d"}, Segment{"處於", "v"}, Segment{"弱勢", "n"}, Segment{"地位", "n"}, Segment{"的", "uj"}, Segment{"消費者", "n"}, Segment{"由此", "c"}, Segment{"將", "d"}, Segment{"陷入", "v"}, Segment{"萬劫不復", "i"}, Segment{"的", "uj"}, Segment{"境地", "s"}, Segment{"。", "x"}, Segment{" ", "x"}},
		[]Segment{Segment{"大", "a"}},
		[]Segment{},
		[]Segment{Segment{"他", "r"}, Segment{"說", "v"}, Segment{"的", "uj"}, Segment{"確實", "ad"}, Segment{"在", "p"}, Segment{"理", "n"}},
		[]Segment{Segment{"長春", "ns"}, Segment{"市長", "n"}, Segment{"春節", "t"}, Segment{"講話", "n"}},
		[]Segment{Segment{"結婚", "v"}, Segment{"的", "uj"}, Segment{"和", "c"}, Segment{"尚未", "d"}, Segment{"結婚", "v"}, Segment{"的", "uj"}},
		[]Segment{Segment{"結合", "v"}, Segment{"成", "n"}, Segment{"分子", "n"}, Segment{"時", "n"}},
		[]Segment{Segment{"旅遊", "vn"}, Segment{"和", "c"}, Segment{"服務", "vn"}, Segment{"是", "v"}, Segment{"最好", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"這件", "mq"}, Segment{"事情", "n"}, Segment{"的確", "d"}, Segment{"是", "v"}, Segment{"我", "r"}, Segment{"的", "uj"}, Segment{"錯", "n"}},
		[]Segment{Segment{"供", "v"}, Segment{"大家", "n"}, Segment{"參考", "v"}, Segment{"指正", "v"}},
		[]Segment{Segment{"哈爾濱", "ns"}, Segment{"政府", "n"}, Segment{"公佈", "v"}, Segment{"塌", "v"}, Segment{"橋", "n"}, Segment{"原因", "n"}},
		[]Segment{Segment{"我", "r"}, Segment{"在", "p"}, Segment{"機場", "n"}, Segment{"入口處", "i"}},
		[]Segment{Segment{"邢永臣", "nr"}, Segment{"攝影", "n"}, Segment{"報道", "v"}},
		[]Segment{Segment{"BP", "eng"}, Segment{"神經網絡", "n"}, Segment{"如何", "r"}, Segment{"訓練", "vn"}, Segment{"才能", "v"}, Segment{"在", "p"}, Segment{"分類", "n"}, Segment{"時", "n"}, Segment{"增加", "v"}, Segment{"區分度", "n"}, Segment{"？", "x"}},
		[]Segment{Segment{"南京市", "ns"}, Segment{"長江大橋", "ns"}},
		[]Segment{Segment{"應", "v"}, Segment{"一些", "m"}, Segment{"使用者", "n"}, Segment{"的", "uj"}, Segment{"建議", "n"}, Segment{"，", "x"}, Segment{"也", "d"}, Segment{"為了", "p"}, Segment{"便於", "v"}, Segment{"利用", "n"}, Segment{"NiuTrans", "eng"}, Segment{"用於", "v"}, Segment{"SMT", "eng"}, Segment{"研究", "vn"}},
		[]Segment{Segment{"長春市", "ns"}, Segment{"長春", "ns"}, Segment{"藥店", "n"}},
		[]Segment{Segment{"鄧穎超", "nr"}, Segment{"生前", "t"}, Segment{"最", "d"}, Segment{"喜歡", "v"}, Segment{"的", "uj"}, Segment{"衣服", "n"}},
		[]Segment{Segment{"胡錦濤", "nr"}, Segment{"是", "v"}, Segment{"熱愛", "a"}, Segment{"世界", "n"}, Segment{"和平", "nz"}, Segment{"的", "uj"}, Segment{"政治局", "n"}, Segment{"常委", "j"}},
		[]Segment{Segment{"程序員", "n"}, Segment{"祝", "v"}, Segment{"海林", "nz"}, Segment{"和", "c"}, Segment{"朱會震", "nr"}, Segment{"是", "v"}, Segment{"在", "p"}, Segment{"孫健", "nr"}, Segment{"的", "uj"}, Segment{"左面", "f"}, Segment{"和", "c"}, Segment{"右面", "f"}, Segment{",", "x"}, Segment{" ", "x"}, Segment{"範凱", "nr"}, Segment{"在", "p"}, Segment{"最", "a"}, Segment{"右面", "f"}, Segment{".", "m"}, Segment{"再往", "d"}, Segment{"左", "f"}, Segment{"是", "v"}, Segment{"李松洪", "nr"}},
		[]Segment{Segment{"一次性", "d"}, Segment{"交", "v"}, Segment{"多少", "m"}, Segment{"錢", "n"}},
		[]Segment{Segment{"兩塊", "m"}, Segment{"五", "m"}, Segment{"一套", "m"}, Segment{"，", "x"}, Segment{"三塊", "m"}, Segment{"八", "m"}, Segment{"一斤", "m"}, Segment{"，", "x"}, Segment{"四塊", "m"}, Segment{"七", "m"}, Segment{"一本", "m"}, Segment{"，", "x"}, Segment{"五塊", "m"}, Segment{"六", "m"}, Segment{"一條", "m"}},
		[]Segment{Segment{"小", "a"}, Segment{"和尚", "nr"}, Segment{"留", "v"}, Segment{"了", "ul"}, Segment{"一個", "m"}, Segment{"像", "v"}, Segment{"大", "a"}, Segment{"和尚", "nr"}, Segment{"一樣", "r"}, Segment{"的", "uj"}, Segment{"和尚頭", "nr"}},
		[]Segment{Segment{"我", "r"}, Segment{"是", "v"}, Segment{"中華人民共和國", "ns"}, Segment{"公民", "n"}, Segment{";", "x"}, Segment{"我", "r"}, Segment{"爸爸", "n"}, Segment{"是", "v"}, Segment{"共和黨", "nt"}, Segment{"黨員", "n"}, Segment{";", "x"}, Segment{" ", "x"}, Segment{"地鐵", "n"}, Segment{"和平門", "ns"}, Segment{"站", "v"}},
		[]Segment{Segment{"張曉梅", "nr"}, Segment{"去", "v"}, Segment{"人民", "n"}, Segment{"醫院", "n"}, Segment{"做", "v"}, Segment{"了", "ul"}, Segment{"個", "q"}, Segment{"B超", "n"}, Segment{"然後", "c"}, Segment{"去", "v"}, Segment{"買", "v"}, Segment{"了", "ul"}, Segment{"件", "q"}, Segment{"T恤", "n"}},
		[]Segment{Segment{"AT&T", "nz"}, Segment{"是", "v"}, Segment{"一件", "m"}, Segment{"不錯", "a"}, Segment{"的", "uj"}, Segment{"公司", "n"}, Segment{"，", "x"}, Segment{"給", "p"}, Segment{"你", "r"}, Segment{"發", "v"}, Segment{"offer", "eng"}, Segment{"了", "ul"}, Segment{"嗎", "y"}, Segment{"？", "x"}},
		[]Segment{Segment{"C++", "nz"}, Segment{"和", "c"}, Segment{"c#", "nz"}, Segment{"是", "v"}, Segment{"什麼", "r"}, Segment{"關係", "n"}, Segment{"？", "x"}, Segment{"11", "m"}, Segment{"+", "x"}, Segment{"122", "m"}, Segment{"=", "x"}, Segment{"133", "m"}, Segment{"，", "x"}, Segment{"是", "v"}, Segment{"嗎", "y"}, Segment{"？", "x"}, Segment{"PI", "eng"}, Segment{"=", "x"}, Segment{"3.14159", "m"}},
		[]Segment{Segment{"你", "r"}, Segment{"認識", "v"}, Segment{"那個", "r"}, Segment{"和", "c"}, Segment{"主席", "n"}, Segment{"握手", "v"}, Segment{"的", "uj"}, Segment{"的哥", "n"}, Segment{"嗎", "y"}, Segment{"？", "x"}, Segment{"他", "r"}, Segment{"開", "v"}, Segment{"一輛", "m"}, Segment{"黑色", "n"}, Segment{"的士", "n"}, Segment{"。", "x"}},
		[]Segment{Segment{"槍桿子", "n"}, Segment{"中", "f"}, Segment{"出", "v"}, Segment{"政權", "n"}},
	}
	noHMMCutResult = [][]Segment{
		[]Segment{Segment{"這", "r"}, Segment{"是", "v"}, Segment{"一個", "m"}, Segment{"伸手不見五指", "i"}, Segment{"的", "uj"}, Segment{"黑夜", "n"}, Segment{"。", "x"}, Segment{"我", "r"}, Segment{"叫", "v"}, Segment{"孫悟空", "nr"}, Segment{"，", "x"}, Segment{"我", "r"}, Segment{"愛", "v"}, Segment{"北京", "ns"}, Segment{"，", "x"}, Segment{"我", "r"}, Segment{"愛", "v"}, Segment{"Python", "eng"}, Segment{"和", "c"}, Segment{"C++", "nz"}, Segment{"。", "x"}},
		[]Segment{Segment{"我", "r"}, Segment{"不", "d"}, Segment{"喜歡", "v"}, Segment{"日本", "ns"}, Segment{"和服", "nz"}, Segment{"。", "x"}},
		[]Segment{Segment{"雷猴", "n"}, Segment{"回歸", "v"}, Segment{"人間", "n"}, Segment{"。", "x"}},
		[]Segment{Segment{"工信處", "n"}, Segment{"女幹事", "n"}, Segment{"每月", "r"}, Segment{"經過", "p"}, Segment{"下屬", "v"}, Segment{"科室", "n"}, Segment{"都", "d"}, Segment{"要", "v"}, Segment{"親口", "n"}, Segment{"交代", "n"}, Segment{"24", "eng"}, Segment{"口", "q"}, Segment{"交換機", "n"}, Segment{"等", "u"}, Segment{"技術性", "n"}, Segment{"器件", "n"}, Segment{"的", "uj"}, Segment{"安裝", "v"}, Segment{"工作", "vn"}},
		[]Segment{Segment{"我", "r"}, Segment{"需要", "v"}, Segment{"廉租房", "n"}},
		[]Segment{Segment{"永和", "nz"}, Segment{"服裝", "vn"}, Segment{"飾品", "n"}, Segment{"有限公司", "n"}},
		[]Segment{Segment{"我", "r"}, Segment{"愛", "v"}, Segment{"北京", "ns"}, Segment{"天安門", "ns"}},
		[]Segment{Segment{"abc", "eng"}},
		[]Segment{Segment{"隱", "n"}, Segment{"馬爾可夫", "nr"}},
		[]Segment{Segment{"雷猴", "n"}, Segment{"是", "v"}, Segment{"個", "q"}, Segment{"好", "a"}, Segment{"網站", "n"}},
		[]Segment{Segment{"「", "x"}, Segment{"Microsoft", "eng"}, Segment{"」", "x"}, Segment{"一", "m"}, Segment{"詞", "n"}, Segment{"由", "p"}, Segment{"「", "x"}, Segment{"MICROcomputer", "eng"}, Segment{"（", "x"}, Segment{"微型", "b"}, Segment{"計算機", "n"}, Segment{"）", "x"}, Segment{"」", "x"}, Segment{"和", "c"}, Segment{"「", "x"}, Segment{"SOFTware", "eng"}, Segment{"（", "x"}, Segment{"軟件", "n"}, Segment{"）", "x"}, Segment{"」", "x"}, Segment{"兩", "m"}, Segment{"部分", "n"}, Segment{"組成", "v"}},
		[]Segment{Segment{"草泥馬", "n"}, Segment{"和", "c"}, Segment{"欺", "vn"}, Segment{"實", "n"}, Segment{"馬", "n"}, Segment{"是", "v"}, Segment{"今年", "t"}, Segment{"的", "uj"}, Segment{"流行", "v"}, Segment{"詞彙", "n"}},
		[]Segment{Segment{"伊", "ns"}, Segment{"藤", "nr"}, Segment{"洋華堂", "n"}, Segment{"總府", "n"}, Segment{"店", "n"}},
		[]Segment{Segment{"中國科學院計算技術研究所", "nt"}},
		[]Segment{Segment{"羅密歐", "nr"}, Segment{"與", "p"}, Segment{"朱麗葉", "nr"}},
		[]Segment{Segment{"我", "r"}, Segment{"購買", "v"}, Segment{"了", "ul"}, Segment{"道具", "n"}, Segment{"和", "c"}, Segment{"服裝", "vn"}},
		[]Segment{Segment{"PS", "eng"}, Segment{":", "x"}, Segment{" ", "x"}, Segment{"我", "r"}, Segment{"覺得", "v"}, Segment{"開源", "n"}, Segment{"有", "v"}, Segment{"一個", "m"}, Segment{"好處", "d"}, Segment{"，", "x"}, Segment{"就是", "d"}, Segment{"能夠", "v"}, Segment{"敦促", "v"}, Segment{"自己", "r"}, Segment{"不斷改進", "l"}, Segment{"，", "x"}, Segment{"避免", "v"}, Segment{"敞", "v"}, Segment{"帚", "ng"}, Segment{"自珍", "b"}},
		[]Segment{Segment{"湖北省", "ns"}, Segment{"石首市", "ns"}},
		[]Segment{Segment{"湖北省", "ns"}, Segment{"十堰市", "ns"}},
		[]Segment{Segment{"總經理", "n"}, Segment{"完成", "v"}, Segment{"了", "ul"}, Segment{"這件", "mq"}, Segment{"事情", "n"}},
		[]Segment{Segment{"電腦", "n"}, Segment{"修好", "v"}, Segment{"了", "ul"}},
		[]Segment{Segment{"做好", "v"}, Segment{"了", "ul"}, Segment{"這件", "mq"}, Segment{"事情", "n"}, Segment{"就", "d"}, Segment{"一了百了", "l"}, Segment{"了", "ul"}},
		[]Segment{Segment{"人們", "n"}, Segment{"審美", "vn"}, Segment{"的", "uj"}, Segment{"觀點", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"我們", "r"}, Segment{"買", "v"}, Segment{"了", "ul"}, Segment{"一個", "m"}, Segment{"美的", "nr"}, Segment{"空調", "n"}},
		[]Segment{Segment{"線程", "n"}, Segment{"初始化", "l"}, Segment{"時", "n"}, Segment{"我們", "r"}, Segment{"要", "v"}, Segment{"注意", "v"}},
		[]Segment{Segment{"一個", "m"}, Segment{"分子", "n"}, Segment{"是", "v"}, Segment{"由", "p"}, Segment{"好多", "m"}, Segment{"原子", "n"}, Segment{"組織", "v"}, Segment{"成", "n"}, Segment{"的", "uj"}},
		[]Segment{Segment{"祝", "v"}, Segment{"你", "r"}, Segment{"馬到功成", "i"}},
		[]Segment{Segment{"他", "r"}, Segment{"掉", "zg"}, Segment{"進", "v"}, Segment{"了", "ul"}, Segment{"無底洞", "ns"}, Segment{"里", "f"}},
		[]Segment{Segment{"中國", "ns"}, Segment{"的", "uj"}, Segment{"首都", "d"}, Segment{"是", "v"}, Segment{"北京", "ns"}},
		[]Segment{Segment{"孫", "zg"}, Segment{"君", "nz"}, Segment{"意", "n"}},
		[]Segment{Segment{"外交部", "nt"}, Segment{"發言人", "l"}, Segment{"馬朝旭", "nr"}},
		[]Segment{Segment{"領導人", "n"}, Segment{"會議", "n"}, Segment{"和", "c"}, Segment{"第四屆", "m"}, Segment{"東亞", "ns"}, Segment{"峰會", "n"}},
		[]Segment{Segment{"在", "p"}, Segment{"過去", "t"}, Segment{"的", "uj"}, Segment{"這", "r"}, Segment{"五年", "t"}},
		[]Segment{Segment{"還", "d"}, Segment{"需要", "v"}, Segment{"很", "zg"}, Segment{"長", "a"}, Segment{"的", "uj"}, Segment{"路", "n"}, Segment{"要", "v"}, Segment{"走", "v"}},
		[]Segment{Segment{"60", "eng"}, Segment{"週年", "t"}, Segment{"首都", "d"}, Segment{"閱兵", "v"}},
		[]Segment{Segment{"你好", "l"}, Segment{"人們", "n"}, Segment{"審美", "vn"}, Segment{"的", "uj"}, Segment{"觀點", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"買", "v"}, Segment{"水果", "n"}, Segment{"然後", "c"}, Segment{"來", "v"}, Segment{"世博園", "nr"}},
		[]Segment{Segment{"買", "v"}, Segment{"水果", "n"}, Segment{"然後", "c"}, Segment{"去", "v"}, Segment{"世博園", "nr"}},
		[]Segment{Segment{"但是", "c"}, Segment{"後來", "t"}, Segment{"我", "r"}, Segment{"才", "d"}, Segment{"知道", "v"}, Segment{"你", "r"}, Segment{"是", "v"}, Segment{"對", "p"}, Segment{"的", "uj"}},
		[]Segment{Segment{"存在", "v"}, Segment{"即", "v"}, Segment{"合理", "vn"}},
		[]Segment{Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"在", "p"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"的", "uj"}, Segment{"就", "d"}, Segment{"以", "p"}, Segment{"和", "c"}, Segment{"和", "c"}, Segment{"和", "c"}},
		[]Segment{Segment{"I", "eng"}, Segment{" ", "x"}, Segment{"love", "eng"}, Segment{"你", "r"}, Segment{"，", "x"}, Segment{"不以為恥", "i"}, Segment{"，", "x"}, Segment{"反", "zg"}, Segment{"以為", "c"}, Segment{"rong", "eng"}},
		[]Segment{Segment{"因", "p"}},
		[]Segment{},
		[]Segment{Segment{"hello", "eng"}, Segment{"你好", "l"}, Segment{"人們", "n"}, Segment{"審美", "vn"}, Segment{"的", "uj"}, Segment{"觀點", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"很", "zg"}, Segment{"好", "a"}, Segment{"但", "c"}, Segment{"主要", "b"}, Segment{"是", "v"}, Segment{"基於", "p"}, Segment{"網頁", "n"}, Segment{"形式", "n"}},
		[]Segment{Segment{"hello", "eng"}, Segment{"你好", "l"}, Segment{"人們", "n"}, Segment{"審美", "vn"}, Segment{"的", "uj"}, Segment{"觀點", "n"}, Segment{"是", "v"}, Segment{"不同", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"為什麼", "r"}, Segment{"我", "r"}, Segment{"不能", "v"}, Segment{"擁有", "v"}, Segment{"想要", "v"}, Segment{"的", "uj"}, Segment{"生活", "vn"}},
		[]Segment{Segment{"後來", "t"}, Segment{"我", "r"}, Segment{"才", "d"}},
		[]Segment{Segment{"此次", "r"}, Segment{"來", "v"}, Segment{"中國", "ns"}, Segment{"是", "v"}, Segment{"為了", "p"}},
		[]Segment{Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解決", "v"}, Segment{"一些", "m"}, Segment{"問題", "n"}},
		[]Segment{Segment{",", "x"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解決", "v"}, Segment{"一些", "m"}, Segment{"問題", "n"}},
		[]Segment{Segment{"其實", "d"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解決", "v"}, Segment{"一些", "m"}, Segment{"問題", "n"}},
		[]Segment{Segment{"好人", "n"}, Segment{"使用", "v"}, Segment{"了", "ul"}, Segment{"它", "r"}, Segment{"就", "d"}, Segment{"可以", "c"}, Segment{"解決", "v"}, Segment{"一些", "m"}, Segment{"問題", "n"}},
		[]Segment{Segment{"是因為", "c"}, Segment{"和", "c"}, Segment{"國家", "n"}},
		[]Segment{Segment{"老年", "t"}, Segment{"搜索", "v"}, Segment{"還", "d"}, Segment{"支持", "v"}},
		[]Segment{Segment{"乾脆", "d"}, Segment{"就", "d"}, Segment{"把", "p"}, Segment{"那", "r"}, Segment{"部", "n"}, Segment{"蒙", "v"}, Segment{"人", "n"}, Segment{"的", "uj"}, Segment{"閒", "n"}, Segment{"法", "j"}, Segment{"給", "p"}, Segment{"廢", "v"}, Segment{"了", "ul"}, Segment{"拉倒", "v"}, Segment{"！", "x"}, Segment{"RT", "eng"}, Segment{" ", "x"}, Segment{"@", "x"}, Segment{"laoshipukong", "eng"}, Segment{" ", "x"}, Segment{":", "x"}, Segment{" ", "x"}, Segment{"27", "eng"}, Segment{"日", "m"}, Segment{"，", "x"}, Segment{"全國人大常委會", "nt"}, Segment{"第三次", "m"}, Segment{"審議", "v"}, Segment{"侵權", "v"}, Segment{"責任法", "n"}, Segment{"草案", "n"}, Segment{"，", "x"}, Segment{"刪除", "v"}, Segment{"了", "ul"}, Segment{"有關", "vn"}, Segment{"醫療", "n"}, Segment{"損害", "v"}, Segment{"責任", "n"}, Segment{"「", "x"}, Segment{"舉證", "v"}, Segment{"倒置", "v"}, Segment{"」", "x"}, Segment{"的", "uj"}, Segment{"規定", "n"}, Segment{"。", "x"}, Segment{"在", "p"}, Segment{"醫患", "n"}, Segment{"糾紛", "n"}, Segment{"中", "f"}, Segment{"本", "r"}, Segment{"已", "d"}, Segment{"處於", "v"}, Segment{"弱勢", "n"}, Segment{"地位", "n"}, Segment{"的", "uj"}, Segment{"消費者", "n"}, Segment{"由此", "c"}, Segment{"將", "d"}, Segment{"陷入", "v"}, Segment{"萬劫不復", "i"}, Segment{"的", "uj"}, Segment{"境地", "s"}, Segment{"。", "x"}, Segment{" ", "x"}},
		[]Segment{Segment{"大", "a"}},
		[]Segment{},
		[]Segment{Segment{"他", "r"}, Segment{"說", "v"}, Segment{"的", "uj"}, Segment{"確實", "ad"}, Segment{"在", "p"}, Segment{"理", "n"}},
		[]Segment{Segment{"長春", "ns"}, Segment{"市長", "n"}, Segment{"春節", "t"}, Segment{"講話", "n"}},
		[]Segment{Segment{"結婚", "v"}, Segment{"的", "uj"}, Segment{"和", "c"}, Segment{"尚未", "d"}, Segment{"結婚", "v"}, Segment{"的", "uj"}},
		[]Segment{Segment{"結合", "v"}, Segment{"成", "n"}, Segment{"分子", "n"}, Segment{"時", "n"}},
		[]Segment{Segment{"旅遊", "vn"}, Segment{"和", "c"}, Segment{"服務", "vn"}, Segment{"是", "v"}, Segment{"最好", "a"}, Segment{"的", "uj"}},
		[]Segment{Segment{"這件", "mq"}, Segment{"事情", "n"}, Segment{"的確", "d"}, Segment{"是", "v"}, Segment{"我", "r"}, Segment{"的", "uj"}, Segment{"錯", "v"}},
		[]Segment{Segment{"供", "v"}, Segment{"大家", "n"}, Segment{"參考", "v"}, Segment{"指正", "v"}},
		[]Segment{Segment{"哈爾濱", "ns"}, Segment{"政府", "n"}, Segment{"公佈", "v"}, Segment{"塌", "v"}, Segment{"橋", "n"}, Segment{"原因", "n"}},
		[]Segment{Segment{"我", "r"}, Segment{"在", "p"}, Segment{"機場", "n"}, Segment{"入口處", "i"}},
		[]Segment{Segment{"邢", "nr"}, Segment{"永", "ns"}, Segment{"臣", "n"}, Segment{"攝影", "n"}, Segment{"報道", "v"}},
		[]Segment{Segment{"BP", "eng"}, Segment{"神經網絡", "n"}, Segment{"如何", "r"}, Segment{"訓練", "vn"}, Segment{"才能", "v"}, Segment{"在", "p"}, Segment{"分類", "n"}, Segment{"時", "n"}, Segment{"增加", "v"}, Segment{"區分度", "n"}, Segment{"？", "x"}},
		[]Segment{Segment{"南京市", "ns"}, Segment{"長江大橋", "ns"}},
		[]Segment{Segment{"應", "v"}, Segment{"一些", "m"}, Segment{"使用者", "n"}, Segment{"的", "uj"}, Segment{"建議", "n"}, Segment{"，", "x"}, Segment{"也", "d"}, Segment{"為了", "p"}, Segment{"便於", "v"}, Segment{"利用", "n"}, Segment{"NiuTrans", "eng"}, Segment{"用於", "v"}, Segment{"SMT", "eng"}, Segment{"研究", "vn"}},
		[]Segment{Segment{"長春市", "ns"}, Segment{"長春", "ns"}, Segment{"藥店", "n"}},
		[]Segment{Segment{"鄧穎超", "nr"}, Segment{"生前", "t"}, Segment{"最", "d"}, Segment{"喜歡", "v"}, Segment{"的", "uj"}, Segment{"衣服", "n"}},
		[]Segment{Segment{"胡錦濤", "nr"}, Segment{"是", "v"}, Segment{"熱愛", "a"}, Segment{"世界", "n"}, Segment{"和平", "nz"}, Segment{"的", "uj"}, Segment{"政治局", "n"}, Segment{"常委", "j"}},
		[]Segment{Segment{"程序員", "n"}, Segment{"祝", "v"}, Segment{"海林", "nz"}, Segment{"和", "c"}, Segment{"朱", "nr"}, Segment{"會", "v"}, Segment{"震", "v"}, Segment{"是", "v"}, Segment{"在", "p"}, Segment{"孫", "zg"}, Segment{"健", "a"}, Segment{"的", "uj"}, Segment{"左面", "f"}, Segment{"和", "c"}, Segment{"右面", "f"}, Segment{",", "x"}, Segment{" ", "x"}, Segment{"範", "nr"}, Segment{"凱", "nr"}, Segment{"在", "p"}, Segment{"最", "d"}, Segment{"右面", "f"}, Segment{".", "x"}, Segment{"再", "d"}, Segment{"往", "zg"}, Segment{"左", "m"}, Segment{"是", "v"}, Segment{"李", "nr"}, Segment{"松", "v"}, Segment{"洪", "nr"}},
		[]Segment{Segment{"一次性", "d"}, Segment{"交", "v"}, Segment{"多少", "m"}, Segment{"錢", "n"}},
		[]Segment{Segment{"兩塊", "m"}, Segment{"五", "m"}, Segment{"一套", "m"}, Segment{"，", "x"}, Segment{"三塊", "m"}, Segment{"八", "m"}, Segment{"一斤", "m"}, Segment{"，", "x"}, Segment{"四塊", "m"}, Segment{"七", "m"}, Segment{"一本", "m"}, Segment{"，", "x"}, Segment{"五塊", "m"}, Segment{"六", "m"}, Segment{"一條", "m"}},
		[]Segment{Segment{"小", "a"}, Segment{"和尚", "nr"}, Segment{"留", "v"}, Segment{"了", "ul"}, Segment{"一個", "m"}, Segment{"像", "v"}, Segment{"大", "a"}, Segment{"和尚", "nr"}, Segment{"一樣", "r"}, Segment{"的", "uj"}, Segment{"和尚頭", "nr"}},
		[]Segment{Segment{"我", "r"}, Segment{"是", "v"}, Segment{"中華人民共和國", "ns"}, Segment{"公民", "n"}, Segment{";", "x"}, Segment{"我", "r"}, Segment{"爸爸", "n"}, Segment{"是", "v"}, Segment{"共和黨", "nt"}, Segment{"黨員", "n"}, Segment{";", "x"}, Segment{" ", "x"}, Segment{"地鐵", "n"}, Segment{"和平門", "ns"}, Segment{"站", "v"}},
		[]Segment{Segment{"張曉梅", "nr"}, Segment{"去", "v"}, Segment{"人民", "n"}, Segment{"醫院", "n"}, Segment{"做", "v"}, Segment{"了", "ul"}, Segment{"個", "q"}, Segment{"B超", "n"}, Segment{"然後", "c"}, Segment{"去", "v"}, Segment{"買", "v"}, Segment{"了", "ul"}, Segment{"件", "zg"}, Segment{"T恤", "n"}},
		[]Segment{Segment{"AT&T", "nz"}, Segment{"是", "v"}, Segment{"一件", "m"}, Segment{"不錯", "a"}, Segment{"的", "uj"}, Segment{"公司", "n"}, Segment{"，", "x"}, Segment{"給", "p"}, Segment{"你", "r"}, Segment{"發", "v"}, Segment{"offer", "eng"}, Segment{"了", "ul"}, Segment{"嗎", "y"}, Segment{"？", "x"}},
		[]Segment{Segment{"C++", "nz"}, Segment{"和", "c"}, Segment{"c#", "nz"}, Segment{"是", "v"}, Segment{"什麼", "r"}, Segment{"關係", "n"}, Segment{"？", "x"}, Segment{"11", "eng"}, Segment{"+", "x"}, Segment{"122", "eng"}, Segment{"=", "x"}, Segment{"133", "eng"}, Segment{"，", "x"}, Segment{"是", "v"}, Segment{"嗎", "y"}, Segment{"？", "x"}, Segment{"PI", "eng"}, Segment{"=", "x"}, Segment{"3", "eng"}, Segment{".", "x"}, Segment{"14159", "eng"}},
		[]Segment{Segment{"你", "r"}, Segment{"認識", "v"}, Segment{"那個", "r"}, Segment{"和", "c"}, Segment{"主席", "n"}, Segment{"握手", "v"}, Segment{"的", "uj"}, Segment{"的哥", "n"}, Segment{"嗎", "y"}, Segment{"？", "x"}, Segment{"他", "r"}, Segment{"開", "v"}, Segment{"一輛", "m"}, Segment{"黑色", "n"}, Segment{"的士", "n"}, Segment{"。", "x"}},
		[]Segment{Segment{"槍桿子", "n"}, Segment{"中", "f"}, Segment{"出", "v"}, Segment{"政權", "n"}},
	}
)

func init() {
	seg.LoadDictionary("../dict.txt")
}

func chanToArray(ch <-chan Segment) []Segment {
	var result []Segment
	for word := range ch {
		result = append(result, word)
	}
	return result
}

func TestCut(t *testing.T) {
	for index, content := range testContents {
		result := chanToArray(seg.Cut(content, true))
		if len(defaultCutResult[index]) != len(result) {
			t.Errorf("default cut for %s length should be %d not %d\n",
				content, len(defaultCutResult[index]), len(result))
			t.Errorf("expect: %v\n", defaultCutResult[index])
			t.Fatalf("got: %v\n", result)
		}
		for i := range result {
			if result[i] != defaultCutResult[index][i] {
				t.Fatalf("expect %s, got %s", defaultCutResult[index][i], result[i])
			}
		}
		result = chanToArray(seg.Cut(content, false))
		if len(noHMMCutResult[index]) != len(result) {
			t.Fatal(content)
		}
		for i := range result {
			if result[i] != noHMMCutResult[index][i] {
				t.Fatal(content)
			}
		}

	}
}

// https://github.com/fxsjy/jieba/issues/132
func TestBug132(t *testing.T) {
	sentence := "又跛又啞"
	cutResult := []Segment{
		Segment{"又", "d"},
		Segment{"跛", "a"},
		Segment{"又", "d"},
		Segment{"啞", "v"},
	}
	result := chanToArray(seg.Cut(sentence, true))
	if len(cutResult) != len(result) {
		t.Fatal(result)
	}
	for i := range result {
		if result[i] != cutResult[i] {
			t.Fatal(result[i])
		}
	}
}

// https://github.com/fxsjy/jieba/issues/137
func TestBug137(t *testing.T) {
	sentence := "前港督衛奕信在八八年十月宣佈成立中央政策研究組"
	cutResult := []Segment{
		Segment{"前", "f"},
		Segment{"港督", "n"},
		Segment{"衛奕", "z"},
		Segment{"信", "n"},
		Segment{"在", "p"},
		Segment{"八八年", "m"},
		Segment{"十月", "t"},
		Segment{"宣佈", "v"},
		Segment{"成立", "v"},
		Segment{"中央", "n"},
		Segment{"政策", "n"},
		Segment{"研究", "vn"},
		Segment{"組", "x"},
	}
	result := chanToArray(seg.Cut(sentence, true))
	if len(cutResult) != len(result) {
		t.Fatal(result)
	}
	for i := range result {
		if result[i] != cutResult[i] {
			t.Fatal(result[i])
		}
	}
}

func TestUserDict(t *testing.T) {
	seg.LoadUserDictionary("../userdict.txt")
	defer seg.LoadDictionary("../dict.txt")
	sentence := "李小福是創新辦主任也是雲計算方面的專家; 什麼是八一雙鹿例如我輸入一個帶「韓玉賞鑒」的標題，在自定義詞庫中也增加了此詞為N類型"

	cutResult := []Segment{
		Segment{"李小福", "nr"},
		Segment{"是", "v"},
		Segment{"創新辦", "i"},
		Segment{"主任", "b"},
		Segment{"也", "d"},
		Segment{"是", "v"},
		Segment{"雲計算", "x"},
		Segment{"方面", "n"},
		Segment{"的", "uj"},
		Segment{"專家", "n"},
		Segment{";", "x"},
		Segment{" ", "x"},
		Segment{"什麼", "r"},
		Segment{"是", "v"},
		Segment{"八一雙鹿", "nz"},
		Segment{"例如", "v"},
		Segment{"我", "r"},
		Segment{"輸入", "v"},
		Segment{"一個", "m"},
		Segment{"帶", "v"},
		Segment{"「", "x"},
		Segment{"韓玉賞鑒", "nz"},
		Segment{"」", "x"},
		Segment{"的", "uj"},
		Segment{"標題", "n"},
		Segment{"，", "x"},
		Segment{"在", "p"},
		Segment{"自定義詞", "n"},
		Segment{"庫中", "nrt"},
		Segment{"也", "d"},
		Segment{"增加", "v"},
		Segment{"了", "ul"},
		Segment{"此", "r"},
		Segment{"詞", "n"},
		Segment{"為", "p"},
		Segment{"N", "eng"},
		Segment{"類型", "n"}}

	result := chanToArray(seg.Cut(sentence, true))
	if len(cutResult) != len(result) {
		t.Fatal(result)
	}
	for i := range result {
		if result[i] != cutResult[i] {
			t.Fatal(result[i])
		}
	}
}

func BenchmarkCutNoHMM(b *testing.B) {
	sentence := "工信處女幹事每月經過下屬科室都要親口交代24口交換機等技術性器件的安裝工作"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		chanToArray(seg.Cut(sentence, false))
	}
}

func BenchmarkCut(b *testing.B) {
	sentence := "工信處女幹事每月經過下屬科室都要親口交代24口交換機等技術性器件的安裝工作"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		chanToArray(seg.Cut(sentence, true))
	}
}
