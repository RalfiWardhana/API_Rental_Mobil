package statusType

import "rental/entity"

type StatusTypeRepository interface {
	CreateStatusType(status_type entity.Status_type) error
	FindAllStatusType() ([]entity.Status_type, error)
	FindByIDStatusType(id string) (entity.Status_type, error)
	UpdateStatusType(id string, status_type entity.Status_type) (error, string)
	DeleteStatusType(id string, status_type entity.Status_type) (error, string)
}
