package download

import (
	"fmt"
	"sync"
	"wget/pkg/utils"
)

func DownloadFileAsync(file string, outputPath string) error {
	urls, err := utils.ReadFileContent(file)
	if err != nil {
		panic("ERROR ON FILE CONTENT READING")
	}

	var (
		wg     sync.WaitGroup
		result []string
	)

	fmt.Printf("content size: %v\n", utils.Display(utils.Sizes(urls)))

	for _, url := range urls {
		wg.Add(1)

		infos, err := utils.GetFileInfo(url)
		if err != nil {
			fmt.Println(err)
			continue
		}

		go func() {
			defer wg.Done()

			err := Download(true, url, outputPath)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("finished " + infos.FileName)
			result = append(result, url)
		}()
	}

	wg.Wait()

	fmt.Printf("\nDownload finished: %v\n", result)

	return nil
}
