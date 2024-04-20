// ResourceCards.tsx
import React from "react";
import "./resource.css";

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
					<div className="card">
						<div className="card-body">
							<h5 className="card-title">Memory Usage</h5>
							<p className="card-text">
								{totalResourceUtil.mem_usage_mb} Megabytes
							</p>
						</div>
					</div>

					{/* Card for net_usage_mb */}
					<div className="card">
						<div className="card-body">
							<h5 className="card-title">Network Usage</h5>
							<p className="card-text">
								{totalResourceUtil.net_usage_mb} Megabytes
							</p>
						</div>
					</div>

					{/* Card for usage_period */}
					<div className="card">
						<div className="card-body">
							<h5 className="card-title">Usage Period</h5>
							<p className="card-text">{totalResourceUtil.usage_period}</p>
						</div>
					</div>
				</div>
			)}
		</div>
	);
};

export default ResourceCards;
