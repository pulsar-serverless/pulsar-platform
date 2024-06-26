import { axiosInstance } from "@/components/interceptors/HttpInterceptor";
import { Paginated } from "@/models/pagination";
import { User } from "@/models/user";

export const userApi = {
  async getUsers(pageSize: number, pageNumber: number, searchQuery?: string) {
    const { data } = await axiosInstance.get<Paginated<User>>("/users", {
      params: { pageSize, pageNumber, searchQuery },
    });
    return data;
  },

  async deleteAllProjects(userId: string) {
    const { data } = await axiosInstance.delete(`/users/${userId}/projects`);
    return data;
  },

  async changeAccountStatus(userId: string, status: string) {
    const { data } = await axiosInstance.put(`/users/${userId}/`, { status });
    return data;
  },

  async getAccountStatus() {
    const { data } = await axiosInstance.get<string>(`/users/status`);
    return data;
  },
};
