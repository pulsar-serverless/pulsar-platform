// ResourceCards.tsx
import React from "react";
import { Card, CardContent, Typography } from "@mui/material";

interface TotalResourceUtil {
	mem_usage_mb: number;
	net_usage_mb: number;
	project_id: string;
	usage_period: string;
}

interface ResourceCardsProps {
	totalResourceUtil: TotalResourceUtil | null; // Explicitly type the prop
}

const ResourceCards: React.FC<ResourceCardsProps> = ({ totalResourceUtil }) => {
	return (
		<div>
			{totalResourceUtil && (
				<div className="parent-card">
					{/* Card for mem_usage_mb */}
					<Card className="card">
						<CardContent className="card-body">
							<Typography variant="h5" component="div" className="card-title">
								Memory Usage
							</Typography>
							<Typography
								variant="body2"
								color="text.secondary"
								className="card-text">
								hi
								{/* {totalResourceUtil.mem_usage_mb} Megabytes */}
							</Typography>
						</CardContent>
					</Card>

					{/* Card for net_usage_mb */}
					<Card className="card">
						<CardContent className="card-body">
							<Typography variant="h5" component="div" className="card-title">
								Network Usage
							</Typography>
							<Typography
								variant="body2"
								color="text.secondary"
								className="card-text">
								hi
								{/* {totalResourceUtil.net_usage_mb} Megabytes */}
							</Typography>
						</CardContent>
					</Card>

					{/* Card for usage_period */}
					<Card className="card">
						<CardContent className="card-body">
							<Typography variant="h5" component="div" className="card-title">
								Usage Period
							</Typography>
							<Typography
								variant="body2"
								color="text.secondary"
								className="card-text">
								hello
								{/* {totalResourceUtil.usage_period} */}
							</Typography>
						</CardContent>
					</Card>
				</div>
			)}
		</div>
	);
};

export default ResourceCards;
