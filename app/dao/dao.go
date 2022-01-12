package dao

import (
	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

//
//func New() (d *Dao) {
//    logger := c.MustMakeLog()
//    logger.Info(c, "request start", nil)
//
//    // 初始化一个orm.DB
//    gormService := c.MustMake(contract.ORMKey).(contract.ORMService)
//    db, err := gormService.GetDB(orm.WithConfigPath("database.default"))
//    if err != nil {
//        logger.Error(c, err.Error(), nil)
//        c.AbortWithError(50001, err)
//        return
//    }
//    db.WithContext(c)
//
//    return &Dao{db:db}
//}
