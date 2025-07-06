import { render, screen, fireEvent } from "@testing-library/react";
import { describe, it, expect, vi } from "vitest";
import { TodoItem } from "./TodoItem";
import type { ITask } from "../../types/task";

describe("TodoItem", () => {
	const mockToggle = vi.fn();
	const baseTodo: ITask = {
		id: "123",
		title: "Write unit tests",
		description: "Cover components with Vitest",
		completed: false,
	};

	it("renders title and description", () => {
		render(<TodoItem todo={baseTodo} toggleTodo={mockToggle} />);
		expect(screen.getByText("Write unit tests")).toBeInTheDocument();
		expect(screen.getByText("Cover components with Vitest")).toBeInTheDocument();
	});

	it("calls toggleTodo on click", () => {
		render(<TodoItem todo={baseTodo} toggleTodo={mockToggle} />);
		const button = screen.getByRole("button");
		fireEvent.click(button);
		expect(mockToggle).toHaveBeenCalledWith("123");
	});

	it("renders completed state with line-through and green styles", () => {
		const completedTodo: ITask = { ...baseTodo, completed: true };
		render(<TodoItem todo={completedTodo} toggleTodo={mockToggle} />);
		const title = screen.getByText("Write unit tests");
		const desc = screen.getByText("Cover components with Vitest");

		expect(title).toHaveClass("line-through");
		expect(desc).toHaveClass("line-through");
		expect(screen.getByRole("button")).toHaveClass("from-green-500");
	});
});
