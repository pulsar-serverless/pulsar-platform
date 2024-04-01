package billing

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/billing"

	"github.com/labstack/echo/v4"
)

// @Summary	Get Pricing Plans
// @ID			get-pricing-plans
// @Accept		json
// @Produce	json
// @Success	200		{object} any
// @Router		/api/projects/plans [get]
// @Param		pageNumber	query		int	true	"Page number"
// @Param		pageSize	query		int	true	"Page size"
// @Security	Bearer
// @Tags		Resources

func GetPricingPlans(billingApi billing.IBillingService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request billing.GetPlansReq

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		plans, err := billingApi.GetPricingPlans(context.TODO(), request)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(http.StatusOK, plans)
	}
}
