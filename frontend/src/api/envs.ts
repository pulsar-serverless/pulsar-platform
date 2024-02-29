import { axiosInstance } from "@/components/interceptors/HttpInterceptor";
import { Project } from "next/dist/build/swc";
import { object, array, string, InferType } from "yup";

export const EnvValidationSchema = object({
  variables: array().required().of(
    object({
      key: string()
        .required("Environment key is required.")
        .matches(/^[a-zA-Z0-9\_]+$/, {
          message:
            "Project must consist of only alpha-numeric characters or -.",
        }),
      value: string().required("Environment value is required."),
    })
  ),
});

export type Envs = InferType<typeof EnvValidationSchema>;

export const EnvVariablesApi = {
  async createEnvVariables(body: Pick<Envs, "variables"> & { projectID: string }) {
    const { data } = await axiosInstance.post<Project>(
      `/projects/envs/${body.projectID}`,
      body
    );
    return data;
  },
};
