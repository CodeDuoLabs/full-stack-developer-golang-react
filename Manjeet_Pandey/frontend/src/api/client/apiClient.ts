export const customInstance = async <T>({
  url,
  method,
  params,
  headers,
  data,
  signal,
}: {
  url: string;
  method: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
  params?: any;
  data?: any;
  headers?: Record<string, any>;
  responseType?: string;
  signal?: AbortSignal;
}): Promise<T> => {
  const requestUrl = `${import.meta.env.VITE_API_BASE_URL}${url}`;
  const options = {
    method,
    headers: {
      "Content-Type": "application/json",
      ...headers,
    },
    ...(data && { body: JSON.stringify(data) }),
    signal
  };
  const queryParams = params
    ? "?" + new URLSearchParams(params).toString()
    : "";

  const fullUrl = `${requestUrl}${queryParams}`;

  const response = await fetch(fullUrl, options);
  return response.json();
};

export default customInstance;
