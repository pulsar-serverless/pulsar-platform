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
  Button
} from "@mui/material";
import SearchRoundedIcon from "@mui/icons-material/SearchRounded";
import HorizontalSplitRoundedIcon from "@mui/icons-material/HorizontalSplitRounded";
import DeleteForeverRoundedIcon from "@mui/icons-material/DeleteForeverRounded";
import LogTable from "./LogTable";
import { keepPreviousData, useMutation, useQuery } from "@tanstack/react-query";
import { LogApi } from "@/api/log";
import { useState } from "react";
import { ConfirmationDialog } from "../modals/ConfirmationDialog";
import { useSnackbar } from "@/hooks/useSnackbar";
import { queryClient } from "../providers/QueryProvider";
import { Log } from "@/models/log";

const ProjectLog: React.FC<{ projectId: string }> = ({ projectId }) => {
  const [page, setPage] = useState(1);
  const [clearLog, setClearLog] = useState(false);

  const [logType, setLogType] = useState<Log['type'][]>(['Info' , 'Warning' , 'Error']);
  const [searchQuery, setSearchQuery] = useState("");

  const snackbar = useSnackbar();

  const { data: logs } = useQuery({
    queryKey: [LogApi.getLogs.name, projectId, page, searchQuery, logType],
    queryFn: () => LogApi.getLogs(projectId, logType, searchQuery, page,),
    placeholderData: keepPreviousData,
  });

  const { mutate: handleClearLogs } = useMutation({
    mutationFn: LogApi.deleteLogs,
    onSuccess: () => {
      snackbar.setSuccessMsg("Project log cleared successfully!");
      queryClient.invalidateQueries({ queryKey: [LogApi.getLogs.name] });
    },
    onError: () => snackbar.setErrorMsg("Unable to clear project logs."),
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
            <FormControl sx={{ minWidth: 180 }} size="small">
              <Select
                displayEmpty
                value={logType}
                multiple
                onChange={(e) => setLogType(e.target.value as any)}
              >
                <MenuItem value={"Warning"}>Warning</MenuItem>
                <MenuItem value={"Error"}>Error</MenuItem>
                <MenuItem value={"Info"}>Info</MenuItem>
              </Select>
            </FormControl>
            <FormControl variant="outlined" size="small" sx={{ flexGrow: 1 }}>
              <OutlinedInput
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                endAdornment={
                  <InputAdornment position="end">
                    <IconButton>
                      <SearchRoundedIcon />
                    </IconButton>
                  </InputAdornment>
                }
              />
            </FormControl>
            <Button
              variant="contained"
              color="error"
              startIcon={<DeleteForeverRoundedIcon />}
              onClick={() => setClearLog(true)}
            >
              Clear all logs
            </Button>
          </Stack>
          {!logs || logs.rows.length == 0 ? (
            <EmptyLogsState />
          ) : (
            <LogTable
              logs={logs.rows!}
              count={logs.totalPages}
              page={logs.page}
              onPaginate={(page) => setPage(page)}
            />
          )}
        </CardContent>
      </Card>

      {clearLog && (
        <ConfirmationDialog
          open={clearLog}
          title="Clear All Logs"
          description="Deleting Project Logs will permanently erase all recorded project logs. "
          handleClose={() => setClearLog(false)}
          handleConfirm={() => {
            handleClearLogs(projectId);
            setClearLog(false);
          }}
        />
      )}
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
