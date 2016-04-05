package cnst
import "time"

const (
	CodeOK = "200"
	CodeRegOK = "201"
	CodeLoginOK = "202"
	CodeUpdateOK = "203"
	CodeQueryDBFail = "301"
	CodeUpdateFail = "302"
	CodeNoEnoughCoins = "501"
	Code = "American Express"
)

const (
	MsgREGOK = "regist success"
	MsgLoginOK = "login success"
	MsgOK = "succ"
	MsgUpdateOK = "update success"
	MsgUpdateFail = "update fail"
	MsgQBFail = "query db occour an error"
)

type Result struct {
	Code       string
	Message    string
	Value      interface{}
	ServerTime int64
}


func Error() *Result {
	ret := &Result{}
	ret.Code= CodeQueryDBFail
	ret.Message = MsgQBFail
	ret.Value=""
	ret.ServerTime = time.Now().Unix()
	return ret
}

func Succ() *Result {
	ret := &Result{}
	ret.Code= CodeOK
	ret.Message = MsgOK
	ret.ServerTime = time.Now().Unix()
	return ret

}
