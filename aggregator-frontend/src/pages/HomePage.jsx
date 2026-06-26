import { Link } from "react-router-dom"

export default function HomePage() {
    return (
        <div className="min-h-screen bg-[#020617] text-[#DAFCE6]">
            <div className="px-10 py-20 grid md:grid-cols-2 gap-10 items-center">
                <div>
                    <h1 className="text-5xl font-bold leading-tight">
                        Relevant <span className="text-[#14a3c7]">Crypto Market</span> Data
                    </h1>

                    <p className="mt-6 text-gray-400 text-lg">
                        Aggregate prices from Binance, Coinbase & Kraken in real time.
                        Visualize OHLC candles, spreads and prices instantly.
                    </p>

                    <div className="flex gap-4 mt-8">
                        <Link
                            to="/charts"
                            className="px-6 py-3 rounded-lg bg-[#14a3c7] text-black font-semibold hover:bg-[#22267C] hover:text-white transition"
                        >
                            View Charts
                        </Link>

                        <Link
                            to="/rates"
                            className="px-6 py-3 rounded-lg border border-[#14a3c7] hover:bg-[#00002B] transition"
                        >
                            Live Rates
                        </Link>
                    </div>
                </div>

                <div className="bg-[#00002B] rounded-xl p-6 border border-[#22267C] shadow-lg">
                    <div className="text-sm text-gray-400 mb-4">
                        Market Overview
                    </div>

                    <img
                        src="/stats.png"
                        alt="Statistic"
                        className="rounded-2xl shadow-2xl"
                    />
                    <div className="absolute inset-0 rounded-2xl bg-[#4988C4]/20 blur-2xl -z-10"></div>
                </div>
            </div>

            <div className="px-10 py-16 grid md:grid-cols-3 gap-6">
                <div className="bg-[#00002B] p-6 rounded-xl border border-[#22267C]">
                    <h3 className="text-[#14a3c7] text-lg font-bold">⚡ Real-time Data</h3>
                    <p className="text-gray-400 mt-2">
                        Stream prices from multiple exchanges simultaneously.
                    </p>
                </div>

                <div className="bg-[#00002B] p-6 rounded-xl border border-[#22267C]">
                    <h3 className="text-[#14a3c7] text-lg font-bold">📊 OHLC Charts</h3>
                    <p className="text-gray-400 mt-2">
                        Candlestick visualization with customizable timeframes.
                    </p>
                </div>

                <div className="bg-[#00002B] p-6 rounded-xl border border-[#22267C]">
                    <h3 className="text-[#14a3c7] text-lg font-bold">🔄 Multi-Exchange</h3>
                    <p className="text-gray-400 mt-2">
                        Binance, Coinbase, Kraken aggregated in one API.
                    </p>
                </div>
            </div>
        </div>
    )
}