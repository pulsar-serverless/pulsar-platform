"use client";

import { setSnackbarMessage } from "@/store/slices/appSlice";
import { useAppDispatch, useAppSelector } from "@/store/store";
import { Alert, Snackbar } from "@mui/material";

export const SnackbarProvider = () => {
  const { snackbarMessage: message } = useAppSelector((state) => state.app);
  const dispatch = useAppDispatch();

  const handleClose = () => {
    dispatch(setSnackbarMessage(undefined));
  };

  return (
    <Snackbar
      open={Boolean(message)}
      autoHideDuration={6000}
      onClose={handleClose}
    >
      <Alert
        onClose={handleClose}
        severity={message?.type != "error" ? "success" : "error" }
        variant="filled"
        sx={{ width: "100%" }}
      >
        {message?.content}
      </Alert>
    </Snackbar>
  );
};
