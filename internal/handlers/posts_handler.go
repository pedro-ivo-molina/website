package handlers

import(
	"net/http"

	"website/internal/dto"
	"website/internal/repository"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	Repo *repository.PostRepository
}

func (h *PostHandler) Create(c *gin.Context) {
	var postDto dto.CreatePostDto

	if err := c.ShouldBindBodyWithJSON(&postDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required field missing: " + err.Error()})

		return
	}

	ctx := c.Request.Context()
	id, err = h.Repo.Create(ctx, postDto.Title, postDto.Content)
	if err != {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
		"message": "Post created"
	})
}
