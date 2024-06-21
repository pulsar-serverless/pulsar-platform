"use client";
import InvocationGraph from "@/components/analytics/InvocationsGraph";
import ResourceCards from "@/components/resources/ResourceCards";
import { Container, Stack, Typography } from "@mui/material";
import { useParams } from "next/navigation";

function Page() {
	const { projectId } = useParams<{ projectId: string }>();

	return (
		<Container sx={{ py: 3 }} maxWidth="md">
			<Typography
				variant="h5"
				sx={{ textTransform: "capitalize" }}
				gutterBottom>
				Analytics
			</Typography>
			<Typography variant="body2">
				Effortlessly monitor executions and track errors.
			</Typography>

			<Stack gap={3} mt={2.5}>
				<ResourceCards projectId={projectId} />
			</Stack>

			<Typography
				variant="subtitle1"
				sx={{ textTransform: "capitalize" }}
				fontWeight={"medium"}
				gutterBottom
				mt={3}>
				Invocations
			</Typography>
			<InvocationGraph projectId={projectId} />
		</Container>
	);
}

export default Page;
