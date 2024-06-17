"use client";

import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  Container,
  Divider,
  Stack,
  Typography,
} from "@mui/material";
import DownloadRoundedIcon from "@mui/icons-material/DownloadRounded";
import AddRoundedIcon from "@mui/icons-material/AddRounded";
import { useParams } from "next/navigation";
import { ProjectApi } from "@/api/projects";
import { useMutation } from "@tanstack/react-query";
import { UploadCode } from "@/components/deployment/UploadCode";
import { UploadAssets } from "@/components/deployment/UploadAssets";

export default function Page() {
  const { projectId } = useParams<{ projectId: string }>();

  const { mutate: handleDownload, isPending } = useMutation({
    mutationFn: ProjectApi.downloadProjectCode,
  });

  return (
    <Container maxWidth="md">
      <Stack gap={5} my={3}>
        <Card>
          <CardContent>
            <Typography mb={2.5} variant="h6" component="div">
              Download Source Code
            </Typography>
            <Typography gutterBottom variant="body2" color="text.secondary">
              Download the latest source code of your serverless app.
            </Typography>
          </CardContent>
          <Divider />
          <CardActions sx={{ justifyContent: "end", p: 1.5 }}>
            <Button
              variant="contained"
              startIcon={<DownloadRoundedIcon />}
              onClick={() => handleDownload(projectId)}
              color="secondary"
              size="medium"
            >
              Download
            </Button>
          </CardActions>
        </Card>

        <UploadCode projectId={projectId} />
        <UploadAssets projectId={projectId} />
      </Stack>
    </Container>
  );
}
