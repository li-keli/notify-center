// app 配置
package db

import (
	"github.com/globalsign/mgo/bson"
	"notify-center/pkg/constant"
)

type AppEntity struct {
	Id             bson.ObjectId         `bson:"_id"`
	JsjUniqueId    int                   `bson:"jsjUniqueId"`
	PushToken      string                `bson:"pushToken"`
	PlatformType   constant.PlatformType `bson:"platformType"`
	TargetType     constant.TargetType   `bson:"targetType"`
	CreateTime     string                `bson:"createTime"`
	CreateTimeUnix int64                 `bson:"createTimeUnix"`
}

const (
	db         = "notification_db"
	collection = "notification-app-register"
)

func (m *AppEntity) InsertAppEntity(appEntity AppEntity) error {
	return Insert(db, collection, appEntity)
}

func (m *AppEntity) FindAllAppEntity() ([]AppEntity, error) {
	var result []AppEntity
	err := FindAll(db, collection, nil, nil, &result)
	return result, err
}

func (m *AppEntity) FindAppEntityByJsjId(jsjUniqueId int) (AppEntity, error) {
	var result AppEntity
	err := FindOne(db, collection, bson.M{"jsjUniqueId": jsjUniqueId}, nil, &result)
	return result, err
}

func (m *AppEntity) UpdateAppEntity(appEntity AppEntity) error {
	return Update(db, collection, bson.M{"_id": appEntity.Id}, appEntity)
}

func (m *AppEntity) RemoveAppEntity(id string) error {
	return Remove(db, collection, bson.M{"_id": bson.ObjectIdHex(id)})
}
