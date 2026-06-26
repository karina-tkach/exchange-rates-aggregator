import {Link, useNavigate} from "react-router-dom"
import {useAuth} from "../contexts/AuthContext.jsx";
import AuthNav from "./AuthNav.jsx";
import AdminNav from "./AdminNav.jsx";

export default function Header() {
    const {user, loading, logout} = useAuth();
    const navigate = useNavigate();
    const isLoggedIn = user && user.role !== null && user.role !== "";
    const isAdmin = user && user?.role.includes("admin");

    const handleLogout = async () => {
        try {
            await logout();
            navigate("/login");
        } catch {
            navigate('/error', {
                state: {
                    message: "Something went wrong",
                    code: 500
                }
            });
        }
    };

    if (loading) {
        return <></>;
    }

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

                    {isAdmin && <AdminNav/>}
                    <AuthNav isLoggedIn={isLoggedIn} onLogout={handleLogout}/>
                </nav>
            </div>
        </header>
    )
}