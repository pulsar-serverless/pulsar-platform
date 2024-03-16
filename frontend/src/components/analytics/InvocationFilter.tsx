import {
  Stack,
  FormControl,
  Select,
  MenuItem,
  InputAdornment,
} from "@mui/material";
import CalendarMonthIcon from "@mui/icons-material/CalendarMonth";

const InvocationFilter: React.FC<{
  status: InvocationStatus;
  setStatus: (status: InvocationStatus) => void;
  graphType: InvocationGraphType;
  setGraphType: (graphType: InvocationGraphType) => void;
}> = ({ status, graphType, setStatus, setGraphType }) => {
  return (
    <Stack direction={"row"} justifyContent={"end"} mb={2} gap={2}>
      <FormControl size="small" sx={{ minWidth: 180 }}>
        <Select
          value={status}
          onChange={(e) => setStatus(e.target.value as InvocationStatus)}
        >
          <MenuItem value={"Success"}>Success</MenuItem>
          <MenuItem value={"Error"}>Error</MenuItem>
        </Select>
      </FormControl>
      <FormControl size="small" sx={{ minWidth: 180 }}>
        <Select
          value={graphType}
          onChange={(e) => setGraphType(e.target.value as InvocationGraphType)}
          startAdornment={
            <InputAdornment position="start">
              <CalendarMonthIcon />
            </InputAdornment>
          }
        >
          <MenuItem value={'Day'}>Last 24 Hours</MenuItem>
          <MenuItem value={'Week'}>Last 7 Days</MenuItem>
          <MenuItem value={'Month'}>Last 30 Days</MenuItem>
        </Select>
      </FormControl>
    </Stack>
  );
};

export default InvocationFilter;
