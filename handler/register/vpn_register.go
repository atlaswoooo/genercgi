package register

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Login处理主函数
//Login处理主函数
func Register(context *gin.Context) {
	var retCode int
	fmt.Printf("context:%+v\r\n", context)
	//前端包处理

	//调用数据库更新状态

	//这个玩意就是给前端的返回值
	retCode = 200
	context.String(retCode, "Login Success.")
}
