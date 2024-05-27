import { Header } from "@/components/layout/Header";
import { Logo } from "@/components/layout/Logo";
import React from "react";

describe("<Header />", () => {
  it("renders", () => {
    cy.mount(<Header />);

    cy.contains("Pulsar");
  });
});

describe("<Logo />", () => {
  it("renders", () => {
    cy.mount(<Logo />);
  });
});
