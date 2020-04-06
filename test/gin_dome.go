package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		g := map[string]string{"type": "his_normal", "sa": "h_1", "q": "命运之夜天之杯2"}
		var ggg []interface{}
		for i := 0; i < 1000; i++ {
			ggg = append(ggg, g)
		}

		c.JSON(200, gin.H{"q": "", "p": false, "bs": "", "csor": "0", "err_no": 0, "errmsg": "",
			"g": ggg})
	})
	r.Run("127.0.0.1:8002") // listen and serve on 0.0.0.0:8080
}
