"use client";

import ChangeTokenDialog from "@/components/settings/ChangeTokenDialog";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  Container,
  Typography,
} from "@mui/material";
import { useParams } from "next/navigation";
import { useState } from "react";

function Page() {
  const [changeToken, setChangeToken] = useState(false);
  const { projectId } = useParams<{ projectId: string }>();

  return (
    <>
      <Container maxWidth="md" sx={{ py: 3 }}>
        <Typography mb={2.5} variant="h5" component="div">
          Settings
        </Typography>

        <Card>
          <CardContent>
            <Typography
              mb={1.5}
              variant="subtitle1"
              fontWeight={"medium"}
              component="div"
            >
              Generate/Change API Key
            </Typography>
            <Typography gutterBottom variant="body2" color="text.secondary">
              Secure your serverless function with API keys for enhanced
              protection.
            </Typography>
          </CardContent>
          <CardActions sx={{ justifyContent: "end", p: 1.5 }}>
            <Button
              variant="contained"
              color="secondary"
              onClick={() => setChangeToken(true)}
            >
              Change Key
            </Button>
          </CardActions>
        </Card>
      </Container>

      {changeToken && (
        <ChangeTokenDialog
          isOpen={changeToken}
          onClose={() => setChangeToken(false)}
          projectId={projectId}
        />
      )}
    </>
  );
}

export default Page;
