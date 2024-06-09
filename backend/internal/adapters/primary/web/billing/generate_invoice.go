package billing

import (
	"net/http"
	"pulsar/internal/core/services/billing"

	"github.com/labstack/echo/v4"
)

// @Summary		Generate Month Invoice
// @ID			generate-project-month-invoice
// @Accept		json
// @Produce		json
// @Success		200 {object} any
// @Router		/api/projects/{projectId}/invoice [post]
// @Param		projectId	path		string	true	"project id"
// @Param		month		query		string	true	"Month"
// @Security	Bearer
// @Tags		Billing
func GenerateInvoice(
	billingApi billing.IBillingService,
) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, "")
	}
}
