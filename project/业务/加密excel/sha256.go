package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	glog "github.com/mao888/mao-glog"
	"github.com/tealeg/xlsx"
)

func main() {
	ctx := context.Background()
	// Excel 文件路径
	inputFilePath := "/Users/betta/GolandProjects/my-project/golang-guide/project/业务/加密excel/JPLAndroid端美国地区付费用户gaid信息.xlsx"
	// 加密后的 Excel 文件路径
	outputFilePath := "/Users/betta/GolandProjects/my-project/golang-guide/project/业务/加密excel/JPLAndroid端美国地区付费用户gaid信息-sha256加密.xlsx"

	// 列索引（从0开始）
	columnIndex := 0

	// 读取 Excel 文件
	file, err := xlsx.OpenFile(inputFilePath)
	if err != nil {
		glog.Errorf(ctx, "读取 Excel 文件失败: %s\n", err)
	}

	// 创建一个新的 Excel 文件
	outputFile := xlsx.NewFile()

	// 创建一个工作表
	sheet, err := outputFile.AddSheet("Sheet1")
	if err != nil {
		glog.Errorf(ctx, "创建工作表失败: %s\n", err)
		return
	}

	// 遍历每一行，并加密指定列的数据，然后写入新文件
	for _, row := range file.Sheets[0].Rows {
		// 确保行中有足够的列
		if len(row.Cells) > columnIndex {
			// 获取指定列的单元格数据
			cellData := row.Cells[columnIndex].String()

			// 对数据进行加密
			encryptedData := sha256Encrypt(cellData)

			// 创建新行
			newRow := sheet.AddRow()

			// 创建新单元格，并将加密后的数据写入新文件
			newCell := newRow.AddCell()
			newCell.SetString(encryptedData)
		}
	}

	// 保存新文件
	err = outputFile.Save(outputFilePath)
	if err != nil {
		glog.Errorf(ctx, "保存 Excel 文件失败: %s\n", err)
		return
	}

	glog.Infof(ctx, "加密后的 Excel 文件已保存到: %s\n", outputFilePath)
}

func sha256Encrypt(input string) string {
	// 创建一个 SHA-256 哈希对象
	hasher := sha256.New()

	// 将字符串转换为字节数组并写入哈希对象
	hasher.Write([]byte(input))

	// 计算哈希值
	hashInBytes := hasher.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashString := hex.EncodeToString(hashInBytes)

	return hashString
}
