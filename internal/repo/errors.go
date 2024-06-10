package repo

//
//import (
//	"errors"
//	"gorm.io/gorm"
//	"team-management/common/logger"
//)
//
//const (
//	InvalidParam   = "invalid param"
//	NotFound       = "not found"
//	InvalidRequest = "invalid request"
//)
//
//type RepoErr struct {
//	Message string
//	Type    string
//}
//
//func (r *RepoErr) Error() string { return r.Message }
//
//func NewRepoErr(Type, Message string) *RepoErr {
//	return &RepoErr{
//		Message: Message,
//		Type:    Type,
//	}
//}
//
//func HandleErr(err error, msg string) *RepoErr {
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		return NewRepoErr(NotFound, msg)
//	}
//	logger.Errorf(err.Error())
//	return NewRepoErr(InvalidRequest, err.Error())
//}
