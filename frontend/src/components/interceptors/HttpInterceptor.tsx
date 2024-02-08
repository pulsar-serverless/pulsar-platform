"use client"
import { ReactNode, useEffect, useState } from "react";
import axios from "axios";
import { useAuth0 } from "@auth0/auth0-react";

export const axiosInstance = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_BASE_URL,
  });

export const HttpInterceptor = ({ children }: { children: ReactNode }) => {
  const { isAuthenticated, getAccessTokenSilently } = useAuth0();
  const [token, setToken] = useState<null | string>(null);

  useEffect(() => {
    (async function () {
      try {
        const newToken = await getAccessTokenSilently();
        setToken(newToken);
      } catch (error) {}
    })();
  }, [getAccessTokenSilently, isAuthenticated]);

  useEffect(() => {
    const interceptor = axiosInstance.interceptors.request.use(
      (config) => {
        if (token) config.headers["Authorization"] = `Bearer  ${token}`;
        return config;
      },
      (error) => Promise.reject(error)
    );

    return () => axiosInstance.interceptors.request.eject(interceptor);
  }, [token]);
  return <>{children}</>;
};
