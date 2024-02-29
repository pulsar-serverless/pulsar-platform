import {
  Button,
  Card,
  CardActions,
  CardContent,
  CircularProgress,
  Divider,
  IconButton,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import RemoveCircleOutlineIcon from "@mui/icons-material/RemoveCircleOutline";
import AddCircleOutlineIcon from "@mui/icons-material/AddCircleOutline";
import { FieldArray, Formik } from "formik";
import { EnvVariablesApi, EnvValidationSchema } from "@/api/envs";
import { useMutation } from "@tanstack/react-query";
import { useSnackbar } from "@/hooks/useSnackbar";

export const EnvVariablesForm = ({ projectID }: { projectID: string }) => {
  const initialValues = { variables: [{ key: "", value: "" }] };

  const snackbar = useSnackbar();

  const { mutate, isPending } = useMutation({
    mutationFn: EnvVariablesApi.createEnvVariables,
    onSuccess: () => {
      snackbar.setSuccessMsg("Environment variables created successfully!");
    },
    onError: () =>
      snackbar.setErrorMsg("Unable to create environment variables"),
  });

  return (
    <>
      <Typography
        variant="h6"
        sx={{ textTransform: "capitalize" }}
        gutterBottom
      >
        Environment Variables
      </Typography>
      <Typography variant="body2">
        Define key-value pairs for your serverless application&apos;s
        environmental variables.
      </Typography>
      <Formik
        onSubmit={(values) => mutate({ ...values, projectID })}
        initialValues={initialValues}
        validationSchema={EnvValidationSchema}
      >
        {({
          values,
          errors,
          touched,
          handleSubmit,
          handleChange,
          handleBlur,
          dirty,
          setValues,
          setTouched,
        }) => (
          <Card sx={{ mt: 4 }} component={"form"} onSubmit={handleSubmit}>
            <CardContent>
              <FieldArray name="variables">
                {({ push, remove }) => {
                  return (
                    <>
                      <Stack direction={"row"} gap={2} sx={{ width: "100%" }}>
                        <Typography variant="body2" sx={{ flex: 3 }}>
                          Key
                        </Typography>
                        <Typography variant="body2" sx={{ flex: 4 }}>
                          Value
                        </Typography>
                      </Stack>
                      <Stack gap={2} alignItems={"start"} my={2}>
                        {values.variables.map(({ key, value }, index) => (
                          <Stack
                            direction={"row"}
                            gap={2}
                            sx={{ width: "100%" }}
                            key={index}
                          >
                            <TextField
                              size="small"
                              sx={{ flex: 3 }}
                              value={key}
                              name={`variables.${index}.key`}
                              placeholder="KEY"
                              onChange={handleChange}
                              onBlur={handleBlur}
                              error={Boolean(
                                touched?.variables?.[index]
                                  ? errors?.variables?.[index]
                                  : null
                              )}
                              helperText={
                                touched?.variables?.[index]
                                  ? (errors.variables?.[index] as any)?.key
                                  : null
                              }
                            />
                            <Stack direction={"row"} sx={{ flex: 4 }} gap={1.5}>
                              <TextField
                                size="small"
                                fullWidth
                                name={`variables.${index}.value`}
                                value={value}
                                onChange={handleChange}
                                onBlur={handleBlur}
                                error={Boolean(
                                  touched?.variables?.[index]
                                    ? errors?.variables?.[index]
                                    : null
                                )}
                                helperText={
                                  touched?.variables?.[index]
                                    ? (errors.variables?.[index] as any)?.value
                                    : null
                                }
                              />
                              <IconButton onClick={() => remove(index)}>
                                <RemoveCircleOutlineIcon />
                              </IconButton>
                            </Stack>
                          </Stack>
                        ))}
                      </Stack>
                      <Button
                        startIcon={<AddCircleOutlineIcon />}
                        variant="outlined"
                        size="small"
                        color="secondary"
                        onClick={() => push({ key: "", value: "" })}
                      >
                        Add another
                      </Button>
                    </>
                  );
                }}
              </FieldArray>
            </CardContent>
            <Divider />
            <CardActions sx={{ justifyContent: "end", p: 1.5 }}>
              {dirty && (
                <Button
                  type="reset"
                  variant="outlined"
                  color="error"
                  onClick={() => {
                    setValues({ variables: [{ key: "", value: "" }] });
                    setTouched({ variables: undefined });
                  }}
                >
                  Cancel
                </Button>
              )}
              <Button
                color="secondary"
                variant="contained"
                type="submit"
                disabled={isPending}
              >
                {!isPending ? (
                  "Save"
                ) : (
                  <CircularProgress color="secondary" size={24} />
                )}
              </Button>
            </CardActions>
          </Card>
        )}
      </Formik>
    </>
  );
};
