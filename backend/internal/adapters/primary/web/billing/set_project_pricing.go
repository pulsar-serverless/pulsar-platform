package billing

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/billing"

	"github.com/labstack/echo/v4"
)

// @Summary		Set Project Pricing Plan
// @ID			set-project-pricing-plan
// @Accept		json
// @Produce		json
// @Success		201
// @Router		/api/projects/{projectId}/plan [post]
// @Param		projectId	path		string	true	"project id"
// @Param		planId		query		string	true	"Plan Id"
// @Security	Bearer
// @Tags		Billing
func SetProjectPricing(billingApi billing.IBillingService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request billing.SetPlanReq

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		err := billingApi.SetProjectPlan(context.TODO(), request)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.NoContent(http.StatusCreated)
	}
}
