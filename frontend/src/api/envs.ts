import { axiosInstance } from "@/components/interceptors/HttpInterceptor";
import { object, array, string, InferType } from "yup";

export const EnvValidationSchema = object({
  variables: array()
    .required()
    .of(
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

type UnArray<T> = T extends Array<infer U> ? U : T;
export type Envs = InferType<typeof EnvValidationSchema>;
export type EnvVariable = UnArray<Envs["variables"]> & { projectID: string };

export const EnvVariablesApi = {
  async createEnvVariables(
    body: Pick<Envs, "variables"> & { projectID: string }
  ) {
    const { data } = await axiosInstance.post<EnvVariable[]>(
      `/projects/envs/${body.projectID}`,
      body
    );
    return data;
  },

  async getEnvVariables(projectID: string) {
    const { data } = await axiosInstance.get<EnvVariable[]>(
      `/projects/envs/${projectID}`
    );
    return data;
  },
};
