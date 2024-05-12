"use client";

import { yellow } from "@mui/material/colors";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import { ReactNode } from "react";

const theme = createTheme({
  palette: {
    mode: "dark",
    primary: {
      main: "#2563eb",
    },
    secondary: {
      main: "#ffffff",
    },
    warning: {
      main: yellow[500],
    },
    text: {
      primary: "#fff",
    },
    divider: "#ffffff24",
    background: {
      default: "#110f0f",
      paper: "#121212",
    },
  },
  typography: {
    fontFamily: "'Inter', sans-serif",
    fontSize: 13,
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: { textTransform: "none" },
      },
    },

    MuiListItemButton: {
      styleOverrides: {
        root: { borderRadius: 12 },
      },
    },
    MuiTab: {
      styleOverrides: {
        root: { textTransform: "none" },
      },
    },
    MuiPaper: {
      defaultProps: {
        elevation: 0,
        variant: "outlined",
      },
    },
  },
});

export const CustomThemeProvider = ({ children }: { children: ReactNode }) => {
  return <ThemeProvider theme={theme}>{children}</ThemeProvider>;
};
