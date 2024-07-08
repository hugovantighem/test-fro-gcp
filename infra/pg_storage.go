package infra

import (
	"context"
	"myproject/app"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PgStorage struct {
	db *gorm.DB
}

func NewPgStorage(
	db *gorm.DB,
) *PgStorage {
	return &PgStorage{
		db: db,
	}
}

func (x *PgStorage) Search(ctx context.Context, criteria app.SearchCriteria) ([]app.Delegation, error) {
	logrus.Debugf("search with criteria: %s", criteria)
	tx := x.db
	if criteria.Year != nil {
		tx = tx.Where(&Delegation{Year: *criteria.Year})
	}

	tx = tx.Order("ts DESC")

	items := []Delegation{}
	tx = tx.Find(&items)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return FromDataModels(items), nil
}

func (x *PgStorage) GetLast(ctx context.Context) (app.Delegation, error) {
	item := Delegation{}
	tx := x.db.Last(&item)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return app.Delegation{}, app.ErrNotFound
		}
		return app.Delegation{}, tx.Error
	}

	return FromDataModel(item), nil
}

func (x *PgStorage) Save(ctx context.Context, items []app.Delegation) error {
	logrus.Debugf("saving: %+v", items)
	tx := x.db.Create(ToDataModels(items))
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
