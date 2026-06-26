export default function LoadingPage() {
    return (
        <div className="min-h-[83vh] flex items-center justify-center bg-[#020617] px-6 relative overflow-hidden text-[#DAFCE6]">
            <div className="absolute w-[400px] h-[400px] bg-[#14a3c7]/20 blur-3xl rounded-full"></div>
            <div className="absolute w-[300px] h-[300px] bg-[#22267C]/30 blur-3xl rounded-full top-20 left-20"></div>

            <div className="relative flex flex-col items-center justify-center bg-[#00002B] px-10 py-12 rounded-2xl shadow-2xl border border-[#22267C] overflow-hidden">
                <div className="absolute inset-0 bg-gradient-to-br from-[#22267C]/20 to-[#14a3c7]/10 blur-xl"></div>

                <div className="relative z-10 flex flex-col items-center">
                    <div className="relative mb-6">
                        <div className="w-14 h-14 border-4 border-[#14a3c7]/20 border-t-[#14a3c7]
                        rounded-full animate-spin shadow-[0_0_15px_#14a3c7]"></div>
                    </div>

                    <span className="text-[#DAFCE6] text-lg font-medium tracking-wide">
                        Loading market data...
                    </span>

                    <span className="text-[#DAFCE6]/40 text-sm mt-2">
                        syncing exchanges...
                    </span>
                </div>
            </div>
        </div>
    )
}