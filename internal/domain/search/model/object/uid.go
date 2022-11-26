package internal_domain_search_model_object

import "github.com/google/uuid"

type UId string

func (u UId) String() string {
	return string(u)
}

func GenerateUId() UId {
	return UId(uuid.NewString())
}

func NewUId(uid string) UId {
	return UId(uid)
}

func (u UId) IsEmpty() bool {
	return u.String() == ""
}

func (u UId) Equals(uid UId) bool {
	return u.String() == uid.String()
}

func (u UId) IsValidUUID() bool {
	_, err := uuid.Parse(u.String())
	return err == nil
}
