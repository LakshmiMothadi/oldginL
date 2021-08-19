package example

import (
	"github.com/gin-gonic/gin"
)
type api struct {
	method string `json:"method"`
	//signal bool `json:"signal"`
}

func getFunction(c *gin.Context){
	apiStatus := api{
		method : "method",
		//signal : signal,
	}
	c.JSON(200, apiStatus)

}
func postFunction(c *gin.Context){


}
func main(){
	router :=gin.Default()
	router.GET("/start", getFunction)
	router.POST("/writing", postFunction)
	router.Run(":8088")

}
