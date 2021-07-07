package controller

const (
	Message             = "Message"
	Data                = "Data"
	Inputrarity         = 400
	InputUnlockarena    = 401
	InputCvc            = 402
	NotFoundSoldier     = 403
	NotFoundInformation = 404
	InputId             = 405
	Success             = 200
)

var statusText = map[int]string{
	Inputrarity:         "请输入Rarity",
	InputUnlockarena:    "请输入Unlockarena",
	InputCvc:            "请输入CVc",
	NotFoundSoldier:     "未找到士兵",
	NotFoundInformation: "未找到相关信息",
	Success:             "成功",
	InputId:             "请输入士兵id",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}
