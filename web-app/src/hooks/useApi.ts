import { useCallback, useEffect, useState } from "react";
import { apiClient, type Method } from "../lib/apiClient";

interface ApiOptions<T> {
	url?: string;
	method?: Method;
	body?: unknown;
	headers?: HeadersInit;
	skip?: boolean;
	onSuccess?: (data: T) => void;
	onError?: (error: any) => void;
}

interface ApiState<T> {
	data: T | null;
	error: string | null;
	isLoading: boolean;
	isSuccess: boolean;
	refetch: () => Promise<void>;
	mutate: (options?: Omit<ApiOptions<T>, "skip">) => Promise<void>;
}

export function useApi<T = any>(url: string, options?: ApiOptions<T>): ApiState<T> {
	const [data, setData] = useState<T | null>(null);
	const [error, setError] = useState<string | null>(null);
	const [isLoading, setLoading] = useState(false);
	const [isSuccess, setSuccess] = useState(false);

	const fetchData = useCallback(async () => {
		if (options?.skip) return;
		setLoading(true);
		setError(null);
		try {
			const result = await apiClient<T>(url, {
				method: options?.method,
				headers: options?.headers,
			});

			setData(result);
			setSuccess(true);
			options?.onSuccess?.(result);
		} catch (err: any) {
			setError(err?.message || "Unexpected error");
			options?.onError?.(err); // send entire object
		} finally {
			setLoading(false);
		}
	}, [url]);

	useEffect(() => {
		fetchData();
	}, [fetchData]);

	const mutate = async (mutateOptions?: Omit<ApiOptions<T>, "skip">) => {
		setLoading(true);
		setError(null);
		try {
			const result = await apiClient<T>(mutateOptions?.url || url, {
				method: mutateOptions?.method || "POST",
				body: mutateOptions?.body,
				headers: mutateOptions?.headers,
			});
			setData(result);
			setSuccess(true);
			mutateOptions?.onSuccess?.(result);
		} catch (err: any) {
			setError(err?.message || "Mutation error");
			mutateOptions?.onError?.(err); // send entire object
		} finally {
			setLoading(false);
		}
	};

	return { data, error, isLoading, isSuccess, refetch: fetchData, mutate };
}
