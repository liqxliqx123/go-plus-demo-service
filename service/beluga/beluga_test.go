package beluga

import (
	"context"
	"my-demo-service/c"
	"my-demo-service/models"
	"my-demo-service/test"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	belugapb "gitlab.xxx.com/xxx-xxx/go-kit/pb/beluga"
)

type BelugaTestSuite struct {
	test.Suite
	MockBelugaClient *MockClient
	srv              GrpcClient
}

func (t *BelugaTestSuite) SetupTest() {
	t.Suite.SetupTest()
	t.MockBelugaClient = NewMockClient(t.Ctrl)
	t.srv = GrpcClient{
		grpcClient:      t.MockBelugaClient,
		serviceName:     "my-demo-service",
		apolloAppID:     "apolloID",
		apolloHost:      "apolloHost",
		apolloCluster:   "apolloCluster",
		apolloNameSpace: "apolloNameSpace",
	}
}

func TestBelugaTestSuite(t *testing.T) {
	suite.Run(t, new(BelugaTestSuite))
}

//func TestBeluga(t *testing.T) {
//	kitConfig.LoadConf(utils.GetProjectPath(), config.GetConfig())
//	cfg := config.GetConfig()
//	err := InitBelugaGRPCClient(cfg)
//	assert.NoError(t, err)
//	beluga := GetInterfaceBeluga()
//	taskID, err := beluga.CreateTask(context.TODO(), nil, "image", "xxx_test", "task1", "task2")
//	assert.NoError(t, err)
//	t.Log(taskID)
//}

func (t *BelugaTestSuite) TestBuildArgoTask() {
	data := &models.xxxKafkaMsg{xxxResultID: 1}
	_, err := t.srv.buildArgoTask(context.TODO(), data, "image_tage", "beluga_task_name", "task1", "task2")
	assert.NoError(t.T(), err)
}

func (t *BelugaTestSuite) TestGetArgoEnvs() {
	_, err := t.srv.getArgoEnvs()
	assert.NoError(t.T(), err)
}

func (t *BelugaTestSuite) TestBuildArgoSteps() {
	data := &models.xxxKafkaMsg{xxxResultID: 1}
	env, _ := t.srv.getArgoEnvs()

	secrets := []*belugapb.Dag_Step_Secret{
		{
			Name: c.ArgoTokenName,
			Key:  c.ApolloToken,
		},
	}
	t.srv.buildArgoSteps(context.TODO(), data, "image", env, secrets, "beluga_task_name", "task1", "task2")
}

func (t *BelugaTestSuite) TestCreateTask() {
	data := &models.xxxKafkaMsg{xxxResultID: 1}
	t.MockBelugaClient.EXPECT().CreateStatefulDagTask(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), maxExecutionTime, gomock.Any()).Return(int64(1), nil)
	_, err := t.srv.CreateTask(context.TODO(), data, "image_tage", "beluga_task_name", "task1", "task2")
	assert.NoError(t.T(), err)
}
