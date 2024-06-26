// ResourceCards.tsx
import React from "react";
import { Card, CardContent, Stack, Typography } from "@mui/material";
import { useQuery } from "@tanstack/react-query";
import { ResourcesApi } from "@/api/resources";

const BYTES_IN_MB = 1048576;

const ResourceCards: React.FC<{ projectId: string }> = ({ projectId }) => {
  const { data: resourceUsage } = useQuery({
    queryKey: [ResourcesApi.getTotalProjectResourceUtil.name],
    queryFn: () => ResourcesApi.getTotalProjectResourceUtil(projectId),
    placeholderData: {
      mem_usage_mb: 0,
      net_usage_mb: 0,
      project_id: projectId,
      usage_period: "",
    },
  });

  return (
    <>
      <Stack direction="row" gap={3}>
        <Card sx={{ flexGrow: 1 }}>
          <CardContent>
            <Typography
              variant="subtitle1"
              sx={{ textTransform: "capitalize" }}
              fontWeight={"medium"}
              component="div"
              gutterBottom
            >
              Memory Usage (MB)
            </Typography>
            <Typography variant="h3">
              {`${((resourceUsage?.mem_usage_mb || 0) / BYTES_IN_MB).toFixed(
                2
              )}`}
            </Typography>
          </CardContent>
        </Card>

        <Card sx={{ flexGrow: 1 }}>
          <CardContent>
            <Typography
              variant="subtitle1"
              sx={{ textTransform: "capitalize" }}
              fontWeight={"medium"}
              gutterBottom
              component="div"
            >
              Network Usage (MB)
            </Typography>
            <Typography variant="h3">
              {`${((resourceUsage?.net_usage_mb || 0) / BYTES_IN_MB).toFixed(
                2
              )}`}
            </Typography>
          </CardContent>
        </Card>

        <Card sx={{ flexGrow: 1 }}>
          <CardContent>
            <Typography
              variant="subtitle1"
              sx={{ textTransform: "capitalize" }}
              fontWeight={"medium"}
              gutterBottom
              component="div"
            >
              Usage Period
            </Typography>
            <Typography variant="h3">{resourceUsage?.usage_period}</Typography>
          </CardContent>
        </Card>
      </Stack>
    </>
  );
};

export default ResourceCards;
