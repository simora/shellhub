package user

import (
	"context"
	"testing"

	"github.com/shellhub-io/shellhub/api/store"
	"github.com/shellhub-io/shellhub/api/store/mocks"
	"github.com/shellhub-io/shellhub/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestUpdateDataUser(t *testing.T) {
	mock := &mocks.Store{}
	s := NewService(store.Store(mock))

	ctx := context.TODO()

	user := &models.User{Name: "name", Email: "email@email.com", Username: "username", Password: "password", TenantID: "tenant"}
	userChangedPassword := &models.User{Name: "name", Email: "", Username: "", Password: "newpassword", TenantID: "tenant"}

	alteredUser := &models.User{Name: "name", Email: "email@email.com", Username: "username", Password: "newpassword", TenantID: "tenant"}

	userChangedData := &models.User{Name: "rename", Email: "new@email.com", Username: "newusername", Password: "", TenantID: "tenant"}

	//Changed Password
	mock.On("GetUserByTenant", ctx, user.TenantID).Return(user, nil).Once()
	mock.On("UpdateUser", ctx, user.Username, user.Email, userChangedPassword.Password, user.TenantID).Return(nil).Once()

	err := s.UpdateDataUser(ctx, userChangedPassword.Username, userChangedPassword.Email, userChangedPassword.Password, userChangedPassword.TenantID)

	assert.NoError(t, err)
	mock.AssertExpectations(t)

	// changed username and email
	mock.On("GetUserByTenant", ctx, user.TenantID).Return(alteredUser, nil).Once()
	mock.On("UpdateUser", ctx, userChangedData.Username, userChangedData.Email, userChangedPassword.Password, userChangedData.TenantID).Return(nil).Once()

	err = s.UpdateDataUser(ctx, userChangedData.Username, userChangedData.Email, userChangedData.Password, userChangedData.TenantID)

}
