package db

import (
	"errors"
	"fmt"

	"github.com/joshm998/drover/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func createPrinter(p *model.Printers) {
	db := GetDBConnection()

	if err := db.Table("printers").Create(&p).Error; err != nil {
		log.Info("failure", model.Printers{}, err)
	}
}

func getAllPrinters() ([]model.Printers, error) {
	var record []model.Printers
	db := GetDBConnection()
	if err := db.Table("printers").Find(&record).Error; err != nil {
		log.Info("failure", []model.Printers{})
		return []model.Printers{}, fmt.Errorf("failed to get printer: %w", err)
	}
	if len(record) == 0 {
		return []model.Printers{}, gorm.ErrRecordNotFound
	}
	return record, nil
}

func getPrinterById(id string) (model.Printers, error) {
	var record []model.Printers
	db := GetDBConnection()
	if err := db.Table("printers").Where("id=?", id).Find(&record).Error; err != nil {
		log.Info("failure", []model.Printers{})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Printers{}, fmt.Errorf("printer with ID %s not found", id)
		}
		return model.Printers{}, fmt.Errorf("failed to get printer: %w", err)
	}
	if len(record) == 0 {
		return model.Printers{}, gorm.ErrRecordNotFound
	}
	return record[0], nil
}

func updatePrinter(id string, p *model.Printers) {
	db := GetDBConnection()
	if err := db.Table("printers").Where("id=?", id).Updates(&p).Error; err != nil {
		log.Info("failure", []model.Printers{})
	}
}

func deletePrinter(id string) (string, error) {
	var p model.Printers
	db := GetDBConnection()
	if err := db.Table("printers").Where("id=?", id).Delete(&p).Error; err != nil {
		log.Info("failure", []model.Printers{})
		return "not able to delete", err
	}

	return "deleted successfully", nil
}
