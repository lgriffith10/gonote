package builders

import "github.com/lgriffith10/gonote/internal/dtos"

type LoginRequestBuilder struct {
	request *dtos.LoginRequest
}

func NewLoginRequestBuilder() *LoginRequestBuilder {
	return &LoginRequestBuilder{
		request: &dtos.LoginRequest{
			Email:    "test@example.com",
			Password: "password",
		},
	}
}

func (b *LoginRequestBuilder) WithEmail(email string) *LoginRequestBuilder {
	b.request.Email = email
	return b
}

func (b *LoginRequestBuilder) WithPassword(password string) *LoginRequestBuilder {
	b.request.Password = password
	return b
}

func (b *LoginRequestBuilder) Get() *dtos.LoginRequest {
	return b.request
}
