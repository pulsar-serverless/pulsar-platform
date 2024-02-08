import { setSnackbarMessage } from "@/store/slices/appSlice";
import { useAppDispatch } from "@/store/store";

export const useSnackbar = () => {
  const dispatch = useAppDispatch();

  const setErrorMsg = (msg: string) =>
    dispatch(setSnackbarMessage({ content: msg, type: "error" }));
    
  const setSuccessMsg = (msg: string) =>
    dispatch(setSnackbarMessage({ content: msg, type: "success" }));

  return { setErrorMsg, setSuccessMsg };
};
