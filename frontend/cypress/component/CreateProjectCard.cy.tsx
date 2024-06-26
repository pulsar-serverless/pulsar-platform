import CreateProjectCard from "@/components/project/CreateProjectCard";
import { CustomThemeProvider } from "@/components/providers/ThemeProvider";
import { CssBaseline, Stack } from "@mui/material";
import React from "react";

describe("<CreateProjectCard />", () => {
  it("renders", () => {
    cy.mount(
      <CustomThemeProvider>
        <CssBaseline>
          <Stack
            direction={"row"}
            sx={{
              position: "fixed",
              top: 64,
              left: 0,
              right: 0,
              bottom: 0,
              overflowY: "scroll",
            }}
          >
            <CreateProjectCard />
          </Stack>
        </CssBaseline>
      </CustomThemeProvider>
    );

    cy.contains("Create a new project");
    // input must exist
    cy.get("input").should("exist");

    // action buttons exist
    cy.get("button").contains("Create").should("exist");
    cy.get("button").contains("Cancel").should("exist");
  });
});
