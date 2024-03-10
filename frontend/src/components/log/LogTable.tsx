import { Log } from "@/models/log";
import {
  Table,
  TableBody,
  TableRow,
  TableCell,
  Typography,
  Pagination,
} from "@mui/material";
import dayjs from "dayjs";
import { FC } from "react";

const LogTable: FC<{ logs: Log[]; page: number; count: number; onPaginate: (page: number) => void }> = ({
  logs,
  page,
  count,
  onPaginate
}) => {
  return (
    <>
      <Table size="small">
        <TableBody sx={{ fontFamily: '"Source Code Pro", monospace' }}>
          {logs.map((log) => (
            <TableRow key={log.id}>
              <TableCell
                component="th"
                sx={{
                  py: 0,
                  borderBottom: 0,
                  verticalAlign: "unset",
                  whiteSpace: "nowrap",
                }}
              >
                <Typography
                  variant="body2"
                  color="secondary.dark"
                  fontFamily='"Source Code Pro", monospace'
                >
                  {dayjs(log.createdAt).format("YYYY-MM-DD HH-mm-ss")}
                </Typography>
              </TableCell>
              <TableCell
                align="left"
                scope="row"
                sx={{ py: 0, borderBottom: 0, width: "100%", lineHeight: 1.5 }}
              >
                <Typography
                  variant="body2"
                  fontFamily='"Source Code Pro", monospace'

                  color={
                    log.type == "Info"
                      ? "secondary"
                      : log.type == "Error"
                      ? "error.light"
                      : "warning.light"
                  }
                >
                  {log.message}
                </Typography>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <Pagination
        count={count}
        page={page}
        onChange={(_, value) =>  onPaginate(value)}
        size="small"
        sx={{ ".MuiPagination-ul": { justifyContent: "end" }, p: 2 }}
        showFirstButton showLastButton 
      />
    </>
  );
};

export default LogTable;
