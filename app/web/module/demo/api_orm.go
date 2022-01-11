package demo

import (
	"github.com/reallovelei/ggg/app/model"
	"github.com/reallovelei/ggg/framework/contract"

	//    "github.com/reallovelei/ggg/app/model"
	"github.com/reallovelei/ggg/framework/gin"
	"github.com/reallovelei/ggg/framework/provider/orm"
)

// DemoOrm Orm的路由方法
func (api *DemoApi) DemoOrm(c *gin.Context) {
	logger := c.MustMakeLog()
	logger.Info(c, "request start", nil)

	// 初始化一个orm.DB
	gormService := c.MustMake(contract.ORMKey).(contract.ORMService)
	db, err := gormService.GetDB(orm.WithConfigPath("database.default"))
	if err != nil {
		logger.Error(c, err.Error(), nil)
		c.AbortWithError(50001, err)
		return
	}
	db.WithContext(c)

	// 将User模型创建到数据库中
	err = db.AutoMigrate(&model.PushLogs{})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	//fmt.Println(c, "migrate ok", nil)
	logger.Info(c, "migrate ok", nil)

	// 插入一条数据
	//email := "foo@gmail.com"
	//name := "foo"
	//age := uint8(25)
	//birthday := time.Date(2001, 1, 1, 1, 1, 1, 1, time.Local)
	//user := &User{
	//    Name:         name,
	//    Email:        &email,
	//    Age:          age,
	//    Birthday:     &birthday,
	//    MemberNumber: sql.NullString{},
	//    ActivatedAt:  sql.NullTime{},
	//    CreatedAt:    time.Now(),
	//    UpdatedAt:    time.Now(),
	//}
	//err = db.Create(user).Error
	//logger.Info(c, "insert user", map[string]interface{}{
	//    "id":  user.ID,
	//    "err": err,
	//})
	// 查询一条数据
	queryPL := &model.PushLogs{PushId: "303235701252063232"}

	err = db.First(queryPL).Error
	logger.Info(c, "query pushlog", map[string]interface{}{
		"err":  err,
		"name": queryPL.Name,
	})

	c.JSON(200, "ok")
}
