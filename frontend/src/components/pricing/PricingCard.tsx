import React, { useState } from "react";
import {
  Card,
  CardContent,
  Typography,
  Button,
  List,
  CardActions,
  ListItem,
  ListItemText,
  Avatar,
  ListItemAvatar,
} from "@mui/material";
import { PricingPlan } from "@/models/PricingPlan";

import NetworkCheckIcon from "@mui/icons-material/NetworkCheck";
import MemoryIcon from "@mui/icons-material/Memory";
import LeakAddIcon from "@mui/icons-material/LeakAdd";
import bytes from "bytes";

const PricingCard: React.FC<{
  plan: PricingPlan;
  onPricePlanSelected?: (plan: PricingPlan) => void;
}> = ({ plan, onPricePlanSelected }) => {
  return (
    <>
      <Card
        sx={{
          flexGrow: 1,
		  width: '50%'
        }}
      >
        <CardContent>
          <Typography
            variant="subtitle1"
            sx={{ textTransform: "capitalize", marginBottom: 1 }}
            fontWeight={"medium"}
            component="div"
            gutterBottom
          >
            {plan.name}
          </Typography>
          <Typography variant="h5" sx={{ marginBottom: 2 }} gutterBottom>
            {plan.price ? `$${plan.price} ETB/month` : "0 ETB/month"}
          </Typography>
          <Typography variant="body2" color="textSecondary">
            {plan.desc}
          </Typography>
          <List sx={{ my: 2 }}>
            <ListItem>
              <ListItemAvatar>
                <Avatar>
                  <LeakAddIcon />
                </Avatar>
              </ListItemAvatar>
              <ListItemText
                primary="Bandwidth"
                secondary={`Includes a quota of ${bytes(
                  plan.allocatedBandwidth
                )} per month for outbound data transfer.`}
              />
            </ListItem>
            <ListItem>
              <ListItemAvatar>
                <Avatar>
                  <MemoryIcon />
                </Avatar>
              </ListItemAvatar>
              <ListItemText
                primary="Memory"
                secondary={`Provides ${bytes(
                  plan.allocatedMemory
                )} per month of dedicated function memory. `}
              />
            </ListItem>
            <ListItem>
              <ListItemAvatar>
                <Avatar>
                  <NetworkCheckIcon />
                </Avatar>
              </ListItemAvatar>
              <ListItemText
                primary="Requests"
                secondary={`Provides a quota of ${plan.allocatedRequests} invocations per month.`}
              />
            </ListItem>
          </List>
        </CardContent>
        <CardActions sx={{ justifyContent: "center" }}>
          {!!onPricePlanSelected ? (
            <Button
              variant="contained"
              color="secondary"
              onClick={() => onPricePlanSelected(plan)}
            >
              Choose Plan
            </Button>
          ) : null}
        </CardActions>
      </Card>
    </>
  );
};

export default PricingCard;
