package billing

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/billing"

	"github.com/labstack/echo/v4"
)

// @Summary	Generate Month Invoice
// @ID			generate-project-month-invoice
// @Accept		json
// @Produce	json
// @Success	200	{object}	GenerateInvoiceResp
// @Router		/api/projects/{projectId}/invoice [post]
// @Param		projectId	path	string	true	"project id"
// @Param		month		query	string	true	"Month"
// @Security	Bearer
// @Tags		Billing
func GenerateInvoice(
	billingApi billing.IBillingService,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectId := c.Param("projectId")
		month := c.QueryParam("month")

		if projectId == "" || month == "" {
			return c.NoContent(http.StatusNoContent)
		}

		invoice, err := billingApi.GenerateInvoice(context.TODO(), billing.GenerateInvoiceReq{ProjectID: projectId, Month: month})
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.File(invoice.FilePath)
	}
}
