package dongming

import (
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Run() {
	err := auth()
	if err != nil {
		color.Redln(err)
	}
	color.Blueln("开始运行")
	err = buildDB()
	if err != nil {
		color.Redln("asda", err)
	}
}

// buildDB
func buildDB() error {
	dir, _ := os.Getwd()
	source := fmt.Sprintf("%s/%s", dir, "dongming/source/source.db")
	dst := fmt.Sprintf("%s/%s/%s", dir, "static/order/", "ajhfgbawg"+".db")
	err := copyAndRenameFile(source, dst)
	if err != nil {
		return fmt.Errorf("复制文件失败：%s", err.Error())
	}
	return nil
}

// 复制文件并改名函数
func copyAndRenameFile(srcPath string, dstPath string) error {
	// 读取源文件内容
	srcData, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return fmt.Errorf("读取源文件失败：%s", err.Error())
	}

	// 获取源文件所在目录
	//srcDir := filepath.Dir(srcPath)

	// 构建新的目标文件路径
	// 注意：这里我们假设目标文件应该在 static/order/ 目录下
	// 因此，我们需要从源文件的完整路径中移除 static/order/ 前缀
	dstBase := filepath.Base(dstPath)
	dstDir := filepath.Dir(dstPath)
	//srcBase := filepath.Base(srcPath)

	// 确保目标目录存在
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dstDir, 0777); err != nil {
			return fmt.Errorf("创建目标目录失败：%s", err.Error())
		}
	}

	// 构建正确的目标文件路径
	dstPath = filepath.Join(dstDir, dstBase)

	// 写入新文件
	err = ioutil.WriteFile(dstPath, srcData, 0777)
	if err != nil {
		return fmt.Errorf("写入新文件失败：%s", err.Error())
	}
	println("新文件已创建：", dstPath)
	return nil
}
