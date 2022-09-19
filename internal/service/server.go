package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kbsonlong/wolf/internal/monitor"
	"github.com/spf13/viper"
)

func Start() {
	r := gin.Default()
	r.GET("/", Ping)
	r.POST("/api/v1/events", SpotEvent)
	r.Run(fmt.Sprintf(":%d", viper.Get("PORT")))
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an accounts
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Ping successful"})
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func SpotEvent(c *gin.Context) {
	var requestBody monitor.Event
	BodyAsArray, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(BodyAsArray, &requestBody); err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, &requestBody)
}
