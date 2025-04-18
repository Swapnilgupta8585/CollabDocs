import "./App.css";
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import { RecoilRoot } from "recoil";
import ProtectedRoute from "./components/ProtectedRoute";
import { Navigate } from "react-router-dom";
import RegisterPage from "./pages/RegisterPage";
import LoginPage from "./pages/LoginPage";
import DashboardPage from "./pages/DashboardPage";

function App() {

  return (
    <>
      <RecoilRoot>
        <BrowserRouter>
          <Routes>
            {/* Public routes */}
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/login" element={<LoginPage/>}/>
      
            {/*Protected routes */}
            <Route element={<ProtectedRoute />}>
              <Route path="/dashboard" element={<DashboardPage/>}/>
            </Route>
            
            
          {/* Redirect to login by default */}
          <Route path="*" element={<Navigate to="/login" replace />} />
          </Routes>
        </BrowserRouter>
      </RecoilRoot>
    </>
  );
}

export default App;
