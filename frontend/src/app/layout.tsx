import { QueryProvider } from "@/components/QueryProvider";
import { store } from "@/store/store";
import { Provider } from "react-redux";
import CssBaseline from "@mui/material/CssBaseline";
import { CustomThemeProvider } from "@/components/ThemeProvider";
import { StoreProvider } from "@/components/StoreProvider";

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
        <StoreProvider>
          <CustomThemeProvider>
            <QueryProvider>{children}</QueryProvider>
            <CssBaseline />
          </CustomThemeProvider>
        </StoreProvider>
      </body>
    </html>
  );
}
