import { axiosInstance } from "@/components/interceptors/HttpInterceptor";
import { Log } from "@/models/log";
import { Paginated } from "@/models/pagination";

export const LogApi = {
  async getLogs(
    projectID: string,
    logTypes: Log['type'][],
    searchQuery: string,
    pageNumber: number = 1,
    pagesize: number = 50
  ) {
    const { data } = await axiosInstance.get<Paginated<Log>>(
      `/projects/logs/${projectID}`,
      {
        params: { pageNumber, pagesize, logTypes, searchQuery },
      }
    );
    return data;
  },

  async deleteLogs(projectID: string) {
    const { data } = await axiosInstance.delete<Paginated<Log>>(
      `/projects/logs/${projectID}`
    );
    return data;
  },
};
