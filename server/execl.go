package server

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"go-crud/model"
	"net/http"
)

func CtrExeclExport(c *gin.Context) {
	ctrs := model.GetCtrs("1=1")
	response := CtrExeclMake(ctrs)
	contentLength := int64(response.Len())
	contentType := "Content-Type:octets/stream"
	extraHeaders := map[string]string{"Content-Disposition": `attachment; filename="广告统计.xlsx"`}
	c.DataFromReader(http.StatusOK, contentLength, contentType, response, extraHeaders)
}

func CtrExeclMake(ctrs *[]model.Ctr) *bytes.Buffer {
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")
	AddCtrHead(f)
	k := 1
	for _, v := range *ctrs {
		k++
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", k), v.StyleId)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", k), v.Show)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", k), v.Click)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", k), v.Crt)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", k), v.CreateDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", k), v.ShowDay)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", k), v.ClickDay)
	}
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	//err := f.SaveAs("./upload/Book1.xlsx")
	buf, err := f.WriteToBuffer()
	if err != nil {
		fmt.Println(err)
	}
	return buf
}

func AddCtrHead(f *excelize.File) {
	f.SetCellValue("Sheet1", fmt.Sprintf("A%d", 1), "编号")
	f.SetCellValue("Sheet1", fmt.Sprintf("B%d", 1), "曝光量")
	f.SetCellValue("Sheet1", fmt.Sprintf("C%d", 1), "点击量")
	f.SetCellValue("Sheet1", fmt.Sprintf("D%d", 1), "比率")
	f.SetCellValue("Sheet1", fmt.Sprintf("E%d", 1), "日期")
	f.SetCellValue("Sheet1", fmt.Sprintf("F%d", 1), "曝光时段详情")
	f.SetCellValue("Sheet1", fmt.Sprintf("G%d", 1), "点击时段详情")
}
