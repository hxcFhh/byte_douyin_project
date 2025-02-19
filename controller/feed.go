package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpleTikTok/model"
	"time"
)

type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed 暂时只用返回，还没有确定后续逻辑
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: Videos,
		NextTime:  time.Now().Unix(),
	})
}
