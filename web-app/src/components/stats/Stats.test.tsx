import { render, screen } from "@testing-library/react";
import { Stats } from "./Stats";
import { vi } from "vitest";
import * as useApiHook from "../../hooks/useApi";

vi.mock("../../hooks/useApi");

describe("Stats Component", () => {
	const mockRefetch = vi.fn();

	beforeEach(() => {
		vi.clearAllMocks();
	});

	it("renders fallback when data is null", () => {
		(useApiHook.useApi as any).mockReturnValue({
			data: null,
			refetch: mockRefetch,
			isLoading: false,
			isSuccess: false,
			error: null,
			mutate: vi.fn(),
		});

		render(<Stats isFetching={false} />);
		expect(screen.getByText(/Total: 0/)).toBeInTheDocument();
		expect(screen.getByText(/Completed: 0/)).toBeInTheDocument();
		expect(screen.getByText(/Remaining: 0/)).toBeInTheDocument();
	});

	it("renders actual data", () => {
		(useApiHook.useApi as any).mockReturnValue({
			data: {
				total_tasks: 12,
				completed_tasks: 7,
				pending_tasks: 5,
			},
			refetch: mockRefetch,
			isLoading: false,
			isSuccess: true,
			error: null,
			mutate: vi.fn(),
		});

		render(<Stats isFetching={false} />);
		expect(screen.getByText(/Total: 12/)).toBeInTheDocument();
		expect(screen.getByText(/Completed: 7/)).toBeInTheDocument();
		expect(screen.getByText(/Remaining: 5/)).toBeInTheDocument();
	});

	it("calls refetch on isFetching change", () => {
		(useApiHook.useApi as any).mockReturnValue({
			data: {
				total_tasks: 4,
				completed_tasks: 2,
				pending_tasks: 2,
			},
			refetch: mockRefetch,
			isLoading: false,
			isSuccess: true,
			error: null,
			mutate: vi.fn(),
		});

		render(<Stats isFetching={true} />);
		expect(mockRefetch).toHaveBeenCalled();
	});
});
