import { Header, Stats, TodoForm, TodoList } from "./components";
import { useTodos } from "./hooks/useTodos";

export default function App() {
	const { tasks, isLoading, toggleTodo, refetchTodos } = useTodos();

	return (
		<div className="min-h-screen bg-gradient-to-br from-purple-50 via-blue-50 to-indigo-100">
			<Header />
			<div className="max-w-2xl mx-auto p-6 bg-white">
				<TodoForm refetchTodo={refetchTodos} />
				<TodoList tasks={tasks ?? []} isLoading={isLoading} toggleTodo={toggleTodo} />

				<Stats isFetching={isLoading} />
			</div>
		</div>
	);
}
