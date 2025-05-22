package builders

import (
	"github.com/google/uuid"
	"github.com/lgriffith10/gonote/internal/dtos"
)

type createClassRequestBuilder struct {
	request *dtos.CreateClassRequest
}

func NewCreateClassRequestBuilder() *createClassRequestBuilder {
	return &createClassRequestBuilder{
		request: &dtos.CreateClassRequest{
			Name:      "Class name",
			Students:  []uuid.UUID{},
			ManagerId: uuid.New(),
		},
	}
}

func (c *createClassRequestBuilder) WithName(name string) *createClassRequestBuilder {
	c.request.Name = name
	return c
}

func (c *createClassRequestBuilder) WithStudents(students []uuid.UUID) *createClassRequestBuilder {
	c.request.Students = students
	return c
}

func (c *createClassRequestBuilder) WithManagerId(managerId uuid.UUID) *createClassRequestBuilder {
	c.request.ManagerId = managerId
	return c
}

func (c *createClassRequestBuilder) Get() *dtos.CreateClassRequest {
	return c.request
}
