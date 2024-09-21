package persistence

import (
	"backend/domain/repository"
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repositories struct {
	Product  repository.ProductRepository
	Category repository.CategoryRepository
	Order    repository.OrderRepository
	Member   repository.MemberRepository
	db       *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	port, _ := strconv.Atoi(DbPort)
	DSN := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		DbUser,
		DbPassword,
		DbHost,
		port,
		DbName,
	)
	db, err := gorm.Open(
		mysql.Open(
			DSN,
		),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	return &Repositories{
		Product: NewProductRepository(db),
		Category: NewCategoryRepository(db),
		Order: NewOrderRepository(db),
		Member: NewMemberRepository(db),
		db: db,
	}, nil
}

// closes the  database connection
func (s *Repositories) Close() error {
	return nil
}
