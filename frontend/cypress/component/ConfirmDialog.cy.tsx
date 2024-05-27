import { AccountAlert } from "@/components/AccountAlert";
import { ConfirmationDialog } from "@/components/modals/ConfirmationDialog";
import { AuthProvider } from "@/components/providers/AuthProvider";
import { CustomThemeProvider } from "@/components/providers/ThemeProvider";
import { CssBaseline, Stack } from "@mui/material";
import React from "react";

describe("<ConfirmDialog />", () => {
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
            <ConfirmationDialog
              open={true}
              title={"Test"}
              description={"Testing Dialog"}
              handleClose={function (): void {
                throw new Error("Function not implemented.");
              }}
              handleConfirm={function (): void {
                throw new Error("Function not implemented.");
              }}
            />
          </Stack>
        </CssBaseline>
      </CustomThemeProvider>
    );

    cy.contains("Test");
    cy.contains("Testing Dialog");
  });
});
