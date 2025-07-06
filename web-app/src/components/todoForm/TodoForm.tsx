import { useState } from "react";
import { Check, X, Plus } from "lucide-react";
import { useApi } from "../../hooks/useApi";

type IProps = {
	refetchTodo: () => void;
};

export const TodoForm = ({ refetchTodo }: IProps) => {
	const [isAdding, setIsAdding] = useState(false);
	const [title, setTitle] = useState("");
	const [description, setDescription] = useState("");

	const { mutate: createTodo } = useApi("/tasks", { skip: true });

	const addTodo = async (event: React.MouseEvent<HTMLButtonElement>) => {
		event.preventDefault();
		if (title.trim()) {
			const newTodo = {
				title: title.trim(),
				description: description.trim(),
			};

			await createTodo({
				method: "POST",
				body: newTodo,
				onSuccess: () => {
					setIsAdding(false);
					refetchTodo();
					setTitle("");
					setDescription("");
				},
				onError: (err) => {
					console.error(err?.message || "Todo creation failed. Please try again.");
				},
			});
		}
	};

	const handleCancel = () => {
		setIsAdding(false);
		setTitle("");
		setDescription("");
	};

	return (
		<div className="mb-6">
			{!isAdding ? (
				<button
					onClick={() => setIsAdding(true)}
					className="flex items-center cursor-pointer gap-2 bg-gradient-to-r from-purple-500 to-blue-500 hover:from-purple-600 hover:to-blue-600 text-white px-6 py-3 rounded-xl transition-all duration-200 transform hover:scale-105 shadow-lg hover:shadow-xl"
				>
					<Plus size={20} />
					Add New Todo
				</button>
			) : (
				<div className="bg-gradient-to-r from-purple-50 to-blue-50 p-6 rounded-xl border-2 border-dashed border-purple-200 shadow-inner">
					<input
						type="text"
						placeholder="Enter todo title..."
						value={title}
						onChange={(e) => setTitle(e.target.value)}
						className="w-full p-3 border border-gray-300 rounded-lg mb-3 focus:outline-none focus:ring-2 focus:ring-blue-500"
						autoFocus
					/>
					<textarea
						placeholder="Enter description (optional)..."
						value={description}
						onChange={(e) => setDescription(e.target.value)}
						className="w-full p-4 border border-purple-200 rounded-xl mb-4 h-24 resize-none focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent transition-all duration-200"
					/>
					<div className="flex gap-2">
						<button
							onClick={addTodo}
							className="bg-gradient-to-r from-green-500 to-emerald-500 hover:from-green-600 hover:to-emerald-600 text-white px-6 py-3 rounded-xl transition-all duration-200 transform hover:scale-105 flex items-center gap-2 shadow-lg"
						>
							<Check size={16} />
							Add Todo
						</button>
						<button
							onClick={handleCancel}
							className="bg-gradient-to-r from-gray-500 to-gray-600 hover:from-gray-600 hover:to-gray-700 text-white px-6 py-3 rounded-xl transition-all duration-200 transform hover:scale-105 flex items-center gap-2 shadow-lg"
						>
							<X size={16} />
							Cancel
						</button>
					</div>
				</div>
			)}
		</div>
	);
};
