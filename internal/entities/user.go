package entities

type User struct {
	DefaultModel
	Address     string `json:"address" gorm:"type:VARCHAR(255) NOT NULL;uniqueIndex"`
	NetworkName string `json:"network_name,omitempty" gorm:"size:32"`
	NetworkURL  string `json:"network_url,omitempty" gorm:"size:512"`
	ChainID     string `json:"chain_id" gorm:"type:VARCHAR(255) NOT NULL"`
	RefreshKey  string `json:"refresh_key,omitempty" gorm:"size:512"`
}
type UserDto struct {
	User
}

func NewDtO(u User) IEntity {
	return &UserDto{
		User: u,
	}
}
func (u UserDto) ToEntity() (interface{}, error) {
	return &u.User, nil
}
func (u UserDto) IsValid() bool {
	return true
}
