package msgController

import (
	"four-seasons-of-Kunming/app/apiExpection"
	"four-seasons-of-Kunming/app/daos/msgDao"
	"four-seasons-of-Kunming/app/models"
	"four-seasons-of-Kunming/app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func GetMsg(c *gin.Context) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	messages, err := msgDao.GetMsg()
	if err != nil {
		log.Println("get msg error" + err.Error())
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, messages, "OK", 200)
}

func CreateMsg(c *gin.Context) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	var msg models.MsgForm

	err := c.ShouldBindJSON(&msg)
	if err != nil {
		log.Println("request parameter error" + err.Error())
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}

	err = msgDao.CreateMsg(models.Msg{
		Content:       msg.Content,
		PublisherName: msg.PublisherName,
		CreateAt:      time.Now(),
	})
	if err != nil {
		log.Println("create msg error" + err.Error())
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil, "OK", 200)
}
