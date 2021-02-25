package users

import "encoding/json"

// PublicUser struct
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// PrivateUser struct
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// Marshall an user
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	userJSON, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}

// Marshall users
func (users *Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(*users))
	for index, user := range *users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}
