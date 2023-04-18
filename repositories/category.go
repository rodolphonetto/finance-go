package repositories

import (
	"fmt"

	"gorm.io/gorm"
	"testes.com/packages/storage"
	"testes.com/packages/types"
)

func CreateCategory(name string) (*gorm.DB, error) {
	if err := storage.DB.AutoMigrate(&types.Category{}); err != nil {
		return nil, err
	}

	id := storage.DB.Create(&types.Category{Name: name})

	var category types.Category
	storage.DB.Where("name = ?", "Bares").First(&category)

	fmt.Println("1 2 3")
	fmt.Println(category)
	return id, nil

}
