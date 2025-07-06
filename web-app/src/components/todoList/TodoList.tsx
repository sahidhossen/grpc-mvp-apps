import type { ITask } from "../../types/task";
import { TodoItem } from "./TodoItem";

type IProps = {
	tasks: ITask[];
	isLoading: boolean;
	toggleTodo: (id: string) => void;
};

export const TodoList = ({ tasks, toggleTodo }: IProps) => {
	return (
		<div className="space-y-3">
			{!tasks || tasks.length === 0 ? (
				<div className="text-center py-16">
					<div className="text-8xl mb-6 animate-bounce">📝</div>
					<p className="text-gray-500 text-xl mb-2">No todos yet</p>
					<p className="text-gray-400">Add your first todo to get started!</p>
				</div>
			) : (
				tasks.map((todo) => <TodoItem key={todo.id} todo={todo} toggleTodo={toggleTodo} />)
			)}
		</div>
	);
};
