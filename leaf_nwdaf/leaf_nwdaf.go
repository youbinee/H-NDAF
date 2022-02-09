package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"os"
	"io"
	"github.com/gin-gonic/gin"
)

type Route struct {
        Name string
        Method string
        Pattern string
        HandlerFunc gin.HandlerFunc
}

type Routes []Route

func requestModelTraining(data string) (respData string){
	fmt.Println("")
	fmt.Println("=== [Leaf-NWDAF] START requestModelTraining ===")
	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = "leaf-NWDAF"
	jsonBody["nfService"] = "training"
	now_t := time.Now().Format("2006-01-02 15:04:05")
	jsonBody["reqTime"] = now_t
	jsonBody["data"] = data
	jsonStr, _ := json.Marshal(jsonBody)

	resp, err := http.Post("http://127.0.0.38:8000/nwdaf-mtlf/:training", "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		respData := jsonData["data"].(string)
		fmt.Println("=== [Leaf-NWDAF] END requestModelTraining ===")
		return respData
	}
	return
}

func requestModelInference(reqNfInstanceId string, data_num string) {
	fmt.Println("=== [Leaf-NWDAF] START requestModelInference ===")
	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = reqNfInstanceId
	jsonBody["nfService"] = "inference"
	now_t := time.Now().Format("2006-01-02 15:04:05")
	jsonBody["reqTime"] = now_t
	jsonBody["data"] = data_num
	jsonStr, _ := json.Marshal(jsonBody)

	resp, err := http.Post("http://127.0.0.18:5005", "application/json; charset=UTF-8", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
	}
	fmt.Println("=== [Leaf-NWDAF] END requestModelTraining ===")
}

func main() {
	n := 1
	var selection int
	var data_num string
	for n < 10 {
		fmt.Println("")
		fmt.Println("====== [Leaf-NWDAF] Select Conditions Start ======")
		fmt.Println("Select Conditions - 1) Network Throughput Prediction, 2) To be added")
		fmt.Scanln(&selection)
		fmt.Println("====== [Leaf-NWDAF] Select Conditions End ======")

		if selection == 1 {
			_, err := os.Stat("./saved_models/throughput_prediction.h5")

			fmt.Println("====== [Leaf-NWDAF] Request Training Start =======")
			if err != nil {
				fmt.Println("====== [Leaf-NWDAF] not exist saved model = go training")
				data := requestModelTraining("Request throughput training")

				origin, _ := os.Open(data)
				defer origin.Close()
				copy, _ := os.Create("./saved_models/throughput_prediction.h5")
				defer copy.Close()
				io.Copy(copy, origin)

			} else if err == nil {
				fmt.Println("====== [Leaf-NWDAF] exist saved model = go inference")
				requestModelInference("Test NF-Function", data_num)
			} else {
				fmt.Println("Occur Training Error")
			}
			fmt.Println("====== [Leaf-NWDAF] Request Training End =======")
		} else {
			fmt.Println("Select number 1")
		}
	}
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
}
