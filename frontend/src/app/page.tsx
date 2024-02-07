"use client";

import { Typography } from "@mui/material";
import Stack from "@mui/material/Stack";
import Button from "@mui/material/Button";
import { useAuth0 } from "@auth0/auth0-react";
import React from "react";

export default function Home() {
  const { loginWithRedirect, isAuthenticated, logout, getAccessTokenSilently } =
    useAuth0();

  return (
    <Stack spacing={2} direction="column" width={100}>
      <Typography variant="h4">Welcome to pulsar</Typography>
      {isAuthenticated ? (
        <Button
          variant="outlined"
          onClick={() =>
            logout({ logoutParams: { returnTo: window.location.origin } })
          }
        >
          Sign Out
        </Button>
      ) : (
        <Button variant="outlined" onClick={() => loginWithRedirect()}>
          Sign In
        </Button>
      )}
    </Stack>
  );
}
