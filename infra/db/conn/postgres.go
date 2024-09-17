package conn

import (
	"cleanarch/pkg/env"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetPSQLConn() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), env.DatabaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	return conn, nil
}
