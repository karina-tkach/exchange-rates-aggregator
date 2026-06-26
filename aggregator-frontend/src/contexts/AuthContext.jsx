import { createContext, useContext, useEffect, useState } from "react";
import { getAccessToken, clearAccessToken, decodeUser, logout as apiLogout } from "../api/auth.js";

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const token = getAccessToken();
        if (token) {
            try {
                setUser(decodeUser(token));
            } catch {
                clearAccessToken();
            }
        }
        setLoading(false);
    }, []);

    const logout = async () => {
        await apiLogout();
        setUser(null);
    };

    return (
        <AuthContext.Provider value={{ user, setUser, loading, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

export const useAuth = () => useContext(AuthContext);
