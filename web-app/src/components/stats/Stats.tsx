import { useEffect } from "react";
import { useApi } from "../../hooks/useApi";

type IProps = {
	isFetching: boolean;
};

export const Stats = ({ isFetching }: IProps) => {
	const { data, refetch } = useApi("/stats");

	useEffect(() => {
		if (isFetching) {
			refetch();
		}
	}, [isFetching]);

	return (
		<div className="mt-8 p-6 bg-gradient-to-r from-purple-100 to-blue-100 rounded-xl shadow-inner border border-purple-200">
			<div className="flex justify-between text-gray-700 font-medium">
				<div className="flex items-center gap-2">
					<span className="w-3 h-3 bg-blue-500 rounded-full"></span>
					<span>Total: {data?.total_tasks || 0}</span>
				</div>
				<div className="flex items-center gap-2">
					<span className="w-3 h-3 bg-green-500 rounded-full"></span>
					<span>Completed: {data?.completed_tasks || 0}</span>
				</div>
				<div className="flex items-center gap-2">
					<span className="w-3 h-3 bg-yellow-500 rounded-full"></span>
					<span>Remaining: {data?.pending_tasks || 0}</span>
				</div>
			</div>
		</div>
	);
};
