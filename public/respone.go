package public

type Result struct {
	Code int
	Data interface{}
	Msg  string
}

// format respone
func (r Result) Format() map[string]interface{} {
	result := make(map[string]interface{})
	result["code"] = r.Code
	result["data"] = r.Data
	result["msg"] = r.Msg
	return result

}
