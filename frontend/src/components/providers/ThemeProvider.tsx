"use client";

import { ThemeProvider, createTheme } from "@mui/material/styles";
import { ReactNode } from "react";

const theme = createTheme({
  palette: {
    primary: {
      main: "#2563eb",
    },
    secondary: {
      main: "#ffffff",
    },
    text: {
      primary: "#fff",
    },
    background: {
      default: "#110f0f",
      paper: "#121212",
    },
  },
});

export const CustomThemeProvider = ({ children }: { children: ReactNode }) => {
  return <ThemeProvider theme={theme}>{children}</ThemeProvider>;
};
