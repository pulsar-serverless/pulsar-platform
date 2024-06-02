import { axiosInstance } from "@/components/interceptors/HttpInterceptor";

interface TotalResourceUtil {
	mem_usage_mb: number;
	net_usage_mb: number;
	project_id: string;
	usage_period: string;
}

export const ResourcesApi = {
	async getMonthlyProjectResource(projectId: string, month: string) {
		const { data } = await axiosInstance.get<[]>(
			`/projects/${projectId}/resources/monthly`,
			{ params: { month } }
		);
		return data;
	},
	async getProjectResourceUtilList(
		projectId: string,
		pageNumber: number,
		pageSize: number,
		month?: string
	) {
		const params = { pageNumber, pageSize, month };
		const { data } = await axiosInstance.get<[]>(
			`/projects/${projectId}/resources`,
			{ params }
		);
		return data;
	},
	async getTotalProjectResourceUtil(projectId: string) {
		const { data } = await axiosInstance.get<TotalResourceUtil>(
			`/projects/${projectId}/resources/total`
		);

		return data;
	},
};
