"use client";

import Stack from "@mui/material/Stack";
import {
  Typography,
  Container,
  Button,
} from "@mui/material";
import { useAuth0 } from "@auth0/auth0-react";
import React from "react";
import GitHubIcon from "@mui/icons-material/GitHub";
import { useEffect } from "react";
import { useRouter } from "next/navigation";

export default function Home() {
  const { user, loginWithRedirect, isAuthenticated, isLoading } = useAuth0();

  const router = useRouter();

  useEffect(() => {
    if (isAuthenticated && !isLoading) {
      router.push(`/${user?.name}`);
    }
  }, [isAuthenticated]);

  return (
    <Container>
      <Stack
        spacing={3}
        alignItems="center"
        justifyContent="center"
        sx={{ height: "100%" }}
      >
        <Typography variant="h4" component="h1" sx={{ mb: 2 }}>
          Login to Pulsar
        </Typography>
        <Button
          variant="contained"
          size="large"
          color="primary"
          startIcon={<GitHubIcon />}
          onClick={() => loginWithRedirect()}
        >
          Continue with Github
        </Button>
      </Stack>
    </Container>
  );
}
