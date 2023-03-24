package dto

type GetAllUser struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	RoleId    uint   `json:"roleId" validate:"required"`
}

type CreateUser struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	RoleId    uint   `json:"roleId" validate:"required"`
}

type UpdateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	RoleId    uint   `json:"roleId"`
}

type UpdateUserRole struct {
	Name string `json:"name" validate:"required"`
}
