"use client";

import { ProjectApi } from "@/api/projects";
import { ConfirmationDialog } from "@/components/modals/ConfirmationDialog";
import ChangeTokenDialog from "@/components/settings/ChangeTokenDialog";
import { useSnackbar } from "@/hooks/useSnackbar";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  Container,
  Stack,
  Typography,
} from "@mui/material";
import { useMutation } from "@tanstack/react-query";
import { useParams, useRouter } from "next/navigation";
import { useState } from "react";

function Page() {
  const [changeToken, setChangeToken] = useState(false);
  const [confirm, setConfirm] = useState<
    false | "REMOVE_API_KEY" | "DELETE_PROJECT"
  >(false);

  const { projectId } = useParams<{ projectId: string }>();

  const snackbar = useSnackbar();
  const router = useRouter();

  const { mutate: handleRemoveAPIKey } = useMutation({
    mutationFn: ProjectApi.removeAPIKey,
    onSuccess: () => {
      snackbar.setSuccessMsg("Project API key removed successfully!");
    },
    onError: () => snackbar.setErrorMsg("Unable to project API key removed."),
  });

  const { mutate: handleDelete } = useMutation({
    mutationFn: ProjectApi.deleteProject,
    onSuccess: () => {
      snackbar.setSuccessMsg("Project deleted successfully!");
      router.push("/username");
    },
    onError: () => snackbar.setErrorMsg("Unable to delete the project."),
  });

  return (
    <>
      <Container maxWidth="md" sx={{ py: 3 }}>
        <Typography mb={2.5} variant="h5" component="div">
          Settings
        </Typography>

        <Stack gap={3}>
          <Card>
            <CardContent>
              <Typography
                mb={1.5}
                variant="subtitle1"
                fontWeight={"medium"}
                component="div"
              >
                Generate/Change API Key
              </Typography>
              <Typography gutterBottom variant="body2" color="text.secondary">
                Secure your serverless function with API keys for enhanced
                protection.
              </Typography>
            </CardContent>
            <CardActions sx={{ justifyContent: "end", p: 1.5 }}>
              <Button
                variant="contained"
                color="secondary"
                onClick={() => setChangeToken(true)}
              >
                Change Key
              </Button>
            </CardActions>
          </Card>

          <Card>
            <CardContent>
              <Typography
                mb={1.5}
                variant="subtitle1"
                fontWeight={"medium"}
                component="div"
              >
                Remove API key authorization
              </Typography>
              <Typography gutterBottom variant="body2" color="text.secondary">
                Removing the API key will leave your project vulnerable,
                allowing unrestricted access by anyone.
              </Typography>
            </CardContent>
            <CardActions sx={{ justifyContent: "end", p: 1.5 }}>
              <Button
                variant="contained"
                color="error"
                onClick={() => setConfirm("REMOVE_API_KEY")}
              >
                Remove API Key
              </Button>
            </CardActions>
          </Card>

          <Card sx={{}}>
            <CardContent>
              <Typography
                mb={1.5}
                variant="subtitle1"
                fontWeight={"medium"}
                component="div"
              >
                Delete Project
              </Typography>
              <Typography gutterBottom variant="body2" color="text.secondary">
                Removing the project will remove all associated files. Users
                will not be able to access the project.
              </Typography>
            </CardContent>
            <CardActions sx={{ justifyContent: "end", p: 1.5 }}>
              <Button
                variant="contained"
                color="error"
                onClick={() => setConfirm("DELETE_PROJECT")}
              >
                Delete Project
              </Button>
            </CardActions>
          </Card>
        </Stack>
      </Container>

      {changeToken && (
        <ChangeTokenDialog
          isOpen={changeToken}
          onClose={() => setChangeToken(false)}
          projectId={projectId}
        />
      )}

      {confirm && (
        <ConfirmationDialog
          open={!!confirm}
          title={
            confirm == "REMOVE_API_KEY" ? "Remove API Key" : "Delete Project"
          }
          description={
            confirm == "REMOVE_API_KEY"
              ? "Are you sure you want to remove the API key? Doing so will leave your project vulnerable,"
              : "Removing the project will remove all associated files. Users will not be able to access the project. This action cannot be undone."
          }
          handleClose={() => setConfirm(false)}
          handleConfirm={() => {
            confirm == "REMOVE_API_KEY"
              ? handleRemoveAPIKey(projectId)
              : handleDelete(projectId);
            setConfirm(false);
          }}
        />
      )}
    </>
  );
}

export default Page;
