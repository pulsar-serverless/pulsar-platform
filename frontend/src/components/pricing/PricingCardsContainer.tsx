import React from "react";
import { Box, Stack } from "@mui/material";
import { useQuery } from "@tanstack/react-query";
import { PricingApi } from "@/api/pricing";
import PricingCard from "./PricingCard";
import { PricingCardSkeleton } from "./PricingCardSkeleton";

const PricingCardsContainer: React.FC = () => {
  const { data, isLoading } = useQuery({
    queryKey: ["pricingPlans"],
    queryFn: () => PricingApi.getPricingPlans(1, 10),
  });

  return (
    <Box sx={{display: 'grid', placeItems: 'center', height: '100%'}}>
      <Stack
        direction="row"
        sx={{
          justifyContent: "center",
          gap: 6,
          width: '100%'
        }}
      >
        {data?.rows.map((plan) => (
          <PricingCard plan={plan} key={plan.id} />
        ))}
        {!data && (
          <>
            <PricingCardSkeleton />
            <PricingCardSkeleton />
          </>
        )}
      </Stack>
    </Box>
  );
};

export default PricingCardsContainer;
