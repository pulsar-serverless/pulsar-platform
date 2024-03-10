import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  Button,
} from "@mui/material";

export const ConfirmationDialog: React.FC<{
  open: boolean;
  title: string;
  description: string;
  handleClose: () => void;
  handleConfirm: () => void;
}> = ({ title, description, open, handleClose, handleConfirm }) => {
  return (
    <Dialog
      open={open}
      onClose={handleClose}
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
    >
      <DialogTitle id="alert-dialog-title">
        {title}
      </DialogTitle>
      <DialogContent>
        <DialogContentText id="alert-dialog-description">
          {description}
        </DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose} variant="outlined">Cancel</Button>
        <Button onClick={handleConfirm} autoFocus color="error" variant="outlined">
          Confirm
        </Button>
      </DialogActions>
    </Dialog>
  );
};
