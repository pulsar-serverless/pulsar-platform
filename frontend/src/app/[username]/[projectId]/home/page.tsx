"use client";

import { ProjectApi } from "@/api/projects";
import { CustomTabPanel } from "@/components/layout/CustomTabPanel";
import { Box, Container, Tab, Tabs } from "@mui/material";
import { useQuery } from "@tanstack/react-query";
import { useParams, useRouter } from "next/navigation";
import { useState } from "react";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import { ProjectDetailsCard } from "@/components/project/ProjectDetailsCard";
import { useSnackbar } from "@/hooks/useSnackbar";
import { EnvVariablesForm } from "@/components/project/EnvVariablesForm";
import ProjectLog from "@/components/log/ProjectLog";

function a11yProps(index: number) {
  return {
    id: `simple-tab-${index}`,
    "aria-controls": `simple-tabpanel-${index}`,
  };
}

export default function Page() {
  dayjs.extend(relativeTime);

  const { projectId: projectName } = useParams<{ projectId: string }>();
  const snackbar = useSnackbar();
  const router = useRouter();

  const {
    data: project,
    isError,
    isLoading,
  } = useQuery({
    queryKey: [ProjectApi.getProject.name, projectName],
    queryFn: () => ProjectApi.getProject(projectName),
  });

  const [value, setValue] = useState(0);

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

  if (isError) {
    snackbar.setErrorMsg("Project not found");
    // TODO: change this
    router.push("/username");
  }

  return (
    <>
      <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
        <Container maxWidth="md" sx={{}}>
          <Tabs
            value={value}
            onChange={handleChange}
            aria-label="basic tabs example"
          >
            <Tab label="Project" {...a11yProps(0)} />
            <Tab label="Logs" {...a11yProps(1)} />
            <Tab label="Analytics" {...a11yProps(2)} />
            <Tab label="Environmental variables" {...a11yProps(3)} />
          </Tabs>
        </Container>
      </Box>
      <Container maxWidth="md" sx={{ my: 3 }}>
        <CustomTabPanel value={value} index={0}>
          <ProjectDetailsCard isLoading={isLoading} project={project} />
        </CustomTabPanel>
        <CustomTabPanel value={value} index={1}>
          <ProjectLog projectId={projectName}/>
        </CustomTabPanel>
        <CustomTabPanel value={value} index={3}>
          <EnvVariablesForm projectID={projectName} />
        </CustomTabPanel>
      </Container>
    </>
  );
}
