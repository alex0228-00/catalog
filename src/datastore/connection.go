//go:generate mockgen -source=cred.go -destination=mock_cred.go -package=datastore

package datastore

import (
	"context"

	"catalog/src/datastore/ent"
	"catalog/src/datastore/ent/system"
)

type ConnStore struct {
	client *ent.Client
}

func NewConnStore(client *ent.Client) *ConnStore {
	return &ConnStore{
		client: client,
	}
}

func (store *ConnStore) GetConnectionById(ctx context.Context, id string) (*ent.System, error) {
	return store.client.System.Query().Where(system.IDEQ(id)).Only(ctx)
}

func (store *ConnStore) CreateConnection(ctx context.Context, sys *ent.System) error {
	_, err := store.client.System.
		Create().
		SetID(sys.ID).
		SetType(sys.Type).
		SetHost(sys.Host).
		SetUniqueIdentifier(sys.UniqueIdentifier).
		SetCredentials(sys.Credentials).
		SetCreatedBy(sys.CreatedBy).
		SetUpdatedBy(sys.UpdatedBy).
		Save(ctx)
	return err
}
