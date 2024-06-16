import { PricingCardSkeleton } from "@/components/pricing/PricingCardSkeleton";

describe('PricingCardSkeleton component', () => {
    it('renders card with skeleton elements', () => {
      cy.mount(<PricingCardSkeleton />);
  
      cy.get('.MuiCard-root').should('be.visible');

      cy.get('.MuiSkeleton-text').should('have.length', 5);
      cy.get('.MuiList-root').should('be.visible');
      cy.get('.MuiListItem-root').should('have.length', 3);

    });
  });
  