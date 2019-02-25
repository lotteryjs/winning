package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/lotteryjs/winning/model"
	"github.com/lotteryjs/winning/service"
	"github.com/lotteryjs/winning/util"
)

func showPlatInfoAction(c *gin.Context) {
	result := util.NewResult()
	defer c.JSON(http.StatusOK, result)

	data := map[string]interface{}{}
	data["version"] = model.Version
	data["database"] = service.Database()
	data["mode"] = model.Conf.RuntimeMode
	data["server"] = model.Conf.Server

	result.Data = data
}
