"use client";

import {
  Card,
  CardContent,
  Typography,
  Divider,
  CardActions,
  Button,
  Box,
  Chip,
} from "@mui/material";
import { FormEvent, useState } from "react";
import FolderZipRoundedIcon from "@mui/icons-material/FolderZipRounded";
import { useMutation } from "@tanstack/react-query";
import { useSnackbar } from "@/hooks/useSnackbar";
import { ProjectApi } from "@/api/projects";

export const UploadAssets = ({ projectId }: { projectId: string }) => {
  const [file, setFile] = useState<File | undefined>();
  const snackbar = useSnackbar();

  const { mutate } = useMutation({
    mutationFn: ProjectApi.uploadProjectStaticSite,
    onSuccess: () => {
      snackbar.setSuccessMsg("Project assets uploaded successfully!");
      setFile(undefined);
    },
    onError: () => snackbar.setErrorMsg("Unable to upload the assets."),
  });

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    const form = new FormData();
    form.set("file", file!);
    form.set("projectId", projectId);

    mutate(form);
  };

  return (
    <Card component="form" onSubmit={handleSubmit}>
      <CardContent>
        <Typography mb={2.5} variant="h6" component="div">
          Upload Static Assets
        </Typography>
        <Typography gutterBottom variant="body2" color="text.secondary">
          Upload your zipped static assets.{" "}
          <Typography component={'span'} variant="body2" color="warning.light">
            Make sure that your 'index.html' file is in the root path.
          </Typography>
        </Typography>
      </CardContent>
      <Divider />
      <CardActions sx={{ justifyContent: "space-between", p: 1.5 }}>
        <Box>
          {file && (
            <Chip
              icon={<FolderZipRoundedIcon />}
              label={file?.name}
              onDelete={() => setFile(undefined)}
            />
          )}
        </Box>
        {!file && (
          <Button
            variant="contained"
            component="label"
            color="secondary"
            size="medium"
          >
            Select File
            <input
              type="file"
              hidden
              onChange={(e) => {
                console.log(e.target.files);
                setFile(e.target.files?.[0]);
              }}
              accept=".zip"
            />
          </Button>
        )}
        {!!file && (
          <Button variant="contained" color="secondary" type="submit">
            Upload
          </Button>
        )}
      </CardActions>
    </Card>
  );
};
