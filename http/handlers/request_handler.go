package handlers

import (
	"github.com/k1rend/RAMS/http/dto"
	"github.com/k1rend/RAMS/request"
	"github.com/labstack/echo/v4"
)

type RequestHandler struct {
	RequestService *request.RequestService
}

func NewRequestHandler(requestService *request.RequestService) *RequestHandler {
	return &RequestHandler{
		RequestService: requestService,
	}
}

func (h *RequestHandler) CreateRequest(c echo.Context) error {
	var req dto.CreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}
	uid := c.Get("user_id").(int32)
	request, err := h.RequestService.CreateRequest(c.Request().Context(), uid, req.ResourceID, req.AccessType, req.AccessReason)	
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create request"})
	}
	return c.JSON(201, request)
}


func (h *RequestHandler) ListUserRequests(c echo.Context) error { 
	uid := c.Get("user_id").(int32)
	requests, err := h.RequestService.ListUserRequests(c.Request().Context(), uid)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to list user requests"})
	}

	var userRequests dto.ListUserRequestsResponse

	for _, r := range requests {
		userRequests.Requests = append(userRequests.Requests, dto.RequestInfo{
			ID:           r.ID,
			ResourceID:   r.ResourceID,
			ResourceName: r.ResourceName,
			AccessType:   r.AccessType,
			AccessReason: r.AccessReason,
			Status:       r.Status,
			Comments:     r.Comments.String, // handle NULL comments
			CreatedAt:    r.CreatedAt.Time.Format("2006-01-02"), // format date as string
		})
	}

	return c.JSON(200, userRequests)
}


func (h *RequestHandler) DeleteRequest(c echo.Context) error {
	return nil
}