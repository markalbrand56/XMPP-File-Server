package controllers

import (
	"XMPP-File-Server/internal/configs"
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
// @Summary Sube un archivo a S3 y lo asocia al usuario que lo subió
// @Description Sube un archivo a S3 y lo asocia al usuario que lo subió.
// @Tags Files
// @Accept  mpfd
// @Produce  json
// @Param directory path string true "Directorio donde se guardará el archivo"
// @Param files formData file true "Archivo a subir"
// @Success 200 {object} responses.UploadSuccessResponse
// @Failure 400 {object} responses.StandardResponse
// @Failure 500 {object} responses.StandardResponse
// @Router /{directory} [post]
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

	paths := make([]string, 0)

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

		paths = append(paths, fmt.Sprintf("%s/files/%s/%s", configs.URL, directory, file.Filename))
	}

	c.JSON(http.StatusOK, responses.UploadSuccessResponse{
		StandardResponse: responses.StandardResponse{
			Code:    http.StatusOK,
			Message: "File uploaded successfully",
		},
		Paths: paths,
	})
}

// GetFile obtiene un archivo de S3
// @Summary Obtiene un archivo de S3
// @Description Obtiene un archivo de S3
// @Tags Files
// @Accept  json
// @Produce  mpfd
// @Param directory path string true "Directorio donde se encuentra el archivo"
// @Param file path string true "Nombre del archivo"
// @Success 200 {object} responses.StandardResponse
// @Failure 400 {object} responses.StandardResponse
// @Failure 500 {object} responses.StandardResponse
// @Router /files/:directory/:file [get]
func GetFile(c *gin.Context) {
	directory := c.Param("directory")
	name := c.Param("file")

	if directory == "" || name == "" {
		c.JSON(http.StatusBadRequest, responses.StandardResponse{
			Code:    http.StatusBadRequest,
			Message: "Directory and name are required",
		})
		return
	}

	path := fmt.Sprintf("%s/%s", directory, name)

	file, err := database.Instance.GetFile(path)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.StandardResponse{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Error getting file: %s", err.Error()),
		})
		return
	}

	contentType := http.DetectContentType(file.Bytes())

	//c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	c.Header("Content-Type", contentType)
	c.Header("Content-Length", fmt.Sprintf("%d", len(file.Bytes())))
	c.Data(http.StatusOK, contentType, file.Bytes())
}
