"use client";
import InvocationGraph from "@/components/analytics/InvocationsGraph";
import { Container, Typography } from "@mui/material";
import { useParams } from "next/navigation";

function Page() {
  const { projectId } = useParams<{ projectId: string }>();

  return (
    <Container sx={{ py: 3 }} maxWidth="md">
      <Typography
        variant="h6"
        sx={{ textTransform: "capitalize" }}
        gutterBottom
      >
        Analytics
      </Typography>
      <Typography variant="body2">
        Effortlessly monitor executions and track errors.
      </Typography>

      <InvocationGraph projectId={projectId}/>
    </Container>
  );
}

export default Page;
