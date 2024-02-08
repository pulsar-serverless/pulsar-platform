import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface AppMessage {
  content: string;
  type: "error" | "success";
}
export interface CounterState {
  snackbarMessage?: AppMessage;
}

const initialState: CounterState = {};

export const appSlice = createSlice({
  name: "app",
  initialState,
  reducers: {
    setSnackbarMessage: (
      state,
      action: PayloadAction<AppMessage | undefined>
    ) => {
      state.snackbarMessage = action.payload;
    },
  },
});

export const { setSnackbarMessage } = appSlice.actions;
