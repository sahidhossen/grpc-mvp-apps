import { Check } from "lucide-react";
import type { ITask } from "../../types/task";

type IProps = {
	todo: ITask;
	toggleTodo: (id: string) => void;
};
export const TodoItem = ({ todo, toggleTodo }: IProps) => {
	return (
		<div
			data-testid="todo-item"
			key={todo.id}
			className={`bg-white border border-gray-200 rounded-xl p-6 shadow-lg hover:shadow-xl transition-all duration-200 transform hover:-translate-y-1 ${
				todo.completed
					? "bg-gradient-to-r from-green-50 to-emerald-50 border-green-200"
					: "hover:bg-gradient-to-r hover:from-purple-50 hover:to-blue-50"
			}`}
		>
			<div className="flex items-start gap-3">
				<button
					onClick={() => toggleTodo(todo.id)}
					className={`flex-shrink-0 w-7 h-7 rounded-full border-2 flex items-center justify-center transition-all duration-200 transform hover:scale-110 ${
						todo.completed
							? "bg-gradient-to-r from-green-500 to-emerald-500 border-green-500 text-white shadow-lg"
							: "border-gray-300 hover:border-purple-500 hover:bg-purple-50"
					}`}
				>
					{todo.completed && <Check size={14} />}
				</button>

				<div className="flex-grow">
					<h3 className={`font-semibold text-gray-800 text-lg ${todo.completed ? "line-through text-gray-500" : ""}`}>
						{todo.title}
					</h3>
					{todo.description && (
						<p className={`text-gray-600 mt-2 ${todo.completed ? "line-through text-gray-400" : ""}`}>
							{todo.description}
						</p>
					)}
				</div>
			</div>
		</div>
	);
};
