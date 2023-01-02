package pkg

import (
	"errors"
	"job-board-api/cmd"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Token struct {
	Hash   string `json:"token"`
	Expire int64  `json:"-"`
	Secret string `json:"-"`
}

func (t *Token) CreateToken(c *fiber.Ctx, userID uint, secret string, exp int) (*Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	t.Expire = int64(exp)
	expiresIn := time.Now().Add(time.Duration(t.Expire) * time.Second).Unix()
	claims["exp"] = expiresIn

	tokenHash, err := token.SignedString([]byte(secret))

	if err != nil {
		return nil, err
	}
	c.Cookie(&fiber.Cookie{
		Name:     "Verify-Rest-Token",
		Value:    tokenHash,
		Secure:   false,
		HTTPOnly: true,
	})
	t.Hash = tokenHash
	t.Expire = expiresIn
	return t, nil
}

func (t *Token) ParseToken(c *fiber.Ctx, secret string) (uint, error) {
	tokenString := c.Cookies("Verify-Rest-Token")

	if tokenString == "" {
		return 0, errors.New("empty auth cookie")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	err2 := claims.Valid()

	if err2 != nil {
		t.DeleteToken(c)
		return 0, err2
	}

	return uint(claims["id"].(float64)), nil
}

func (t *Token) DeleteToken(c *fiber.Ctx) {
	c.ClearCookie("Verify-Rest-Token")
}

func (t *Token) RefreshToken(c *fiber.Ctx, secret string) (*Token, error) {
	u, err := t.ParseToken(c, secret)

	if err != nil {
		return nil, nil
	}

	return t.CreateToken(c, u, secret, cmd.Http.Jwt.Expire)
}
