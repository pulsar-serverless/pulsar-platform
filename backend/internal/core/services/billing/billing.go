package billing

import (
	"context"
	"pulsar/internal/core/domain/billing"
	"pulsar/internal/core/domain/common"
	aS "pulsar/internal/core/services/analytics"
	"pulsar/internal/core/services/project"
	"pulsar/internal/ports"
)

type IBillingService interface {
	GetPlanResource(ctx context.Context, req GetPlanResReq) (*billing.PlanResources, error)
	GetPricingPlans(ctx context.Context, req GetPlansReq) (*common.Pagination[GetPlansResp], error)
	SetProjectPlan(ctx context.Context, req SetPlanReq) error
	CheckPlanLimit(ctx context.Context, projId, planId string) error
	GenerateInvoice(ctx context.Context, req GenerateInvoiceReq) (*GenerateInvoiceResp, error)
}

type BillingService struct {
	repo             ports.IBillingRepository
	fileRepo         ports.IFileRepository
	projectService   project.IProjectService
	analyticsService aS.IAnalyticsService
	resourceService  aS.IResourceService
}

func NewBillingService(repo ports.IBillingRepository, fileRepo ports.IFileRepository, projectService project.IProjectService, analyticsService aS.IAnalyticsService, resourceService aS.IResourceService) *BillingService {
	return &BillingService{
		repo:             repo,
		fileRepo:         fileRepo,
		projectService:   projectService,
		analyticsService: analyticsService,
		resourceService:  resourceService,
	}
}
