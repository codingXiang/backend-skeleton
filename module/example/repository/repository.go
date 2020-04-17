package repository

import (
	"github.com/codingXiang/backend-skeleton/model"
	"github.com/codingXiang/backend-skeleton/module/example"
	"github.com/jinzhu/gorm"
)

type ExampleRepository struct {
	orm *gorm.DB
}

func NewExampleRepository(orm *gorm.DB) example.Repository {
	return &ExampleRepository{
		orm: orm,
	}
}

func (d *ExampleRepository) CreateDepartment(data model.DepartmentInterface) (*model.Department, error) {
	var (
		err error
		in  = data.(*model.Department)
	)
	err = d.orm.Create(&in).Error
	return in, err
}

func (d *ExampleRepository) CreateUser(data model.UserInterface) (*model.User, error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = d.orm.Create(&in).Error
	return in, err
}

func (d *ExampleRepository) GetDepartment(data model.DepartmentInterface) (*model.Department, error) {
	var (
		err error
		in  = data.(*model.Department)
	)
	err = d.orm.
		Preload("PreviousDepartment").
		First(&in).Error
	return in, err
}

func (d *ExampleRepository) GetUser(data model.UserInterface) (*model.User, error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = d.orm.
		Preload("Department").
		First(&in).Error
	return in, err
}

func (d *ExampleRepository) ModifyDepartment(data model.DepartmentInterface) (*model.Department, error) {
	var (
		err error
		in  = data.(*model.Department)
	)
	err = d.orm.Model(&in).Updates(in).Error
	return in, err
}

func (d *ExampleRepository) ModifyUser(data model.UserInterface) (*model.User, error) {
	var (
		err error
		in  = data.(*model.User)
	)
	err = d.orm.Model(&in).Updates(in).Error
	return in, err
}

func (d *ExampleRepository) DeleteDepartment(data model.DepartmentInterface) (error) {
	var (
		err error
		in  = data.(*model.Department)
	)
	err = d.orm.Delete(&in).Error
	return err
}

func (d *ExampleRepository) DeleteUser(data model.UserInterface) (error) {

	var (
		err error
		in  = data.(*model.User)
	)
	err = d.orm.Delete(&in).Error
	return err
}

func (d *ExampleRepository) GetDepartmentList() ([]*model.Department, error) {
	var (
		datas = make([]*model.Department, 0)
		err   error
	)
	err = d.orm.Find(&datas).Error
	return datas, err
}

func (d *ExampleRepository) GetUserList() ([]*model.User, error) {
	var (
		datas = make([]*model.User, 0)
		err   error
	)
	err = d.orm.Find(&datas).Error
	return datas, err
}
