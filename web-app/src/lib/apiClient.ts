export type Method = "GET" | "POST" | "PUT" | "PATCH" | "DELETE";

interface RequestOptions {
	method?: Method;
	body?: unknown;
	headers?: HeadersInit;
}

const BASE_URL = import.meta.env.VITE_API_URL || "";

/**
 * Central API client to handle all requests
 */
export async function apiClient<T>(url: string, options: RequestOptions = {}): Promise<T> {
	const res = await fetch(`${BASE_URL}${url}`, {
		method: options.method || "GET",
		headers: {
			"Content-Type": "application/json",
			...(options.headers || {}),
		},
		body: options.body ? JSON.stringify(options.body) : undefined,
	});

	const json = await res.json();

	if (!res.ok || json.success === false) {
		throw {
			status: res.status,
			message: json.error || "API Error",
			success: false,
			data: json || null,
		};
	}

	return json;
}
