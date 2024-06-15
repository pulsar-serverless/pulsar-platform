import React, { useState } from "react";
import {
	Card,
	CardContent,
	Typography,
	Button,
	List,
	ListItem,
	ListItemText,
} from "@mui/material";
import PricingPlanModal from "./PricingPlanModal";

const PricingCard: React.FC<{ plan: any }> = ({ plan }) => {
	const [isModalOpen, setIsModalOpen] = useState(false);

	const handleChoosePlan = (planId: string) => {
		console.log(`Selected Plan ID: ${planId}`);
		setIsModalOpen(true);
	};

	const handleCloseModal = () => {
		setIsModalOpen(false);
	};

	return (
		<>
			<Card
				sx={{
					flexGrow: 1,
					minWidth: 300,
					maxWidth: 350,
					minHeight: 450,
					marginBottom: 2,
					display: "flex",
					flexDirection: "column",
					justifyContent: "center",
					alignItems: "center",
				}}>
				<CardContent>
					<Typography
						variant="subtitle1"
						sx={{ textTransform: "capitalize", marginBottom: 1 }}
						fontWeight={"medium"}
						component="div"
						gutterBottom>
						{plan.Name}
					</Typography>
					<Typography variant="h5" sx={{ marginBottom: 2 }}>
						{plan.price ? `$${plan.price}/month` : "Free"}
					</Typography>
					<Typography
						variant="body2"
						color="textSecondary"
						sx={{ marginBottom: 3 }}>
						{plan.Description}
					</Typography>
					<List sx={{ marginBottom: 3 }}>
						{plan.functionalities.map((func: any) => (
							<ListItem
								key={func.functionalityId}
								sx={{ padding: 0, marginBottom: 1 }}>
								<ListItemText primary={`• ${func.functionalityValue}`} />
							</ListItem>
						))}
					</List>
				</CardContent>
				<Button
					variant="contained"
					color="secondary"
					sx={{
						alignSelf: "center",
						marginBottom: 2,
						border: "1px solid white",
						"&:hover": {
							backgroundColor: "black",
							color: "white",
						},
					}}
					onClick={() => handleChoosePlan(plan.ID)}>
					Choose Plan
				</Button>
			</Card>
			<PricingPlanModal
				isOpen={isModalOpen}
				onClose={handleCloseModal}
				planId={plan.ID}
			/>
		</>
	);
};

export default PricingCard;
