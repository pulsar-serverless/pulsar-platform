"use client";

import { ProjectApi } from "@/api/projects";
import { ConfirmationDialog } from "@/components/modals/ConfirmationDialog";
import ChangeTokenDialog from "@/components/settings/ChangeTokenDialog";
import { useSnackbar } from "@/hooks/useSnackbar";
import {
  Alert,
  Button,
  Card,
  CardActions,
  CardContent,
  Container,
  Stack,
  Typography,
} from "@mui/material";
import { useMutation } from "@tanstack/react-query";
import { useParams } from "next/navigation";
import { useState } from "react";

function Page() {
  const [changeToken, setChangeToken] = useState(false);
  const [removeAPIKey, setRemoveAPIkey] = useState(false);

  const { projectId } = useParams<{ projectId: string }>();

  const snackbar = useSnackbar();

  const { mutate: handleRemoveAPIKey } = useMutation({
    mutationFn: ProjectApi.removeAPIKey,
    onSuccess: () => {
      snackbar.setSuccessMsg("Project API key removed successfully!");
    },
    onError: () => snackbar.setErrorMsg("Unable to project API key removed."),
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
              <Alert severity="warning">
                Removing the API key will leave your project vulnerable,
                allowing unrestricted access by anyone.
              </Alert>
            </CardContent>
            <CardActions sx={{ justifyContent: "end", p: 1.5 }}>
              <Button
                variant="contained"
                color="error"
                onClick={() => setRemoveAPIkey(true)}
              >
                Remove API Key
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

      {removeAPIKey && (
        <ConfirmationDialog
          open={removeAPIKey}
          title="Remove API Key"
          description="Are you sure you want to remove the API key? Doing so will leave your project vulnerable,"
          handleClose={() => setRemoveAPIkey(false)}
          handleConfirm={() => {
            handleRemoveAPIKey(projectId);
            setRemoveAPIkey(false)
          }}
        />
      )}
    </>
  );
}

export default Page;
