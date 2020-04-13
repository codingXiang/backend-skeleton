package repository

import (
	"github.com/codingXiang/backend-skeleton/model"
	"github.com/codingXiang/backend-skeleton/module/demo"
	"github.com/jinzhu/gorm"
)

type DemoRepository struct {
	orm *gorm.DB
}

func NewDemoRepository(orm *gorm.DB) demo.Repository {
	return &DemoRepository{
		orm: orm,
	}
}

func (d *DemoRepository) CreateDepartment(data *model.Department) (model.DepartmentModelInterface, error) {
	var (
		err error
	)
	err = d.orm.Create(&data).Error
	return data, err
}

func (d *DemoRepository) CreateUser(data *model.User) (model.UserModelInterface, error) {
	var (
		err error
	)
	err = d.orm.Create(&data).Error
	return data, err
}

func (d *DemoRepository) GetDepartment(data *model.Department) (model.DepartmentModelInterface, error) {
	var (
		err error
	)
	err = d.orm.
		Preload("PreviousDepartment").
		First(&data, "id = ?", data.GetID()).Error
	return data, err
}

func (d *DemoRepository) GetUser(data *model.User) (model.UserModelInterface, error) {
	var (
		err error
	)
	err = d.orm.
		Preload("Department").
		Where("id = ?", data.GetID()).
		First(&data).Error
	return data, err
}

func (d *DemoRepository) ModifyDepartment(data *model.Department) (model.DepartmentModelInterface, error) {
	var (
		model = new(model.Department)
		err   error
	)
	model.ID = data.GetID()
	if err = d.orm.Model(&model).Updates(data).Error; err != nil {
		return nil, err
	}
	return d.GetDepartment(model)
}

func (d *DemoRepository) ModifyUser(data *model.User) (model.UserModelInterface, error) {
	var (
		model = new(model.User)
		err   error
	)
	model.ID = data.GetID()
	if err = d.orm.Model(&model).Updates(data).Error; err != nil {
		return nil, err
	}
	return d.GetUser(model)
}

func (d *DemoRepository) DeleteDepartment(data *model.Department) (error) {
	var (
		err error
	)
	err = d.orm.Delete(&data).Error
	return err
}

func (d *DemoRepository) DeleteUser(data *model.User) (error) {

	var (
		err error
	)
	err = d.orm.Delete(&data).Error
	return err
}

func (d *DemoRepository) GetDepartmentList() ([]model.DepartmentModelInterface, error) {
	var (
		datas  []*model.Department
		result = make([]model.DepartmentModelInterface, 0)
		err    error
	)
	if err = d.orm.Find(&datas).Error; err != nil {
		return nil, err
	}
	for _, data := range datas {
		result = append(result, data)
	}
	return result, nil
}

func (d *DemoRepository) GetUserList() ([]model.UserModelInterface, error) {
	var (
		datas  []*model.User
		result = make([]model.UserModelInterface, 0)
		err    error
	)
	if err = d.orm.Find(&datas).Error; err != nil {
		return nil, err
	}
	for _, data := range datas {
		result = append(result, data)
	}
	return result, nil
}
