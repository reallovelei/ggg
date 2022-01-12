package dao

import (
	"github.com/reallovelei/ggg/app/model"
	"gorm.io/gorm"
)

const (
	_selectByUpdate = "SELECT push_id, updated_at FROM push_logs where updated_at < ?"
	_deleteByUpdate = "DELETE FROM push_logs where updated_at < ?"
)

func (d *Dao) GetPushLogs(db *gorm.DB, update string) (list []model.PushLogs) {
	db.Raw(_selectByUpdate, update).Scan(&list)
	return
}
