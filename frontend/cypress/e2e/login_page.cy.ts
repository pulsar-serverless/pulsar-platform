describe("Authentication Page Redirect", () => {
  it("successfully redirects user to Github OAuth", () => {
    // Visit home page
    cy.visit("http://localhost:1324");

    // Redirect to OAuth page
    cy.origin(Cypress.env("auth0_domain"), () => {
      // click continue with github button
      cy.contains("Continue with Github").click();

      // redirect after authenticating with github
      cy.url().should("equal", Cypress.env("auth0_redirect_uri"));
    });
  });
});
