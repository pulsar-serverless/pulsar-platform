import {
  Card,
  CardActionArea,
  CardContent,
  Stack,
  Typography,
} from "@mui/material";
import AddRoundedIcon from "@mui/icons-material/AddRounded";
import Link from "next/link";

function CreateProjectCard() {
  return (
    <Card>
      <CardActionArea LinkComponent={Link} href="?action=create-project">
        <CardContent
          sx={{
            display: "grid",
            placeItems: "center",
            p: 3,
            minHeight: 232,
          }}
        >
          <Stack alignItems={"center"}>
            <AddRoundedIcon sx={{ fontSize: 48, mb: 2 }} />
            <Typography variant="body2">Create a Project</Typography>
          </Stack>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}

export default CreateProjectCard;
