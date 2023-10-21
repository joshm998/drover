package handler

import (
	"github.com/joshm998/drover/db"
	"github.com/joshm998/drover/model"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	CreateRecordCoreTeam(bl *model.Printers)
	GetRecordSetPost(id string) (model.Printers, error)
	UpdatePrinter(id string, bl *model.Printers)
	DeletePrinter(id string) (string, error)
}

func NewService(sqlDB db.SqlClient) Service {
	return &service{
		sqlDB: sqlDB,
	}
}

type service struct {
	sqlDB db.SqlClient
}

func (s *service) CreateRecordCoreTeam(p *model.Printers) {
	// err :=
	s.sqlDB.CreatePrinterRecord(p)
	// if err != nil {
	// 	log.Info("Failure: mot getting data from Table", err)
	// 	return model.BlogData{}, err
	// }
	// return record, nil
}

func (s *service) GetRecordSetPost(id string) (model.Printers, error) {
	record, err := s.sqlDB.GetPrinterById(id)
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return model.Printers{}, err
	}
	return record, nil
}

func (s *service) UpdatePrinter(id string, bl *model.Printers) {
	// record, err :=
	s.sqlDB.UpdatePrinter(id, bl)
	// if err != nil {
	// 	log.Info("Failure: mot getting data from Table", err)
	// 	return model.BlogData{}, err
	// }
	// return record, nil
}

func (s *service) DeletePrinter(id string) (string, error) {
	record, err := s.sqlDB.DeletePrinter(id)
	if err != nil {
		log.Info("Failure: not getting data from Table", err)
		return "record is not available in the system", err
	}
	return record, nil

}
