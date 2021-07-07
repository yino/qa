package interfaces

// response code
const (
	//200 ~ 299 数据正常
	StatusSuccess = 200
	// 4000 ~ 4999
	ErrorUserNotFound = 4000
	// database err 5000~5999
	ErrorRegister = 5001
	// params error 6000~60001
	ErrorParams = 60001

)

var StatusText = map[int]string{
	StatusSuccess:     "success",
	ErrorUserNotFound: "user not fond",
	ErrorRegister:     "register fail",
	ErrorParams:     "params error",
}
