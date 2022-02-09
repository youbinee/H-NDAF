package anlf

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

func requestModelinference(c *gin.Context) {

	reqBody, err := c.GetRawData()
	if err != nil {
		log.Println(err)
	}

	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = "root-nwdaf"
        jsonBody["nfService"] = "mtlf-inference"
        now_t := time.Now().Format("2006-01-02 15:04:05")
        jsonBody["reqTime"] = now_t
        jsonBody["data"] = "training model data"
	json.Unmarshal(reqBody, &jsonBody)
	log.Println(jsonBody)

	jsonStr, _ := json.Marshal(jsonBody)
	transport := &http.Transport{
		ForceAttemptHTTP2: false,
	}
	http := &http.Client{Transport: transport}
	resp, err := http.Post("http://127.0.0.38:5005", "application/json; charset=UTF-8", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println("======== [ROOT-AnLF] START requestModelInference ======")
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)

		c.JSON(200, gin.H{
			"nfService":     "root-nwdaf",
			"reqNFInstance": "anlf-inference",
			"reqTime":       jsonData["reqTime"],
			"data":          jsonData["data"],
		})
		fmt.Println("======== [ROOT-AnLF] END requestModelInference ======")
	}

}
func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nwdaf-anlf")

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
	c.String(http.StatusOK, "Connect Root NWDAF - AnLF")
}

var routes = Routes{
	{
		"Index",
		"POST",
		"/",
		Index,
	},

	{
		"anlf",
		strings.ToUpper("Post"),
		"/:inference",
		requestModelinference,
	},
}
