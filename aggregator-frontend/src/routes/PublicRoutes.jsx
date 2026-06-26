import ErrorPage from "../pages/generic/ErrorPage.jsx";
import HomePage from "../pages/HomePage.jsx";
import {Route} from "react-router-dom";
import RatesPage from "../pages/RatesPage.jsx";
import SpreadPage from "../pages/SpreadPage.jsx";
import ChartPage from "../pages/ChartPage.jsx";

const PublicRoutes = [
    <Route path="/" element={<HomePage />} key="home"/>,
    <Route path="/rates" element={<RatesPage/>} key="rates"/>,
    <Route path="/spreads" element={<SpreadPage/>} key="spreads"/>,
    <Route path="/charts" element={<ChartPage/>} key="charts"/>,


    <Route path="/error" element={<ErrorPage />} key="error"/>
];

export default PublicRoutes;