import { ProjectApi } from "@/api/projects";
import { useSnackbar } from "@/hooks/useSnackbar";
import { Project } from "@/models/project";
import {
  Dialog,
  DialogTitle,
  DialogContent,
  Alert,
  Stack,
  DialogActions,
  Button,
  CircularProgress,
  TextField,
} from "@mui/material";
import { useMutation } from "@tanstack/react-query";
import { useFormik } from "formik";
import { object, string } from "yup";

const schema = object({
  subdomain: string()
    .min(3, "Subdomain must be longer than 6 characters.")
    .matches(/^[a-zA-Z0-9\-]+$/, {
      message: "Subdomain must consist of only alpha-numeric characters or -.",
    })
});

export const ChangeSubdomain: React.FC<{
  onClose: () => void;
  project?: Project;
}> = ({ project, onClose }) => {
  const snackbar = useSnackbar();

  const { handleSubmit, handleBlur, handleChange, errors, values, touched } =
    useFormik({
      initialValues: { subdomain: project?.subdomain || "" },
      validationSchema: schema,
      onSubmit: (values, {}) => {
        mutate(values.subdomain);
      },
    });

  const { mutate, isPending } = useMutation({
    mutationFn: (subdomain: string) =>
      ProjectApi.ChangeSubdomain(project?.id || "", subdomain),
    onSuccess: (data) => {
      snackbar.setSuccessMsg("Project sud-domain updated.");
      onClose();
    },
    onError: () => snackbar.setErrorMsg("Unable to change project sub-domain."),
  });

  return (
    <>
      {" "}
      <Dialog onClose={onClose} open={Boolean(true)}>
        <DialogTitle>Change subdomain</DialogTitle>
        <DialogContent>
          <Alert severity="error" sx={{ mb: 2.5 }}>
            Switching subdomains may cause temporary downtime and affect website
            visibility.
          </Alert>
          <Stack direction={"row"} gap={1} alignItems={"center"}>
            <TextField
              fullWidth
              size="small"
              name="subdomain"
              value={values.subdomain}
              onChange={handleChange}
              onBlur={handleBlur}
              error={!!(touched.subdomain && errors.subdomain)}
              helperText={errors.subdomain}
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
            onClick={() => handleSubmit()}
          >
            {!isPending ? (
              "Change"
            ) : (
              <CircularProgress color="secondary" size={24} />
            )}
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};
