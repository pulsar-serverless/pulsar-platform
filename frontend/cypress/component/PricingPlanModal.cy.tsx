import PricingPlanModal from "@/components/pricing/PricingPlanModal";

describe('PricingPlanModal component', () => {
  it('renders modal with basic structure', () => {

    const mockProps = {
      isOpen: true,
      onClose: () => {}, 
      planId: 'some-plan-id',
    };

    cy.mount(<PricingPlanModal {...mockProps} />);

    cy.get('.MuiDialog-root').should('be.visible');

    cy.get('.MuiDialogTitle-root').should('be.visible');

    cy.get('.MuiDialogContent-root').should('be.visible');

  });
});
