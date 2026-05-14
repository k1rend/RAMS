package handlers

import (
	"github.com/k1rend/RAMS/http/dto"
	"github.com/k1rend/RAMS/resource"
	"github.com/labstack/echo/v4"
)

type ResourceHandler struct {
	ResourceService *resource.ResourceService
}

func NewResourceHandler(resourceService *resource.ResourceService) *ResourceHandler {
	return &ResourceHandler{
		ResourceService: resourceService,
	}
}


func (h *ResourceHandler) CreateResource(c echo.Context) error { // add owner_id to access grants table
	var req dto.CreateResourceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}
	uid := c.Get("user_id").(int32)
	return h.ResourceService.CreateResource(c.Request().Context(), req.Name, req.Description, req.Resource, req.ResourceType, uid)

}


func (h *ResourceHandler) GetAvailableResources(c echo.Context) { // return GetAvailableResourcesResponse
	// needs access_grants table in db
}

func (h *ResourceHandler) ListResources(c echo.Context) { // return ListResourcesResponse
	resources, err := h.ResourceService.ListResources(c.Request().Context())
	if err != nil {
		c.JSON(500, map[string]string{"error": "Failed to list resources"})
		return
	}

	var response dto.ListResourcesResponse
	for _, r := range resources {
		response.Resources = append(response.Resources, dto.BaseResourceInfo{
			ID:           r.ID,
			Name:         r.Name,
			Description:  r.Description,
			ResourceType: r.ResourceType,
		})
	}
	c.JSON(200, response)
}

func (h *ResourceHandler) UpdateResource(c echo.Context) error { // needs access grants check
	// needs access_grants table in db
	return nil
}

func (h *ResourceHandler) DeleteResource(c echo.Context) error { // needs access grants check
	// needs access_grants table in db
	return nil
}