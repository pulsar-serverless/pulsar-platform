import { axiosInstance } from "@/components/interceptors/HttpInterceptor";

export const AnalyticsApi = {
  async getLast24HoursInvocations(projectId: string, status: InvocationStatus) {
    const { data } = await axiosInstance.get<[]>(
      `/projects/${projectId}/analytics/hourly`,
      { params: { status } }
    );
    return data;
  },
  async getLast7DaysInvocations(projectId: string, status: InvocationStatus) {
    const { data } = await axiosInstance.get<[]>(
      `/projects/${projectId}/analytics/weekly`,
      { params: { status } }
    );
    return data;
  },
  async getLast30DaysInvocations(projectId: string, status: InvocationStatus) {
    const { data } = await axiosInstance.get<[]>(
      `/projects/${projectId}/analytics/monthly`,
      { params: { status } }
    );
    return data;
  },
};
