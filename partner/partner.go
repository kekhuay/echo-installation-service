package partner

import "github.com/kekhuay/echo-installation-service/model"

// Store PartnerStore interface
type Store interface {
	CreatePartner(*model.Partner) error
	List(offset, limit int64) ([]model.Partner, int64, error)
	UpdatePartner(*model.Partner) error
	GetByID(string) (*model.Partner, error)
	DeletePartner(*model.Partner) error
}
