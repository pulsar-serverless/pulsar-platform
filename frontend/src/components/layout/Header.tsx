"use client";
import {
  useTheme,
  AppBar,
  Toolbar,
  Stack,
  Typography,
  Button,
  Avatar,
  IconButton,
  Menu,
  MenuItem,
  Box,
  Divider,
  ListItemIcon,
  ListItemText,
  Tooltip,
} from "@mui/material";
import AddRoundedIcon from "@mui/icons-material/AddRounded";
import LogoutRoundedIcon from "@mui/icons-material/LogoutRounded";
import PersonIcon from "@mui/icons-material/Person";
import { Logo } from "./Logo";
import { useState } from "react";
import Link from "next/link";

export const Header = () => {
  const theme = useTheme();

  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);
  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  const username = "dagem";

  return (
    <AppBar
      elevation={0}
      sx={{
        borderBottom: `1px solid ${theme.palette.divider}`,
        zIndex: 1000,
      }}
    >
      <Toolbar>
        <Button LinkComponent={Link} href="/" color="secondary">
          <Stack direction={"row"} alignItems={"center"} gap={2}>
            <Logo />
            <Typography variant="h5">Pulsar</Typography>
          </Stack>
        </Button>

        <Stack
          direction="row"
          sx={{ ml: "auto" }}
          gap={3}
          alignItems={"center"}
        >
          <Button LinkComponent={Link} href="/docs" color="secondary">
            Docs
          </Button>
          <Button
            LinkComponent={Link}
            href={`/${username}/projects`}
            color="secondary"
          >
            Projects
          </Button>
          <Link shallow href="?action=create-project">
            <Button variant="contained" startIcon={<AddRoundedIcon />}>
              Create Project
            </Button>
          </Link>
          <IconButton
            id="profile-button"
            aria-controls={open ? "profile-menu" : undefined}
            aria-haspopup="true"
            aria-expanded={open ? "true" : undefined}
            onClick={handleClick}
          >
            <Avatar sx={{ width: 24, height: 24 }} />
          </IconButton>

          <Menu
            id="profile-menu"
            anchorEl={anchorEl}
            open={open}
            onClose={handleClose}
            MenuListProps={{
              "aria-labelledby": "profile-button",
            }}
            sx={{
              "& .MuiList-root": { p: 0 },
              "& .MuiPaper-root": { minWidth: 240, maxWidth: 290 },
            }}
          >
            <Stack
              direction={"row"}
              gap={5}
              alignItems={"center"}
              sx={{ width: "100%", justifyContent: "space-between", p: 2 }}
            >
              <Box>
                <Typography
                  sx={{ whiteSpace: "nowrap" }}
                  variant="subtitle1"
                  gutterBottom
                >
                  Dagem Tadesse
                </Typography>
                <Typography variant="body2">Developer</Typography>
              </Box>
              <Tooltip title="Logout">
                <IconButton>
                  <LogoutRoundedIcon />
                </IconButton>
              </Tooltip>
            </Stack>
            <Divider />
            <MenuItem sx={{ py: 2 }}>
              <ListItemIcon>
                <PersonIcon />
              </ListItemIcon>
              <ListItemText primary={"View Profile"} />
            </MenuItem>
          </Menu>
        </Stack>
      </Toolbar>
    </AppBar>
  );
};
