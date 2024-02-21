"use client";

import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  Divider,
  Typography,
} from "@mui/material";
import DownloadRoundedIcon from "@mui/icons-material/DownloadRounded";
import { useParams } from "next/navigation";
import { ProjectApi } from "@/api/projects";
import { useMutation } from "@tanstack/react-query";

export default function Page() {
  const { projectId } = useParams<{ projectId: string }>();

  const { mutate: handleDownload, isPending } = useMutation({
    mutationFn: ProjectApi.downloadProjectCode,
  });

  return (
    <Box my={3}>
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
    </Box>
  );
}
