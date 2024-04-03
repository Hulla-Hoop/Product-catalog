package catalog

import (
	"fmt"
	"testing"
	db "testinhousead/internal/DB"
	"testinhousead/internal/logger"
	"testinhousead/internal/model"
)

type Mock struct {
}

func (m *Mock) CreateCategory(reqId string, name string) (*model.Category, error) {

	return nil, fmt.Errorf("")
}
func (m *Mock) DeleteCategory(reqId string, id int) (*model.Category, error) {
	return nil, fmt.Errorf("")
}
func (m *Mock) UpdateCategory(reqId string, id int, name string) (*model.Category, error) {
	return nil, fmt.Errorf("")
}
func (m *Mock) AllCategories(reqID string) ([]model.Category, error) {
	return nil, fmt.Errorf("")
}
func (m *Mock) GoodsOnCateory(reqID string, category string) ([]model.Product, error) {
	return nil, fmt.Errorf("")
}
func (m *Mock) CreateGoods(reqId string, name string) (*model.Goods, error) {
	return nil, fmt.Errorf("")
}
func (m *Mock) DeleteGoods(reqId string, id int) (*model.Goods, error) {
	return nil, fmt.Errorf("")
}
func (m *Mock) UpdateGoods(reqId string, id int, name string) (*model.Goods, error) {
	return nil, fmt.Errorf("")
}
func (m *Mock) CreateRelation(reqId string, goods_id int, category_id int) (*model.Product, error) {
	return nil, fmt.Errorf("")
}
func (m *Mock) Close() error {
	return nil
}

func Test_catalog_CreateCategory(t *testing.T) {
	type fields struct {
		logger *logger.Logger
		db     db.DB
	}
	l := logger.New()
	db := new(Mock)

	type args struct {
		reqId string
		name  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "One",
			fields: fields{logger: l,
				db: db},
			args:    args{reqId: "", name: "Fresh"},
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &catalog{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			_, err := s.CreateCategory(tt.args.reqId, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("catalog.CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_catalog_DeleteCategory(t *testing.T) {
	type fields struct {
		logger *logger.Logger
		db     db.DB
	}

	l := logger.New()
	db := new(Mock)

	type args struct {
		reqId string
		id    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "One",
			fields: fields{logger: l,
				db: db},
			args:    args{reqId: "", id: "1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &catalog{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			_, err := s.DeleteCategory(tt.args.reqId, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("catalog.DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_catalog_UpdateCategory(t *testing.T) {
	type fields struct {
		logger *logger.Logger
		db     db.DB
	}

	l := logger.New()
	db := new(Mock)

	type args struct {
		reqId string
		id    string
		name  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "One",
			fields: fields{logger: l,
				db: db},
			args:    args{reqId: "", id: "1", name: "Fresh"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &catalog{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			_, err := s.UpdateCategory(tt.args.reqId, tt.args.id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("catalog.DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_catalog_CreateGoods(t *testing.T) {
	type fields struct {
		logger *logger.Logger
		db     db.DB
	}
	l := logger.New()
	db := new(Mock)

	type args struct {
		reqId    string
		name     string
		category string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "One",
			fields: fields{logger: l,
				db: db},
			args:    args{reqId: "", name: "Fresh"},
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &catalog{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			_, err := s.CreateGoods(tt.args.reqId, tt.args.name, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("catalog.CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_catalog_DeleteGoods(t *testing.T) {
	type fields struct {
		logger *logger.Logger
		db     db.DB
	}

	l := logger.New()
	db := new(Mock)

	type args struct {
		reqId string
		id    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "One",
			fields: fields{logger: l,
				db: db},
			args:    args{reqId: "", id: "1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &catalog{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			_, err := s.DeleteGoods(tt.args.reqId, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("catalog.DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_catalog_UpdateGoods(t *testing.T) {
	type fields struct {
		logger *logger.Logger
		db     db.DB
	}

	l := logger.New()
	db := new(Mock)

	type args struct {
		reqId string
		id    string
		name  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "One",
			fields: fields{logger: l,
				db: db},
			args:    args{reqId: "", id: "1", name: "Fresh"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &catalog{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			_, err := s.UpdateGoods(tt.args.reqId, tt.args.id, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("catalog.DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
