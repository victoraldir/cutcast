package app

import (
	"github.com/victoraldir/cutcast/internal/app/infra/config"
	"github.com/victoraldir/cutcast/internal/app/record/infra/adapters"
	"github.com/victoraldir/cutcast/internal/app/record/infra/controllers"
	"github.com/victoraldir/cutcast/internal/app/record/usecases"
)

type Application struct {
	RecordGroupController controllers.RecordController
	TrimGroupController   controllers.TrimController
}

func NewApplication(cfg config.Configuration) *Application {

	// Repositories
	recordFileRepository := adapters.NewRecordFileFFMPEGRepository()
	recordDbRepository := adapters.NewRecordDbMemoryRepository()
	trimDbRepository := adapters.NewTrimDbMemoryRepository()

	// UseCases
	createRecordGroupUseCase := usecases.NewCreateRecordGroup(
		recordFileRepository,
		recordDbRepository,
	)

	finishRecordGroupUseCase := usecases.NewFinishRecordGroup(
		recordDbRepository,
	)

	trimRecordUseCase := usecases.NewTrimRecordGroup(
		recordFileRepository,
		trimDbRepository,
	)

	listRecordGroupUseCase := usecases.NewListRecordGroup(
		recordDbRepository,
	)

	listTrimRecordGroupUseCase := usecases.NewListTrimRecordGroup(
		trimDbRepository,
	)

	return &Application{
		RecordGroupController: *controllers.NewRecordController(
			createRecordGroupUseCase,
			finishRecordGroupUseCase,
			listRecordGroupUseCase,
		),
		TrimGroupController: *controllers.NewTrimController(
			listTrimRecordGroupUseCase,
			trimRecordUseCase,
		),
	}
}
