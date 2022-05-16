package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"simpleTikTok/DB"
	"simpleTikTok/common"
	"simpleTikTok/model"
)

// 与config/application.yml下面的server 内容对应，可以写读文件init这个参数，但是暂时没有必要
const preUrl = "http://127.0.0.1:8090/static/"

var count int = 1

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

func Publish(ctx *gin.Context) {

	// 鉴权
	token := ctx.PostForm("token")
	_, c, err := common.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "token错误",
		})
		return
	}

	// 通过解析token得到用户，后续应该会用到线程池来保证效率，暂时这么写
	db := DB.GetDB()
	var user = model.User{}
	// 通过解析道德的claims.UserId进行查询
	db.Where("id = ?", c.UserId).First(&user)

	// 解析文件
	data, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "没有传数据文件",
		})
		return
	}

	// 接收文件，data的最大值为 32 Mib, gin.Engine，也就是default出来的r可以设置
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", c.UserId, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := ctx.SaveUploadedFile(data, saveFile); err != nil {
		ctx.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 简单的实现非持久化的视频存储，但是文件会持久化进服务器
	count++
	Videos = append(Videos, model.Video{
		Id:            int64(count),
		Author:        user,
		PlayUrl:       preUrl + finalName, // 拼接url
		CoverUrl:      "",                 // 暂时不设置封面
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	})

	ctx.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: Videos, // 这里需要查询到数据库中Author属于该ID的视频，暂时的理解
	})
}
