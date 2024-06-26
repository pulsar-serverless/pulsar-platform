import { axiosInstance } from "@/components/interceptors/HttpInterceptor";
import { ChangeSubdomain } from "@/components/settings/ChangeSubdomain";
import { Paginated } from "@/models/pagination";
import { Project } from "@/models/project";
import { object, string, InferType } from "yup";

export const createProjectSchema = object({
  name: string()
    .required("Project name is required.")
    .min(3, "Project name must be longer than 3 characters.")
    .matches(/^[a-zA-Z0-9\-]+$/, {
      message: "Project must consist of only alpha-numeric characters or -.",
    }),
  subdomain: string()
    .min(3, "Subdomain must be longer than 6 characters.")
    .matches(/^[a-zA-Z0-9\-]+$/, {
      message: "Subdomain must consist of only alpha-numeric characters or -.",
    }),
});

export const CreateProjectSchema = createProjectSchema;

export const ProjectApi = {
  async createProject(body: InferType<typeof createProjectSchema>) {
    const { data } = await axiosInstance.post<Project>("/projects", body);
    return data;
  },

  async getProject(projectName: string) {
    const { data } = await axiosInstance.get<Project>(
      `/projects/${encodeURIComponent(projectName)}`
    );
    return data;
  },

  async getProjects(
    pageNumber: number = 0,
    pagesize: number = 20,
    userId?: string
  ) {
    const { data } = await axiosInstance.get<Paginated<Project>>(`/projects`, {
      params: { userId, pageNumber, pagesize },
    });
    return data;
  },

  async downloadProjectCode(projectId: string) {
    const { data } = await axiosInstance.get(`/projects/code/${projectId}`, {
      responseType: "blob",
    });
    const blobUrl = URL.createObjectURL(data);
    window.open(blobUrl);
  },

  async uploadProjectCode(body: FormData) {
    const { data } = await axiosInstance.put(
      `/projects/code/${body.get("projectId")}`,
      body
    );
    return data;
  },

  async generateApiKey(projectId: string) {
    const { data } = await axiosInstance.put<{ token: string }>(
      `/projects/${projectId}/api-token`
    );
    return data;
  },

  async removeAPIKey(projectId: string) {
    const { data } = await axiosInstance.delete(
      `/projects/${projectId}/api-token`
    );
    return data;
  },

  async deleteProject(projectId: string) {
    const { data } = await axiosInstance.delete(`/projects/${projectId}`);
    return data;
  },

  async ChangeSubdomain(projectId: string, subdomain: string) {
    const { data } = await axiosInstance.put(`/projects/${projectId}`, {
      subdomain: subdomain,
    });
    return data;
  },
  
  async downloadInvoice(projectId: string) {
    const currentDate = new Date();
    const month = currentDate.getMonth() + 1;
    const year = currentDate.getFullYear();
    const currentMonth = month.toString() + year.toString();

    const {data} = await axiosInstance.get(
      `/projects/${projectId}/invoice?month=${currentMonth}`,
      {
        responseType: "blob",
      }
    )
    const blobUrl = URL.createObjectURL(data);
    window.open(blobUrl);
  },
};
