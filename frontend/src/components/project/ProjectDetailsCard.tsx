import {
  Card,
  CardMedia,
  Skeleton,
  CardContent,
  Typography,
  Box,
  List,
  ListItem,
  ListItemText,
  Stack,
  IconButton,
  useTheme,
} from "@mui/material";
import LaunchRoundedIcon from "@mui/icons-material/LaunchRounded";
import dayjs from "dayjs";
import { Project } from "@/models/project";

export const ProjectDetailsCard = ({
  project,
  isLoading,
}: {
  project?: Project | undefined;
  isLoading: boolean;
}) => {
  const theme = useTheme();

  const url = new URL(process.env.NEXT_PUBLIC_SERVERLESS_URL!);
  if (project) url.hostname = project?.subdomain + "." + url.hostname;

  return (
    <>
      <Typography
        variant="h6"
        sx={{ textTransform: "capitalize" }}
        gutterBottom
      >
        Project Details
      </Typography>
      <Card
        sx={{
          width: "100%",
          display: "flex",
          mt: 4,
        }}
      >
        <CardMedia
          sx={{
            minWidth: 350,
            aspectRatio: 1.6,
          }}
        >
          <Skeleton variant="rectangular" width={"100%"} height={"100%"} />
        </CardMedia>

        <CardContent sx={{ width: "100%" }}>
          {!isLoading ? (
            <Typography variant="h6" sx={{ textTransform: "capitalize" }}>
              {project?.name}
            </Typography>
          ) : (
            <Skeleton
              variant="text"
              sx={{ fontSize: theme.typography.h6.fontSize, width: 240 }}
            />
          )}
          <Box sx={{ mt: 2 }}>
            <List sx={{ width: "100%", p: 0, ".MuiListItem-root": { px: 0 } }}>
              <ListItem>
                <ListItemText
                  primaryTypographyProps={{ gutterBottom: true }}
                  primary="Created"
                  secondary={
                    !isLoading ? (
                      dayjs(project?.createdAt).fromNow()
                    ) : (
                      <Skeleton
                        variant="text"
                        sx={{
                          fontSize: theme.typography.body2.fontSize,
                          width: 240,
                        }}
                      />
                    )
                  }
                />
              </ListItem>
              <ListItem>
                <ListItemText
                  primaryTypographyProps={{ gutterBottom: true }}
                  primary="Domain"
                  secondary={
                    !isLoading ? (
                      <Stack direction="row" alignItems={"center"} gap={1}>
                        <Typography variant="body2">{url.href}</Typography>{" "}
                        <IconButton
                          aria-label="delete"
                          LinkComponent={"a"}
                          target="_blank"
                          href={url.href}
                        >
                          <LaunchRoundedIcon fontSize="small" />
                        </IconButton>
                      </Stack>
                    ) : (
                      <Skeleton
                        variant="text"
                        sx={{
                          fontSize: theme.typography.body2.fontSize,
                          width: 240,
                        }}
                      />
                    )
                  }
                />
              </ListItem>
              <ListItem>
                <ListItemText
                  primaryTypographyProps={{ gutterBottom: true }}
                  primary={"Status"}
                  secondary={
                    !isLoading ? (
                      <Stack direction="row" alignItems={"center"} gap={1.5}>
                        <Box
                          sx={{
                            width: "1em",
                            height: "1em",
                            bgcolor: statusColor(project!.deploymentStatus),
                            borderRadius: "50%",
                          }}
                        ></Box>
                        <Typography
                          variant="body2"
                          textTransform={"capitalize"}
                        >
                          {project?.deploymentStatus}
                        </Typography>
                      </Stack>
                    ) : (
                      <Skeleton
                        variant="text"
                        sx={{
                          fontSize: theme.typography.body2.fontSize,
                          width: 240,
                        }}
                      />
                    )
                  }
                />
              </ListItem>
            </List>
          </Box>
        </CardContent>
      </Card>
    </>
  );
};

export const statusColor = (status: Project["deploymentStatus"]) => {
  return status == "failed"
    ? "error.main"
    : status == "building"
    ? "primary.main"
    : status == "done"
    ? "success.main"
    : "grey.500";
};
