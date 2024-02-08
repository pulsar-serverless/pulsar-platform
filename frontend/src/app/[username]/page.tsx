"use client";

import { Container, Typography } from "@mui/material";
import { useRouter } from "next/navigation";
import { useAuth0 } from "@auth0/auth0-react";
import { useEffect } from "react";

export default function Page() {
  const { isAuthenticated, isLoading } = useAuth0();
  const router = useRouter()

  useEffect(() => {
    if (!isAuthenticated && !isLoading) {
      router.push("/")
    }
  });

  return (
    <Container>
      <Typography>Projects</Typography>
    </Container>
  );
}
