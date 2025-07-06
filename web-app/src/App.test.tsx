import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { vi } from "vitest";
import App from "./App";
import * as apiHook from "./hooks/useApi";

vi.mock("./components/TodoForm/TodoForm", () => ({
	TodoForm: ({ refetchTodo }: any) => <button onClick={refetchTodo}>Mock TodoForm</button>,
}));

describe("App", () => {
	const mockTasks = [
		{ id: "1", title: "Task 1", completed: false },
		{ id: "2", title: "Task 2", completed: true },
	];

	const mockRefetch = vi.fn(() => Promise.resolve());
	const mockMutate = vi.fn(() => Promise.resolve());

	beforeEach(() => {
		vi.clearAllMocks();
		vi.spyOn(apiHook, "useApi").mockImplementation((url: string) => {
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

	it("renders header, todo form, todo list, and stats", () => {
		render(<App />);

		expect(screen.getByText("My Todo List")).toBeInTheDocument();
		expect(screen.getByText("Mock TodoForm")).toBeInTheDocument();
		expect(screen.getAllByTestId("todo-item").length).toBe(mockTasks.length);
		expect(screen.getByText("Total: 0")).toBeInTheDocument();
	});

	it("calls refetchTodo when TodoForm button is clicked", async () => {
		render(<App />);
		const refetchButton = screen.getByText("Mock TodoForm");
		fireEvent.click(refetchButton);

		await waitFor(() => {
			expect(mockRefetch).toHaveBeenCalled();
		});
	});
});
