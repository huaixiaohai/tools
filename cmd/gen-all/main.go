package main

import (
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"github.com/huaixiaohai/tools/util"
	"io/ioutil"
	"os"
	"strings"
)

const (
	basePath = "./internal/app/"
	servicePath = basePath+"service/"
	schemaPath = basePath+"schema/"
	repoPath = basePath+"dao/"
	modelPath = repoPath+"model/"
)

const (
	service = "service"
	repo = "repo"
	schema = "schema"
	model = "model"
)

func main() {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "-h")
	}

	n := flag.String("n", "", "请输入名称")

	flag.Parse()
	name := *n

	mTemp := getTemplate()

	genFile(servicePath+util.SnakeString(name)+".srv.go", name, mTemp[service])
	genFile(schemaPath+util.SnakeString(name)+".go", name, mTemp[schema])
	genFile(repoPath+util.SnakeString(name)+".repo.go", name, mTemp[repo])
	genFile(modelPath+util.SnakeString(name)+".model.go", name, mTemp[model])

	fmt.Println()
}

func genServiceFile(name string) error {
	fileName := servicePath+ util.SnakeString(name)+".srv.go"
	if checkFileIsExist(fileName) {
		return nil
	}

	buf, err := ioutil.ReadFile("./template/service")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	content := string(buf)
	content = strings.Replace(content, "Test", name,-1)

	err = os.WriteFile(fileName, []byte(content), 0666)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func genFile(fileName, name, template string) error {
	//fileName := servicePath+SnakeString(name)+".srv.go"
	if checkFileIsExist(fileName) {
		return nil
	}

	//buf, err := ioutil.ReadFile(template)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return err
	//}
	//content := string(buf)
	//content = strings.Replace(content, "Test", name,-1)
	template = strings.Replace(template, "Test", name,-1)

	err := os.WriteFile(fileName, []byte(template), 0666)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func getTemplate() map[string]string {
	res := make(map[string]string, 0)
	box := packr.New("box", "./template")
	content, _ := box.FindString(service)
	res[service] = content
	content, _ = box.FindString(schema)
	res[schema] = content
	content, _ = box.FindString(repo)
	res[repo] = content
	content, _ = box.FindString(model)
	res[model] = content
	return res
}