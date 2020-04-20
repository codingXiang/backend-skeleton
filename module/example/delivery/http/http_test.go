package http_test

import (
	"fmt"
	"github.com/codingXiang/backend-skeleton/model"
	"github.com/codingXiang/backend-skeleton/module/example/delivery"
	http2 "github.com/codingXiang/backend-skeleton/module/example/delivery/http"
	"github.com/codingXiang/backend-skeleton/module/example/mock"
	delivery2 "github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/delivery/http"
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/codingXiang/cxgateway/pkg/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	http3 "net/http"
	"testing"
)

//Suite 為設定 mock repository
type Suite struct {
	suite.Suite
	gateway     delivery2.HttpHandler
	tester      util.HttpTesterInterface
	httpHandler delivery.HttpHandler
}

//測試變數
var (
	testDepartmentData = []*model.Department{
		&model.Department{
			ID:                   "CX001",
			Name:                 "CX",
			PreviousDepartmentID: "",
		},
	}
	testUserData = []*model.User{
		&model.User{
			ID:           "11634",
			Username:     "xiang",
			DepartmentID: "test",
			Email:        "xianglai6658@digiwin.com",
			Location:     "",
		},
	}
)

//初始化 Suite
func (s *Suite) SetupSuite() {
	//建立 Api Gateway，自定義 config
	s.gateway = http.NewApiGatewayWithData("config", []byte(`
application:
  timeout:
    read: 1000
    write: 1000
  port: 8080
  mode: "release"
  log:
    level: "error"
    format: "text"
  appId: "app"
  appToken: ""
  apiBaseRoute: "/api"
i18n:
  defaultLanguage: "zh_Hant"
  file:
    path: "../../../../i18n"
    type: "yaml"
`))
	s.tester = util.NewHttpTester(s.gateway.GetEngine())
	// 建立 mock controller
	ctrl := gomock.NewController(s.T())
	// 透過 mock 建立 service
	service := mock.NewMockService(ctrl)
	//設定 mock 資料
	{
		//Get
		{
			service.EXPECT().GetDepartment(&model.Department{ID: "CX001",}).Return(testDepartmentData[0], nil)
			service.EXPECT().GetUser(&model.User{ID: "11634"}).Return(testUserData[0], nil)
		}
		//GetList
		{
			service.EXPECT().GetDepartmentList().Return(testDepartmentData, nil)
			service.EXPECT().GetUserList().Return(testUserData, nil)
		}
		//Modify
		{
			service.EXPECT().ModifyDepartment(testDepartmentData[0]).Return(testDepartmentData[0], nil)
			service.EXPECT().ModifyUser(testUserData[0]).Return(testUserData[0], nil)
		}
		//Create
		{
			service.EXPECT().CreateDepartment(testDepartmentData[0]).Return(testDepartmentData[0], nil)
			service.EXPECT().CreateUser(testUserData[0]).Return(testUserData[0], nil)
		}
		//Delete
		{
			service.EXPECT().DeleteDepartment(&model.Department{
				ID: "CX001",
			}).Return(nil)
			service.EXPECT().DeleteUser(&model.User{
				ID: "11634",
			}).Return(nil)
		}
	}
	// 設定 http handler
	s.httpHandler = http2.NewExampleHandler(s.gateway, service)
}

//TestStart 為測試程式進入點
func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}

//TestGetDepartmentList 為測試 HttpHandler 中的 GetDepartmentList 方法
func (s *Suite) TestGetDepartmentList() {
	uri := "/api/v1/department"
	statusCode, response := s.tester.GET(uri)
	require.Equal(s.T(), http3.StatusOK, statusCode)
	require.Equal(s.T(), e.SUCCESS, response.Code)
}

//TestGetDepartment 為測試 HttpHandler 中的 GetDepartment 方法
func (s *Suite) TestGetDepartment() {
	uri := "/api/v1/department/CX001"
	statusCode, response := s.tester.GET(uri)
	require.Equal(s.T(), http3.StatusOK, statusCode)
	require.Equal(s.T(), e.SUCCESS, response.Code)
}

//TestGetUserList 為測試 HttpHandler 中的 GetUserList 方法
func (s *Suite) TestGetUserList() {
	uri := "/api/v1/user"
	statusCode, response := s.tester.GET(uri)
	fmt.Println(response)
	require.Equal(s.T(), http3.StatusOK, statusCode)
	require.Equal(s.T(), e.SUCCESS, response.Code)
}

//TestGetUser 為測試 HttpHandler 中的 GetUser 方法
func (s *Suite) TestGetUser() {
	uri := "/api/v1/user/11634"
	statusCode, response := s.tester.GET(uri)
	require.Equal(s.T(), http3.StatusOK, statusCode)
	require.Equal(s.T(), e.SUCCESS, response.Code)
}

//TestCreateDepartment 為測試 HttpHandler 中的 CreateDepartment 方法
func (s *Suite) TestCreateDepartment() {
	uri := "/api/v1/department"
	statusCode, response := s.tester.POST(uri, testDepartmentData[0])
	require.Equal(s.T(), http3.StatusCreated, statusCode)
	require.Equal(s.T(), e.CREATED, response.Code)
}

//TestCreateUser 為測試 HttpHandler 中的 CreateUser 方法
func (s *Suite) TestCreateUser() {
	uri := "/api/v1/user"
	statusCode, response := s.tester.POST(uri, testUserData[0])
	require.Equal(s.T(), http3.StatusCreated, statusCode)
	require.Equal(s.T(), e.CREATED, response.Code)
}

//TestModifyDepartment 為測試 HttpHandler 中的 ModifyDepartment 方法
func (s *Suite) TestModifyDepartment() {
	uri := "/api/v1/department/CX001"
	statusCode, _ := s.tester.PATCH(uri, testDepartmentData[0])
	require.Equal(s.T(), http3.StatusCreated, statusCode)
}

//TestModifyUser 為測試 HttpHandler 中的 ModifyUser 方法
func (s *Suite) TestModifyUser() {
	uri := "/api/v1/user/11634"
	statusCode, response := s.tester.PATCH(uri, testUserData[0])
	require.Equal(s.T(), http3.StatusCreated, statusCode)
	fmt.Println(response)
	require.Equal(s.T(), e.CREATED, response.Code)
}

//TestDeleteDepartment 為測試 HttpHandler 中的 DeleteDepartment 方法
func (s *Suite) TestDeleteDepartment() {
	uri := "/api/v1/department/CX001"
	statusCode, _ := s.tester.DELETE(uri, nil)
	require.Equal(s.T(), http3.StatusNoContent, statusCode)
}

//TestDeleteUser 為測試 HttpHandler 中的 DeleteUser 方法
func (s *Suite) TestDeleteUser() {
	uri := "/api/v1/user/11634"
	statusCode, _ := s.tester.DELETE(uri, nil)
	require.Equal(s.T(), http3.StatusNoContent, statusCode)

}
