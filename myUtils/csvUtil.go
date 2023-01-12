package myUtils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// ReadCsv 读取csv文件
func ReadCsv(csvPath string) (records [][]string) {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		fmt.Printf("打开文件出错:%v", err)
		return
	}
	defer func(csvFile *os.File) {
		err = csvFile.Close()
		if err != nil {
			fmt.Printf("文件关闭失败:%v", err)
		}
	}(csvFile)

	reader := csv.NewReader(csvFile)

	records, err = reader.ReadAll()
	if err != nil {
		fmt.Printf("csv读取文件出错:%v", err)
		return
	}
	log.Println("csv 文件读取完成")
	return records
}

// WriteCsv 写入csv文件
func WriteCsv(filePath string, records [][]string) {
	//读写模式、追加写入、不存在就创建文件
	csvFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("文件创建错误,%v", err)
		return
	}

	defer func(csvFile *os.File) {
		err = csvFile.Close()
		if err != nil {
			fmt.Printf("文件关闭失败:%v", err)
		}
	}(csvFile)

	writer := csv.NewWriter(csvFile)
	err = writer.WriteAll(records)
	if err != nil {
		fmt.Printf("文件写入错误,%v", err)
		return
	}
	writer.Flush() //刷新
	fmt.Println("csv 文件写入完成")
}

type Columns struct {
	Column0 []string
	Column1 []string
	Column2 []string
	Column3 []string
}

func GetColumns(path string) *Columns {
	//email,password,uid,token
	var c0 []string
	var c1 []string
	var c2 []string
	var c3 []string
	records := ReadCsv(path)
	for _, record := range records {
		c0 = append(c0, record[0])
		c1 = append(c1, record[1])
		c2 = append(c2, record[2])
		c3 = append(c3, record[3])
	}
	columns := &Columns{
		Column0: c0,
		Column1: c1,
		Column2: c2,
		Column3: c3,
	}
	return columns
}
