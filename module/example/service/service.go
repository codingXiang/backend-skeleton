package service

import (
	"github.com/codingXiang/backend-skeleton/model"
	"github.com/codingXiang/backend-skeleton/module/example"
)

type ExampleService struct {
	demoRepo example.Repository
}

func NewExampleService(demoRepo example.Repository) example.Service {
	return &ExampleService{
		demoRepo: demoRepo,
	}
}

func (d *ExampleService) CreateDepartment(data model.DepartmentInterface) (*model.Department, error) {
	if result, err := d.demoRepo.CreateDepartment(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *ExampleService) CreateUser(data model.UserInterface) (*model.User, error) {
	if result, err := d.demoRepo.CreateUser(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *ExampleService) DeleteDepartment(data model.DepartmentInterface) (error) {
	if err := d.demoRepo.DeleteDepartment(data); err != nil {
		return err
	}
	return nil
}

func (d *ExampleService) DeleteUser(data model.UserInterface) (error) {
	if err := d.demoRepo.DeleteUser(data); err != nil {
		return err
	}
	return nil
}

func (d *ExampleService) GetDepartment(data model.DepartmentInterface) (*model.Department, error) {
	if result, err := d.demoRepo.GetDepartment(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *ExampleService) GetUser(data model.UserInterface) (*model.User, error) {
	if result, err := d.demoRepo.GetUser(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *ExampleService) ModifyDepartment(data model.DepartmentInterface) (*model.Department, error) {
	if result, err := d.demoRepo.ModifyDepartment(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *ExampleService) ModifyUser(data model.UserInterface) (*model.User, error) {
	if result, err := d.demoRepo.ModifyUser(data); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *ExampleService) GetDepartmentList() ([]*model.Department, error) {
	if result, err := d.demoRepo.GetDepartmentList(); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func (d *ExampleService) GetUserList() ([]*model.User, error) {
	if result, err := d.demoRepo.GetUserList(); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}
