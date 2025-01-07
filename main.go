package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"wget/pkg/download"
	"wget/pkg/mirror"
	ratelimit "wget/pkg/rate-limit"
	"wget/pkg/utils"
)

func main() {
	flags, urls, err := utils.ParseFlag()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Usage: ./wget [options...] [<url1> ...]")
		os.Exit(0)
	}

	if strings.HasPrefix(flags["P"], "/") {
		flags["P"] = "." + flags["P"]
	}

	err = utils.ValidateFlag(flags)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	utils.SetOutputDir(flags["P"])

	if flags["i"] != "" {
		err = download.DownloadFileAsync(flags["i"], flags["P"])
		if err != nil {
			fmt.Println(err)
		}
	} else {
		for _, url := range urls {
			infos, err := utils.GetFileInfo(url)
			if err != nil {
				fmt.Println("IMPOSIBLE TO GET FILE INFO")
				continue
			}
			if flags["B"] == "true" {
				err = download.DownloadBackground(url, flags["P"])
			} else if flags["rate-limit"] != "" {
				err = ratelimit.DownloadRateLimit(url, flags["P"], flags["rate-limit"])
			} else if flags["mirror"] == "true" {
				mirror.RootURL = strings.TrimSuffix(url, "/")
				err = mirror.MirrorSite(url, flags["P"])
			} else {
				fmt.Println("start at", utils.FormatDate(time.Now()))
				fmt.Printf("sending request, awaiting response... ")
				fmt.Printf("status %v\n", utils.HttpCode(infos.StatusCode))
				fmt.Printf("content size: %s [~%s]\n", infos.ContentLength, utils.ToMega(infos.ContentLength))
				// var (
				// 	name string
				// 	path string
				// )
				if flags["O"] != "" {
					if !strings.HasPrefix(flags["P"], "./") {
						fmt.Printf("saving file to: %s\n", filepath.Join(filepath.Clean(flags["P"]), flags["O"]))
					} else {
						fmt.Printf("saving file to: ./%s\n", filepath.Join(filepath.Clean(flags["P"]), flags["O"]))
					}
					// fmt.Printf("saving file to: %s\n", filepath.Join(filepath.Clean(flags["P"]), flags["O"]))
					//fmt.Printf("saving file to: %s\n", filepath.Clean(flags["P"]))
				} else {
					if !strings.HasPrefix(flags["P"], "./") {
						fmt.Printf("saving file to: %s\n", filepath.Join(filepath.Clean(flags["P"]), infos.FileName))
					} else {
						fmt.Printf("saving file to: ./%s\n", filepath.Join(filepath.Clean(flags["P"]), infos.FileName))
					}
					//fmt.Printf("saving file to: %s\n", filepath.Base(flags["P"]))
				}
				err = download.Download(false, url, flags["P"])
				if err == nil {
					fmt.Printf("\n\nDownloaded [%s]\nfinished at %v\n", url, utils.FormatDate(time.Now()))
				}
			}

			if err != nil {
				fmt.Println("\nError:", err)
			}

			if flags["O"] != "" {
				err = utils.RenameFile(infos.FileName, flags["O"])
				if err != nil {
					fmt.Println("Error:", err)
				}
			}
		}
	}
}
