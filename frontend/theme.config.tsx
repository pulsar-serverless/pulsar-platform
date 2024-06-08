import React from "react";
import { DocsThemeConfig } from "nextra-theme-docs";
import { Button, Stack, Typography } from "@mui/material";
import Link from "next/link";
import { Logo } from "@/components/layout/Logo";

const config: DocsThemeConfig = {
  logo: (
    <Button color="primary" style={{textDecoration: 'none'}}>
      <Stack direction={"row"} alignItems={"center"} gap={2}>
        <Logo />
        <Typography variant="h5">Pulsar</Typography>
      </Stack>
    </Button>
  ),
  project: {
    link: "https://github.com/pulsar-serverless/pulsar-platform",
  },
  docsRepositoryBase: "https://github.com/pulsar-serverless/pulsar-platform",
  footer: {
    text: "Pulsar Documentation",
  },
};

export default config;
