package routes

import (
	"fmt"
	CO "golang_login/config"

	"github.com/kataras/iris"
)

// APILogin function
func APILogin(ctx iris.Context) {

	type Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type Member struct {
		ID 		 string `json:"id"`
		Username string `json:"username"`
	}

	var TextLogin Login

	ctx.ReadJSON(&TextLogin)

	db := CO.DB()

	strSQL := "SELECT id, username FROM `members` WHERE username = ? AND password = ?"
	stmt, _ := db.Prepare(strSQL)
	rows, err := stmt.Query(TextLogin.Username, TextLogin.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer stmt.Close()
	
	var m Member

	rows.Next()
	err = rows.Scan(&m.ID, &m.Username)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx.JSON(m)
}


