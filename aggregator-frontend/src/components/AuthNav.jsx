import { Link } from "react-router-dom";
import { LogIn, LogOut } from "lucide-react";

export default function AuthNav({ isLoggedIn, onLogout }) {
    return isLoggedIn ? (
        <button onClick={onLogout} className="flex cursor-pointer items-center gap-1 hover:text-blue-200 transition">
            <LogOut className="w-5 h-5" />
            Logout
        </button>
    ) : (
        <>
            <Link to="/login" className="flex cursor-pointer items-center gap-1 hover:text-blue-200 transition">
                <LogIn className="w-5 h-5" />
                Login
            </Link>
        </>
    );
}