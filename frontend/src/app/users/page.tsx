"use client";

import {
  Button,
  Container,
  FormControl,
  IconButton,
  InputAdornment,
  Menu,
  MenuItem,
  OutlinedInput,
  Paper,
  Stack,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TablePagination,
  TableRow,
  Typography,
} from "@mui/material";
import { useEffect, useMemo, useState } from "react";

import SearchRoundedIcon from "@mui/icons-material/SearchRounded";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import React from "react";
import { useQuery, keepPreviousData } from "@tanstack/react-query";
import { userApi } from "@/api/user";
import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { User } from "@/models/user";

const columnHelper = createColumnHelper<User>();

const TableMenu = () => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(
    null
  );
  const open = Boolean(anchorEl);
  const handleClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };
  return (
    <>
      <IconButton onClick={handleClick} size="small">
        <MoreVertIcon />
      </IconButton>
      <Menu
        id="long-menu"
        MenuListProps={{
          "aria-labelledby": "long-button",
        }}
        anchorEl={anchorEl}
        open={open}
        onClose={handleClose}
      >
        <MenuItem onClick={handleClose}>View projects</MenuItem>
        <MenuItem onClick={handleClose} color="error">Remove All projects</MenuItem>
        <MenuItem onClick={handleClose} color="error.light">Pause All projects</MenuItem>
      </Menu>
    </>
  );
}

const Page = () => {
  const [pageCount, setPageCount] = useState(0);
  const [searchQuery, setSearchQuery] = useState('');
  const [data, setData] = useState<User[]>([]);

  const columns = useMemo(
    () => [
      columnHelper.accessor("userId", {
        cell: (row) => (
          <Button variant="text" size="small" color="secondary">
            {row.getValue()}
          </Button>
        ),
        header: "Identifier",
      }),
      columnHelper.accessor("email", {
        header: "Email",
      }),
      columnHelper.accessor("projectCount", {
        header: "Projects",
      }),
      columnHelper.display({
        id: "actions",
        header: "Actions",
        cell: () => <TableMenu/>
      }),
    ],
    []
  );

  const table = useReactTable({
    columns,
    data,
    initialState: {
      pagination: {
        pageIndex: 0,
        pageSize: 10,
      },
    },
    manualPagination: true,
    getCoreRowModel: getCoreRowModel(),
    pageCount,
  });

  const {
    pagination: { pageIndex, pageSize },
  } = table.getState();

  const { data: users } = useQuery({
    queryKey: [userApi.getUsers.name, pageSize, pageIndex, searchQuery],
    queryFn: () => userApi.getUsers(pageSize, pageIndex + 1, searchQuery),
    placeholderData: keepPreviousData,
  });

  useEffect(() => {
    if (users) {
      setData(users.rows);
      setPageCount(users.totalPages);
    }
  }, [users]);
  return (
    <>
      <Container maxWidth="md" sx={{ py: 3 }}>
        <Typography mb={2.5} variant="h5" component="div">
          Users
        </Typography>
        <Typography variant="body2">
          Effortlessly monitor executions and track errors.
        </Typography>

        <Stack alignItems="end" mb={1.5} mt={4}>
          <FormControl variant="outlined" size="small" sx={{ flexGrow: 1 }}>
            <OutlinedInput
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              placeholder="search by user id"
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
        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              {table.getHeaderGroups().map((headerGroup) => (
                <TableRow key={headerGroup.id}>
                  {headerGroup.headers.map((header) => (
                    <TableCell key={header.id} sx={{ p: 2 }}>
                      {flexRender(
                        header.column.columnDef.header,
                        header.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))}
            </TableHead>
            <TableBody>
              {table.getRowModel().rows.map((row) => {
                return (
                  <TableRow key={row.id}>
                    {row.getVisibleCells().map((cell) => (
                      <TableCell key={cell.id} sx={{ px: 2, py: 1.5 }}>
                        {flexRender(
                          cell.column.columnDef.cell,
                          cell.getContext()
                        )}
                      </TableCell>
                    ))}
                  </TableRow>
                );
              })}
            </TableBody>
          </Table>
          <TablePagination
            component="div"
            count={pageCount}
            page={pageIndex}
            onPageChange={(_, page) => table.setPageIndex(page)}
            rowsPerPage={pageSize}
            onRowsPerPageChange={(e) => table.setPageSize(parseInt(e.target.value, 10))}
          />
        </TableContainer>
      </Container>
    </>
  );
};

export default Page;
