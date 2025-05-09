//go:build database

package service

import (
	"context"
	"testing"

	"catalog/src"

	"github.com/stretchr/testify/require"
)

func TestConnectionStore(t *testing.T) {
	rq := require.New(t)
	ctx := context.Background()

	mockCred, cleanup := src.MockCredentials(t, 1)
	defer cleanup()

	svc, err := NewConnectionService(src.CipherKey, DbClient)
	rq.NoError(err)

	t.Run("create and get", func(t *testing.T) {
		rq.NoError(err)

		toCreate := Connection{
			Type:             "test-type",
			Host:             "test-host",
			UniqueIdentifier: "test-unique-id",
			Credentials:      mockCred[0],
			CreatedBy:        "test-creator",
			UpdatedBy:        "test-updater",
		}

		expected, err := svc.CreateConnection(ctx, toCreate)
		rq.NoError(err)
		rq.NotEmpty(expected.ID)

		conn, err := svc.GetConnectionById(ctx, expected.ID)
		rq.NoError(err)

		rq.EqualExportedValues(expected, conn)
	})
}
