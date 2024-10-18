package services

type IService interface {
	HandleRequest(method string, params interface{}, id int) (interface{}, error)
}
