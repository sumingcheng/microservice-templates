package config

type CustomError struct {
	Name  string `json:"name"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Error string `json:"error"`
}

func (ce *CustomError) BadParameter(err error) *CustomError {
	ce.Name = "BadParameter"
	ce.Code = 400
	ce.Msg = "参数传递不合法"
	ce.Error = err.Error()
	return ce
}
