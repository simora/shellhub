package routes

import (
	"net/http"

	"crypto/sha256"
	"encoding/hex"
	"github.com/shellhub-io/shellhub/api/apicontext"
	"github.com/shellhub-io/shellhub/api/user"
)

const (
	UpdateUserURL = "/user"
)

func UpdateUser(c apicontext.Context) error {

	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return err
	}

	tenant := ""
	if v := c.Tenant(); v != nil {
		tenant = v.ID
	}
	if req.Password != "" {
		sum := sha256.Sum256([]byte(req.Password))
		var sum_byte []byte = sum[:]
		req.Password = hex.EncodeToString(sum_byte)
	}

	svc := user.NewService(c.Store())

	if err := svc.UpdateDataUser(c.Ctx(), req.Username, req.Email, req.Password, tenant); err != nil {
		if err == user.ErrUnauthorized {
			return c.NoContent(http.StatusForbidden)
		}

		return err
	}

	return c.JSON(http.StatusOK, nil)
}
