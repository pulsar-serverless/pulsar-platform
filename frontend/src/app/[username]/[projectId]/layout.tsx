"use client";
import {
  Box,
  Container,
  Drawer,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Stack,
} from "@mui/material";
import { ReactNode } from "react";

import HomeRoundedIcon from "@mui/icons-material/HomeRounded";
import TopicRoundedIcon from "@mui/icons-material/TopicRounded";
import CreditCardRoundedIcon from "@mui/icons-material/CreditCardRounded";
import SettingsRoundedIcon from "@mui/icons-material/SettingsRounded";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";

export default function Layout({ children }: { children: ReactNode }) {
  const navLinks: { link: string; label: string; icon: ReactNode }[] = [
    { link: "projects", label: "Home", icon: <HomeRoundedIcon /> },
    { link: "assets", label: "Assets", icon: <TopicRoundedIcon /> },
    { link: "billing", label: "Billing", icon: <CreditCardRoundedIcon /> },
    { link: "deployment", label: "Deployment", icon: <CloudUploadIcon /> },
  ];

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
              <ListItemButton>
                <ListItemIcon>{link.icon}</ListItemIcon>
                <ListItemText primary={link.label} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>

        <List sx={{ mt: "auto" }}>
          <ListItem disablePadding>
            <ListItemButton>
              <ListItemIcon>{<SettingsRoundedIcon />}</ListItemIcon>
              <ListItemText primary={"Settings"} />
            </ListItemButton>
          </ListItem>
        </List>
      </Drawer>
      <Box component="main" sx={{ flexGrow: 1 }}>
        <Container>{children}</Container>
      </Box>
    </>
  );
}
