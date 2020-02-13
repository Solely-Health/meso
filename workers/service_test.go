package workers

import (
	"testing"

	"github.com/meso-org/meso/repository"
)

type mockWorkersRepository struct {
	worker *repository.Worker
}

func (mockWorker *mockWorkersRepository) Store(worker *repository.Worker) error {
	mockWorker.worker = worker
	return nil
}

func TestRegisterNewWorker(t *testing.T) {
	// create mock variables
	var (
		email      = "mcnultymikkal@gmail.com"
		firstName  = "Mike"
		lastName   = "McNut"
		occupation = "Respitory Therapist"
		license    = "ASD123"
	)
	var workers mockWorkersRepository

	s := NewService(&workers)
	id, err := s.RegisterNewWorker(email, firstName, lastName, occupation, license)
	if err != nil {
		t.Fatal(err)
	}
	if id == "" {
		t.Fatal(id)
	}
}
