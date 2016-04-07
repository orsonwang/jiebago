package jiebago_test

import (
	"fmt"

	"jiebago"
)

func Example() {
	var seg jiebago.Segmenter
	seg.LoadDictionary("dict.txt")

	print := func(ch <-chan string) {
		for word := range ch {
			fmt.Printf(" %s /", word)
		}
		fmt.Println()
	}

	fmt.Print("【全模式】：")
	print(seg.CutAll("我來到北京清華大學"))

	fmt.Print("【精確模式】：")
	print(seg.Cut("我來到北京清華大學", false))

	fmt.Print("【新詞識別】：")
	print(seg.Cut("他來到了網易杭研大廈", true))

	fmt.Print("【搜索引擎模式】：")
	print(seg.CutForSearch("小明碩士畢業於中國科學院計算所，後在日本京都大學深造", true))
	// Output:
	// 【全模式】： 我 / 來到 / 北京 / 清華 / 清華大學 / 華大 / 大學 /
	// 【精確模式】： 我 / 來到 / 北京 / 清華大學 /
	// 【新詞識別】： 他 / 來到 / 了 / 網易 / 杭研 / 大廈 /
	// 【搜索引擎模式】： 小明 / 碩士 / 畢業 / 於 / 中國 / 科學 / 學院 / 科學院 / 中國科學院 / 計算 / 計算所 / ， / 後 / 在 / 日本 / 京都 / 大學 / 日本京都大學 / 深造 /
}

func Example_suggestFrequency() {
	var seg jiebago.Segmenter
	seg.LoadDictionary("dict.txt")

	print := func(ch <-chan string) {
		for word := range ch {
			fmt.Printf(" %s /", word)
		}
		fmt.Println()
	}
	sentence := "超敏C反應蛋白是什麼？"
	fmt.Print("Before:")
	print(seg.Cut(sentence, false))
	word := "超敏C反應蛋白"
	oldFrequency, _ := seg.Frequency(word)
	frequency := seg.SuggestFrequency(word)
	fmt.Printf("%s current frequency: %f, suggest: %f.\n", word, oldFrequency, frequency)
	seg.AddWord(word, frequency)
	fmt.Print("After:")
	print(seg.Cut(sentence, false))

	sentence = "如果放到post中將出錯"
	fmt.Print("Before:")
	print(seg.Cut(sentence, false))
	word = "中將"
	oldFrequency, _ = seg.Frequency(word)
	frequency = seg.SuggestFrequency("中", "將")
	fmt.Printf("%s current frequency: %f, suggest: %f.\n", word, oldFrequency, frequency)
	seg.AddWord(word, frequency)
	fmt.Print("After:")
	print(seg.Cut(sentence, false))

	sentence = "今天天氣不錯"
	fmt.Print("Before:")
	print(seg.Cut(sentence, false))
	word = "今天天氣"
	oldFrequency, _ = seg.Frequency(word)
	frequency = seg.SuggestFrequency("今天", "天氣")
	fmt.Printf("%s current frequency: %f, suggest: %f.\n", word, oldFrequency, frequency)
	seg.AddWord(word, frequency)
	fmt.Print("After:")
	print(seg.Cut(sentence, false))
	// Output:
	// Before: 超敏 / C / 反應 / 蛋白 / 是 / 什麼 / ？ /
	// 超敏C反應蛋白 current frequency: 0.000000, suggest: 1.000000.
	// After: 超敏C反應蛋白 / 是 / 什麼 / ？ /
	// Before: 如果 / 放到 / post / 中將 / 出錯 /
	// 中將 current frequency: 763.000000, suggest: 494.000000.
	// After: 如果 / 放到 / post / 中 / 將 / 出錯 /
	// Before: 今天天氣 / 不錯 /
	// 今天天氣 current frequency: 3.000000, suggest: 0.000000.
	// After: 今天 / 天氣 / 不錯 /
}

func Example_loadUserDictionary() {
	var seg jiebago.Segmenter
	seg.LoadDictionary("dict.txt")

	print := func(ch <-chan string) {
		for word := range ch {
			fmt.Printf(" %s /", word)
		}
		fmt.Println()
	}
	sentence := "李小福是創新辦主任也是雲計算方面的專家"
	fmt.Print("Before:")
	print(seg.Cut(sentence, true))

	seg.LoadUserDictionary("userdict.txt")

	fmt.Print("After:")
	print(seg.Cut(sentence, true))
	// Output:
	// Before: 李小福 / 是 / 創新 / 辦 / 主任 / 也 / 是 / 雲 / 計算 / 方面 / 的 / 專家 /
	// After: 李小福 / 是 / 創新辦 / 主任 / 也 / 是 / 雲計算 / 方面 / 的 / 專家 /
}
