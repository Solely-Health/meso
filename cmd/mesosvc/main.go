package main

import (
	inmem "github.com/meso-org/meso/inmemorydb"
	repo "github.com/meso-org/meso/repository"
	workers "github.com/meso-org/meso/workers"
)

func main() {
	var (
		inmemorydb = true
	)

	// Repository Registration here
	var (
		workersRepo repo.WorkerRepository
	)

	// For development purposes we will just use in memory db (for now, can be configured)
	if inmemorydb {
		workersRepo = inmem.NewWorkerRepository()
	} else {
		// we can pick and choose what kind of db we want to use here
	}

	// Service Registration here
	workersService := workers.NewService(workersRepo)
}
