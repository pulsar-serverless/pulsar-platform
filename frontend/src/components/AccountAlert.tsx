"use client"

import { userApi } from "@/api/user";
import { Box, Alert } from "@mui/material";
import { useQuery } from "@tanstack/react-query";

export const AccountAlert: React.FC = () => {
  const { data: status } = useQuery({
    queryKey: [userApi.getAccountStatus.name],
    queryFn: userApi.getAccountStatus,
  });

  if (status && status != "Active") {
    return (
      <>
        <Box sx={{ position: "absolute", bottom: 0, left: 0, right: 0 }}>
          <Alert severity="error">
            Your account is suspended. Please contact to your administrator.
          </Alert>
        </Box>
      </>
    );
  }

  return <></>;
};
