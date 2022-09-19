package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gin-gonic/gin"
	"github.com/kbsonlong/wolf/internal/cmdb/aliyun"
	"github.com/kbsonlong/wolf/internal/monitor"
	"github.com/spf13/viper"
)

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
	results := GetInstance()
	for _, instance := range results.Body.Instances.Instance {
		fmt.Println(instance.HostName)
	}
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
	results := GetInstance()
	fmt.Println(results.Body.Instances)
	for instance := range results.Body.Instances.Instance {
		fmt.Println(instance)
	}
	// fmt.Println(results.Body.Instances)
	c.JSON(http.StatusOK, results.Body.Instances)
}

func GetInstance() *ecs20140526.DescribeInstancesResponse {
	ak := viper.Get("ACCESS_KEY")
	sk := viper.Get("SECRET_KEY")
	client, _err := aliyun.CreateClient(tea.String(fmt.Sprintf("%v", ak)), tea.String(fmt.Sprintf("%v", sk)))
	if _err != nil {
		fmt.Println(_err)
	}

	describeInstancesRequest := &ecs20140526.DescribeInstancesRequest{
		RegionId: tea.String("cn-shenzhen"),
	}
	runtime := &util.RuntimeOptions{}
	results, _err := client.DescribeInstancesWithOptions(describeInstancesRequest, runtime)

	fmt.Println(results.Body.Instances.Instance)
	for index, value := range results.Body.Instances.Instance {
		fmt.Printf("%d\n", index)
		fmt.Printf("%+v\n", value.InstanceId)
	}
	return results
}
