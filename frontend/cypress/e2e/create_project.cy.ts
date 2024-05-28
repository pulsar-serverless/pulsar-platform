describe("Create Project Flow", () => {
  it("successfully creates a project on dash page", () => {
    // Visit home page
    cy.visit("http://localhost:1324");

    // Redirect to OAuth page
    cy.origin(Cypress.env("auth0_domain"), () => {
      // click continue with github button
      cy.contains("Continue with Github").click();

      // redirect after authenticating with github
      cy.url().should("equal", Cypress.env("auth0_redirect_uri"));

      cy.contains("Recent Projects");

      // find create project button
      cy.contains("Create a Project").click();

      // get the input and type 'new-project'
      cy.get("input").type("new-project");

      // click on create button
      cy.contains('Create').click()

      // verify it's created
      cy.contains('new-project')
    });
  });
});
