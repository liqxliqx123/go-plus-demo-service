package beluga

import (
	"context"
	"my-demo-service/c"
	"my-demo-service/models"
	"fmt"

	belugapb "gitlab.xxx.com/xxx-xxx/go-kit/pb/beluga"
)

//buildArgoTask makes stateful day task spec
func (g *GrpcClient) buildArgoTask(ctx context.Context, data *models.xxxKafkaMsg, imageTag, belugaTaskName string, tasks ...string) (*belugapb.StatefulDagTaskSpec, error) {
	envs, err := g.getArgoEnvs()
	if err != nil {
		return nil, err
	}
	secrets := []*belugapb.Dag_Step_Secret{
		{
			Name: c.ArgoTokenName,
			Key:  c.ApolloToken,
		},
		{
			Name: c.ArgoTokenName,
			Key:  "ELASTIC_APM_SECRET_TOKEN",
		},
	}

	if imageTag == "local" || imageTag == "" {
		imageTag = "latest"
	}

	return &belugapb.StatefulDagTaskSpec{
		Dag: &belugapb.Dag{
			Steps: g.buildArgoSteps(ctx, data, imageTag, envs, secrets, belugaTaskName, tasks...),
		},
	}, nil
}

func (g *GrpcClient) getArgoEnvs() ([]*belugapb.Dag_Step_Env, error) {
	return []*belugapb.Dag_Step_Env{
		{
			Name:  c.ApolloAppID,
			Value: g.apolloAppID,
		},
		{
			Name:  c.ApolloCLuster,
			Value: g.apolloCluster,
		},
		{
			Name:  c.ApolloHost,
			Value: g.apolloHost,
		},
		{
			Name:  c.ApolloNameSpace,
			Value: g.apolloNameSpace,
		},
		{
			Name:  "SERVICE.NAME",
			Value: g.serviceName,
		},
		{
			Name:  "ENABLE_APOLLO",
			Value: "TRUE",
		},
		{
			Name:  "ENV",
			Value: g.env,
		},
	}, nil
}

func (g *GrpcClient) buildArgoSteps(ctx context.Context, data *models.xxxKafkaMsg, image string, envs []*belugapb.Dag_Step_Env, secrets []*belugapb.Dag_Step_Secret, belugaTaskName string, tasks ...string) []*belugapb.Dag_Step {
	var steps []*belugapb.Dag_Step
	//var lenTasks = len(tasks)
	var dependence *string
	//var final = false
	for taskOrder, task := range tasks {
		//if taskOrder == lenTasks-1 {
		//	final = true
		//}
		//sequence := taskOrder + 1
		var step = belugapb.Dag_Step{
			Name:  task,
			Image: image,
			//Commands: []string{fmt.Sprintf(c.TaskCommand, task, sequence, final, belugaTaskName)},
			Commands: []string{fmt.Sprintf(c.TaskCommand, data.xxxResultID)},
			Envs:     envs,
			Secrets:  secrets,
			Depends:  nil,
			//RetryPolicy: nil,
			//FanOut:      nil,
		}
		if taskOrder > 0 {
			step.Depends = dependence
		}
		dependence = &tasks[taskOrder]
		steps = append(steps, &step)
	}
	return steps
}
