package service

import (
	"github.com/codingXiang/backend-skeleton/model"
	"github.com/codingXiang/backend-skeleton/module/demo"
)

type DemoService struct {
	demoRepo demo.Repository
}

func NewDemoService(demoRepo demo.Repository) demo.Service {
	return &DemoService{
		demoRepo: demoRepo,
	}
}

func (d *DemoService) CreateDepartment(data *model.Department) (model.DepartmentModelInterface, error) {
	if result, err := d.demoRepo.CreateDepartment(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *DemoService) CreateUser(data *model.User) (model.UserModelInterface, error) {
	if result, err := d.demoRepo.CreateUser(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *DemoService) DeleteDepartment(data *model.Department) (error) {
	if err := d.demoRepo.DeleteDepartment(data); err != nil {
		return err
	}
	return nil
}

func (d *DemoService) DeleteUser(data *model.User) (error) {
	if err := d.demoRepo.DeleteUser(data); err != nil {
		return err
	}
	return nil
}

func (d *DemoService) GetDepartment(data *model.Department) (model.DepartmentModelInterface, error) {
	if result, err := d.demoRepo.GetDepartment(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *DemoService) GetUser(data *model.User) (model.UserModelInterface, error) {
	if result, err := d.demoRepo.GetUser(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *DemoService) ModifyDepartment(data *model.Department) (model.DepartmentModelInterface, error) {
	if result, err := d.demoRepo.ModifyDepartment(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *DemoService) ModifyUser(data *model.User) (model.UserModelInterface, error) {
	if result, err := d.demoRepo.ModifyUser(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *DemoService) GetDepartmentList() ([]model.DepartmentModelInterface, error) {
	if result, err := d.demoRepo.GetDepartmentList(); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *DemoService) GetUserList() ([]model.UserModelInterface, error) {
	if result, err := d.demoRepo.GetUserList(); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
