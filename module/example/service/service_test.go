package service_test

import (
	"errors"
	"github.com/codingXiang/backend-skeleton/model"
	"github.com/codingXiang/backend-skeleton/module/example"
	"github.com/codingXiang/backend-skeleton/module/example/mock"
	"github.com/codingXiang/backend-skeleton/module/example/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

//Suite 為設定 mock repository
type Suite struct {
	suite.Suite
	service example.Service
}

//測試變數
var (
	testDepartmentData = []*model.Department{
		&model.Department{
			ID:                   "CX001",
			Name:                 "CX",
			PreviousDepartmentID: "",
		},
		&model.Department{},
	}
	testUserData = []*model.User{
		&model.User{
			ID:           "11634",
			Username:     "xiang",
			DepartmentID: "CX001",
			Email:        "xianglai6658@digiwin.com",
			Location:     "台中",
			EXT:          4665,
		},
		&model.User{},
	}
)

//初始化 Suite
func (s *Suite) SetupSuite() {
	// 建立 mock controller
	ctrl := gomock.NewController(s.T())
	// 透過 mock 建立 repository
	repo := mock.NewMockRepository(ctrl)
	/*
		建立 repository mock data
	*/
	// Create
	{
		//Department
		repo.EXPECT().CreateDepartment(testDepartmentData[0]).Return(testDepartmentData[0], nil)
		repo.EXPECT().CreateDepartment(testDepartmentData[1]).Return(nil, errors.New("error"))
		//User
		repo.EXPECT().CreateUser(testUserData[0]).Return(testUserData[0], nil)
		repo.EXPECT().CreateUser(testUserData[1]).Return(nil, errors.New("error"))
	}
	//GetList
	{
		//Department
		repo.EXPECT().GetDepartmentList().Return(testDepartmentData, nil)
		repo.EXPECT().GetDepartmentList().Return(nil, errors.New("error")).Times(1)
		//User
		repo.EXPECT().GetUserList().Return(testUserData, nil)
		repo.EXPECT().GetUserList().Return(nil, errors.New("error")).Times(1)
	}
	//Get
	{
		//Department
		repo.EXPECT().GetDepartment(testDepartmentData[0]).Return(testDepartmentData[0], nil)
		repo.EXPECT().GetDepartment(testDepartmentData[1]).Return(nil, errors.New("error"))
		//Department
		repo.EXPECT().GetUser(testUserData[0]).Return(testUserData[0], nil)
		repo.EXPECT().GetUser(testUserData[1]).Return(nil, errors.New("error"))
	}
	//Modify
	{
		//Department
		repo.EXPECT().ModifyDepartment(testDepartmentData[0]).Return(testDepartmentData[0], nil)
		repo.EXPECT().ModifyDepartment(testDepartmentData[1]).Return(nil, errors.New("error"))
		//Department
		repo.EXPECT().ModifyUser(testUserData[0]).Return(testUserData[0], nil)
		repo.EXPECT().ModifyUser(testUserData[1]).Return(nil, errors.New("error"))
	}
	//Delete
	{
		//Department
		repo.EXPECT().DeleteDepartment(testDepartmentData[0]).Return(nil)
		repo.EXPECT().DeleteDepartment(testDepartmentData[1]).Return(errors.New("error"))
		//Department
		repo.EXPECT().DeleteUser(testUserData[0]).Return(nil)
		repo.EXPECT().DeleteUser(testUserData[1]).Return(errors.New("error"))
	}
	// 初始化 demoService
	s.service = service.NewExampleService(repo)
}

//TestStart 為測試程式進入點
func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

//TestCreateDepartment 用於測試 Service 中的 CreateDepartment
func (s *Suite) TestCreateDepartment() {
	// 情境一
	{
		data, err := s.service.CreateDepartment(testDepartmentData[0])
		require.NoError(s.T(), err)
		require.NotNil(s.T(), data)
		require.IsType(s.T(), testDepartmentData[0], data)
	}
	// 情境二
	{
		data, err := s.service.CreateDepartment(testDepartmentData[1])
		require.Error(s.T(), err)
		require.Nil(s.T(), data)
	}
}

//TestCreateUser 用於測試 Service 中的 CreateUser
func (s *Suite) TestCreateUser() {
	// 情境一
	{
		data, err := s.service.CreateUser(testUserData[0])
		require.NoError(s.T(), err)
		require.NotNil(s.T(), data)
		require.IsType(s.T(), testUserData[0], data)
	}
	// 情境二
	{
		data, err := s.service.CreateUser(testUserData[1])
		require.Error(s.T(), err)
		require.Nil(s.T(), data)
	}
}

//TestGetDepartmentList 用於測試 Service 中的 GetDepartmentList
func (s *Suite) TestGetDepartmentList() {
	// 情境一
	{
		data, err := s.service.GetDepartmentList()
		require.NoError(s.T(), err)
		require.NotNil(s.T(), data)
		require.IsType(s.T(), testDepartmentData, data)
	}
	// 情境二
	{
		data, err := s.service.GetDepartmentList()
		require.Error(s.T(), err)
		require.Nil(s.T(), data)
	}
}

//TestGetUserList 用於測試 Service 中的 GetUserList
func (s *Suite) TestGetUserList() {
	// 情境一
	{
		data, err := s.service.GetUserList()
		require.NoError(s.T(), err)
		require.NotNil(s.T(), data)
		require.IsType(s.T(), testUserData, data)
	}
	// 情境二
	{
		data, err := s.service.GetUserList()
		require.Error(s.T(), err)
		require.Nil(s.T(), data)
	}
}

//TestGetDepartment 用於測試 Service 中的 GetDepartment
func (s *Suite) TestGetDepartment() {
	// 情境一
	{
		data, err := s.service.GetDepartment(testDepartmentData[0])
		require.NoError(s.T(), err)
		require.NotNil(s.T(), data)
		require.IsType(s.T(), testDepartmentData[0], data)
	}
	// 情境二
	{
		data, err := s.service.GetDepartment(testDepartmentData[1])
		require.Error(s.T(), err)
		require.Nil(s.T(), data)
	}
}

//TestGetUser 用於測試 Service 中的 GetUser
func (s *Suite) TestGetUser() {
	// 情境一
	{
		data, err := s.service.GetUser(testUserData[0])
		require.NoError(s.T(), err)
		require.NotNil(s.T(), data)
		require.IsType(s.T(), testUserData[0], data)
	}
	// 情境二
	{
		data, err := s.service.GetUser(testUserData[1])
		require.Error(s.T(), err)
		require.Nil(s.T(), data)
	}
}

//TestModifyDepartment 用於測試 Service 中的 ModifyDepartment
func (s *Suite) TestModifyDepartment() {
	// 情境一
	{
		data, err := s.service.ModifyDepartment(testDepartmentData[0])
		require.NoError(s.T(), err)
		require.NotNil(s.T(), data)
		require.IsType(s.T(), testDepartmentData[0], data)
	}
	// 情境二
	{
		data, err := s.service.ModifyDepartment(testDepartmentData[1])
		require.Error(s.T(), err)
		require.Nil(s.T(), data)
	}
}

//TestModifyUser 用於測試 Service 中的 ModifyUser
func (s *Suite) TestModifyUser() {
	// 情境一
	{
		data, err := s.service.ModifyUser(testUserData[0])
		require.NoError(s.T(), err)
		require.NotNil(s.T(), data)
		require.IsType(s.T(), testUserData[0], data)
	}
	// 情境二
	{
		data, err := s.service.ModifyUser(testUserData[1])
		require.Error(s.T(), err)
		require.Nil(s.T(), data)
	}
}

//TestDeleteDepartment 用於測試 Service 中的 DeleteDepartment
func (s *Suite) TestDeleteDepartment() {
	// 情境一
	{
		err := s.service.DeleteDepartment(testDepartmentData[0])
		require.NoError(s.T(), err)
	}
	// 情境二
	{
		err := s.service.DeleteDepartment(testDepartmentData[1])
		require.Error(s.T(), err)
	}
}

//TestDeleteUser 用於測試 Service 中的 DeleteUser
func (s *Suite) TestDeleteUser() {
	// 情境一
	{
		err := s.service.DeleteUser(testUserData[0])
		require.NoError(s.T(), err)
	}
	// 情境二
	{
		err := s.service.DeleteUser(testUserData[1])
		require.Error(s.T(), err)
	}
}