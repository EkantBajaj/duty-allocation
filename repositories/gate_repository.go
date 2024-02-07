package repositories

import (
	"github.com/ekantbajaj/duty-allocation/models"
	"gorm.io/gorm"
)

type GateRepository interface {
	GetGateByID(id uint) (*models.Gate, error)
	CreateGateEntry(gateEntry *models.GateEntry) error
	GetGateCountByGate() ([]models.GateCountWithName, error)
	GetTotalGateCount() (count int, err error)
	UpdateGateCount(gateID uint, count int) error
	GetGates() ([]models.Gate, error)
}

type gateRepository struct {
	db *gorm.DB
}

func NewGateRepository(db *gorm.DB) GateRepository {
	return &gateRepository{
		db: db,
	}
}

func (r *gateRepository) GetGateByID(id uint) (*models.Gate, error) {
	var gate models.Gate
	if err := r.db.First(&gate, id).Error; err != nil {
		return nil, err
	}
	return &gate, nil
}

func (r *gateRepository) CreateGateEntry(gateEntry *models.GateEntry) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(gateEntry).Error; err != nil {
			return err
		}
		
		var gateCount models.GateCount
		if err := tx.Set("gorm:query_option", "FOR UPDATE").Where("active = ? AND gate_id =?", true, gateEntry.GateID).First(&gateCount).Error; err != nil {
			return err
		}
		gateCount.Count += gateEntry.Count
		if err := tx.Save(&gateCount).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *gateRepository) GetGateCountByGate() ([]models.GateCountWithName, error) {
	
	var gateCounts []models.GateCountWithName

	if err := r.db.Table("gate_counts").Select("gate_counts.*, gates.name").Joins("JOIN gates ON gates.id = gate_counts.gate_id").Where("gate_counts.active = ?", true).Find(&gateCounts).Error; err != nil {
        return nil, err
    }
	return gateCounts, nil
}

func (r *gateRepository) GetTotalGateCount() (count int, err error) {
	err = r.db.Raw("SELECT SUM(count) FROM gate_counts WHERE active = ?", true).Scan(&count).Error
    return count, err
}

func (r *gateRepository) UpdateGateCount(gateID uint, count int) error {
	return r.db.Model(&models.GateCount{}).Where("gate_id = ?", gateID).Update("count", count).Error
}


func (r *gateRepository) GetGates() ([]models.Gate, error) {
	var gates []models.Gate
	if err := r.db.Find(&gates).Error; err != nil {
		return nil, err
	}
	return gates, nil
}
