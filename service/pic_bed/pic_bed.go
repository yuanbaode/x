package pic_bed

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

func PushGitee(owner, repo, token, branch, path2 string) (string, string, string) {
	content := ImagesToBase64(path2)

	var (
		filePath string = "upload/"
	)
	filePath += GetRandomString(10) + path.Ext(path2)
	url := "https://gitee.com/api/v5/repos/" + owner + "/" + repo + "/contents/" + filePath

	args := make(map[string]string)
	args["access_token"] = token
	args["content"] = content
	args["message"] = "upload pic for repo-image-hosting"
	args["branch"] = branch

	jsonBytes, _ := json.Marshal(args)
	data := bytes.NewReader(jsonBytes)
	// 初始化请求与响应
	req, err := http.NewRequest(http.MethodPost, url, data)
	if err != nil {
		log.Fatal(err)
	}

	// 设置请求方法
	req.Header.Set("Content-Type", "application/json")

	c := http.Client{}
	// 发起请求
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// 获取响应的数据实体
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var mapResult map[string]interface{}
	readAll, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(readAll, &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	d := ""
	p := ""
	s := ""

	_, ok := mapResult["content"]

	if ok {
		if mapResult["content"] != nil {
			d = mapResult["content"].(map[string]interface{})["download_url"].(string)
			p = mapResult["content"].(map[string]interface{})["path"].(string)
			s = mapResult["content"].(map[string]interface{})["sha"].(string)
		}
	}

	return d, p, s
}

func ImagesToBase64(str_images string) string {
	f, err := os.Open(str_images)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)
	return encoded
}

func GetRandomString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
