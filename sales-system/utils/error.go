package utils

type CustomError struct {
	Name  string `json:"name"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Error string `json:"error"`
}

func (ce *CustomError) Success() *CustomError {
	ce.Name = "success"
	ce.Code = 0
	ce.Msg = "ok"
	ce.Error = ""
	return ce
}

func (ce *CustomError) BadParameter(err error) *CustomError {
	ce.Name = "BadParameter"
	ce.Code = 1001
	ce.Msg = "参数传递不合法"
	ce.Error = TranslateErrors(err)

	return ce
}

func (ce *CustomError) CreateDataFailed(err error) *CustomError {
	ce.Name = "CREATE_DATA_FAILED"
	ce.Code = 1002
	ce.Msg = "数据创建失败"
	ce.Error = err.Error()
	return ce
}

func (ce *CustomError) QueryDataFailed(err error) *CustomError {
	ce.Name = "QUERY_DATA_FAILED"
	ce.Code = 1003
	ce.Msg = "数据查询失败"
	ce.Error = err.Error()
	return ce
}

func (ce *CustomError) UpdateDataFailed(err error) *CustomError {
	ce.Name = "UPDATE_DATA_FAILED"
	ce.Code = 1004
	ce.Msg = "数据更新失败"
	ce.Error = err.Error()
	return ce
}

func (ce *CustomError) DeleteDataFailed(err error) *CustomError {
	ce.Name = "DELETE_DATA_FAILED"
	ce.Code = 1005
	ce.Msg = "数据删除失败"
	ce.Error = err.Error()
	return ce
}

func (ce *CustomError) InvalidId(err error) *CustomError {
	ce.Name = "INVALID_ID"
	ce.Code = 1006
	ce.Msg = "id 不存在"
	ce.Error = "nil"
	return ce
}
