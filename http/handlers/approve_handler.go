package handlers

import (
	"github.com/k1rend/RAMS/approve"
	"github.com/labstack/echo/v4"
)

type ApproveHandler struct {
	ApprovalService *approve.ApprovalService
}

func NewApproveHandler(approvalService *approve.ApprovalService) *ApproveHandler {
	return &ApproveHandler{
		ApprovalService: approvalService,
	}
}