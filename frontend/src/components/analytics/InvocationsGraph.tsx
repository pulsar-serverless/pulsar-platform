"use client";

import {
  Card,
  CardContent,
  Box,
  alpha,
  useTheme,
  Typography,
} from "@mui/material";
import {
  ResponsiveContainer,
  AreaChart,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  Area,
} from "recharts";
import InvocationFilter from "./InvocationFilter";
import { AnalyticsApi } from "@/api/analytics";
import { useQuery } from "@tanstack/react-query";
import { FC, useState } from "react";
import dayjs from "dayjs";
import localizedFormat from "dayjs/plugin/localizedFormat";
dayjs.extend(localizedFormat);

const InvocationGraph: React.FC<{ projectId: string }> = ({ projectId }) => {
  const theme = useTheme();

  const [status, setStatus] = useState<InvocationStatus>("Success");
  const [graphType, setGraphType] = useState<InvocationGraphType>("Day");

  const { data: hourlyInvocations } = useQuery({
    queryKey: [AnalyticsApi.getLast24HoursInvocations.name, status],
    queryFn: () => AnalyticsApi.getLast24HoursInvocations(projectId, status),
    enabled: graphType == "Day",
  });

  const { data: weeklyInvocations } = useQuery({
    queryKey: [AnalyticsApi.getLast7DaysInvocations.name, status],
    queryFn: () => AnalyticsApi.getLast7DaysInvocations(projectId, status),
    enabled: graphType == "Week",
  });

  const { data: monthlyInvocations } = useQuery({
    queryKey: [AnalyticsApi.getLast30DaysInvocations.name, status],
    queryFn: () => AnalyticsApi.getLast30DaysInvocations(projectId, status),
    enabled: graphType == "Month",
  });

  const data =
    graphType == "Day"
      ? hourlyInvocations
      : graphType == "Week"
      ? weeklyInvocations
      : monthlyInvocations;

  return (
    <>
      <Card sx={{ mt: 3 }}>
        <CardContent>
          <InvocationFilter
            {...{ status, setStatus, graphType, setGraphType }}
          />
          <Box height={400}>
            <ResponsiveContainer width="100%" height="100%">
              <AreaChart width={500} height={400} data={data}>
                <CartesianGrid
                  strokeDasharray="3 0"
                  vertical={false}
                  stroke={theme.palette.divider}
                />
                <XAxis dataKey={"timestamp"} hide={true} />
                <YAxis tickLine={false} axisLine={false} />
                <Tooltip content={customTooltip(graphType == 'Day' ? "LT" : "ll")} />
                <Area
                  type="linear"
                  dataKey="count"
                  strokeWidth={2.5}
                  stroke={
                    status == "Success"
                      ? theme.palette.primary.main
                      : theme.palette.error.main
                  }
                  fillOpacity={0.15}
                />
              </AreaChart>
            </ResponsiveContainer>
          </Box>
        </CardContent>
      </Card>
    </>
  );
};

const customTooltip = (format: string) => {
  const TooltipCard = ({ active, payload, label }: any) => {
    if (active && payload && payload.length) {
      const data = payload.at(0).payload;
      console.log(data)
      return (
        <Card>
          <CardContent>
            <Typography variant="subtitle2" gutterBottom>
              {data.count || 0} Invocations
            </Typography>
            <Typography variant="body2" color="text.secondary">
              {dayjs(data.timestamp).format(format)}
            </Typography>
          </CardContent>
        </Card>
      );
    }

    return null;
  };

  return TooltipCard;
};
export default InvocationGraph;
