"use client";

import { User, useAuth0 } from "@auth0/auth0-react";
import { useRouter } from "next/navigation";
import { ReactNode, useEffect } from "react";

export const AuthGuard: React.FC<{
  children: ReactNode;
  role?: string;
}> = ({ children, role }) => {
  const { isAuthenticated, isLoading, user } = useAuth0<
    User & { roleType: string[] }
  >();
  const router = useRouter();

  console.log(user)

  useEffect(() => {
    if (!isAuthenticated && !isLoading) {
      router.push(`/`);
    } else{
      if (role && !(user?.roleType.includes(role))) router.push("/");
    }
  }, [isAuthenticated, router, isLoading, role, user?.roleType]);

  return <>{children}</>;
};
