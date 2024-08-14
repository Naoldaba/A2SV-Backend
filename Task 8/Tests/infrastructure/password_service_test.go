package infrastructure_test


import (
	"testing"

	"golang.org/x/crypto/bcrypt"
	"github.com/stretchr/testify/suite"
	"task_manager_api/Domain"
	"task_manager_api/Infrastructure"
)

type PasswordHashingTestSuite struct {
	suite.Suite
}

func (suite *PasswordHashingTestSuite) TestHashPasswordSuccess() {
	user := &domain.User{Password: "securepassword"}

	hashedPassword, err := infrastructure.PasswordHasher(user)
	suite.NoError(err, "Hashing password should not produce an error")
	suite.NotEmpty(hashedPassword, "Hashed password should not be empty")

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	suite.NoError(err, "Hashed password should match the original password")
}

func (suite *PasswordHashingTestSuite) TestHashPasswordErrorHandling() {
	user := &domain.User{Password: ""}

	hashedPassword, err := infrastructure.PasswordHasher(user)
	suite.NoError(err, "Hashing an empty password should not produce an error")
	suite.NotEmpty(hashedPassword, "Hashed password should not be empty even for an empty input")
}

func TestPasswordHashingTestSuite(t *testing.T) {
	suite.Run(t, new(PasswordHashingTestSuite))
}
