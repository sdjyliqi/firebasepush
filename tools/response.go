package tools

import (
	"encoding/json"
	"strings"
)

type Payload struct {
	Errcode string
	Meg     string
	Body    string
}

var ErrCode = map[string][]string{
	"ERR_USR_NOT_EXIST":         {"0x1001", "Error:The username is not existed."},
	"ERR_USR_LOGIN_WAIT":        {"0x1005", "Error:The client should wait for 30s for the next login."},
	"ERR_USR_IS_EXIST":          {"0x1002", "Error:The username is already existed."},
	"ERR_USR_TOKEN_INVALID":     {"0x1003", "Error:The token of the user is invalid."},
	"ERR_USR_PASSWD_INVALID":    {"0x1004", "Error:The passwd of the user is invalid."},
	"ERR_USR_IS_FORBID":         {"0x1002", "Error:The username is deactivated."},
	"ERR_ARGS_INVALID":          {"0x4000", "Error:The args in the request invalid."},
	"ERR_DBMYSQL_CONN_FAILED":   {"0x3001", "Error:Connect to the mysql failed."},
	"ERR_DBMYSQL_SELECT_FAILED": {"0x3002", "Error:Select the db failed."},
	"ERR_UNKNOW":                {"0xFFFF", "Error:Unexpected error."},
}

func CreateResponse(resIndex string, body interface{}) interface{} {
	var resBody = make(map[string]interface{})
	if strings.ToLower(resIndex) == "succ" {
		resBody["code"] = "0"
		resBody["data"] = body
	} else {
		infoList, err := ErrCode[resIndex]
		if err == true {
			resBody["code"] = infoList[0]
			resBody["msg"] = infoList[1]
			resBody["data"] = body
		} else {
			resBody["code"] = "0xffff"
			resBody["msg"] = "Error:Unknow"
		}
	}
	return resBody
}

func CreateResponseHttpBody(resIndex string, body interface{}) string {
	var resBody = make(map[string]interface{})
	if strings.ToLower(resIndex) == "succ" {
		resBody["code"] = "0"
		resBody["data"] = body
	} else {
		infoList, err := ErrCode[resIndex]
		if err == true {
			resBody["code"] = infoList[0]
			resBody["msg"] = infoList[1]
			resBody["data"] = body
		} else {
			resBody["code"] = "0xffff"
			resBody["msg"] = "Error:Unknow"
		}
	}
	strhttpBody, _ := json.Marshal(resBody)
	return string(strhttpBody)
}
