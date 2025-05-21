package builders

import "github.com/lgriffith10/gonote/internal/dtos"

type RegisterRequestBuilder struct {
	request *dtos.RegisterRequest
}

func NewRegisterRequestBuilder() *RegisterRequestBuilder {
	return &RegisterRequestBuilder{
		request: &dtos.RegisterRequest{
			Email:     "test@test.com",
			Password:  "password",
			Firstname: "Luciano",
			Lastname:  "Griffith",
		},
	}
}

func (b *RegisterRequestBuilder) WithEmail(email string) *RegisterRequestBuilder {
	b.request.Email = email
	return b
}

func (b *RegisterRequestBuilder) WithPassword(password string) *RegisterRequestBuilder {
	b.request.Password = password
	return b
}

func (b *RegisterRequestBuilder) WithFirstname(firstname string) *RegisterRequestBuilder {
	b.request.Firstname = firstname
	return b
}

func (b *RegisterRequestBuilder) WithLastname(lastname string) *RegisterRequestBuilder {
	b.request.Lastname = lastname
	return b
}

func (b *RegisterRequestBuilder) Get() *dtos.RegisterRequest {
	return b.request
}
