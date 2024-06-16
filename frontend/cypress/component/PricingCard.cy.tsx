import PricingCard from "@/components/pricing/PricingCard";
import { PricingPlan } from "@/models/PricingPlan";
import bytes from "bytes";

describe('PricingCard component', () => {
    it('renders card content with plan details', () => {
      const mockPlan: PricingPlan = {
        id: 'basic-plan',
        name: 'Starter',
        desc: 'Ideal for low-traffic applications',
        tier: 'basic',
        allocatedMemory: 536870912, 
        allocatedBandwidth: 1073741824, 
        allocatedRequests: 10000,
        price: 10,
      };
  
      cy.mount(<PricingCard plan={mockPlan} />);
  
      cy.get('.MuiTypography-subtitle1').should('contain.text', mockPlan.name);
  
      cy.get('.MuiTypography-h5').should('contain.text', `${mockPlan.price} ETB/month`);
  
      cy.get('.MuiTypography-body2').should('contain.text', mockPlan.desc);
  
      cy.get('.MuiList-root').should('be.visible');
  
      cy.get('.MuiListItemText-primary').first().should('contain.text', 'Bandwidth');
      cy.get('.MuiListItemText-secondary').first().should('contain.text', `${bytes(mockPlan.allocatedBandwidth)} per month`);
  
      cy.get('.MuiListItemText-primary').eq(1).should('contain.text', 'Memory');
      cy.get('.MuiListItemText-secondary').eq(1).should('contain.text', `${bytes(mockPlan.allocatedMemory)} per month`);

    });
  });