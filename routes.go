package main

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	r.GET("/mngnt", ManganatoHandler)
	r.GET("/chapters-list/:provider/:titleId", ChaptersListHandler)
	r.GET("/search/:provider/:titleName", SearchHandler)
	r.GET("/get-chapter/:provider/:chapterId", GetChapterHandler)
}
