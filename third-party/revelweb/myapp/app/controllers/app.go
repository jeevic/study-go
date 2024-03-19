package controllers

import (
	"fmt"

	"github.com/jeevic/study-go/third-party/revelweb/myapp/app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/revel/revel"
)

var (
	Db *gorm.DB
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Hello() revel.Result {
	Db, err := gorm.Open("mysql", "root:123456@(172.16.122.131:3306)/liaowang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("db connect error")
	}
	defer Db.Close()
	var user models.Author
	Db.Where("author_id = ?", 1).Find(&user)
	return c.RenderJSON(user)

}
