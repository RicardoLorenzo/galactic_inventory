package inventory

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/RicardoLorenzo/server-test/go/config"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidPage     = errors.New("invalid page")
	ErrInvalidPageSize = errors.New("invalid page size")
)

type Handler struct {
	store *Store
}

func NewHandler(c *config.Config) (*Handler, error) {
	s, err := NewStore(c)

	return &Handler{
		store: s,
	}, err

}

func (h *Handler) processError(c *gin.Context, returnCode int, err error) {
	c.JSON(returnCode, gin.H{
		"error": err.Error(),
	})

}

func (h *Handler) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.processError(c, http.StatusBadRequest, ErrInvalidPage)
	}

	pagesSize, err := strconv.Atoi(c.DefaultQuery("page_size", "100"))
	if err != nil {
		h.processError(c, http.StatusBadRequest, ErrInvalidPageSize)
	}

	spacecraft := &Spacecraft{
		Name:   c.Query("name"),
		Class:  c.Query("class"),
		Status: c.Query("status"),
	}

	list, err := h.store.List(page, pagesSize, spacecraft)
	if err != nil {
		h.processError(c, http.StatusBadRequest, err)
	}

	c.JSON(200, gin.H{
		"data": list,
	})
}

func (h *Handler) Create(c *gin.Context) {
	spacecraft := &Spacecraft{
		Name:     c.PostForm("name"),
		Class:    c.PostForm("class"),
		Armament: c.PostForm("armament"),
		Crew:     c.PostForm("crew"),
		Image:    c.PostForm("image"),
		Value:    c.PostForm("value"),
		Status:   c.PostForm("status"),
	}

	err := h.store.Create(spacecraft)
	if err != nil {
		h.processError(c, http.StatusBadRequest, err)
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}

func (h *Handler) Update(c *gin.Context) {
	spacecraft := &Spacecraft{
		Name:     c.PostForm("name"),
		Class:    c.PostForm("class"),
		Armament: c.PostForm("armament"),
		Crew:     c.PostForm("crew"),
		Image:    c.PostForm("image"),
		Value:    c.PostForm("value"),
		Status:   c.PostForm("status"),
	}

	err := h.store.Update(spacecraft)
	if err != nil {
		h.processError(c, http.StatusBadRequest, err)
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	spacecraft := &Spacecraft{
		Name:     c.PostForm("name"),
		Class:    c.PostForm("class"),
		Armament: c.PostForm("armament"),
		Crew:     c.PostForm("crew"),
		Image:    c.PostForm("image"),
		Value:    c.PostForm("value"),
		Status:   c.PostForm("status"),
	}

	err := h.store.Delete(spacecraft)
	if err != nil {
		h.processError(c, http.StatusBadRequest, err)
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}
