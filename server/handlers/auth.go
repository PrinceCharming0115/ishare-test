package handlers

import (
	"ishare-test/responses"
	s "ishare-test/server"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type HandlerAuth struct {
	Server       *s.Server
	OAuth2Config oauth2.Config
}

func NewHandlerAuth(server *s.Server) *HandlerAuth {
	return &HandlerAuth{
		Server: server,
		OAuth2Config: oauth2.Config{
			RedirectURL:  server.Config.ClientCallbackUrl,
			ClientID:     server.Config.ClientID,
			ClientSecret: server.Config.ClientSecret,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.google.com/o/oauth2/auth",
				TokenURL: "https://accounts.google.com/o/oauth2/token",
			},
		},
	}
}

func (h *HandlerAuth) Login(c *fiber.Ctx) error {
	url := h.OAuth2Config.AuthCodeURL("state")
	return c.Redirect(url)
}

func (h *HandlerAuth) Callback(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := h.OAuth2Config.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Failed to exchange token")
	}

	return responses.Response(c, fiber.StatusOK, fiber.Map{
		"access_token": token.AccessToken,
	})
}
