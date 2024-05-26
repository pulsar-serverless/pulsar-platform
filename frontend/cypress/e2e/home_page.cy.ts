describe("Home Page Loads", () => {
  it("successfully loads the default home page", () => {
    cy.visit("http://localhost:1324");

    cy.contains("Pulsar");
    cy.contains("Docs");
    cy.contains("Pricing");
  });
});
