import {Building} from "lucide-react";
import {Link} from "react-router-dom";

export default function AdminNav() {
    return (
        <>
            <Link to="/admin/exchanges" className="flex cursor-pointer items-center gap-1 text-[#14a3c7] hover:text-[#DAFCE6] transition">
                <Building className="w-5 h-5" />
                Exchanges
            </Link>
        </>
    );
}