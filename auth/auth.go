package auth

type Service struct {
	BaseUrl string
}

type Auth struct {
	JWT   string `json:"jwt"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Meta MetaLoginResponse `json:"meta"`
}

type MetaLoginResponse struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refreshToken"`
	User         *UserMeta `json:"user"`
}

type UserMeta struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"username"`
}

func (s *Service) Login() (*Auth, error) {

}
