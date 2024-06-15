import React from "react";
import {
	Stack
} from "@mui/material";
import { useQuery } from "@tanstack/react-query";
import { PricingApi } from "@/api/pricing";
import PricingCard from "./PricingCard";

const PricingCardsContainer: React.FC = () => {
  const { data,  isLoading } = useQuery({
    queryKey: ["pricingPlans"],
    queryFn: () => PricingApi.getPricingPlans(1, 10),
  });

  return (
    <Stack
      direction="row"
      sx={{
        justifyContent: "center",
        gap: 6,
      }}
    >
      {data?.rows.map((plan) => (
        <PricingCard plan={plan} key={plan.id} />
      ))}
    </Stack>
  );
};

export default PricingCardsContainer;
