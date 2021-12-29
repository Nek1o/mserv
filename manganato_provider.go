package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
)

const manganatoUri = "https://readmanganato.com"

type ManganatoProvider struct {
}

func (p *ManganatoProvider) GetChapter(id string) (Chapter, error) {
	panic("not implemented") // TODO: Implement
}

func (p *ManganatoProvider) ListChapters(id string) ([]Chapter, error) {
	panic("not implemented") // TODO: Implement
}

func (p *ManganatoProvider) Search(name string) ([]Title, error) {
	panic("not implemented") // TODO: Implement
}

func (p *ManganatoProvider) ToChapterModel() Chapter {
	panic("not implemented") // TODO: Implement
}

func (p *ManganatoProvider) ToTitleModel() Title {
	panic("not implemented") // TODO: Implement
}

func ManganatoHandler(ctx *gin.Context) {
	url := "https://readmanganato.com/manga-ng952689/chapter-1"
	r, err := http.Get(url)
	defer r.Body.Close()
	if err != nil {
		fmt.Printf("err: %v", err)
		ctx.JSON(500, Status{StatusErr, err.Error()})
		return
	}

	node, err := html.Parse(r.Body)
	if err != nil {
		fmt.Printf("err: %v", err)
		ctx.JSON(500, Status{StatusErr, err.Error()})
		return
	}

	imgs := make([]string, 0)
	var f func(*html.Node)
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "img" {
			for _, attr := range node.Attr {
				if attr.Key == "src" {
					imgs = append(imgs, attr.Val)
				}
			}
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)

	for i := range imgs {
		fmt.Printf("%s\n", imgs[i])
	}
	client := http.DefaultClient
	for i := range imgs {
		req, err := http.NewRequest("GET", imgs[i], nil)
		if err != nil {
			fmt.Printf("err: %v", err)
			ctx.JSON(500, Status{StatusErr, err.Error()})
			return
		}

		req.Header.Add("Referer", "https://readmanganato.com/")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("err: %v", err)
			ctx.JSON(500, Status{StatusErr, err.Error()})
			return
		}
		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("err: %v", err)
			ctx.JSON(500, Status{StatusErr, err.Error()})
			return
		}

		err = ioutil.WriteFile(strconv.Itoa(i)+".jpeg", data, fs.ModeAppend)
		if err != nil {
			fmt.Printf("err: %v", err)
			ctx.JSON(500, Status{StatusErr, err.Error()})
			return
		}
	}

	// ctx.JSON(200, Status{StatusOk, ""})
	ctx.JSON(200, Status{StatusOk, ""})
}
