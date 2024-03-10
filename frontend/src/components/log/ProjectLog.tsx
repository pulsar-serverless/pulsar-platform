import {
  Box,
  Card,
  CardContent,
  FormControl,
  IconButton,
  InputAdornment,
  MenuItem,
  OutlinedInput,
  Select,
  Stack,
  Typography,
} from "@mui/material";
import SearchRoundedIcon from "@mui/icons-material/SearchRounded";
import HorizontalSplitRoundedIcon from "@mui/icons-material/HorizontalSplitRounded";
import LogTable from "./LogTable";
import { keepPreviousData, useQuery } from "@tanstack/react-query";
import { LogApi } from "@/api/log";
import { useState } from "react";

const ProjectLog: React.FC<{ projectId: string }> = ({ projectId }) => {
  const [page, setPage] = useState(1);

  const { data: logs } = useQuery({
    queryKey: [LogApi.getLogs.name, projectId, page],
    queryFn: () => LogApi.getLogs(projectId, page),
    placeholderData: keepPreviousData,
  });

  return (
    <>
      <Typography
        variant="h6"
        sx={{ textTransform: "capitalize" }}
        gutterBottom
      >
        Project logs
      </Typography>
      <Typography variant="body2">
        Effortlessly monitor executions and track errors.
      </Typography>

      <Card sx={{ mt: 4 }}>
        <CardContent sx={{ p: 1, pb: 0, "&:last-child": { pb: 0 } }}>
          <Stack
            component="form"
            direction={"row"}
            alignItems={"center"}
            sx={{ mb: 2 }}
            gap={2}
          >
            <FormControl sx={{ minWidth: 120 }} size="small">
              <Select displayEmpty value="AllLogs">
                <MenuItem value={"AllLogs"}>All Logs</MenuItem>
                <MenuItem value={"Warning"}>Warning</MenuItem>
                <MenuItem value={"Error"}>Error</MenuItem>
                <MenuItem value={"Info"}>Info</MenuItem>
              </Select>
            </FormControl>
            <FormControl variant="outlined" size="small" sx={{ flexGrow: 1 }}>
              <OutlinedInput
                endAdornment={
                  <InputAdornment position="end">
                    <IconButton>
                      <SearchRoundedIcon />
                    </IconButton>
                  </InputAdornment>
                }
              />
            </FormControl>
          </Stack>
          {!logs ? (
            <EmptyLogsState />
          ) : (
            <LogTable
              logs={logs.rows!}
              count={logs.totalPages}
              page={logs.page}
              onPaginate={page => setPage(page)}
            />
          )}
        </CardContent>
      </Card>
    </>
  );
};

const EmptyLogsState = () => {
  return (
    <Box sx={{ p: 6, display: "grid", placeItems: "center" }}>
      <HorizontalSplitRoundedIcon sx={{ fontSize: 48, mb: 2 }} />
      <Typography variant="body2">There are no logs at this time</Typography>
    </Box>
  );
};

export default ProjectLog;
