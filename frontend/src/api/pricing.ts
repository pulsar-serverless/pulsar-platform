import { axiosInstance } from "@/components/interceptors/HttpInterceptor";
import { PricingPlan } from "@/models/PricingPlan";
import { Paginated } from "@/models/pagination";

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

	async getPricingPlans(pageNumber: number, pageSize: number): Promise<Paginated<PricingPlan>> {
		const { data } = await axiosInstance.get<any>(`/projects/plans`, {
			params: { pageNumber, pageSize },
		});
		return data;
	},
};
