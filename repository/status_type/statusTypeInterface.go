package statusType

import "rental/domain"

type StatusTypeRepository interface {
	CreateStatusType(status_type domain.Status_type) error
	FindAllStatusType() ([]domain.Status_type, error)
	FindByIDStatusType(id string) (domain.Status_type, error)
	UpdateStatusType(id string, status_type domain.Status_type) (error, string)
	DeleteStatusType(id string, status_type domain.Status_type) (error, string)
}
