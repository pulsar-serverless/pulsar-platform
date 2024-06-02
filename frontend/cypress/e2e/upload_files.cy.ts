describe("Upload Files Flow", () => {
  it("successfully uploads code", () => {
    // Visit home page
    cy.visit("http://localhost:1324");

    // ... authenticated already
    cy.contains("Recent Projects");

    // click on previously created project
    cy.contains("new-project").click();

    // click on deployment tab
    cy.contains("Deployment").click();

    // click on upload code button
    cy.contains("Upload").click();

    const filePath = "/Users/yoni/Desktop/test/test-app";
    // get input for file upload
    cy.get('input[type="file"]').selectFile(filePath);

    // click upload button
    cy.contains("Upload").click();
  });
});
