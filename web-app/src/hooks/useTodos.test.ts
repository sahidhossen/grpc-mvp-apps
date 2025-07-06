import { renderHook, act } from "@testing-library/react";
import { vi } from "vitest";
import { useTodos } from "./useTodos"; // Adjust import path
import * as apiHook from "./useApi"; // Path to your useApi hook

describe("useTodos hook", () => {
	const mockTasks = [
		{ id: "1", title: "Test todo 1", completed: false },
		{ id: "2", title: "Test todo 2", completed: true },
	];

	let mockRefetch: any;
	let mockMutate: any;

	beforeEach(() => {
		mockRefetch = vi.fn();
		mockMutate = vi.fn();

		// Mock useApi hook calls
		vi.spyOn(apiHook, "useApi").mockImplementation((url) => {
			if (url === "/tasks") {
				return {
					data: mockTasks,
					error: null,
					isLoading: false,
					isSuccess: true,
					refetch: mockRefetch,
					mutate: () => Promise.resolve(),
				};
			}
			if (url === "/tasks/complete") {
				return {
					data: null,
					error: null,
					isLoading: false,
					isSuccess: false,
					refetch: () => Promise.resolve(),
					mutate: mockMutate,
				};
			}
			return {
				data: null,
				error: null,
				isLoading: false,
				isSuccess: false,
				refetch: () => Promise.resolve(),
				mutate: () => Promise.resolve(),
			};
		});
	});

	afterEach(() => {
		vi.restoreAllMocks();
	});

	it("should return tasks and isLoading state", () => {
		const { result } = renderHook(() => useTodos());

		expect(result.current.tasks).toEqual(mockTasks);
		expect(result.current.isLoading).toBe(false);
	});

	it("should call mutate with correct params on toggleTodo", () => {
		const { result } = renderHook(() => useTodos());

		act(() => {
			result.current.toggleTodo("1");
		});

		expect(mockMutate).toHaveBeenCalledWith({
			url: "/tasks/1/toggle-task-complete",
			method: "PATCH",
			onSuccess: expect.any(Function),
			onError: expect.any(Function),
		});
	});

	it("should call refetch after toggleTodo success", () => {
		const { result } = renderHook(() => useTodos());

		mockMutate.mockImplementation(({ onSuccess }: any) => {
			onSuccess?.();
		});

		act(() => {
			result.current.toggleTodo("1");
		});

		expect(mockRefetch).toHaveBeenCalled();
	});

	it("should handle refetchTodos correctly", async () => {
		const { result } = renderHook(() => useTodos());

		await act(async () => {
			await result.current.refetchTodos();
		});

		expect(mockRefetch).toHaveBeenCalled();
	});
});
