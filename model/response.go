package model

//{
//    "status":"fail",
//    "reason":"签名验证失败!"
//}
type ResponseFail struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

//{
//    "status":"succ",
//    "info":"XXXX"
//}
type ResponseSucc struct {
	Status string      `json:"status"`
	Info   interface{} `json:"info"`
}
