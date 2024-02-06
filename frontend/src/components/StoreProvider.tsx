"use client"
import { store } from "@/store/store";
import { ReactNode } from "react";
import { Provider } from "react-redux";

export const StoreProvider = ({ children }: { children: ReactNode }) => (
  <Provider store={store}>{children}</Provider>
);
