export default function Pagination({
                                       currentPage,
                                       lastPage,
                                       onChange,
                                   }) {
    if (lastPage <= 1) return null

    return (
        <div className="flex justify-center gap-3 mt-8">
            <button
                disabled={currentPage === 1}
                onClick={() => onChange(currentPage - 1)}
                className="px-4 py-2 rounded-lg bg-[#00002B]
                border border-[#22267C]
                disabled:opacity-40 hover:bg-[#10104f]
                hover:cursor-pointer"
            >
                Previous
            </button>

            <div className="px-4 py-2">
                {currentPage} / {lastPage}
            </div>

            <button
                disabled={currentPage === lastPage}
                onClick={() => onChange(currentPage + 1)}
                className="px-4 py-2 rounded-lg bg-[#00002B]
                border border-[#22267C]
                disabled:opacity-40 hover:bg-[#10104f]
                hover:cursor-pointer"
            >
                Next
            </button>
        </div>
    )
}