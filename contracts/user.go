package contracts

type User interface {
	Create(UserCreateInput) (*UserCreateOutput, error)
	AcquireAccessTokenUsingEmailAndPassword(AcquireAccessTokenUsingEmailAndPasswordInput) (*AcquireAccessTokenOutput, error)
}

type UserCreateInput struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsSysAdmin bool   `json:"is_sysadmin"`
}

type UserCreateOutput struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

type AcquireAccessTokenUsingEmailAndPasswordInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AcquireAccessTokenOutput struct {
	Token string `json:"token" db:"token"`
}
