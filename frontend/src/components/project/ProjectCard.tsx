import { Project } from "@/models/project";
import {
  Box,
  Card,
  CardActionArea,
  CardContent,
  Stack,
  Typography,
} from "@mui/material";
import Link from "next/link";
import { statusColor } from "./ProjectDetailsCard";
import { useAuth0 } from "@auth0/auth0-react";

function ProjectCard({ project }: { project: Project }) {
  const { isAuthenticated, logout, user } = useAuth0();
  const username = user?.sub || ""
  return (
    <Card>
      <CardActionArea LinkComponent={Link} href={`${username}/${project.id}/home`}>
        <CardContent
          sx={{
            px: 3,
            py: 2,
          }}
        >
          <Stack sx={{ minHeight: 200 }}>
            <Typography
              variant="h6"
              gutterBottom
              sx={{ fontWeight: "medium", textTransform: "capitalize" }}
            >
              {project.name}
            </Typography>
            <Typography
              variant="subtitle2"
              color={"secondary"}
              sx={{ fontWeight: "light" }}
            >
              {project.id}
            </Typography>
            <Box sx={{ flexGrow: 1 }}></Box>
            <Stack direction="row" alignItems={"center"} gap={1.5}>
              <Box
                sx={{
                  width: "1em",
                  height: "1em",
                  bgcolor: statusColor(project!.deploymentStatus),
                  borderRadius: "50%",
                }}
              ></Box>
              <Typography
                variant="body2"
                textTransform={"capitalize"}
              >
                {project?.deploymentStatus}
              </Typography>
            </Stack>
          </Stack>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}

export default ProjectCard;
