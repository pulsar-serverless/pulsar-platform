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
import { useQuery, keepPreviousData, useMutation } from "@tanstack/react-query";
import { userApi } from "@/api/user";
import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { User } from "@/models/user";
import { AuthGuard } from "@/components/providers/AuthGuard";
import { useRouter } from "next/navigation";
import { useSnackbar } from "@/hooks/useSnackbar";
import { queryClient } from "@/components/providers/QueryProvider";
import { ConfirmationDialog } from "@/components/modals/ConfirmationDialog";

const columnHelper = createColumnHelper<User>();

const TableMenu: React.FC<{ user: User; onDelete: () => void }> = ({ user, onDelete }) => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);
  const handleClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };
  const router = useRouter();
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
        <MenuItem
          onClick={() => {
            router.push(`/${user.userId}`);
            handleClose();
          }}
        >
          View projects
        </MenuItem>
        <MenuItem onClick={() => {onDelete(); handleClose();}} color="error">
          Remove All projects
        </MenuItem>
        <MenuItem onClick={handleClose} color="error.light">
          Pause All projects
        </MenuItem>
      </Menu>
    </>
  );
};

const Page = () => {
  const [pageCount, setPageCount] = useState(0);
  const [searchQuery, setSearchQuery] = useState("");
  const [data, setData] = useState<User[]>([]);

  const [confirmDeleteAllProject, setConfirmDeleteAllProject] = useState<undefined | string>();

  const snackbar = useSnackbar();
  const { mutate: handleDeleteAllProjects } = useMutation({
    mutationFn: userApi.deleteAllProjects,
    onSuccess: () => {
      snackbar.setSuccessMsg("All Projects of the user deleted successfully!");
      queryClient.invalidateQueries({ queryKey: [userApi.getUsers.name] });
    },
    onError: () => snackbar.setErrorMsg("Unable to delete user's projects."),
  });

  const columns = useMemo(
    () => [
      columnHelper.accessor("userId", {
        cell: (row) => <>{row.getValue()}</>,
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
        cell: (props) => <TableMenu user={props.row.original} onDelete={() => setConfirmDeleteAllProject(props.row.original.userId)}/>,
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
    <AuthGuard role="Admin">
      <Container maxWidth="md" sx={{ py: 3 }}>
        <Typography mb={2.5} variant="h5" component="div">
          Users
        </Typography>
        <Typography variant="body2">
          Effortlessly manage users and their projects.
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
            onRowsPerPageChange={(e) =>
              table.setPageSize(parseInt(e.target.value, 10))
            }
          />
        </TableContainer>
      </Container>
      {confirmDeleteAllProject && (
        <ConfirmationDialog
          open={!!confirmDeleteAllProject}
          title="Delete All Projects"
          description="Deleting all projects will permanently erase all records of all projects of a user."
          handleClose={() => setConfirmDeleteAllProject(undefined)}
          handleConfirm={() => {
            handleDeleteAllProjects(confirmDeleteAllProject);
            setConfirmDeleteAllProject(undefined);
          }}
        />
      )}
    </AuthGuard>
  );
};

export default Page;
