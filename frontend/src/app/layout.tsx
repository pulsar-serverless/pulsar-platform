import { QueryProvider } from "@/components/providers/QueryProvider";
import CssBaseline from "@mui/material/CssBaseline";
import { CustomThemeProvider } from "@/components/providers/ThemeProvider";
import { StoreProvider } from "@/components/providers/StoreProvider";
import { AuthProvider } from "@/components/providers/AuthProvider";
import { Header } from "@/components/layout/Header";
import { Alert, Box, Stack } from "@mui/material";
import { HttpInterceptor } from "@/components/interceptors/HttpInterceptor";
import { SnackbarProvider } from "@/components/providers/SnackbarProvider";
import { AccountAlert } from "@/components/AccountAlert";

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
          <HttpInterceptor>
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
                      overflowY: "scroll",
                    }}
                  >
                    <AccountAlert />
                    {children}
                  </Stack>
                </QueryProvider>
                <SnackbarProvider />
              </CustomThemeProvider>
            </StoreProvider>
          </HttpInterceptor>
        </AuthProvider>
      </body>
    </html>
  );
}
