export default function PairSelect({value, onChange, pairs = [] }) {
    return (
        <div className="flex items-center gap-3">
            <span className="text-[#DAFCE6]/60 text-2xl">Pair:</span>

            <select
                value={value}
                onChange={(e) => onChange(e.target.value)}
                className="bg-[#00002B] border border-[#22267C] text-[#DAFCE6]
                px-4 py-2 rounded-lg outline-none text-lg
                focus:border-[#14a3c7] transition"
            >
                {pairs.map((p) => (
                    <option key={p} value={p}>
                        {p}
                    </option>
                ))}
            </select>
        </div>
    )
}