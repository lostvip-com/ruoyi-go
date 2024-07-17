package jobrunner

import (
	"github.com/winc-link/hummingbird/internal/dtos"
	coreContainer "github.com/winc-link/hummingbird/internal/hummingbird/core/container"
	"github.com/winc-link/hummingbird/internal/models"
	"github.com/winc-link/hummingbird/internal/pkg/container"

	"github.com/winc-link/hummingbird/internal/pkg/timer/jobs"
	"time"

	"github.com/winc-link/hummingbird/internal/pkg/di"
)

type JobRunFunc func(jobId string, job jobs.JobSchedule)

func NewJobRunFunc(dic *di.Container) JobRunFunc {
	logger := container.LoggingClientFrom(dic.Get)

	return func(jobId string, job jobs.JobSchedule) {
		start := time.Now()
		defer func() {
			logger.Infof("JobRunFunc cost: %v ms", time.Since(start).Milliseconds())
		}()

		logger.Infof("JobId: %v, job in: %+v", jobId, job.RuntimeJobStu)

		//调用驱动
		s := job.JobData
		deviceApp := coreContainer.DeviceItfFrom(dic.Get)
		res := deviceApp.DeviceAction(dtos.JobAction{
			ProductId:   s.ActionData[0].ProductId,
			ProductName: s.ActionData[0].ProductName,
			DeviceId:    s.ActionData[0].DeviceId,
			DeviceName:  s.ActionData[0].DeviceName,
			Code:        s.ActionData[0].Code,
			DateType:    s.ActionData[0].DateType,
			Value:       s.ActionData[0].Value,
		})

		dbClient := coreContainer.DBClientFrom(dic.Get)
		_, err := dbClient.AddSceneLog(models.SceneLog{
			SceneId: job.JobID,
			Name:    job.JobName,
			ExecRes: res.ToString(),
		})
		if err != nil {
			logger.Errorf("add sceneLog err %v", err.Error())
		}

	}
}
