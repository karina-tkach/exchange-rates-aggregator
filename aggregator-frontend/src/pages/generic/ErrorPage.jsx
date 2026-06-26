import { useLocation, useNavigate } from "react-router-dom"
import { AlertCircle } from "lucide-react"

export default function ErrorPage() {
    const location = useLocation()
    const navigate = useNavigate()

    const { message, code } = location.state || {
        message: "Unexpected error occurred. Please try again later.",
        code: 500,
    }

    return (
        <div className="min-h-[83vh] flex items-center justify-center bg-[#020617] px-6 text-[#DAFCE6]">
            <div className="absolute w-[500px] h-[500px] bg-[#14a3c7]/20 blur-3xl rounded-full"></div>
            <div className="absolute w-[300px] h-[300px] bg-[#22267C]/30 blur-3xl rounded-full top-20 left-20"></div>

            <div className="relative w-full max-w-xl bg-[#00002B] rounded-2xl shadow-2xl p-10 text-center border border-[#22267C] overflow-hidden">
                <div className="absolute inset-0 bg-gradient-to-br from-[#22267C]/20 to-[#14a3c7]/10 blur-xl"></div>

                <div className="relative z-10">
                    <div className="flex justify-center mb-6 text-[#14a3c7]">
                        <AlertCircle className="h-14 w-14 drop-shadow-[0_0_10px_#14a3c7]" />
                    </div>

                    <h1 className="text-5xl font-extrabold mb-2">
                        Error <span className="text-[#14a3c7]">{code}</span>
                    </h1>

                    <p className="text-[#DAFCE6]/70 mb-8 text-lg">
                        {message}
                    </p>

                    <div className="flex justify-center gap-4 flex-wrap">

                        <button
                            onClick={() => navigate(-1)}
                            className="bg-[#14a3c7] px-6 py-3 rounded-xl text-black font-semibold
                            hover:bg-[#22267C] hover:text-white transition shadow-lg"
                        >
                            ← Go Back
                        </button>

                        <button
                            onClick={() => navigate("/")}
                            className="border border-[#14a3c7] px-6 py-3 rounded-xl text-[#DAFCE6]
                            hover:bg-[#00002B] hover:border-[#22267C] transition"
                        >
                            Home
                        </button>
                    </div>

                </div>
            </div>
        </div>
    )
}