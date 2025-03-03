import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

const fiveMinutesinMs = 1000 * 60 * 5;
export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: fiveMinutesinMs,
      gcTime: fiveMinutesinMs,
      refetchOnWindowFocus: false,
      retry: false,
    },
  },
});
createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <App />
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  </StrictMode>,
);
