package repositories

import (
	"context"

	"github.com/animesh_30/TicketWave/models"
	"gorm.io/gorm"
)

type EventRepostitory struct {
	db *gorm.DB
}

func(r *EventRepostitory) GetMany(ctx context.Context) ([]*models.Event,error){
	event := []*models.Event{}

	res := r.db.Model(&models.Event{}).Order("updated_at desc").Find(&event)

	if res.Error != nil {
		return nil,res.Error
	}
	return event,nil

}
func(r *EventRepostitory) GetOne(ctx context.Context, eventId uint) (*models.Event,error){
	event := &models.Event{}
	res := r.db.Model(event).Where("id=?",eventId).First(event)

	if res.Error != nil {
		return nil,res.Error
	}
	return event,nil
}

func(r *EventRepostitory) CreateOne(ctx context.Context,event *models.Event) (*models.Event,error){
	
	res := r.db.Model(event).Create(event)

	if res.Error != nil {
		return nil,res.Error
	}
	return event,nil
}

func (r *EventRepostitory) UpdateOne(ctx context.Context , eventId uint , updateData map[string]interface{}) (*models.Event,error){
	event := &models.Event{}

	updateRes := r.db.Model(event).Where ("id=?", eventId).Updates(updateData)

	if updateRes.Error != nil {
		return nil,updateRes.Error
	}

	getRes := r.db.Where("id = ?" , eventId).First(event)
	if getRes.Error != nil{
		return nil,getRes.Error
	}
	return event,nil
}

func (r *EventRepostitory) DeleteOne(ctx context.Context , eventId uint) error{
	res := r.db.Delete(&models.Event{} , eventId)
	
	return res.Error
}

func NewEventRepository(db *gorm.DB) models.EventRepostitory{
	return & EventRepostitory{
		db:db,
	} 
}