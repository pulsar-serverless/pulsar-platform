// import React from "react";
// import { Bar } from "react-chartjs-2";
// import { ChartOptions } from "chart.js";
// import "./resource.css";

// interface TotalResourceUtil {
// 	mem_usage_mb: number;
// 	net_usage_mb: number;
// 	project_id: string;
// 	usage_period: string;
// }

// interface ResourceGraphProps {
// 	totalResourceUtil: TotalResourceUtil | null;
// }

// const ResourceGraph: React.FC<ResourceGraphProps> = ({ totalResourceUtil }) => {
// 	if (!totalResourceUtil) return null;

// 	const data = {
// 		labels: ["Memory Usage (MB)", "Network Usage (MB)"],
// 		datasets: [
// 			{
// 				label: "Usage",
// 				backgroundColor: ["#36A2EB", "#FFCE56"],
// 				borderColor: "rgba(0,0,0,1)",
// 				borderWidth: 2,
// 				data: [totalResourceUtil.mem_usage_mb, totalResourceUtil.net_usage_mb],
// 			},
// 		],
// 	};
// 	const options: ChartOptions<"bar"> = {
// 		plugins: {
// 			title: {
// 				display: true,
// 				text: "Resource Utilization",
// 			},
// 		},
// 		scales: {
// 			y: {
// 				beginAtZero: true,
// 			},
// 		},
// 	};

// 	return (
// 		<div className="chart-container">
// 			<Bar
// 				data={data}
// 				options={options} // Pass the options object here
// 			/>
// 		</div>
// 	);
// };

// export default ResourceGraph;
