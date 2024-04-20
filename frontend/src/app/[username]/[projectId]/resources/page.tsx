"use client";
import ResourceApp from "@/components/resources/cards";
import { Container, Typography } from "@mui/material";
import { useParams } from "next/navigation";

function Page() {
	const { projectId } = useParams<{ projectId: string }>();

	return (
		<Container sx={{ py: 3 }} maxWidth="md">
			<Typography
				variant="h6"
				sx={{ textTransform: "capitalize" }}
				gutterBottom>
				Resources
			</Typography>
			<Typography variant="body2">
				Effortlessly track your resource usage.
			</Typography>

			<ResourceApp projectId={projectId} />
		</Container>
	);
}

export default Page;
