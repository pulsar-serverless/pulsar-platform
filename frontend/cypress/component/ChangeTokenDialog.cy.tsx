import { CustomThemeProvider } from "@/components/providers/ThemeProvider";
import ChangeTokenDialog from "@/components/settings/ChangeTokenDialog";
import { CssBaseline, Stack } from "@mui/material";
import React from "react";

describe("<ChangeTokenDialog />", () => {
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
          ></Stack>
          <ChangeTokenDialog
            isOpen={true}
            onClose={function (): void {
              throw new Error("Function not implemented.");
            }}
            projectId={""}
          />
        </CssBaseline>
      </CustomThemeProvider>
    );
  });
});
