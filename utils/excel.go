package utils

import (
	"fmt"
	"github.com/marshhu/gin-frame/log"
	"github.com/spf13/cast"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
	"io"
	"net/http"
	"strings"
	"time"
)

// ReadXlsx xlsx解析
func ReadXlsx(fileContent []byte, sheetName string) ([][]string, error) {
	xlFile, err := xlsx.OpenBinary(fileContent)
	if err != nil {
		return nil, err
	}
	sheet := xlFile.Sheet[sheetName]
	if sheet == nil {
		return nil, fmt.Errorf("%s sheet页不存在", sheetName)
	}
	table := make([][]string, len(sheet.Rows))
	for k := 0; k < len(sheet.Rows); k++ {
		table[k] = []string{}
		for i := 0; i < len(sheet.Rows[k].Cells); i++ {
			table[k] = append(table[k], sheet.Rows[k].Cells[i].String()) //注意，有时间列的时候要转
		}
	}
	return table, nil
}

type UploadFunc func(fileName string, data []byte) string

func ReadExcelWithImage(reader io.Reader, sheetName string, imageColIndex int, upload UploadFunc) ([][]string, error) {
	f, err := excelize.OpenReader(reader)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 遍历所有单元格
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 定义二维数组
	var data [][]string
	maxCol := len(rows[0])
	for i, row := range rows {
		var rowData []string
		for _, colCell := range row {
			rowData = append(rowData, colCell)
		}
		if len(rowData) < maxCol { // 补齐空白缺失列
			for j := len(rowData); j < maxCol; j++ {
				rowData = append(rowData, "")
			}
		}

		if i != 0 {
			imageUrl := getAndSavePicture(f, sheetName, i, imageColIndex, upload)
			rowData[imageColIndex] = imageUrl
		}
		data = append(data, rowData)
	}

	// 关闭 Excel 文件
	f.Close()
	return data, nil
}

func getAndSavePicture(f *excelize.File, sheetName string, row, col int, upload UploadFunc) string {
	imageColumnName, _ := excelize.ColumnNumberToName(col + 1)
	imageCellName := imageColumnName + fmt.Sprintf("%d", row+1)
	if pictures, err := f.GetPictures(sheetName, imageCellName); err == nil {
		// 如果单元格中有图片，保存图片并复制链接到当前单元格
		for _, picture := range pictures {
			//imagePath := "\\mockups\\images\\"
			//path := RootDir() + imagePath
			//if err := MkDir(path); err != nil {
			//	return ""
			//}
			//filename := strings.ReplaceAll(NewUUID(), "-", "") + picture.Extension
			//err = os.WriteFile(path+filename, picture.File, 0644)
			//if err != nil {
			//	fmt.Println(err)
			//	return ""
			//}
			//
			//fileAddr := lib.GetStringConf("base.http.file_addr") + "/"
			//url := fileAddr + "images/" + filename
			filename := strings.ReplaceAll(NewUUID(), "-", "") + picture.Extension
			newFileName := upload(filename, picture.File)
			return newFileName
		}
	}
	return ""
}

// ExcelDateFormat 读取Excel里面的时间格式
func ExcelDateFormat(xlsxTime string, format string) time.Time {
	stamp, err := time.ParseInLocation(format, xlsxTime, time.Local)
	if err == nil {
		return stamp
	}
	stamp, err = cast.StringToDateInDefaultLocation(xlsxTime, time.Local)
	if err == nil {
		return stamp
	}
	log.Error("invalid time %s", xlsxTime)
	return stamp
}

func WriteXLSXFileStream(w http.ResponseWriter, fileName string, file []byte) error {
	w.Header().Add("Content-Type",
		"application/application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	w.Header().Add("Content-Transfer-Encoding", "binary")
	_, err := w.Write(file)
	return err
}

func WriteFileStream(w http.ResponseWriter, fileName string, file []byte) error {
	//w.Header().Add("Content-Type",
	//	"application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	w.Header().Add("Content-Type",
		"application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	w.Header().Add("Content-Transfer-Encoding", "binary")
	_, err := w.Write(file)
	return err
}
