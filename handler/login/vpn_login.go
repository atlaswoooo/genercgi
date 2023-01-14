package login

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//Login处理主函数
//Login处理主函数
func Login(context *gin.Context) {
	var retCode int
	//前端包处理

	//调用数据库更新状态
	DB, _ := sql.Open("mysql", "atlaswu:microread321@tcp(database-2.c0obwovzevnt.us-west-2.rds.amazonaws.com:3306)/atlaswu")
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")

	//这个玩意就是给前端的返回值
	retCode = 200
	context.String(retCode, "Login Success.")
}
