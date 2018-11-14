package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


func main() {

	image2Base64txt("a.jpg")
	base64txt2Image("a.jpg.base64.txt",".jpg")
	imgBase64 := imgUrl2Base64("http://10.231.20.42:8080/v4/photos/a9-AAABZqm5-TBLTtHqAAAAAg==/data")
	fmt.Println(imgBase64)
}

/**
	本地图片转base64，并写入 path.base64.txt 文件
 */
func image2Base64txt(path string){
	//读原图片
	ff, _ := os.Open(path)
	defer ff.Close()
	sourceBuffer := make([]byte, 500000)
	n, _ := ff.Read(sourceBuffer)
	//base64压缩
	sourceString := base64.StdEncoding.EncodeToString(sourceBuffer[:n])
	//写入临时文件
	name := path + ".base64.txt"
	ioutil.WriteFile(name, []byte(sourceString), 0667)
}

/**
	base64转图片，存为图片
	extension为文件拓展名，例如 .jpg
 */
func base64txt2Image(path,extension string){
	//读取临时文件
	cc, _ := ioutil.ReadFile(path)
	//解压
	dist, _ := base64.StdEncoding.DecodeString(string(cc))
	//写入新文件
	name := path + extension
	f, _ := os.OpenFile(name, os.O_RDWR | os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
}


/**
	图片URL转base64，返回base64字符串
 */
func imgUrl2Base64(url string) string{
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	imgBase64 := base64.StdEncoding.EncodeToString(body)
	fmt.Println(imgBase64)
	return imgBase64
}
