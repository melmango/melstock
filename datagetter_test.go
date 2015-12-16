package melstock
import (
	"testing"
	"fmt"
	"github.com/melman-go/aliopengo/util"
)



func TestTokenValidate(t *testing.T) {
	code := "sz002378"
	//	realTime := GetRealtime(code)
	//	pk:= GetPK(code)
	//	funFlow := GetFundFlow(code)
	//	info := GetInfo(code)
	//	fmt.Println(util.JsonEncodeS(realTime))
	//	fmt.Println(util.JsonEncodeS(pk))
	//	fmt.Println(util.JsonEncodeS(funFlow))
	//	fmt.Println(util.JsonEncodeS(info))
//	list := GetDaily(code, 15)
	list := GetWeekly(code)
	fmt.Println(util.JsonEncodeS(list))
}