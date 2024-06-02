"use client";
import { CreateProjectModal } from "@/components/project/CreateProjectModal";
import { AuthGuard } from "@/components/providers/AuthGuard";
import { ReactNode } from "react";

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <AuthGuard>
      {<CreateProjectModal />}
      {children}
    </AuthGuard>
  );
}
