import React, { useState } from "react";
import {
	Dialog,
	DialogTitle,
	DialogContent,
	DialogActions,
	Button,
	Typography,
	CircularProgress,
} from "@mui/material";

const PricingPlanModal: React.FC<{
	isOpen: boolean;
	onClose: () => void;
	planId: string;
}> = ({ isOpen, onClose, planId }) => {
	const [loading, setLoading] = useState(false);
	const [success, setSuccess] = useState(false);

	const handleConfirm = () => {
		setLoading(true);
		// Mock post request with setTimeout
		setTimeout(() => {
			setLoading(false);
			setSuccess(true);
			setTimeout(() => {
				setSuccess(false);
				onClose();
			}, 2000);
		}, 2000);
	};

	return (
		<Dialog open={isOpen} onClose={onClose}>
			<DialogTitle>Confirm Plan</DialogTitle>
			<DialogContent>
				{loading ? (
					<CircularProgress />
				) : success ? (
					<Typography variant="body1" color="green">
						Success! Your plan has been selected.
					</Typography>
				) : (
					<Typography variant="body1">
						Are you sure you want to choose this plan?
					</Typography>
				)}
			</DialogContent>
			{!loading && !success && (
				<DialogActions>
					<Button onClick={onClose} color="primary">
						Cancel
					</Button>
					<Button onClick={handleConfirm} color="secondary" variant="contained">
						Confirm
					</Button>
				</DialogActions>
			)}
		</Dialog>
	);
};

export default PricingPlanModal;
