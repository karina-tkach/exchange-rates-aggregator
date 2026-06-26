import {Route, Routes} from "react-router-dom";
import NotFound from "./pages/generic/NotFoundPage.jsx";
import PublicRoutes from "./routes/PublicRoutes.jsx";
import Footer from "./components/Footer.jsx";
import Header from "./components/Header.jsx";

function App() {
  return (
      <div className="min-h-screen flex flex-col bg-[#0F2854] text-white">
          <Header/>
          <main className="flex-1">
              <Routes>
                  {PublicRoutes}
                  <Route path="*" element={<NotFound/>}/>
              </Routes>
          </main>
          <Footer/>
      </div>
    );
}

export default App
