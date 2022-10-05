package repository

import (
	"awesomeProject/internal/app/ds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetProductByID(id uint) (*ds.Product, error) {
	product := &ds.Product{}
	err := r.db.First(product, "id = ?", id).Error // find product with code D42
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *Repository) CreateProduct(product *ds.Product) error {
	err := r.db.Create(product).Error
	//id := r.db.Last(product, "price = ?", product.)
	return err
}
