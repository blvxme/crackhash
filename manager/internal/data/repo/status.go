package repo

import "manager/internal/common/dto"

type RequestStatusRepo interface {
	Create(uuid string, requestStatus dto.RequestStatus, opts ...interface{}) (err error)
	Read(uuid string, opts ...interface{}) (requestStatus dto.RequestStatus, err error)
	Update(uuid string, requestStatus dto.RequestStatus, opts ...interface{}) (err error)
	Delete(uuid string, opts ...interface{}) (err error)
}
