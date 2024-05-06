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
};
