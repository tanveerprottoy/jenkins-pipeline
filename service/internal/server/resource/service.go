package resource

import (
	"context"
	"net/http"

	"github.com/tanveerprottoy/basic-go-server/internal/pkg/constant"
	"github.com/tanveerprottoy/basic-go-server/pkg/config"
	"github.com/tanveerprottoy/basic-go-server/pkg/errorpkg"
	"github.com/tanveerprottoy/basic-go-server/pkg/timepkg"
)

type Service struct {
	repository *Repository
}

func NewService(r *Repository) *Service {
	s := new(Service)
	s.repository = r
	return s
}

func (s Service) readOneInternal(id string) (Resource, error) {
	return s.repository.ReadOne(id)
}

func (s Service) GetBasicData(ctx context.Context) map[string]any {
	m := make(map[string]any)
	m["message"] = config.GetEnvValue("MESSAGE")
	return m
}

func (s Service) Create(d *CreateUpdateResourceDto, ctx context.Context) (Resource, *errorpkg.HTTPError) {
	// convert dto to entity
	b := Resource{}
	b.Name = d.Name
	n := timepkg.NowUnixMilli()
	b.CreatedAt = n
	b.UpdatedAt = n
	err := s.repository.Create(&b)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	return b, nil
}

func (s Service) ReadMany(limit, page int, ctx context.Context) (map[string]any, *errorpkg.HTTPError) {
	m := make(map[string]any)
	m["items"] = make([]Resource, 0)
	m["limit"] = limit
	m["page"] = page
	offset := limit * (page - 1)
	d, err := s.repository.ReadMany(limit, offset)
	if err != nil {
		return m, errorpkg.HandleDBError(err)
	}
	m["items"] = d
	return m, nil
}

func (s Service) ReadOne(id string, ctx context.Context) (Resource, *errorpkg.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	return b, nil
}

func (s Service) Update(id string, d *CreateUpdateResourceDto, ctx context.Context) (Resource, *errorpkg.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	b.Name = d.Name
	b.UpdatedAt = timepkg.NowUnixMilli()
	rows, err := s.repository.Update(id, &b)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	if rows > 0 {
		return b, nil
	}
	return b, &errorpkg.HTTPError{Code: http.StatusBadRequest, Err: errorpkg.NewError(constant.OperationNotSuccess)}
}

func (s Service) Delete(id string, ctx context.Context) (Resource, *errorpkg.HTTPError) {
	b, err := s.readOneInternal(id)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	rows, err := s.repository.Delete(id)
	if err != nil {
		return b, errorpkg.HandleDBError(err)
	}
	if rows > 0 {
		return b, nil
	}
	return b, &errorpkg.HTTPError{Code: http.StatusBadRequest, Err: errorpkg.NewError(constant.OperationNotSuccess)}
}
