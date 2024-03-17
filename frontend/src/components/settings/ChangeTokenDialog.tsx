import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  Stack,
  IconButton,
  OutlinedInput,
  InputAdornment,
  useTheme,
  Alert,
  CircularProgress,
} from "@mui/material";
import ContentCopyRoundedIcon from "@mui/icons-material/ContentCopyRounded";
import VisibilityOffRoundedIcon from "@mui/icons-material/VisibilityOffRounded";
import VisibilityRoundedIcon from "@mui/icons-material/VisibilityRounded";
import { FC, useState } from "react";
import { useMutation } from "@tanstack/react-query";
import { ProjectApi } from "@/api/projects";
import { useSnackbar } from "@/hooks/useSnackbar";

const ChangeTokenDialog: FC<{
  isOpen: boolean;
  onClose: () => void;
  projectId: string;
}> = ({ onClose, isOpen, projectId }) => {
  const [show, setShow] = useState(false);
  const [key, setKey] = useState("");
  const snackbar = useSnackbar();

  const { mutate, isPending } = useMutation({
    mutationFn: () => ProjectApi.generateApiKey(projectId),
    onSuccess: (data) => {
      snackbar.setSuccessMsg("Project API key updated.");
      console.log(data);
      setKey(data.token || "");
    },
    onError: () => snackbar.setErrorMsg("Unable to change project API key"),
  });

  return (
    <Dialog onClose={onClose} open={Boolean(isOpen)}>
      <DialogTitle>Change API Key</DialogTitle>
      <DialogContent>
        <Alert severity="warning" sx={{ mb: 2.5 }}>
          The API key will only be displayed once. Make sure to save it after
          changing.
        </Alert>
        <Stack direction={"row"} gap={1} alignItems={"center"}>
          <OutlinedInput
            placeholder="*********"
            type={!show ? "password" : "text"}
            fullWidth
            size="small"
            name="name"
            disabled={true}
            value={key}
            endAdornment={
              <InputAdornment position="end">
                <Stack direction={"row"} gap={1}>
                  <IconButton
                    size="small"
                    color="secondary"
                    disabled={key == ""}
                    onClick={() => setShow((val) => !val)}
                  >
                    {show ? (
                      <VisibilityOffRoundedIcon />
                    ) : (
                      <VisibilityRoundedIcon />
                    )}
                  </IconButton>
                  <IconButton
                    size="small"
                    color="secondary"
                    disabled={key == ""}
                    onClick={() => navigator.clipboard.writeText(key)}
                  >
                    <ContentCopyRoundedIcon />
                  </IconButton>
                </Stack>
              </InputAdornment>
            }
          />
        </Stack>
      </DialogContent>
      <DialogActions>
        <Button variant="outlined" color="secondary" onClick={onClose}>
          Cancel
        </Button>
        <Button
          variant="contained"
          type="submit"
          color="error"
          onClick={() => mutate()}
        >
          {!isPending ? (
            "Change"
          ) : (
            <CircularProgress color="secondary" size={24} />
          )}
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export default ChangeTokenDialog;
