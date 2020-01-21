package slack

type AuthService struct {
	api *SlackClient
}

type User struct {
	Id 	string  `json:"id"`
	Name string `json:"name"`
}
type Auth struct {
	User User	`json:"user"`
}

func (s *AuthService) Test() (*Auth, error) {

	req, _ := s.api.NewRequest(_GET, "users.identity", nil)

	auth := new(Auth)

	_, err := s.api.Do(req, auth)

	if err != nil {
		return nil, err
	}

	return auth, nil
}
