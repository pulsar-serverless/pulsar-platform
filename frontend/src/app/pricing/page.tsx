"use client";

import React from "react";
import PricingCardsContainer from "@/components/pricing/PricingCardsContainer";
import { Container } from "@mui/material";

const Page: React.FC = () => {
  return (
    <Container maxWidth="md" sx={{ py: 3 }}>
      <PricingCardsContainer />
    </Container>
  );
};

export default Page;
