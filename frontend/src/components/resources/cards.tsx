import React, { useEffect, useState } from "react";
import { ResourcesApi } from "@/api/resources";
import ResourceCards from "./ResourceCards"; // Import ResourceCards component

interface TotalResourceUtil {
  mem_usage_mb: number;
  net_usage_mb: number;
  project_id: string;
  usage_period: string;
}

const ResourceApp: React.FC<{ projectId: string }> = ({ projectId }) => {
  const [totalResourceUtil, setTotalResourceUtil] =
    useState<TotalResourceUtil | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const currentDate = new Date();
        const currentMonth = currentDate.getMonth();
        const monthlyResource = await ResourcesApi.getMonthlyProjectResource(
          projectId,
          currentMonth.toString()
        );
        console.log("Monthly Project Resource:", monthlyResource);

        const resourceUtilList = await ResourcesApi.getProjectResourceUtilList(
          projectId,
          1,
          10,
          currentMonth.toString()
        );
        console.log("Project Resource Util List:", resourceUtilList);

        const totalResourceUtil =
          await ResourcesApi.getTotalProjectResourceUtil(projectId);
        console.log("Total Project Resource Util:", totalResourceUtil);
        setTotalResourceUtil(totalResourceUtil);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, [projectId]);

  return <ResourceCards totalResourceUtil={totalResourceUtil} />;
};

export default ResourceApp;
