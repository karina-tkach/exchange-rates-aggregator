import { Navigate } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext.jsx";
import LoadingPage from "../pages/generic/LoadingPage.jsx";

const ProtectedRoute = ({ children, roles }) => {
    const { user, loading } = useAuth();

    if (loading)
        return (
            <LoadingPage/>
        );

    if (!user) return <Navigate to="/login" />;

    if (!roles.some(role => user?.role.includes(role))) {
        return (
            <Navigate
                to="/"
            />
        );
    }

    return children;
};

export default ProtectedRoute;
