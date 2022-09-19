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
	fmt.Println(requestBody.Content.InstanceID)
	for _, instance := range results.Body.Instances.Instance {
		fmt.Println(instance.HostName)
	}
	// fmt.Println(results.Body.Instances)
	c.JSON(http.StatusOK, results.Body.Instances)
}

func GetInstance() *ecs20140526.DescribeInstancesResponse {
	client, _err := aliyun.CreateClient(tea.String("LTAI5tB7P844wg5Ks3AE5xQe"), tea.String("P2ao5wPJYQk2sDknjPfqGiNZZmx7Ts"))
	if _err != nil {
		fmt.Println(_err)
	}

	describeInstancesRequest := &ecs20140526.DescribeInstancesRequest{
		RegionId: tea.String("cn-shenzhen"),
	}
	runtime := &util.RuntimeOptions{}
	results, _err := client.DescribeInstancesWithOptions(describeInstancesRequest, runtime)
	fmt.Println(results.Body.Instances.Instance[0].InstanceId)
	return results
}
