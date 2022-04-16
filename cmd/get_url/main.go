package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	url := "http://3962999.cn:32080/iot-platform/assets/iot/device/20220402/239995_20220402094132_33_0.jpg?Expires=1649727690&OSSAccessKeyId=LTAI5tCAN7ygM1LH8RNkAd4P&Signature=uSOXTh8sxWPWSHfqoOubhiAKJ64%3D&x-oss-process=image%2Findexcrop%2Cy_1440%2Ci_0"
	url = ""
	url = strings.Split(url, "?")[0]
	names := strings.Split(url, "/")
	name := names[len(names)-1]
	filename := fmt.Sprintf("%s_%d_%s", now.Format("20060102150405"), now.Nanosecond()/10e6, name)
	fmt.Println(filename)
}
