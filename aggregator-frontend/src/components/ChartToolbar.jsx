export default function ChartToolbar({
                                         pairs,
                                         exchanges,
                                         pair,
                                         exchange,
                                         period,
                                         timeframe,
                                         timeframeOptions,
                                         onPairChange,
                                         onExchangeChange,
                                         onPeriodChange,
                                         onTimeframeChange,
                                     }) {
    return (
        <div className="flex flex-wrap gap-3 items-center bg-[#00002B] border border-[#22267C] rounded-2xl p-4">
            <select
                value={pair}
                onChange={(e) => onPairChange(e.target.value)}
                className="bg-[#020617] text-[#DAFCE6] border border-[#22267C] px-3 py-2 rounded-xl outline-none"
            >
                {pairs.map((p) => (
                    <option key={p} value={p}>
                        {p}
                    </option>
                ))}
            </select>

            <select
                value={exchange}
                onChange={(e) => onExchangeChange(e.target.value)}
                className="bg-[#020617] text-[#DAFCE6] border border-[#22267C] px-3 py-2 rounded-xl outline-none"
            >
                <option value="">ALL EXCHANGES</option>
                {exchanges.map((ex) => (
                    <option key={ex} value={ex}>
                        {ex}
                    </option>
                ))}
            </select>

            <select
                value={period}
                onChange={(e) => onPeriodChange(e.target.value)}
                className="bg-[#020617] text-[#14a3c7] border border-[#22267C] px-3 py-2 rounded-xl"
            >
                <option value="1h">1H</option>
                <option value="24h">24H</option>
                <option value="7d">7D</option>
            </select>

            <select
                value={timeframe}
                onChange={(e) => onTimeframeChange(e.target.value)}
                className="bg-[#020617] text-[#14a3c7] border border-[#22267C] px-3 py-2 rounded-xl"
            >
                {timeframeOptions.map((t) => (
                    <option key={t.value} value={t.value}>
                        {t.label}
                    </option>
                ))}
            </select>

            <div className="ml-auto text-[#DAFCE6]/50">
                LIVE • 30s refresh
            </div>
        </div>
    )
}