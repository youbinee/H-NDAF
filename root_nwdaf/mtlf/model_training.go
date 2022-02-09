package mtlf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

type ResponseInfo struct {
	nfService       string `json:"nfService"`
	reqNFInstanceID string `json:"reqNFInstance"`
	reqTime         string `json:"reqTime"`
	data            string `json:"data"`
}

func requestModelTraining(c *gin.Context) {
	reqBody := map[string]interface{}{}
	log.Print("====== [ROOT-MTLF] START requestModelTraining ======")

	c.JSON(http.StatusOK, gin.H{
		"nfService":     "root-nwdaf",
		"reqNFInstance": "root-mtlf",
		"reqTime":       reqBody["reqTime"],
		"data":          "../root_nwdaf/pythonmodule/saved_models/throughput_prediction.h5",
	})

	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = "root-nwdaf"
	jsonBody["nfService"] = "mtlf-training"
	now_t := time.Now().Format("2006-01-02 15:04:05")
	jsonBody["reqTime"] = now_t
	jsonBody["data"] = "Training success"
	jsonStr, _ := json.Marshal(jsonBody)

	transport := &http.Transport{
		ForceAttemptHTTP2: false,
	}
	http := &http.Client{Transport: transport}
	resp, err := http.Post("http://127.0.0.38:5005", "application/json; charset=UTF-8", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println("======== [ROOT-MTLF] END requestModelTraining ======")
		fmt.Println()
	}
	return
}
func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nwdaf-mtlf")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		case "PATCH":
			group.PATCH(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Connect Root NWDAF - MTLF")
}

var routes = Routes{
	{
		"Index",
		"POST",
		"/",
		Index,
	},

	{
		"mtlf",
		strings.ToUpper("Post"),
		"/:training",
		requestModelTraining,
	},
}
