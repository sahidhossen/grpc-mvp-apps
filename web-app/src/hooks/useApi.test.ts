import { type Mock,  } from "vitest";
import { renderHook, act } from "@testing-library/react";
import { useApi } from "./useApi";
import { apiClient } from "../lib/apiClient";

vi.mock("../lib/apiClient");

const mockClient = apiClient as Mock

describe("useApi", () => {
	afterEach(() => {
		vi.clearAllMocks();
	});

	it("should fetch data successfully", async () => {
		const mockResponse = { message: "success" };
		mockClient.mockResolvedValue(mockResponse);

		const { result } = renderHook(() => useApi("/api/test"));

		// wait for useEffect hook
		await act(async () => {
			await Promise.resolve();
		});

		mockClient.mockClear();
		mockClient.mockResolvedValueOnce(mockResponse);

		// Trigger refetchData
		await act(async () => {
			await result.current.refetch();
		});

		expect(result.current.data).toEqual(mockResponse);
		expect(result.current.error).toBeNull();
		expect(result.current.isSuccess).toBe(true);
		expect(mockClient).toHaveBeenCalledWith("/api/test", {
			method: undefined,
			headers: undefined,
		});
	});

	it("should handle fetch error", async () => {
		mockClient.mockRejectedValue(new Error("Failed to load"));

		const { result } = renderHook(() => useApi("/api/test"));

		await act(async () => {
			await Promise.resolve();
		});

		mockClient.mockClear();
		mockClient.mockRejectedValueOnce(new Error("Failed to load"));

		await act(async () => {
			await result.current.refetch();
		});

		expect(result.current.data).toBeNull();
		expect(result.current.error).toBe("Failed to load");
		expect(result.current.isSuccess).toBe(false);
	});

	it("should mutate data", async () => {
		const mockData = { message: "created" };
		mockClient.mockResolvedValueOnce(mockData);

		const { result } = renderHook(() => useApi("/api/create", { skip: true }));

		await act(async () => {
			await result.current.mutate({
				method: "POST",
				body: { foo: "bar" },
			});
		});

		expect(result.current.data).toEqual(mockData);
		expect(result.current.error).toBeNull();
		expect(result.current.isSuccess).toBe(true);
		expect(mockClient).toHaveBeenCalledWith("/api/create", {
			method: "POST",
			body: { foo: "bar" },
			headers: undefined,
		});
	});
});
