package utility

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// 查找json
func find_json(root, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return nil
	}
	defer func() {
		cerr := file.Close()
		if cerr == nil {
			err = cerr
		}
	}()
	//创建缓冲写入器，提高写入效率
	writer := bufio.NewWriter(file)
	//延迟刷新缓冲区，确保文件全部写入
	defer func() {
		ferr := writer.Flush()
		if err == nil {
			err = ferr
		}
	}()
	walkFn := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			_, werr := writer.WriteString(path + "\n")
			if werr != nil {
				return werr
			}
		}
		return nil
	}
	return filepath.Walk(root, walkFn)
}

// 读取文件内容
func ReadFile(number int) {
	file, err := os.Open("json_files.txt")
	if err != nil {
		log.Println("打开文件失败：", err)
		return
	}
	defer file.Close()
	var allPaths []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		allPaths = append(allPaths, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时出错:", err)
		return
	}
	for i := 0; ; i++ {
		// 筛选包含 "\621\" 的路径
		var paths621 []string
		for _, path := range allPaths {
			pathString := "\\621\\" + strconv.Itoa(i) + "\\"
			if strings.Contains(path, pathString) {
				paths621 = append(paths621, path)
			}
		}
		// 检查是否找到符合条件的路径
		if len(paths621) == 0 {
			fmt.Println("没有找到属于 621/" + strconv.Itoa(i) + " 文件夹的路径")
			return
		}
		log.Println(len(paths621))
		// 选择 n 个路径（如果不足 n 个，则选择全部）
		selectedPaths := paths621
		// 存储所有 JSON 内容
		var jsonContents []interface{}
		if len(paths621) > number {
			selectedPaths = paths621[:]
		}

		// 遍历选中的路径，读取 JSON 文件内容
		for _, path := range selectedPaths {
			// 读取 JSON 文件
			jsonFile, err := os.Open(path)
			if err != nil {
				fmt.Printf("打开 JSON 文件失败: %s, 错误: %v\n", path, err)
				continue
			}
			defer jsonFile.Close()

			// 读取文件内容
			byteValue, err := io.ReadAll(jsonFile)
			if err != nil {
				fmt.Printf("读取 JSON 文件内容失败: %s, 错误: %v\n", path, err)
				continue
			}

			// 解析 JSON 内容
			var jsonData interface{}
			err = json.Unmarshal(byteValue, &jsonData)
			if err != nil {
				fmt.Printf("解析 JSON 文件失败: %s, 错误: %v\n", path, err)
				continue
			}

			// 将解析后的 JSON 数据添加到切片中
			jsonContents = append(jsonContents, jsonData)
		}

		// 将所有 JSON 内容合并为一个 JSON 数组
		mergedJSON, err := json.Marshal(jsonContents)
		if err != nil {
			fmt.Println("合并 JSON 内容失败:", err)
			return
		}
		jsonFile := strconv.Itoa(i) + ".json"
		// 将合并后的 JSON 写入新文件
		outputFile, err := os.Create(jsonFile)
		if err != nil {
			fmt.Println("创建输出文件失败:", err)
			return
		}
		defer outputFile.Close()
		_, err = outputFile.Write(mergedJSON)

		if err != nil {
			fmt.Println("写入输出文件失败:", err)
			return
		}

		fmt.Printf("成功将选中的 JSON 文件内容合并到 %d.json\n", i)
		cmd := exec.Command(
			"mongoimport",
			"--db", "english",
			"--collection", strconv.Itoa(i),
			"--file", "W:\\go_gf\\find_word_json\\"+jsonFile,
			"--jsonArray",
		)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("导入失败: %v\n输出: %s\n", err, output)
			return
		}
		fmt.Printf("导入成功：\n%s\n", output)
	}

}

/*func main() {
	err := find_json("W:\\go_gf\\proxima\\utility", "json_files.txt")
	if err != nil {
		return
	} else {
		fmt.Println("JSON files list written to json_files.txt")
	}
	//ReadFile(1)
}*/
