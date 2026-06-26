import { Link } from "react-router-dom"
import { AlertTriangle } from "lucide-react"

export default function NotFound() {
    return (
        <div className="min-h-[83vh] flex items-center justify-center bg-[#020617] px-6 text-[#DAFCE6] relative overflow-hidden">
            <div className="absolute w-[500px] h-[500px] bg-[#14a3c7]/20 blur-3xl rounded-full top-10 left-10"></div>
            <div className="absolute w-[400px] h-[400px] bg-[#22267C]/30 blur-3xl rounded-full bottom-10 right-10"></div>

            <div className="relative w-full max-w-xl bg-[#00002B] rounded-2xl shadow-2xl p-10 text-center border border-[#22267C] overflow-hidden">
                <div className="absolute inset-0 bg-gradient-to-br from-[#22267C]/20 to-[#14a3c7]/10 blur-xl"></div>

                <div className="relative z-10">
                    <div className="flex justify-center mb-6 text-[#14a3c7]">
                        <AlertTriangle className="h-14 w-14 drop-shadow-[0_0_10px_#14a3c7]" />
                    </div>

                    <h1 className="text-7xl font-extrabold text-[#14a3c7] mb-2">
                        404
                    </h1>

                    <h2 className="text-2xl font-semibold text-[#DAFCE6] mb-4">
                        Page not found
                    </h2>

                    <p className="text-[#DAFCE6]/70 mb-8 text-lg">
                        The page you're looking for doesn’t exist or may have been moved.
                    </p>

                    <Link
                        to="/"
                        className="inline-block bg-[#14a3c7] px-6 py-3 rounded-xl font-semibold text-black
                        hover:bg-[#22267C] hover:text-white transition shadow-lg"
                    >
                        ← Back to Market
                    </Link>
                </div>
            </div>
        </div>
    )
}