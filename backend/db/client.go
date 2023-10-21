package db

import (
	"github.com/joshm998/drover/model"
)

type SqlClient interface {
	CreatePrinterRecord(bl *model.Printers)
	GetPrinters() ([]model.Printers, error)
	GetPrinterById(string) (model.Printers, error)
	UpdatePrinter(id string, bl *model.Printers)
	DeletePrinter(id string) (string, error)
}

func NewClient(config *Config) SqlClient {
	return &sqlClient{
		config: config,
	}
}

type Config struct {
	DBConnection string
}

type sqlClient struct {
	config *Config
}

func (c *sqlClient) CreatePrinterRecord(p *model.Printers) {
	createPrinter(p)
}

func (c *sqlClient) GetPrinters() ([]model.Printers, error) {
	return getAllPrinters()
}

func (c *sqlClient) GetPrinterById(id string) (model.Printers, error) {
	return getPrinterById(id)
}

func (c *sqlClient) UpdatePrinter(id string, p *model.Printers) {
	updatePrinter(id, p)
}

func (c *sqlClient) DeletePrinter(id string) (string, error) {
	return deletePrinter(id)
}
