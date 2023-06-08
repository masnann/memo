package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct{}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R *gin.Engine
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{} // currently has no properties

	// Create an account group
	g := c.R.Group(os.Getenv("ACCOUNT_API_URL"))

	g.GET("/me", h.Me)
	g.POST("/signup", h.SignUp)
	g.POST("/signin", h.SignIn)
	g.POST("/signout", h.SignOut)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)

}

func (h *Handler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

func (h *Handler) SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "its sign up",
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "its SignIn",
	})
}

func (h *Handler) SignOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "its SignOut",
	})
}

func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "its Tokens",
	})
}

func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "its Image",
	})
}

func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "its DeleteImage",
	})
}

func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "its Details",
	})
}
