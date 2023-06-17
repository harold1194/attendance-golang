package migrations

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/harold/attendance-golang/database/models"
	"gorm.io/gorm"
)

type User struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) Register(context *fiber.Ctx) error {
	user := User{}

	err := context.BodyParser(&user)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&user).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not register the user"})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Successfully Registered",
	})
	return nil
}

func (r *Repository) GetUsers(context *fiber.Ctx) error {
	userModel := &[]models.User{}

	err := r.DB.Find(userModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get users data"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user fetch successfully",
		"data":    userModel,
	})
	return nil
}

func (r *Repository) GetUserByID(context *fiber.Ctx) error {
	id := context.Params("id")
	userModel := &models.User{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be found",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(userModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get user",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user id successfully fetched",
		"data":    userModel,
	})
	return nil
}

func (r *Repository) DeleteUser(context *fiber.Ctx) error {
	userModel := models.User{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	err := r.DB.Delete(userModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete user"})
		return nil
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user delete successfully",
	})
	return nil
}

// GenerateJWT generates a JSON Web Token (JWT) with the provided user ID.
func GenerateJWT(userID uint) (string, error) {
	// Define the claims for the JWT
	claims := jwt.MapClaims{
		"user_id": userID,
		// Add other desired claims...
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate the token string
	jwtSecret := []byte("your_secret_key") // Replace with your secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (r *Repository) Login(context *fiber.Ctx) error {
	user := User{}

	err := context.BodyParser(&user)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	// Perform authentication logic
	// Example:
	userModel := &models.User{}
	err = r.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(userModel).Error
	if err != nil {
		context.Status(http.StatusUnauthorized).JSON(
			&fiber.Map{"message": "invalid email or password"})
		return err
	}

	// Assuming authentication is successful, generate a token
	// Example:
	token, err := GenerateJWT(userModel.ID)
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "failed to generate token"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login successful",
		"token":   token,
	})
	return nil
}
