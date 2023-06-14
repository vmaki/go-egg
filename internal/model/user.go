package model

type User struct {
	BaseModel

	Name     string `json:"name,omitempty"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	CommonTimestampsField
}
