import React from "react";
import {
  Card,
  CardContent,
  Button,
  List,
  CardActions,
  ListItem,
  ListItemText,
  ListItemAvatar,
  Skeleton,
} from "@mui/material";
import { PricingPlan } from "@/models/PricingPlan";

import NetworkCheckIcon from "@mui/icons-material/NetworkCheck";
import MemoryIcon from "@mui/icons-material/Memory";
import LeakAddIcon from "@mui/icons-material/LeakAdd";
import bytes from "bytes";

export const PricingCardSkeleton: React.FC<{
  onPricePlanSelected?: (plan: PricingPlan) => void;
}> = ({ onPricePlanSelected }) => {
  return (
    <>
      <Card
        sx={{
          flexGrow: 1,
          width: "50%",
        }}
      >
        <CardContent>
          <Skeleton variant="text" sx={{ fontSize: "1.75rem" }} />
          <Skeleton variant="text" sx={{ fontSize: "1.75rem" }} />

          <Skeleton variant="rounded" sx={{ height: 55 }} />

          <List sx={{ my: 2 }}>
            <ListItem>
              <ListItemAvatar>
                <Skeleton variant="circular" width={40} height={40} />
              </ListItemAvatar>
              <ListItemText
                primary={<Skeleton variant="text" height={22} sx={{ mb: 1 }} />}
                secondary={<Skeleton variant="rounded" height={37} />}
              />
            </ListItem>
            <ListItem>
              <ListItemAvatar>
                <Skeleton variant="circular" width={40} height={40} />
              </ListItemAvatar>
              <ListItemText
                primary={<Skeleton variant="text" height={22} sx={{ mb: 1 }} />}
                secondary={<Skeleton variant="rounded" height={37} />}
              />
            </ListItem>
            <ListItem>
              <ListItemAvatar>
                <Skeleton variant="circular" width={40} height={40} />
              </ListItemAvatar>
              <ListItemText
                primary={<Skeleton variant="text" height={22} sx={{ mb: 1 }} />}
                secondary={<Skeleton variant="rounded" height={37} />}
              />
            </ListItem>
          </List>
        </CardContent>
        <CardActions sx={{ justifyContent: "center" }}>
          {!!onPricePlanSelected ? (
            <Button variant="contained" color="secondary">
              Choose Plan
            </Button>
          ) : null}
        </CardActions>
      </Card>
    </>
  );
};
