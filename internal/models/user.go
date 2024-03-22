package models

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
}

type UserCreate struct {
	Name     string `json:"name"`
	Division string `json:"division"`
}

type UserUpdate struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
}

type UserGet struct {
	ID int `json:"id"`
}

type UserDelete struct {
	ID int `json:"id"`
}
