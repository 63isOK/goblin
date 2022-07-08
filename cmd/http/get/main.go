package main

import (
	"io"
	"net/http"
)

func main() {
	r, err := http.Get("https://linguipark-bucket.oss-cn-guangzhou.aliyuncs.com/imgs%2F68515f7a-e211efe2%2F20220617%2F1655430729_377529_68515f7a-e211efe2_0_110582531_T6_full.jpg?Expires=1655437937\u0026OSSAccessKeyId=LTAI5tQRN7YgABwHN9Sj5Rec\u0026Signature=WhuY10HQrXS04gD2Oa1vOs90hQo%3D")
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	println(len(b))
}
