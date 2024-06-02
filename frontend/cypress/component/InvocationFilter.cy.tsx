import InvocationFilter from "@/components/analytics/InvocationFilter";
import React from "react";

describe("<InvocationFilter />", () => {
  it("renders", () => {
    cy.mount(
      <InvocationFilter
        status={"Success"}
        setStatus={function (status: InvocationStatus): void {
          throw new Error("Function not implemented.");
        }}
        graphType={"Day"}
        setGraphType={function (graphType: InvocationGraphType): void {
          throw new Error("Function not implemented.");
        }}
      />
    );

    cy.contains("Success");
    cy.contains("Last 24 Hours").click().contains("Last");
  });
});
