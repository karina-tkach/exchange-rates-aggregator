export default function EmptyState({ title, description }) {
    return (
        <div className="flex flex-col items-center justify-center h-full text-center bg-[#020617] border border-[#22267C] rounded-2xl p-10">

            <div className="w-14 h-14 rounded-full bg-[#22267C] flex items-center justify-center mb-4">
                <span className="text-[#14a3c7] text-xl">⚠</span>
            </div>

            <h3 className="text-xl font-semibold text-[#DAFCE6]">
                {title}
            </h3>

            <p className="text-[#DAFCE6]/60 mt-2 max-w-md">
                {description}
            </p>
        </div>
    )
}