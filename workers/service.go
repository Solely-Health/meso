package workers

import (
	"fmt"
	"github.com/meso-org/meso/repository"
)

type Service interface {
	RegisterNewWorker(firstName, lastName, occupation, license string) (repository.WorkerID, error)
}

type service struct {
	workers repository.WorkerRepository
}

func (s *service) RegisterNewWorker(firstName, lastName, occupation, license string) (repository.WorkerID, error) {
	if firstName == "" || lastName == "" || occupation == "" {
		return "", fmt.Errorf("in RegisterNewWorker, provided arguments are invalid")
	}

	workerID := repository.GenerateWorkerID()

	worker := repository.NewWorker(workerID, firstName, lastName, occupation, license)

	if err := s.workers.Store(worker); err != nil {
		return "", err
	}

	// we can trigger a "NewWorkerRegistered" to other services from here

	return worker.WorkerID, nil
}

// NewService - pass this function a repository instance,
// and it will return a new service that has access to that repository
func NewService(workersRepo repository.WorkerRepository) Service {
	return &service{
		workers: workersRepo,
	}
}
