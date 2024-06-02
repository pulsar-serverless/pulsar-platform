"use client";
import {
	Box,
	Drawer,
	List,
	ListItem,
	ListItemButton,
	ListItemIcon,
	ListItemText,
} from "@mui/material";
import { ReactNode } from "react";

import HomeRoundedIcon from "@mui/icons-material/HomeRounded";
import SettingsRoundedIcon from "@mui/icons-material/SettingsRounded";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
import InsertChartRoundedIcon from "@mui/icons-material/InsertChartRounded";
import  MemoryRoundedIcon  from "@mui/icons-material/MemoryRounded";
import Link from "next/link";
import { usePathname } from "next/navigation";

export default function Layout({ children }: { children: ReactNode }) {
  const navLinks: { link: string; label: string; icon: ReactNode }[] = [
    { link: "home", label: "Home", icon: <HomeRoundedIcon /> },
    { link: "analytics", label: "Analytics", icon: <InsertChartRoundedIcon /> },
    { link: "resources", label: "Resources", icon: <MemoryRoundedIcon /> },
    { link: "deployment", label: "Deployment", icon: <CloudUploadIcon /> },
  ];
  const url = usePathname();

  return (
    <>
      <Drawer
        variant="permanent"
        anchor="left"
        sx={{
          width: 265,
          flexShrink: 0,
          "& .MuiDrawer-paper": {
            p: 1,
            position: "initial",
            width: 265,
            boxSizing: "border-box",
          },
        }}
      >
        <List>
          {navLinks.map((link, index) => (
            <ListItem key={link.link} disablePadding>
              <ListItemButton
                LinkComponent={Link}
                href={link.link}
                selected={url.includes(link.link)}
                color="secondary"
              >
                <ListItemIcon>{link.icon}</ListItemIcon>
                <ListItemText primary={link.label} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>

        <List sx={{ mt: "auto" }}>
          <ListItem disablePadding>
            <ListItemButton
              LinkComponent={Link}
              href={"settings"}
              selected={url.includes("settings")}
              color="secondary"
            >
              <ListItemIcon>{<SettingsRoundedIcon />}</ListItemIcon>
              <ListItemText primary={"Settings"} />
            </ListItemButton>
          </ListItem>
        </List>
      </Drawer>
      <Box component="main" sx={{ flexGrow: 1, overflow: "scroll" }}>
        {children}
      </Box>
    </>
  );
}
