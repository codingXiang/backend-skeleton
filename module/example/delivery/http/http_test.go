package http_test

import (
	"github.com/codingXiang/backend-skeleton/module/example/delivery"
	"github.com/stretchr/testify/suite"
	"testing"
)

//Suite 為設定 mock repository
type Suite struct {
	suite.Suite
	httpHandler delivery.HttpHandler
}

//初始化 Suite
func (s *Suite) SetupSuite() {}

//TestStart 為測試程式進入點
func TestStart(t *testing.T) {
	suite.Run(t, new(Suite))
}