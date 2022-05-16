package controller

import (
	"github.com/gin-gonic/gin"
	"simpleTikTok/model"
)

type CommentListResponse struct {
	model.Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
/*func CommentAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
*/

// CommentAction 评论逻辑，后续补充
func CommentAction(ctx *gin.Context) {

}
