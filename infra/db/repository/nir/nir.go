package nirRepository

import (
	"cleanarch/infra/db/schemas/nir/results/results"
	"cleanarch/services/nir"
	"context"

	"github.com/jackc/pgx/v5"
)

type NirRepository struct {
	// Repository *INirRepository
	db *pgx.Conn
}

var _ nir.INirRepository = (*NirRepository)(nil)

func NewINirRepository(db *pgx.Conn) *NirRepository {
	return &NirRepository{
		db: db,
	}
}

func (r *NirRepository) GetNirResults() ([]results.NirResult, error) {
	queries := results.New(r.db)
	nirResults, err := queries.GetResults(context.Background())
	if err != nil {
		return nil, err
	}

	return nirResults, nil
}

func (r *NirRepository) GetNIRResultsByID(id int64) (results.NirResult, error) {
	queries := results.New(r.db)
	nirResult, err := queries.GetResult(context.Background(), id)
	if err != nil {
		return results.NirResult{}, err
	}

	return nirResult, nil
}
