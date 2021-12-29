package main

import (
	"github.com/gin-gonic/gin"
)

type Provider interface {
	GetChapter(id string) (Chapter, error)
	ListChapters(id string) ([]Chapter, error)
	Search(name string) ([]Title, error)
	ToChapterModel() Chapter
	ToTitleModel() Title
}

type Title struct {
	// path component
	ID           string `json:"id"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	ChaptersUrls string `json:"chapters_urls"`
}

type Chapter struct {
	// path component
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Pages []string `json:"pages"`
}

func GetRightProvider(providerName string) Provider {
	switch providerName {
	case "manganato":
		return &ManganatoProvider{}
	case "mangadex":
		return &MangadexProvider{}
	default:
		return nil
	}
}

func ChaptersListHandler(ctx *gin.Context) {
	var uriData struct {
		Provider string `uri:"provider" binding:"required"`
		TitleID  string `uri:"titleId" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uriData); err != nil {
		ctx.JSON(400, Status{StatusErr, "err: " + err.Error()})
		return
	}

	p := GetRightProvider(uriData.Provider)
	if p == nil {
		ctx.JSON(400, Status{StatusErr, "wrong provider"})
		return
	}

	chapters, err := p.ListChapters(uriData.TitleID)
	if err != nil {
		ctx.JSON(500, Status{StatusErr, "err: " + err.Error()})
		return
	}

	ctx.JSON(200, chapters)
}

func SearchHandler(ctx *gin.Context) {
	var uriData struct {
		Provider  string `uri:"provider" binding:"required"`
		TitleName string `uri:"titleName" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uriData); err != nil {
		ctx.JSON(400, Status{StatusErr, "err: " + err.Error()})
		return
	}

	p := GetRightProvider(uriData.Provider)
	if p == nil {
		ctx.JSON(400, Status{StatusErr, "wrong provider"})
		return
	}

	titles, err := p.Search(uriData.TitleName)
	if err != nil {
		ctx.JSON(500, Status{StatusErr, "err: " + err.Error()})
		return
	}

	ctx.JSON(200, titles)
}

func GetChapterHandler(ctx *gin.Context) {
	var uriData struct {
		Provider  string `uri:"provider" binding:"required"`
		ChapterID string `uri:"chapterId" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uriData); err != nil {
		ctx.JSON(400, Status{StatusErr, "err: " + err.Error()})
		return
	}

	p := GetRightProvider(uriData.Provider)
	if p == nil {
		ctx.JSON(400, Status{StatusErr, "wrong provider"})
		return
	}

	chapter, err := p.GetChapter(uriData.ChapterID)
	if err != nil {
		ctx.JSON(500, Status{StatusErr, "err: " + err.Error()})
		return
	}

	ctx.JSON(200, chapter)
}
