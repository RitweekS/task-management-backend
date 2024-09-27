
package common

type Response struct {
	ResponseMessage string `json:"response_message"`
	ResponseCode int `json:"response_code"`
	ResponseData interface{} `json:"response_data"`

}