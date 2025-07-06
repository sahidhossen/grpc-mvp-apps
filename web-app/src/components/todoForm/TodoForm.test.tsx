import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { describe, vi, it, expect, beforeEach } from "vitest";
import * as apiHook from "../../hooks/useApi";
import { TodoForm } from "./TodoForm";

vi.mock("../../hooks/useApi");

describe("TodoForm", () => {
	const mockMutate = vi.fn();
	const mockRefetch = vi.fn();

	beforeEach(() => {
		vi.resetAllMocks();
		vi.spyOn(apiHook, "useApi").mockReturnValue({
			data: null,
			error: null,
			isLoading: false,
			isSuccess: false,
			refetch: vi.fn(),
			mutate: mockMutate,
		});
	});

	it("renders Add New Todo button initially", () => {
		render(<TodoForm refetchTodo={mockRefetch} />);
		expect(screen.getByText("Add New Todo")).toBeInTheDocument();
	});

	it("opens form on click", () => {
		render(<TodoForm refetchTodo={mockRefetch} />);
		fireEvent.click(screen.getByText("Add New Todo"));
		expect(screen.getByPlaceholderText("Enter todo title...")).toBeInTheDocument();
		expect(screen.getByText("Cancel")).toBeInTheDocument();
	});

	it("submits form when title is filled", async () => {
		mockMutate.mockImplementation(async ({ onSuccess }) => {
			onSuccess?.();
		});

		render(<TodoForm refetchTodo={mockRefetch} />);
		fireEvent.click(screen.getByText("Add New Todo"));

		fireEvent.change(screen.getByPlaceholderText("Enter todo title..."), {
			target: { value: "New Task" },
		});
		fireEvent.change(screen.getByPlaceholderText("Enter description (optional)..."), {
			target: { value: "Details about the task" },
		});

		fireEvent.click(screen.getByText("Add Todo"));

		await waitFor(() => {
			expect(mockMutate).toHaveBeenCalledWith(
				expect.objectContaining({
					method: "POST",
					body: {
						title: "New Task",
						description: "Details about the task",
					},
				})
			);
			expect(mockRefetch).toHaveBeenCalled();
		});
	});

	it("does not submit if title is empty", async () => {
		render(<TodoForm refetchTodo={mockRefetch} />);
		fireEvent.click(screen.getByText("Add New Todo"));

		fireEvent.click(screen.getByText("Add Todo"));

		await waitFor(() => {
			expect(mockMutate).not.toHaveBeenCalled();
		});
	});
});
