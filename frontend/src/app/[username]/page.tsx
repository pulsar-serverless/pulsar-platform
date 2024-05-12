"use client";

import {
  Box, Container,
  Grid,
  Pagination, Typography
} from "@mui/material";
import { useState } from "react";

import CreateProjectCard from "@/components/project/CreateProjectCard";
import ProjectCard from "@/components/project/ProjectCard";
import { useQuery } from "@tanstack/react-query";
import { ProjectApi } from "@/api/projects";
import { useParams } from "next/navigation";

export default function Page() {
  const [page, setPage] = useState(1);
  const handleChange = (event: React.ChangeEvent<unknown>, value: number) => {
    setPage(value);
  };

  const { username } = useParams<{ username: string }>();

  const { data: projects } = useQuery({
    queryKey: [ProjectApi.getProjects.name, page, username],
    queryFn: () => ProjectApi.getProjects(page , 10, username),
  });

  return (
    <Container maxWidth="md" sx={{ py: 3,overflow: 'auto' }}>
      <Typography variant="h6" sx={{ textTransform: "capitalize" }}>
        Recent Projects
      </Typography>

      <Grid container sx={{ mt: 0.5 }} spacing={4}>
        <Grid item lg={4} sm={6} xs={12}>
          <CreateProjectCard />
        </Grid>

        {projects?.rows.map((project, index) => (
          <Grid item lg={4} sm={6} xs={12} key={`project-${index}`}>
            <ProjectCard project={project} />
          </Grid>
        ))}
      </Grid>

      <Box mt={6}>
        <Pagination
          count={projects?.totalPages || 0}
          page={page}
          onChange={handleChange}
          variant="outlined"
          color="primary"
        />
      </Box>
    </Container>
  );
}
