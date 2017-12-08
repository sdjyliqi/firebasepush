package typany

import (
	//"encoding/json"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tools"

	"google.golang.org/api/option"

	firebase "github.com/acoshift/go-firebase-admin"
	"github.com/bitly/go-simplejson"
)

type TyPanySetting struct {
	ProjectID string
	APIKey    string
	SdkPath   string
}

type DevideId struct {
	Id string
}

type DevideIds []DevideId

var typanysetting = &TyPanySetting{
	ProjectID: "test-d84e7",
	APIKey:    "AAAAVCwyRcY:APA91bHoyuvVVB1NFfMAA4UAFnTkuk2kobVhndMQfHHm31IBUk_-nf0fHiasOxc4SSKokPA6o2dODZValcyVZykVGtDqzvjan-aJPwdyPAkLWU27onoUIavw3RziBs6R4GwGor8LLa78",
	SdkPath:   "./typany/typany-firebase.json",
}

type PNotification struct {
	Vername   string   `json:"vername,omitempty"`
	Vercode   string   `json:"vercode,omitempty"`
	Title     string   `json:"title,omitempty"`
	Content   string   `json:"content,omitempty"`
	Targeturl string   `json:"targeturl,omitempty"`
	Home      string   `json:"home,omitempty"`
	Icon      string   `json:"icon,omitempty"`
	Pics      []string `json:"pics,omitempty"`
	Disp      string   `json:"disp,omitempty"`
	Popfrom   string   `json:"popfrom,omitempty"`
	Popto     string   `json:"popto,omitempty"`
}

//定义typany自定义推送的信息Data负载定义
type PushData struct {
	Vername   string   `json:"vername,omitempty"`
	Vercode   string   `json:"vercode,omitempty"`
	Title     string   `json:"title,omitempty"`
	Content   string   `json:"content,omitempty"`
	Targeturl string   `json:"targeturl,omitempty"`
	Home      string   `json:"home,omitempty"`
	Icon      string   `json:"icon,omitempty"`
	Pics      []string `json:"pics,omitempty"`
	Disp      string   `json:"disp,omitempty"`
	Popfrom   string   `json:"popfrom,omitempty"`
	Popto     string   `json:"popto,omitempty"`
}

//
//定义typany自定义推送的信息Data中那些为忽略或者为空字段，那些不能为空。
//如果字典中的某些
var PushDataValid = map[string]interface{}{
	"vername":   map[string]interface{}{"datatype": "string", "allownull": true, "minLen": 1, "maxLen": 32},
	"vercode":   map[string]interface{}{"datatype": "string", "allownull": true, "minLen": 1, "maxLen": 32},
	"title":     map[string]interface{}{"datatype": "string", "allownull": false, "minLen": 1, "maxLen": 128},
	"targeturl": map[string]interface{}{"datatype": "string", "allownull": false, "minLen": 1, "maxLen": 1024},
	"content":   map[string]interface{}{"datatype": "string", "allownull": false, "minLen": 1, "maxLen": 32},
	"home":      map[string]interface{}{"datatype": "string", "allownull": false, "minLen": 1, "maxLen": 32},
	"icon":      map[string]interface{}{"datatype": "string", "allownull": false, "minLen": 8, "maxLen": 128},
	"pics":      map[string]interface{}{"datatype": "list", "allownull": true, "minLen": 0, "maxLen": 4},
	"disp":      map[string]interface{}{"datatype": "string", "allownull": true, "minLen": 0, "maxLen": 16},
	"popfrom":   map[string]interface{}{"datatype": "string", "allownull": false, "minLen": 14, "maxLen": 14},
	"popto":     map[string]interface{}{"datatype": "string", "allownull": false, "minLen": 14, "maxLen": 14},
}

var MsgNode = PNotification{
	Vername:   "1.0",
	Vercode:   "1.0.0",
	Title:     "欢天地喜过大年",
	Content:   "乐声起来了，一场欢天喜地的乡间舞开始，这是世界上所有跳舞里边最好的一场跳舞",
	Targeturl: "http://typany.com",
	Home:      "typany",
	Icon:      "http://d2ezgnxmilyqe4.cloudfront.net/media/official/T40.png",
	Disp:      "small",
	Pics:      []string{"http://d2ezgnxmilyqe4.cloudfront.net/media/official/T144.png"},
	Popfrom:   "20171208120000",
	Popto:     "20171208130000",
}

var jsonStr = `  
{  
   "vername":"1.0",
   "vercode":"1.0.0",
   "title":"欢天地喜过大年",
   "content":"乐声起来了，一场欢天喜地的乡间舞开始，这是世界上所有跳舞里边最好的一场跳舞",
   "targeturl":"http://typany.com",
   "home":"typany",
   "icon":"http://typany.com",
   "disp":"small",
   "pics":["http://typany.com"],
   "popfrom": "20171206120000",
   "popto":   "20171206130000"
	}`

func ChkPushMsgValid(msg string) bool {
	js, err := simplejson.NewJson([]byte(jsonStr))
	fmt.Println(js, err)
	if err != nil {
		return false
	}
	for k, v := range PushDataValid {
		fmt.Println(k, "   :", v)
		if v == true {
			jsVal := js.Get(k)
			fmt.Println("need check:", *jsVal)
		}
	}
	fmt.Println("icon is :", js.Get("icon").MustString())
	fmt.Println("title is:", js.Get("title").MustString())
	return true
}

func ChkReqValid(deviceId, pushRes, pushType string) bool {
	if len(deviceId) < 10 || len(pushRes) < 10 || len(pushType) <= 0 {
		return false
	}
	typeMap := map[string]bool{
		"0": true,
		"1": true,
	}
	_, ok := typeMap[pushType]
	if ok != true {
		return false
	}
	return true
}

func ChkReqValidForPushMulDevices(deviceIds interface{}, pushRes, pushType string) bool {

	if deviceIds == nil || len(pushRes) < 10 || len(pushType) <= 0 {
		return false
	}
	typeMap := map[string]bool{
		"0": true,
		"1": true,
	}
	_, ok := typeMap[pushType]
	if ok != true {
		return false
	}

	return true
}

//
var firApp *firebase.App
var firFCM *firebase.FCM

func init() {
	firApp, _ = firebase.InitializeApp(context.Background(), firebase.AppOptions{
		ProjectID: typanysetting.ProjectID,
		APIKey:    typanysetting.APIKey,
	}, option.WithCredentialsFile(typanysetting.SdkPath))
	firFCM = firApp.FCM()
}

func PushDataToClientByFirebase(deviceId string, pushRes string, pushType string) interface{} {

	resp, err := firFCM.SendToDevice(context.Background(), deviceId,
		firebase.Message{Data: MsgNode,
			TimeToLive: 600,
		})
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return resp
}

//函数功能：推送一个消息给多个设备
//参数列表：deviceIds 为待接收设备的tokenid 列表
//         pushRes   为字符串格式的消息体
//         pushType  推送消息的类型，0表示通知消息类型，1表示自定义消息类型
//返回值：interface{}数据类型数据

func PushDataToClientsByFirebase(deviceIds []string, pushRes string, pushType string) interface{} {
	resp, err := firFCM.SendToDevices(context.Background(), deviceIds,
		firebase.Message{Data: MsgNode,
			TimeToLive: 600,
		})
	if err != nil {
		return err
	}
	return resp
}

func ConvertStringArray(res []interface{}) []string {
	var slist = []string{}
	for _, v := range res {
		slist = append(slist, v.(string))
	}
	return slist

}

func PushdataToDevice(w http.ResponseWriter, req *http.Request) {
	//如果是POST请求，更新一下提交的数据
	//strPushType 0表示通知类消息，即Firebase中的Nodification
	//            1表示自定义消息，表示推送信息中的
	//strDeviceId 为设备firebase，
	req.ParseForm()
	//这样获取到的是列表deviceIdList := req.Form["to"]
	strDeviceId := req.PostFormValue("to")
	strPushData := req.PostFormValue("content")
	strPushType := req.PostFormValue("type")
	bReqValid := ChkReqValid(strDeviceId, strPushData, strPushType)
	if bReqValid != true {
		httpBody := tools.CreateResponseHttpBody("ERR_ARGS_INVALID", "")
		io.WriteString(w, httpBody)
	} else {
		pushResult := PushDataToClientByFirebase(strDeviceId, strPushData, strPushType)
		fmt.Println(strDeviceId, strPushData, strPushType)
		fmt.Println(pushResult)
		if pushResult != nil {
			strPushResult, _ := json.Marshal(pushResult)
			io.WriteString(w, string(strPushResult))
		} else {
			httpBody := tools.CreateResponseHttpBody("ERR_UNKNOW", "")
			io.WriteString(w, httpBody)
		}
	}
}

func PushdataToDevices(w http.ResponseWriter, req *http.Request) {
	//如果是POST请求，更新一下提交的数据
	//strPushType 0表示通知类消息，即Firebase中的Nodification
	//            1表示自定义消息，表示推送信息中的
	//strDeviceId 为设备firebase，
	req.ParseForm()
	bReqValid := true
	//这样获取到的是列表deviceIdList := req.Form["to"]
	strDeviceIds := req.PostFormValue("to")
	strPushData := req.PostFormValue("content")
	strPushType := req.PostFormValue("type")
	js, err := simplejson.NewJson([]byte(strDeviceIds))
	var cliTokenIds []interface{}
	if js != nil {
		cliTokenIds, err = js.Array()
		if err != nil {
			bReqValid = false
		}
	} else {
		bReqValid = false
	}
	if bReqValid == true {
		bReqValid = ChkReqValidForPushMulDevices(cliTokenIds, strPushData, strPushType)
	}

	if bReqValid != true {
		httpBody := tools.CreateResponseHttpBody("ERR_ARGS_INVALID", "")
		io.WriteString(w, httpBody)
	} else {
		tokenIds := ConvertStringArray(cliTokenIds)
		pushResult := PushDataToClientsByFirebase(tokenIds, strPushData, strPushType)
		if pushResult != nil {
			strPushResult, _ := json.Marshal(pushResult)
			io.WriteString(w, string(strPushResult))
		} else {
			httpBody := tools.CreateResponseHttpBody("ERR_UNKNOW", "")
			io.WriteString(w, httpBody)
		}
	}

}
