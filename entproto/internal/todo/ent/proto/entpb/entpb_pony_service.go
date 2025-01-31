// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	entproto "entgo.io/contrib/entproto"
	ent "entgo.io/contrib/entproto/internal/todo/ent"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// PonyService implements PonyServiceServer
type PonyService struct {
	client *ent.Client
	UnimplementedPonyServiceServer
}

// NewPonyService returns a new PonyService
func NewPonyService(client *ent.Client) *PonyService {
	return &PonyService{
		client: client,
	}
}

// ToProtoPony transforms the ent type to the pb type
func ToProtoPony(e *ent.Pony) (*Pony, error) {
	v := &Pony{}
	id := int64(e.ID)
	v.Id = id
	name := e.Name
	v.Name = name
	return v, nil
}

// ToProtoPonyList transforms a list of ent type to a list of pb type
func ToProtoPonyList(e []*ent.Pony) ([]*Pony, error) {
	var pbList []*Pony
	for _, entEntity := range e {
		pbEntity, err := ToProtoPony(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// BatchCreate implements PonyServiceServer.BatchCreate
func (svc *PonyService) BatchCreate(ctx context.Context, req *BatchCreatePoniesRequest) (*BatchCreatePoniesResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.PonyCreate, len(requests))
	for i, req := range requests {
		pony := req.GetPony()
		var err error
		bulk[i], err = svc.createBuilder(pony)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Pony.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := ToProtoPonyList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreatePoniesResponse{
			Ponies: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *PonyService) createBuilder(pony *Pony) (*ent.PonyCreate, error) {
	m := svc.client.Pony.Create()
	ponyName := pony.GetName()
	m.SetName(ponyName)
	return m, nil
}
