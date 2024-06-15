import React, { useState } from "react";
import { Dialog, DialogTitle, DialogContent, Stack } from "@mui/material";
import { useMutation, useQuery } from "@tanstack/react-query";
import { PricingApi } from "@/api/pricing";
import PricingCard from "../pricing/PricingCard";
import { ConfirmationDialog } from "../modals/ConfirmationDialog";
import { useSnackbar } from "@/hooks/useSnackbar";

const SetPricingPlanModal = ({
	open,
	onClose,
	projectId,
}: {
	projectId: string;
	open: boolean;
	onClose: () => void;
}) => {
	const snackbar = useSnackbar();

	const [confirmChange, setConfirmChange] = useState<string | undefined>(
		undefined
	);

	const { data, isLoading } = useQuery({
		queryKey: ["pricingPlans"],
		queryFn: () => PricingApi.getPricingPlans(1, 10),
	});

	const { mutate } = useMutation({
		mutationFn: (planId: string) =>
			PricingApi.changePricingPlan(planId, projectId),
		onSuccess: () => {
			snackbar.setSuccessMsg("Project pricing plan changed.");
			onClose();
		},
		onError: () => snackbar.setErrorMsg("Unable to change pricing plan."),
	});

	const handleChangePricingPlan = () => {
		mutate(confirmChange!!);
	};

	return (
		<>
			<Dialog open={open} onClose={onClose}>
				<DialogTitle>Change Pricing plan</DialogTitle>
				<DialogContent>
					<Stack
						direction="row"
						sx={{
							justifyContent: "center",
							gap: 1,
						}}>
						{data?.rows.map((plan) => (
							<PricingCard
								plan={plan}
								key={plan.id}
								onPricePlanSelected={(plan) => setConfirmChange(plan.id)}
							/>
						))}
					</Stack>
				</DialogContent>
			</Dialog>
			{confirmChange && (
				<ConfirmationDialog
					open={!!confirmChange}
					title="You are about to change your serverless hosting plan"
					description="Changing your serverless plan may affect costs & resource limits.  Review details & confirm."
					handleClose={() => setConfirmChange(undefined)}
					handleConfirm={handleChangePricingPlan}
				/>
			)}
		</>
	);
};

export default SetPricingPlanModal;
