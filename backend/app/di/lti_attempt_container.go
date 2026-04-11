package di

import (
	"gitlab.com/a10869/api-modules/backend/app/usecases"
)

func (di *DIContainer) LTIAttemptCreateUC() *usecases.LTIAttemptCreateUC {
	return &usecases.LTIAttemptCreateUC{
		LTIAttemptQueries:     di.Queries.LTIAttemptQueries,
		LTIRoomQueries:        di.Queries.LTIRoomQueries,
		LaunchData:            di.Queries.LTILaunchDataQueries,
		RoundQueuePoolQueries: di.Queries.RoundQueuePoolQueries,
		LTIRoutingQueries:     di.Queries.LTIRoutingQueries,
		PNETServerQueries:     di.Queries.PNETServerQueries,
		UserQueries:           di.Queries.UserQueries,
	}
}

func (di *DIContainer) LTIAttemptEditUC() *usecases.LTIAttemptEditUC {
	return &usecases.LTIAttemptEditUC{
		LTIAttemptQueries: di.Queries.LTIAttemptQueries,
	}
}

func (di *DIContainer) LTIAttemptEditBulkUC() *usecases.LTIAttemptEditBulkUC {
	return &usecases.LTIAttemptEditBulkUC{
		LTIAttemptQueries:   di.Queries.LTIAttemptQueries,
		PNETServerQueries:   di.Queries.PNETServerQueries,
		TaskLTISyncResultUC: di.TaskLTISyncResultUC(),
	}
}

func (di *DIContainer) LTIAttemptGetUC() *usecases.LTIAttemptGetUC {
	return &usecases.LTIAttemptGetUC{
		LTIAttemptQueries: di.Queries.LTIAttemptQueries,
	}
}

func (di *DIContainer) LTIAttemptListUC() *usecases.LTIAttemptListUC {
	return &usecases.LTIAttemptListUC{
		LTIAttemptQueries: di.Queries.LTIAttemptQueries,
	}
}

func (di *DIContainer) LTIAttemptDeleteUC() *usecases.LTIAttemptDeleteUC {
	return &usecases.LTIAttemptDeleteUC{
		LTIAttemptQueries: di.Queries.LTIAttemptQueries,
	}
}
