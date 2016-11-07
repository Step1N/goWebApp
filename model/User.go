package model

import "fmt"

//FBUser user details
type FBUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Interest string `json:"interest"`
	Likes    string `json:"likes,omitempty"`
}

func (u FBUser) String() string {
	return fmt.Sprintf("User<%d %s %v %v>", u.ID, u.Name, u.Interest, u.Likes)
}
