import { useRouter, useSearchParams } from "next/navigation";

export const useURlState = (key: string) => {
    const queries = useSearchParams();
    const router = useRouter();
  
    const removeState = () => {
      const url = new URL(window.location.href);
      url.searchParams.delete(key);
      router.replace(url.toString());
    };
  
    return { state: queries.get(key), removeState };
  };