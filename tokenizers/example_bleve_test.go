package tokenizers_test

import (
	"fmt"
	"log"
	"os"

	"github.com/blevesearch/bleve"
	"jiebago/tokenizers"
)

func Example_beleveSearch() {
	// open a new index
	indexMapping := bleve.NewIndexMapping()

	err := indexMapping.AddCustomTokenizer("jieba",
		map[string]interface{}{
			"file": "../dict.txt",
			"type": "jieba",
		})
	if err != nil {
		log.Fatal(err)
	}

	// create a custom analyzer
	err = indexMapping.AddCustomAnalyzer("jieba",
		map[string]interface{}{
			"type":      "custom",
			"tokenizer": "jieba",
			"token_filters": []string{
				"possessive_en",
				"to_lower",
				"stop_en",
			},
		})

	if err != nil {
		log.Fatal(err)
	}

	indexMapping.DefaultAnalyzer = "jieba"
	cacheDir := "jieba.beleve"
	os.RemoveAll(cacheDir)
	index, err := bleve.New(cacheDir, indexMapping)

	if err != nil {
		log.Fatal(err)
	}

	docs := []struct {
		Title string
		Name  string
	}{
		{
			Title: "Doc 1",
			Name:  "This is the first document we’ve added",
		},
		{
			Title: "Doc 2",
			Name:  "The second one 你 中文測試中文 is even more interesting! 吃水果",
		},
		{
			Title: "Doc 3",
			Name:  "買水果然後來世博園。",
		},
		{
			Title: "Doc 4",
			Name:  "工信處女幹事每月經過下屬科室都要親口交代24口交換機等技術性器件的安裝工作",
		},
		{
			Title: "Doc 5",
			Name:  "咱倆交換一下吧。",
		},
	}
	// index docs
	for _, doc := range docs {
		index.Index(doc.Title, doc)
	}

	// search for some text
	for _, keyword := range []string{"水果世博園", "你", "first", "中文", "交換機", "交換"} {
		query := bleve.NewQueryStringQuery(keyword)
		search := bleve.NewSearchRequest(query)
		search.Highlight = bleve.NewHighlight()
		searchResults, err := index.Search(search)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Result of \"%s\": %d matches:\n", keyword, searchResults.Total)
		for i, hit := range searchResults.Hits {
			rv := fmt.Sprintf("%d. %s, (%f)\n", i+searchResults.Request.From+1, hit.ID, hit.Score)
			for fragmentField, fragments := range hit.Fragments {
				rv += fmt.Sprintf("%s: ", fragmentField)
				for _, fragment := range fragments {
					rv += fmt.Sprintf("%s", fragment)
				}
			}
			fmt.Printf("%s\n", rv)
		}
	}
	// Output:
	// Result of "水果世博園": 2 matches:
	// 1. Doc 3, (1.099550)
	// Name: 買<span class="highlight">水果</span>然後來<span class="highlight">世博</span>園。
	// 2. Doc 2, (0.031941)
	// Name: The second one 你 中文測試中文 is even more interesting! 吃<span class="highlight">水果</span>
	// Result of "你": 1 matches:
	// 1. Doc 2, (0.391161)
	// Name: The second one <span class="highlight">你</span> 中文測試中文 is even more interesting! 吃水果
	// Result of "first": 1 matches:
	// 1. Doc 1, (0.512150)
	// Name: This is the <span class="highlight">first</span> document we’ve added
	// Result of "中文": 1 matches:
	// 1. Doc 2, (0.553186)
	// Name: The second one 你 <span class="highlight">中文</span>測試<span class="highlight">中文</span> is even more interesting! 吃水果
	// Result of "交換機": 2 matches:
	// 1. Doc 4, (0.608495)
	// Name: 工信處女幹事每月經過下屬科室都要親口交代24口<span class="highlight">交換機</span>等技術性器件的安裝工作
	// 2. Doc 5, (0.086700)
	// Name: 咱倆<span class="highlight">交換</span>一下吧。
	// Result of "交換": 2 matches:
	// 1. Doc 5, (0.534158)
	// Name: 咱倆<span class="highlight">交換</span>一下吧。
	// 2. Doc 4, (0.296297)
	// Name: 工信處女幹事每月經過下屬科室都要親口交代24口<span class="highlight">交換</span>機等技術性器件的安裝工作
}
