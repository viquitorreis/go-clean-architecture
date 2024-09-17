package nir

import "cleanarch/infra/db/schemas/nir/results/results"

type NirUseCases struct {
	Repository INirRepository
}

var _INirUseCases = (*NirUseCases)(nil)

func NewNirUseCases(repository INirRepository) *NirUseCases {
	return &NirUseCases{Repository: repository}
}

func (n *NirUseCases) GetNirResults() ([]results.NirResult, error) {
	nirResults, err := n.Repository.GetNirResults()
	if err != nil {
		return nil, err
	}

	return nirResults, nil
}

func (n *NirUseCases) GetNIRResultsByID(id int64) (results.NirResult, error) {
	nirResult, err := n.Repository.GetNIRResultsByID(id)
	if err != nil {
		return results.NirResult{}, err
	}

	return nirResult, nil
}
