package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	urlToScrape := "https://slackmojis.com/"
	request, err := http.NewRequest("GET", urlToScrape, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Not Firefox")
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP Response body. ", err)
	}
	var count int
	var imgSrcs []string
	document.Find("img").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("src")
		if exists {
			count++
			imgSrcs = append(imgSrcs, imgSrc)
		}
	})

	var countDownloader int
	var downloadNames []string
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		if class, exists := element.Attr("class"); exists && class == "downloader" {
			countDownloader++
			downloadName, exists := element.Attr("download")
			if exists {
				downloadNames = append(downloadNames, downloadName)
			}

		}
	})

	if len(imgSrcs) != len(downloadNames) {
		fmt.Println("lengths of images Src and Names are different.")
	}

	downloadNameToImgSrc := mapDownloadNamesToImgSrc(downloadNames, imgSrcs)
	printMap(downloadNameToImgSrc)

	fmt.Printf("Total of %d images found.", len(imgSrcs))
	fmt.Printf("Total of %d downloadNames found.", len(downloadNames))
}

func printMap(givenMap map[string]string) {
	fmt.Printf("Size of the map is %d.\n", len(givenMap))
	for k, v := range givenMap {
		fmt.Printf("DownloadName: {%s}, and ImgSrc: {%s}\n", k, v)
	}
}

func mapDownloadNamesToImgSrc(downloadNames, imgSrcs []string) map[string]string {
	downloadNameToImgSrc := make(map[string]string)
	for _, dName := range downloadNames {
		for _, imgSrc := range imgSrcs {
			if !strings.Contains(imgSrc, dName) {
				continue
			}
			downloadNameToImgSrc[dName] = imgSrc
		}
	}
	return downloadNameToImgSrc
}

func interleaveSlices(imgSrcs, emojiIDNames []string) []string {
	l1 := len(imgSrcs)
	l2 := len(emojiIDNames)

	intermixedSlice := make([]string, l1+l2)
	var cnt int
	var i int
	for ; i < l1 && i < l2; i++ {
		intermixedSlice[cnt] = emojiIDNames[i]
		cnt++
		intermixedSlice[cnt] = imgSrcs[i]
		cnt++
	}
	for ; i < l1 && i < l2; i++ {
		intermixedSlice[cnt] = emojiIDNames[i]
		cnt++
	}

	for ; i < l2; i++ {
		intermixedSlice[cnt] = imgSrcs[i]
		cnt++
	}
	return intermixedSlice
}

func printSlice(slice []string) {
	for _, str := range slice {
		fmt.Println(str)
	}
}
