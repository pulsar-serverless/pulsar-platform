import { QueryProvider } from "@/components/providers/QueryProvider";
import { store } from "@/store/store";
import { Provider } from "react-redux";
import CssBaseline from "@mui/material/CssBaseline";
import { CustomThemeProvider } from "@/components/providers/ThemeProvider";
import { StoreProvider } from "@/components/providers/StoreProvider";
import { AuthProvider } from "@/components/providers/AuthProvider";
import { Header } from "@/components/layout/Header";
import { Stack } from "@mui/material";

export const metadata = {
  title: "Pulsar",
  description: "Serverless Web Platform",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <AuthProvider>
          <StoreProvider>
            <CustomThemeProvider>
              <CssBaseline />
              <QueryProvider>
                <Header />
                <Stack
                  direction={"row"}
                  sx={{
                    position: "fixed",
                    top: 64,
                    left: 0,
                    right: 0,
                    bottom: 0,
                    overflowY: 'scroll'
                  }}
                >
                  {children}
                </Stack>
              </QueryProvider>
            </CustomThemeProvider>
          </StoreProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
