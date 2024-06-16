"use client";
import {
  Button,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField,
  Typography,
} from "@mui/material";
import { useFormik } from "formik";
import { useURlState } from "@/hooks/userUrlState";
import { ProjectApi, createProjectSchema } from "@/api/projects";
import { useMutation } from "@tanstack/react-query";
import { useSnackbar } from "@/hooks/useSnackbar";
import { useRouter } from "next/navigation";

export const CreateProjectModal = () => {
  const { state: isOpen, removeState } = useURlState("action");

  const router = useRouter();
  const snackbar = useSnackbar();

  const { mutate, isPending } = useMutation({
    mutationFn: ProjectApi.createProject,
    onSuccess: (data) => {
      snackbar.setSuccessMsg("Project created successfully!");
      removeState();
      router.push(`/username/${data.id}/home`);
    },
    onError: () => snackbar.setErrorMsg("Unable to create a project"),
  });

  const { handleSubmit, handleBlur, handleChange, errors, values, touched } =
    useFormik({
      initialValues: { name: "" },
      validationSchema: createProjectSchema,
      onSubmit: (values, {}) => {
        mutate(values);
      },
    });

  return (
    <>
      <Dialog
        onClose={removeState}
        open={Boolean(isOpen)}
        PaperProps={{ component: "form", onSubmit: handleSubmit }}
      >
        <DialogTitle>Create a new project</DialogTitle>
        <DialogContent>
          <Typography variant="body2" gutterBottom mb={3}>
            Please type in the name of the project to you want to create.
          </Typography>
          <TextField
            placeholder="Project name"
            variant="outlined"
            fullWidth
            size="small"
            name="name"
            onChange={handleChange}
            onBlur={handleBlur}
            value={values.name}
            error={!!(touched.name && errors.name)}
            helperText={errors.name}
          />
        </DialogContent>
        <DialogActions>
          <Button variant="outlined" color="secondary" onClick={removeState}>
            Cancel
          </Button>
          <Button variant="contained" type="submit">
            {!isPending ? "Create" : <CircularProgress color="secondary" size={24}/>}
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};
