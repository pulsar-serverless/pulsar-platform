import LogTable  from "@/components/log/LogTable";
import { Log } from "@/models/log";


describe("LogTable", () => {
  it("renders logs and pagination", () => {
    const mockLogs: Log[] = [
        { id: "1", type: "Info", message: "Informational message", createdAt: new Date("2024-05-27T00:00:00.000Z").toDateString() },
        { id: "2", type: "Error", message: "Error message", createdAt: new Date("2024-05-26T23:59:59.000Z").toDateString() },
    ];

    cy.mount(
      <LogTable logs={mockLogs} page={1} count={10} onPaginate={() => {}} /> // Mock onPaginate function
    );

    cy.get("table tbody").children().should("have.length", 2);
    cy.get("table tbody").within(() => {
      cy.get("tr").each(($row) => {
        cy.wrap($row).find("th p").should("exist"); 
        cy.wrap($row).find("td p").should("exist"); 
      });
    });
  });
});
