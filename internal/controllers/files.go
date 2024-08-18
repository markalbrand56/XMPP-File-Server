package controllers

import (
	"XMPP-File-Server/internal/database"
	"XMPP-File-Server/internal/responses"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type Form struct {
	Files []*multipart.FileHeader `form:"files" binding:"required"`
}

// Upload sube un archivo a S3 y lo asocia al usuario que lo subió
// @Summary (Requiere autentificación) Sube un archivo a S3 y lo asocia al usuario que lo subió
// @Description (Requiere autentificación) Sube un archivo a S3 y lo asocia al usuario que lo subió. Necesita un token de autentificación
// @Tags Documents
// @Accept  mpfd
// @Produce  json
// @Param documentType path string true "Tipo de documento"
// @Param files formData file true "Archivo a subir"
// @Success 200 {object} StandardResponse
// @Failure 400 {object} StandardResponse
// @Failure 401 {object} StandardResponse
// @Failure 500 {object} StandardResponse
// @Router /files/{directory} [post]
func Upload(c *gin.Context) {
	directory := c.Param("directory")

	if directory == "" {
		c.JSON(http.StatusBadRequest, responses.StandardResponse{
			Code:    http.StatusBadRequest,
			Message: "Directory is required",
		})
		return
	}

	var form Form

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, responses.StandardResponse{
			Code:    http.StatusBadRequest,
			Message: "File is required",
		})
		return
	}

	for _, file := range form.Files {
		// Upload to S3

		path := fmt.Sprintf("%s/%s", directory, file.Filename)

		err := database.Instance.Insert(path, file)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.StandardResponse{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("Error uploading file: %s", err.Error()),
			})
			return
		}
	}

	c.JSON(http.StatusOK, responses.StandardResponse{
		Code:    http.StatusOK,
		Message: "Files uploaded successfully",
	})
}
