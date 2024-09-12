package middlewares

import (
	"encoding/json"
	"fmt"
	"ishare-test/responses"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TokenInfo struct {
	Audience      string `json:"aud"`
	ExpiresIn     string `json:"expires_in"`
	Scope         string `json:"scope"`
	Email         string `json:"email"`
	VerifiedEmail string `json:"email_verified"`
}

func validateToken(token string) (*TokenInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://oauth2.googleapis.com/tokeninfo?access_token=%s", token))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid token")
	}

	var tokenInfo TokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return nil, err
	}

	return &tokenInfo, nil
}

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return responses.ErrorResponse(c, fiber.StatusUnauthorized, "Authorization header is missing")
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	tokenInfo, err := validateToken(token)
	if err != nil {
		return responses.ErrorResponse(c, fiber.StatusUnauthorized, "Access token is invalid")
	}

	if tokenInfo.VerifiedEmail != "true" {
		return responses.ErrorResponse(c, fiber.StatusUnauthorized, "Email address is not verified")
	}

	if tokenInfo.ExpiresIn == "0" {
		return responses.ErrorResponse(c, fiber.StatusUnauthorized, "Access token has expired")
	}

	return c.Next()
}
