package dto

// DTOs for authentication requests and responses
// Auth DTOs
type RegisterRequest struct {
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	DepartmentId int     `json:"department_id"`
	ManagerId    int     `json:"manager_id"`
	Roles        []int32 `json:"roles_id_array"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Resource DTOs

type Resource struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Resource     string `json:"resource"`
	ResourceType string `json:"resource_type"`
	OwnerID      int32  `json:"owner_id"`
	CreatedAt    string `json:"created_at"`
}

type BaseResourceInfo struct {
	ID           int32  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ResourceType string `json:"resource_type"`
}

type CreateResourceRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Resource     string `json:"resource"`
	ResourceType string `json:"resource_type"`
} // owner_id will be determined from JWT claims

type GetAvailableResourcesResponse struct {
	Resources []Resource `json:"resources"`
}

type ListResourcesResponse struct {
	Resources []BaseResourceInfo `json:"resources"`
}

// Requests DTOs

type CreateRequest struct {
	ResourceID   int32  `json:"resource_id"`
	AccessType   string `json:"access_type"` //"read", "write", "create", "delete"
	AccessReason string `json:"access_reason"`
}

type ListUserRequestsResponse struct {
	Requests []RequestInfo `json:"requests"`
}

type RequestInfo struct {
	ID           int32  `json:"id"`
	ResourceID   int32  `json:"resource_id"`
	ResourceName string `json:"resource_name"`
	AccessType   string `json:"access_type"`
	AccessReason string `json:"access_reason"`
	Status       string `json:"status"`
	Comments     string `json:"comments"`
	CreatedAt    string `json:"created_at"`
}