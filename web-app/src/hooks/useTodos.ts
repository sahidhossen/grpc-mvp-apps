import { useCallback } from "react";
import toast from "react-hot-toast";
import { useApi } from "./useApi";

export function useTodos() {
	const { data: tasks, refetch, isLoading } = useApi("/tasks");
	const { mutate: completeTask } = useApi("/tasks/complete", { skip: true });

	const toggleTodo = useCallback(
		(id: string) => {
			completeTask({
				url: `/tasks/${id}/toggle-task-complete`,
				method: "PATCH",
				onSuccess: () => refetch(),
				onError: (err) => {
					const message = err?.message || "Toggle task failed. Please try again.";
					toast.error(message);
					console.error(message);
				},
			});
		},
		[completeTask, refetch]
	);

	const refetchTodos = useCallback(async () => {
		await refetch();
	}, [refetch]);

	return {
		tasks,
		isLoading,
		toggleTodo,
		refetchTodos,
	};
}
