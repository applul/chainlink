package migration1549496047

import (
	"github.com/jinzhu/gorm"
	"github.com/smartcontractkit/chainlink/store/models"
)

type Migration struct{}

func (m Migration) Migrate(tx *gorm.DB) error {
	return tx.AutoMigrate(&models.Key{}).Error
}
