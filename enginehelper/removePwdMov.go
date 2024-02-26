package enginehelper

import (
	"io/fs"
	"os"
	"path/filepath"
)

func get() {
	//req.SetOutputFile("test.html").Get("https://v.douyin.com/iNdHQB4x/")
	//res, err := http.Get("http://metalsucks.net")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer res.Body.Close()
	//if res.StatusCode != 200 {
	//	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	//}

	//// Load the HTML document
	//doc, err := goquery.NewDocumentFromReader(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Find the review items
	//doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
	//	// For each item found, get the title
	//	title := s.Find("a").Text()
	//	fmt.Printf("Review %d: %s\n", i, title)
	//})
}
func removePwdMov() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("pwd fail", err.Error())
	}
	var waitPaths []string
	if err = filepath.Walk(pwd, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(info.Name()) == ".mov" {
			waitPaths = append(waitPaths, path)
		}
		return nil
	}); err != nil {
		log.Fatal("walk fail", err.Error())
	}
	for _, path := range waitPaths {
		if err = os.Remove(path); err != nil {
			log.Println("remove fail", path)
		}
	}
	log.Println("finished")
}
