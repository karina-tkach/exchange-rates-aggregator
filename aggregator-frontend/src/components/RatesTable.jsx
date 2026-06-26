export default function RatesTable({ data }) {
    return (
        <div className="w-full mt-6 border border-[#22267C] rounded-xl overflow-hidden">
            <div className="grid grid-cols-2 bg-[#00002B] text-[#DAFCE6]/60 text-lg px-4 py-3 border-b border-[#22267C]">
                <div>Exchange</div>
                <div className="text-right">Price</div>
            </div>

            <div className="bg-[#020617] text-2xl">
                {data.map((r, i) => (
                    <div
                        key={i}
                        className="grid grid-cols-2 px-4 py-3 border-b border-[#22267C]/40 hover:bg-[#00002B] transition"
                    >
                        <div className="text-[#DAFCE6] flex items-center gap-2">
                            <span className="w-2 h-2 rounded-full bg-[#14a3c7]" />
                            {r.exchange}
                        </div>

                        <div className="text-right text-[#14a3c7] font-mono">
                            {r.price}
                        </div>
                    </div>
                ))}
            </div>
        </div>
    )
}