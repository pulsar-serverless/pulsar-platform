import { axiosInstance } from "@/components/interceptors/HttpInterceptor";
import { Project } from "@/models/project";
import { object, string, InferType } from "yup";

export const createProjectSchema = object({
  name: string()
    .required("Project name is required.")
    .min(3, "Project name must be longer than 3 characters.")
    .matches(/^[a-zA-Z0-9]+$/, {
      message: "Project name must be only alpha-numeric.",
    }),
});

export const CreateProjectSchema = createProjectSchema;

export const ProjectApi = {
  async createProject(body: InferType<typeof createProjectSchema>) {
    const { data } = await axiosInstance.post<Project>("/projects", body);
    return data;
  },
};
