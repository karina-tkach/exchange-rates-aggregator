import ProtectedRoute from "./ProtectedRoute.jsx";
import AdminExchangesPage from "../pages/AdminExchangesPage.jsx";
import {Route} from "react-router-dom";

const AdminRoutes = [
    <Route path="/admin/exchanges" key="exchanges" element={<ProtectedRoute roles={["admin"]}>
        <AdminExchangesPage/>
    </ProtectedRoute>}/>,
];

export default AdminRoutes;