package nir

import "cleanarch/infra/db/schemas/nir/results/results"

type INirRepository interface {
	GetNirResults() ([]results.NirResult, error)
	GetNIRResultsByID(id int64) (results.NirResult, error)
}

type INirUseCase interface {
	GetNirResults() ([]results.NirResult, error)
	GetNIRResultsByID(id int64) (results.NirResult, error)
}
