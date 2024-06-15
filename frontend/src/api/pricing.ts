import { axiosInstance } from "@/components/interceptors/HttpInterceptor";

export const PricingApi = {
	async setProjectPricingPlan(
		projectId: string,
		planId: string
	): Promise<number> {
		const { status } = await axiosInstance.post<number>(
			`/api/projects/${projectId}/plan`,
			null,
			{ params: { planId } }
		);
		return status;
	},

	async getPricingPlans(pageNumber: number, pageSize: number): Promise<any> {
		const { data } = await axiosInstance.get<any>(`/api/projects/plans`, {
			params: { pageNumber, pageSize },
		});
		return data;
	},
};
