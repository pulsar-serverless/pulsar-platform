import { Project } from "@/models/project";
import { Card, CardActionArea, CardContent, Typography } from "@mui/material";
import Link from "next/link";

function ProjectCard({ project }: { project: Project }) {
  return (
    <Card>
      <CardActionArea LinkComponent={Link} href={`username/${project.id}/home`}>
        <CardContent
          sx={{
            px: 3,
            py: 2,
            minHeight: 200,
          }}
        >
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
        </CardContent>
      </CardActionArea>
    </Card>
  );
}

export default ProjectCard;
