package inmemorydb

import (
	"fmt"
	"sync"

	repository "github.com/meso-org/meso/repository"
)

type workerRepository struct {
	mtx     sync.RWMutex
	workers map[repository.WorkerID]*repository.Worker
}

// Store - A instance of the Store() definition in the repository interface
// Locates a worker via WorkerID in the workers map
func (r *workerRepository) Store(w *repository.Worker) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.workers[w.WorkerID] = w
	return nil
}

func (r *workerRepository) Find(id repository.WorkerID) (*repository.Worker, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	worker := r.workers[id]
	if worker == nil {
		return worker, fmt.Errorf("Could not find worker by id: %v", id)
	}
	return worker, nil
}

func (r *workerRepository) FindAll() ([]*repository.Worker, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	workers := []*repository.Worker{}
	for _, worker := range r.workers {
		workers = append(workers, worker)
	}

	return workers, nil
}

// NewWorkerRepository returns a new instance of a in-memory cargo repository.
func NewWorkerRepository() repository.WorkerRepository {
	return &workerRepository{
		workers: make(map[repository.WorkerID]*repository.Worker),
	}
}

//  ---------------- //
type facilityRepository struct {
	mtx        sync.RWMutex
	facilities map[repository.FacilityID]*repository.Facility
}

// Store - A instance of the Store() definition in the repository interface
// Locates a worker via WorkerID in the workers map
func (r *facilityRepository) Store(f *repository.Facility) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.facilities[f.FacilityID] = f
	return nil
}
