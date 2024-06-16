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
import LogoutRoundedIcon from "@mui/icons-material/LogoutRounded";
import PersonIcon from "@mui/icons-material/Person";
import { Logo } from "./Logo";
import { useState } from "react";
import Link from "next/link";
import { useAuth0 } from "@auth0/auth0-react";

export const Header = () => {
  const theme = useTheme();
  const { isAuthenticated, logout, user } = useAuth0();

  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);
  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  const username = user?.sub || ""

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
            Documentation
          </Button>
          <Button LinkComponent={Link} href="/pricing" color="secondary">
            Pricing
          </Button>
          {isAuthenticated && (
            <>
              <Button
                LinkComponent={Link}
                href={`/${username}`}
                color="secondary"
              >
                Projects
              </Button>
              <IconButton
                id="profile-button"
                aria-controls={open ? "profile-menu" : undefined}
                aria-haspopup="true"
                aria-expanded={open ? "true" : undefined}
                onClick={handleClick}
              >
                <Avatar sx={{ width: 24, height: 24 }} src={user?.picture} />
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
                      {user?.name}
                    </Typography>
                    <Typography variant="body2">
                      {user?.roleType.includes("Admin") ? "Admin" : "Developer"}
                    </Typography>
                  </Box>
                  <Tooltip title="Logout">
                    <IconButton
                      onClick={() =>
                        logout({
                          logoutParams: {
                            returnTo: process.env.NEXT_PUBLIC_redirectUri,
                          },
                        })
                      }
                    >
                      <LogoutRoundedIcon />
                    </IconButton>
                  </Tooltip>
                </Stack>
              </Menu>
            </>
          )}
        </Stack>
      </Toolbar>
    </AppBar>
  );
};
