import { Link } from "react-router-dom"

export default function Header() {
    return (
        <header className="bg-[#00002B] border-b border-[#22267C] text-[#DAFCE6] shadow-lg relative">
            <div className="absolute inset-0 bg-gradient-to-r from-[#22267C]/10 via-[#14a3c7]/10 to-[#22267C]/10 pointer-events-none" />

            <div className="relative mx-10 px-6 py-4 flex justify-between items-center">
                <h1 className="text-2xl font-bold tracking-wide">
                    <Link
                        to="/"
                        className="text-[#14a3c7] hover:text-[#DAFCE6] transition"
                    >
                        DeMarket
                    </Link>
                </h1>

                <nav className="flex items-center text-lg gap-6 font-medium">
                    <Link
                        to="/charts"
                        className="text-[#14a3c7] hover:text-[#DAFCE6] transition"
                    >
                        Charts
                    </Link>

                    <Link
                        to="/rates"
                        className="text-[#14a3c7] hover:text-[#DAFCE6] transition"
                    >
                        Rates
                    </Link>

                    <Link
                        to="/spreads"
                        className="text-[#14a3c7] hover:text-[#DAFCE6] transition"
                    >
                        Spread
                    </Link>
                </nav>
            </div>
        </header>
    )
}