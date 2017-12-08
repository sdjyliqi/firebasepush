package main

import (
	"log"
	"net/http"
	"typany"
)

func main() {
	//https://console.firebase.google.com/project/sdjyliqi/settings/cloudmessaging/?hl=zh-cn
	//在上述的管理页面中获取服务器端的秘药。
	//秘钥分新版和旧版本的，如果有新的就用新的，
	// SendToDevice
	//curl -d "to=%5b%22fxSxVRT7UpE%3aAPA91bESK7KrFAjyXgfeYIr2dYYNcOr_zBK7JCmASDoQAgf_Bdx5Gbx-JbJ8EJI8QU9fPKsiqgn06itajKIuS2e1l-oK7iLKG5URC3oscFEZg0TmZigSe2ZSzteS2YGvdnC7lGZnIUiI%22%5d&content=23aaaaaaaaaaaaaaa232&type=1"   "10.129.156.234:8888/pushdevices"
	http.HandleFunc("/push", typany.PushdataToDevice)
	http.HandleFunc("/pushdevices", typany.PushdataToDevices)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
