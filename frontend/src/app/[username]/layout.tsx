"use client"
import { CreateProjectModal } from "@/components/project/CreateProjectModal";
import { useAuth0 } from "@auth0/auth0-react";
import { useRouter } from "next/navigation";
import { ReactNode, useEffect } from "react";

export default function Layout({ children }: { children: ReactNode }) {
  const { isAuthenticated, isLoading } = useAuth0();
  const router = useRouter();

  useEffect(() => {
    if (!isAuthenticated && !isLoading) {
      router.push("/");
    }
  });

  return (
    <>
      {<CreateProjectModal />}
      {children}
    </>
  );
}
