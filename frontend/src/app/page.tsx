"use client";

import Stack from "@mui/material/Stack";
import { Typography, Container, Button, Grid, Box } from "@mui/material";
import { User, useAuth0 } from "@auth0/auth0-react";
import React from "react";
import GitHubIcon from "@mui/icons-material/GitHub";
import { useEffect } from "react";
import { useRouter } from "next/navigation";

export default function Home() {
  const { isAuthenticated, isLoading, user, loginWithRedirect } = useAuth0<
    User & { roleType: string[] }
  >();
  const router = useRouter();

  useEffect(() => {
    if (isAuthenticated && !isLoading && user) {
      if (user.roleType.includes("Admin")) router.push("/users");
      else router.push(`/${user.sub}`);
    }
  }, [isAuthenticated, isLoading, router, user]);

  return (
    <Container>
      <Box
        sx={{
          height: "100%",
          width: "100%",
          display: "grid",
          placeItems: "center",
        }}
      >
        <Grid container spacing={6} sx={{ alignItems: "center" }}>
          <Grid item xs={12} md={6}>
            <Typography variant="h2" gutterBottom fontWeight={"bold"}>
              Serverless <br /> That Fits in Your Pocket.
            </Typography>
            <Typography variant="h5" fontWeight={'light'}>
              Deploy code instantly without the weight of traditional server
              management.
            </Typography>
          </Grid>
          <Grid item xs={12} md={6} sx={{}}>
            <Stack alignItems={"center"}>
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
          </Grid>
        </Grid>
      </Box>
    </Container>
  );
}
