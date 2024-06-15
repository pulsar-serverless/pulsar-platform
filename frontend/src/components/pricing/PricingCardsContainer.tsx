import React, { useState, useEffect } from "react";
import { Container, Typography, CircularProgress, Box } from "@mui/material";
import { useQuery } from "@tanstack/react-query";
import { PricingApi } from "@/api/pricing";
import PricingCard from "./PricingCard";

// Fallback data
const fallbackPlans = {
	pageNumber: 1,
	pageSize: 10,
	totalPages: 1,
	totalItems: 4,
	plans: [
		{
			ID: "b3b1e20f-0f91-4d21-8bb4-0f3f3d3e3c3a",
			Name: "Free Plan",
			Description: "Basic plan with limited resources",
			PlanTeir: "free",
			NotifyThreshold: 80,
			functionalities: [
				{
					functionalityId: "1",
					functionalityValue:
						"Includes 1 GB of storage space and basic computational resources suitable for small projects.",
				},
				{
					functionalityId: "2",
					functionalityValue:
						"Allows up to 10,000 requests per month with minimal downtime.",
				},
				{
					functionalityId: "3",
					functionalityValue:
						"Provides basic support through community forums and online resources.",
				},
			],
		},
		{
			ID: "d5d3f32f-2f13-5f43-aad6-2f5f5g5h5e5d",
			Name: "Pro Plan",
			Description: "Advanced plan for professional use",
			PlanTeir: "pro",
			NotifyThreshold: 90,
			functionalities: [
				{
					functionalityId: "7",
					functionalityValue:
						"Provides 50 GB of storage space and high-performance computational resources for demanding applications.",
				},
				{
					functionalityId: "8",
					functionalityValue:
						"Allows up to 1,000,000 requests per month with guaranteed uptime.",
				},
				{
					functionalityId: "9",
					functionalityValue:
						"Priority email support with a response time of up to 12 hours.",
				},
			],
		},
	],
};

const PricingCardsContainer: React.FC = () => {
	const { data, isError, isLoading } = useQuery({
		queryKey: ["pricingPlans"],
		queryFn: () => PricingApi.getPricingPlans(1, 10),
	});

	const [plans, setPlans] = useState<any[]>([]);

	useEffect(() => {
		if (!isLoading) {
			if (isError || !data || data.plans.length === 0) {
				setPlans(fallbackPlans.plans);
			} else {
				setPlans(data.plans);
			}
		}
	}, [data, isError, isLoading]);

	if (isLoading) {
		return (
			<Box
				display="flex"
				justifyContent="center"
				alignItems="center"
				height="100vh"
				width="100vw">
				<CircularProgress />
			</Box>
		);
	}

	if (isError && plans.length === 0) {
		return <Typography>Error loading pricing plans</Typography>;
	}

	// Divide the plans into two groups
	const firstColumnPlans = plans.slice(0, 2);
	const secondColumnPlans = plans.slice(2, 4);

	return (
		<Container
			sx={{
				display: "flex",
				flexDirection: "column",
				alignItems: "center",
				justifyContent: "center",
				marginTop: 4,
				marginBottom: 4,
				width: "100vw",
			}}>
			<Box
				sx={{
					display: "flex",
					justifyContent: "center",
					gap: 3,
					flexWrap: "wrap",
					width: "100%",
				}}>
				{firstColumnPlans.map((plan) => (
					<Box key={plan.ID} sx={{ flex: "1 0 45%", maxWidth: "45%" }}>
						<PricingCard plan={plan} />
					</Box>
				))}
			</Box>
			<Box
				sx={{
					display: "flex",
					justifyContent: "center",
					gap: 3,
					flexWrap: "wrap",
					width: "100%",
					marginTop: 3,
				}}>
				{secondColumnPlans.map((plan) => (
					<Box key={plan.ID} sx={{ flex: "1 0 45%", maxWidth: "45%" }}>
						<PricingCard plan={plan} />
					</Box>
				))}
			</Box>
		</Container>
	);
};

export default PricingCardsContainer;
