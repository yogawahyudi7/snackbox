package product

import (
	"github.com/furqonzt99/snackbox/models"
	"gorm.io/gorm"
)

type ProductInterface interface {
	AddProduct(product models.Product) (models.Product, error)
	FindProduct(productId, partnerId int) (models.Product, error)
	DeleteProduct(productId, partnerId int) error
	GetAllProduct() ([]models.Product, error)
	SearchProduct(product string) ([]models.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (p *ProductRepository) AddProduct(product models.Product) (models.Product, error) {

	err := p.db.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *ProductRepository) FindProduct(productId, partnerId int) (models.Product, error) {
	var product models.Product
	err := p.db.Where("id = ? AND partner_id = ?", productId, partnerId).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *ProductRepository) DeleteProduct(productId, partnerId int) error {

	var delete models.Product
	err := p.db.Where("id = ? AND partner_id = ?", productId, partnerId).Delete(&delete).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) GetAllProduct() ([]models.Product, error) {
	var products []models.Product

	err := p.db.Preload("Partner").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) SearchProduct(product string) ([]models.Product, error) {

	var products []models.Product
	err := p.db.Where("title LIKE ?", "%"+product+"%").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
