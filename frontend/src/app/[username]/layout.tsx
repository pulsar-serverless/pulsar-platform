import { CreateProjectModal } from "@/components/project/CreateProjectModal";
import { ReactNode } from "react";

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <>
      {<CreateProjectModal />}
      {children}
    </>
  );
}
