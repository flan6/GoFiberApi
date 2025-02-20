package repository

import (
	"api/db"
	"api/models"

	"gorm.io/gorm"
)

type FornecedorRepo[T any] interface {
	Create(T) error
	ListAll() ([]T, error)
	Find(id uint) (T, error)
	Update(T) error
	Delete(T) error
	Search(string) ([]models.Fornecedor, error)
}

type fornecedor struct {
	db *gorm.DB
}

func NewFornecedorRepo() (FornecedorRepo[models.Fornecedor], error) {
	db, err := db.Connect()
	if err != nil {
		return nil, err
	}
	return &fornecedor{db.Model(&models.Fornecedor{})}, nil
}

func (f *fornecedor) Create(fornecedor models.Fornecedor) error {

	return f.db.Create(&fornecedor).Error

}

func (f *fornecedor) Find(id uint) (models.Fornecedor, error) {

	var fornecedor models.Fornecedor
	if err := f.db.Where("ID = ?", id).Find(&fornecedor).Error; err != nil {
		return models.Fornecedor{}, err
	}

	return fornecedor, nil
}

func (f *fornecedor) ListAll() ([]models.Fornecedor, error) {

	var fornecedores []models.Fornecedor
	if err := f.db.Find(&fornecedores).Error; err != nil {
		return nil, err
	}

	return fornecedores, nil
}

func (f *fornecedor) Search(filter string) ([]models.Fornecedor, error) {

	var fornecedores []models.Fornecedor
	if err := f.db.Find(&fornecedores, "NOME LIKE ?", filter+"%").Error; err != nil {
		return nil, err
	}

	return fornecedores, nil
}

func (f *fornecedor) Update(fornecedor models.Fornecedor) error {

	return f.db.Where("id=?", fornecedor.ID).Updates(fornecedor).Error
}

func (f *fornecedor) Delete(fornecedor models.Fornecedor) error {
	return f.db.Delete(&fornecedor).Error

}
