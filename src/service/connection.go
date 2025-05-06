package service

import (
	"context"

	"catalog/src"
	"catalog/src/datastore"
	"catalog/src/datastore/ent"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Map to ent.System
type Connection struct {
	ID               string     `json:"id"`
	Type             string     `json:"type"`
	Host             string     `json:"host"`
	UniqueIdentifier string     `json:"unique_identifier"`
	Credentials      Credential `json:"credentials,omitempty"`
	CreatedBy        string     `json:"created_by"`
	UpdatedBy        string     `json:"updated_by"`
}

type ConnectionService struct {
	connStore *datastore.ConnStore
	credCodec *SecureCredCodec
	logger    *zap.Logger
}

func NewConnectionService(key []byte, dbClient *ent.Client) (*ConnectionService, error) {
	codec, err := NewSecureCredCodec(key)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create secure credential codec")
	}

	store := &ConnectionService{
		credCodec: codec,
		logger:    src.DefaultLogger,
		connStore: datastore.NewConnStore(dbClient),
	}
	return store, nil
}

func (svc *ConnectionService) CreateConnection(ctx context.Context, conn Connection) (*Connection, error) {
	conn.ID = uuid.New().String()
	logger := svc.logger.With(zap.String(src.LogConnectionId, conn.ID))

	raw, err := svc.convertToEntSystem(&conn)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to transform connection for database")
	}

	logger.Info("Creating new connection")
	if err = svc.connStore.CreateConnection(ctx, raw); err != nil {
		if src.AsError[*ent.ConstraintError](err) {
			return nil, src.ErrorDuplicatedSystem
		}
		return nil, errors.Wrap(err, "Failed to create connection in database")
	}

	return &conn, nil
}

func (svc *ConnectionService) GetConnectionById(ctx context.Context, id string) (*Connection, error) {
	logger := svc.logger.With(zap.String(src.LogConnectionId, id))

	logger.Debug("Fetching connection from database")
	raw, err := svc.connStore.GetConnectionById(ctx, id)
	if err != nil {
		if src.AsError[*ent.NotFoundError](err) {
			return nil, src.ErrorSystemNotFound
		}
		return nil, errors.Wrap(err, "Failed to fetch connection from database")
	}

	conn, err := svc.convertToConnection(raw)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode connection from database result")
	}

	logger.Debug("Successfully fetch connection from database")
	return conn, nil
}

/*
 * Private functions
 */
func (svc *ConnectionService) convertToConnection(raw *ent.System) (*Connection, error) {
	conn := new(Connection)

	conn.ID = raw.ID
	conn.Type = raw.Type
	conn.Host = raw.Host
	conn.UniqueIdentifier = raw.UniqueIdentifier
	conn.CreatedBy = raw.CreatedBy
	conn.UpdatedBy = raw.UpdatedBy

	if len(raw.Credentials) > 0 {
		cred, err := svc.credCodec.Decode(raw.Credentials)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to decode credentials from database result")
		}
		conn.Credentials = cred
	}

	return conn, nil
}

func (svc *ConnectionService) convertToEntSystem(conn *Connection) (*ent.System, error) {
	cred, err := svc.credCodec.Encode(conn.Credentials)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode credentials for database")
	}

	raw := &ent.System{
		ID:               conn.ID,
		Type:             conn.Type,
		Host:             conn.Host,
		UniqueIdentifier: conn.UniqueIdentifier,
		Credentials:      cred,
		CreatedBy:        conn.CreatedBy,
		UpdatedBy:        conn.UpdatedBy,
	}

	return raw, nil
}
