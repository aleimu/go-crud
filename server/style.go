package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"go-crud/cache"
	"go-crud/model"
	. "go-crud/util"
	"net/http"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var onetime = 3600 * time.Second
var aa = "[{\"ID\":2,\"CreatedAt\":\"2019-08-29T02:31:38Z\",\"UpdatedAt\":\"2019-08-29T02:31:38Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T02:31:38Z\",\"DownTime\":\"2019-08-29T02:31:38Z\"},{\"ID\":3,\"CreatedAt\":\"2019-08-29T02:31:39Z\",\"UpdatedAt\":\"2019-08-29T02:31:39Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T02:31:39Z\",\"DownTime\":\"2019-08-29T02:31:39Z\"},{\"ID\":4,\"CreatedAt\":\"2019-08-29T02:31:40Z\",\"UpdatedAt\":\"2019-08-29T02:31:40Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T02:31:40Z\",\"DownTime\":\"2019-08-29T02:31:40Z\"},{\"ID\":5,\"CreatedAt\":\"2019-08-29T02:31:41Z\",\"UpdatedAt\":\"2019-08-29T02:31:41Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T02:31:41Z\",\"DownTime\":\"2019-08-29T02:31:41Z\"},{\"ID\":6,\"CreatedAt\":\"2019-08-29T02:31:57Z\",\"UpdatedAt\":\"2019-08-29T02:31:57Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T02:31:57Z\",\"DownTime\":\"2019-08-29T02:31:57Z\"},{\"ID\":7,\"CreatedAt\":\"2019-08-29T02:32:18Z\",\"UpdatedAt\":\"2019-08-29T02:32:18Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T02:32:18Z\",\"DownTime\":\"2019-08-29T02:32:18Z\"},{\"ID\":8,\"CreatedAt\":\"2019-08-29T02:34:58Z\",\"UpdatedAt\":\"2019-08-29T02:34:58Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T02:34:58Z\",\"DownTime\":\"2019-08-29T02:34:58Z\"},{\"ID\":9,\"CreatedAt\":\"2019-08-29T02:43:08Z\",\"UpdatedAt\":\"2019-08-29T02:43:08Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T02:43:08Z\",\"DownTime\":\"2019-08-29T02:43:08Z\"},{\"ID\":10,\"CreatedAt\":\"2019-08-29T02:43:34Z\",\"UpdatedAt\":\"2019-08-29T02:43:34Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T02:43:34Z\",\"DownTime\":\"2019-08-29T02:43:34Z\"},{\"ID\":11,\"CreatedAt\":\"2019-08-29T02:43:54Z\",\"UpdatedAt\":\"2019-08-29T02:43:54Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T02:43:54Z\",\"DownTime\":\"2019-08-29T02:43:54Z\"},{\"ID\":12,\"CreatedAt\":\"2019-08-29T02:45:17Z\",\"UpdatedAt\":\"2019-08-29T02:45:17Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T02:45:17Z\",\"DownTime\":\"2019-08-29T02:45:17Z\"},{\"ID\":13,\"CreatedAt\":\"2019-08-29T03:22:01Z\",\"UpdatedAt\":\"2019-08-29T03:22:01Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T03:22:01Z\",\"DownTime\":\"2019-08-29T03:22:01Z\"},{\"ID\":14,\"CreatedAt\":\"2019-08-29T03:22:25Z\",\"UpdatedAt\":\"2019-08-29T03:22:25Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T03:22:25Z\",\"DownTime\":\"2019-08-29T03:22:25Z\"},{\"ID\":15,\"CreatedAt\":\"2019-08-29T03:23:39Z\",\"UpdatedAt\":\"2019-08-29T03:23:39Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T03:23:39Z\",\"DownTime\":\"2019-08-29T03:23:39Z\"},{\"ID\":16,\"CreatedAt\":\"2019-08-29T03:23:40Z\",\"UpdatedAt\":\"2019-08-29T03:23:40Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T03:23:40Z\",\"DownTime\":\"2019-08-29T03:23:40Z\"},{\"ID\":17,\"CreatedAt\":\"2019-08-29T03:25:00Z\",\"UpdatedAt\":\"2019-08-29T03:25:00Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T03:25:00Z\",\"DownTime\":\"2019-08-29T03:25:00Z\"},{\"ID\":18,\"CreatedAt\":\"2019-08-29T03:39:54Z\",\"UpdatedAt\":\"2019-08-29T03:39:54Z\",\"group_id;index\":2,\"image_id\":2,\"ImageUrl\":\"www\",\"ImageName\":\"qqqq\",\"Url\":\"qqwqw\",\"OperId\":123,\"OperName\":\"321\",\"Status\":1,\"Close\":1,\"Mode\":1,\"Frequency\":\"2\",\"Position\":1,\"System\":1,\"Note\":\"\",\"UpTime\":\"2019-08-29T03:39:54Z\",\"DownTime\":\"2019-08-29T03:39:54Z\"}]"

func GetStyle(c *gin.Context) {
	id := c.Query("id")
	style, err := model.GetStyleById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": OK, "data": style})
	}

}

func StyleList(c *gin.Context) {
	//var data map[string]interface{}
	data := make(map[string]interface{})
	id := c.Query("id")
	status := c.Query("status")
	system := c.Query("system")
	fmt.Println("Query:", id, status, system)
	if id != "" {
		data["id"] = id
	}
	if status != "" {
		data["status"] = status
	}
	if system != "" {
		data["system"] = system
	}
	fmt.Println("data:", data)
	style, err := model.GetStyleList(data, 0, 10, " id desc")
	count := model.GetStyleTotal(data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": OK, "data": style, "count": count})
	}

}

func AddStyle(c *gin.Context) {
	var sql model.StyleForm
	err := c.ShouldBind(&sql) // form 必须这样绑定,json 也可以用这个方式校验,看源码可以看出是依据c.Request.Method, c.ContentType()推断出合适的类型
	// err := c.ShouldBindJSON(&newImage)	// 只有json可以
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Sprintln("styleform", sql)
	err = model.AddNewStyle(sql)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "add success!"})
	go FreshRedis(sql.GroupId) // 同步下

}

func UpdateStyle(c *gin.Context) {
	var sql model.StyleForm
	var ok int64 = 1
	err := c.ShouldBind(sql)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Sprintln("sql", sql)
	result := model.DB.Model(&model.Style{}).Where("id = ?", sql.Id).Updates(&model.Style{ImageName: sql.ImageName, Url: sql.Url, GroupId: sql.GroupId, Status: sql.Status}) // model式批量更新
	fmt.Println("result:", result, result.Error, result.RowsAffected)
	if result.RowsAffected != ok || result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"code": ERR, "data": "update style err!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "data": sql})
	go FreshRedis(sql.GroupId) // 同步下
}

func DelStyle(c *gin.Context) {
	id := c.Query("id")
	//status := c.DefaultQuery("status", "1")
	//strconv.Atoi(status)
	//DB.Delete(&email)
	model.DB.Where("id = ?", id).Delete(model.Style{})
	c.JSON(http.StatusOK, gin.H{"code": OK, "msg": "OK"})
}

func SearchStyle(c *gin.Context) {
	// 依据分组查询此分组下的style并加入缓存中,二次查询尝试从缓存中获取数据
	id := c.Query("system")
	groupName := "system:" + id
	rds, err := cache.RedisClient.Get(groupName).Result()
	if err != nil || rds == "" {
		rds1 := GetFreshStyle(id)
		c.JSON(http.StatusOK, gin.H{"code": OK, "source": "db", "data": rds1})
		rds = Map2Str(rds1)
		go AsyncRedis(groupName, rds)
		// FIXME Str2Map不能解析[{map}],应该使用Str2Slice
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": OK, "source": "rds", "data": Str2Slice(rds)})

}

func GetFreshStyle(group_id interface{}) []model.Style {
	data := make(map[string]interface{})
	data["group_id"] = group_id
	fmt.Println("data:", data)
	styles, err := model.GetStyleList(data, 0, 0, " id desc")
	if err != nil {
		fmt.Println(err.Error())
		panic("GetFreshStyle err: " + err.Error())
	}
	return styles
}

func AsyncRedis(groupName string, data interface{}) {
	// 按分组group同步修改缓存中的信息
	_, err := cache.RedisClient.Set(groupName, data, onetime).Result()
	if err != nil {
		fmt.Println(err.Error())
		panic("AsyncRedis err: " + err.Error())
	}
}

func FreshRedis(group_id interface{}) {
	styles := GetFreshStyle(group_id)
	groupName := fmt.Sprintf("system:%s", group_id)
	AsyncRedis(groupName, Map2Str(styles))
}

func Str2Map(jsonData string) (result map[string]interface{}) {
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		fmt.Println(err.Error())
		panic("str2Map err: " + err.Error())
	}
	return result
}

func Str2Slice(jsonData string) (result []interface{}) {
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		fmt.Println(err.Error())
		panic("Str2Slice err: " + err.Error())
	}
	return result
}

func Map2Str(mapData interface{}) (result string) {
	resultByte, err := json.Marshal(mapData)
	result = string(resultByte)
	if err != nil {
		fmt.Println(err.Error())
		panic("map2Str err: " + err.Error())
	}
	return result
}

func Byte2Map(jsonData []byte) (result map[string]interface{}) {
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		fmt.Println(err.Error())
		panic("Byte2Map err: " + err.Error())
	}
	return result
}
func Map2Byte(mapData interface{}) (result []byte) {
	resultByte, err := json.Marshal(mapData)
	if err != nil {
		fmt.Println(err.Error())
		panic("Map2Byte err: " + err.Error())
	}
	return resultByte
}
