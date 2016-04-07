#結巴分詞 Go 語言版：Jiebago


[![Build Status](https://travis-ci.org/wangbin/jiebago.png?branch=master)](https://travis-ci.org/wangbin/jiebago) [![GoDoc](https://godoc.org/github.com/wangbin/jiebago?status.svg)](https://godoc.org/github.com/wangbin/jiebago)

[結巴分詞](https://github.com/fxsjy/jieba) 是由 [@fxsjy](https://github.com/fxsjy) 使用 Python 編寫的中文分詞組件，Iiebago 是結巴分詞的 Golang 語言實現。


## 安裝

```
go get github.com/wangbin/jiebago/...
```

## 使用

```
package main

import (
        "fmt"

        "github.com/wangbin/jiebago"
)

var seg jiebago.Segmenter

func init() {
        seg.LoadDictionary("dict.txt")
}

func print(ch <-chan string) {
        for word := range ch {
                fmt.Printf(" %s /", word)
        }
        fmt.Println()
}

func Example() {
        fmt.Print("【全模式】：")
        print(seg.CutAll("我來到北京清華大學"))

        fmt.Print("【精確模式】：")
        print(seg.Cut("我來到北京清華大學", false))

        fmt.Print("【新詞識別】：")
        print(seg.Cut("他來到了網易杭研大廈", true))

        fmt.Print("【搜索引擎模式】：")
        print(seg.CutForSearch("小明碩士畢業於中國科學院計算所，後在日本京都大學深造", true))
}
```
輸出結果：

```
【全模式】： 我 / 來到 / 北京 / 清華 / 清華大學 / 華大 / 大學 /

【精確模式】： 我 / 來到 / 北京 / 清華大學 /

【新詞識別】： 他 / 來到 / 了 / 網易 / 杭研 / 大廈 /

【搜索引擎模式】： 小明 / 碩士 / 畢業 / 於 / 中國 / 科學 / 學院 / 科學院 / 中國科學院 / 計算 / 計算所 / ， / 後 / 在 / 日本 / 京都 / 大學 / 日本京都大學 / 深造 /
```

更多信息請參考[文檔](https://godoc.org/github.com/wangbin/jiebago)。

## 分詞速度

 - 2MB / Second in Full Mode
 - 700KB / Second in Default Mode
 - Test Env: AMD Phenom(tm) II X6 1055T CPU @ 2.8GHz; 《金庸全集》 

## 許可證

MIT: http://wangbin.mit-license.org
