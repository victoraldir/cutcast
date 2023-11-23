package app

import (
	"github.com/victoraldir/cutcast/internal/app/infra/config"
	"github.com/victoraldir/cutcast/internal/app/record/infra/adapters"
	"github.com/victoraldir/cutcast/internal/app/record/infra/controllers"
	"github.com/victoraldir/cutcast/internal/app/record/infra/watchers"
	"github.com/victoraldir/cutcast/internal/app/record/usecases"
	"github.com/victoraldir/cutcast/pkg/command"
)

type Application struct {
	RecordGroupController controllers.RecordController
	TrimGroupController   controllers.TrimController
	FsWatcher             watchers.RealTimeStreamWatcher
}

func NewApplication(cfg config.Configuration) *Application {

	// Watcher
	fsWatcher, err := watchers.NewRealTimeStreamWatcher()
	if err != nil {
		panic(err)
	}

	// Start listening for events.
	defer fsWatcher.Listen()

	// Repositories
	recordDbRepository := adapters.NewRecordDbMemoryRepository()
	trimDbRepository := adapters.NewTrimDbMemoryRepository()
	fsWatcherRepository := adapters.NewWatcherDbMemoryRepository(fsWatcher.Watch, fsWatcher.Unwatch)
	recordFileRepository := adapters.NewRecordFileFFMPEGRepository(command.NewCommandBuilder())

	// UseCases
	createRecordGroupUseCase := usecases.NewCreateRecordGroup(
		recordFileRepository,
		recordDbRepository,
		fsWatcherRepository,
		cfg.Media.Dir,
	)

	finishRecordGroupUseCase := usecases.NewFinishRecordGroup(
		recordDbRepository,
		fsWatcherRepository,
	)

	trimRecordUseCase := usecases.NewTrimRecordGroup(
		recordFileRepository,
		trimDbRepository,
		cfg.Media.Dir,
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
		FsWatcher: *fsWatcher,
	}
}
