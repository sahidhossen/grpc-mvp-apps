export const Header = () => {
	return (
		<div className="relative overflow-hidden bg-gradient-to-r from-purple-600 via-blue-600 to-indigo-600 text-white">
			<div className="absolute inset-0 bg-black opacity-10"></div>
			<div className="absolute inset-0 bg-gradient-to-r from-purple-600/20 via-transparent to-blue-600/20"></div>

			<div className="absolute top-0 left-0 w-40 h-40 bg-white/10 rounded-full -translate-x-20 -translate-y-20 animate-pulse"></div>
			<div className="absolute top-10 right-10 w-24 h-24 bg-white/5 rounded-full animate-bounce"></div>
			<div className="absolute bottom-0 right-0 w-32 h-32 bg-white/10 rounded-full translate-x-16 translate-y-16"></div>

			<div className="relative max-w-2xl mx-auto px-6 py-12">
				<div className="text-center">
					<div className="inline-flex items-center justify-center w-16 h-16 bg-white/20 rounded-full mb-4 backdrop-blur-sm">
						<span className="text-2xl">✨</span>
					</div>
					<h1 className="text-4xl md:text-5xl font-bold mb-4 bg-gradient-to-r from-white to-blue-100 bg-clip-text text-transparent">
						My Todo List
					</h1>
					<p className="text-xl text-blue-100 mb-2">Stay organized and get things done</p>
					<div className="flex items-center justify-center gap-2 text-blue-200">
						<span className="w-2 h-2 bg-blue-300 rounded-full animate-pulse"></span>
						<span className="text-sm">Your productivity companion</span>
						<span className="w-2 h-2 bg-blue-300 rounded-full animate-pulse"></span>
					</div>
				</div>
			</div>
		</div>
	);
};
