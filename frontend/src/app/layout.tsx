import { QueryProvider } from "@/components/providers/QueryProvider";
import { store } from "@/store/store";
import { Provider } from "react-redux";
import CssBaseline from "@mui/material/CssBaseline";
import { CustomThemeProvider } from "@/components/providers/ThemeProvider";
import { StoreProvider } from "@/components/providers/StoreProvider";
import { AuthProvider } from "@/components/providers/AuthProvider";

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
              <QueryProvider>{children}</QueryProvider>
              <CssBaseline />
            </CustomThemeProvider>
          </StoreProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
