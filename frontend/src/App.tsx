import './App.css';
import { Routes, Route } from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import LandingPage from "./pages/LandingPage";
import Login from "./pages/Login";
import Layout from "./pages/Layout";
import SelectTemplate from "./pages/SelectTemplate";
import PartyTemplate from "./pages/PartyTemplate";
import PlanningTemplate from "./pages/PlanningTemplate";
import WeddingTemplate from "./pages/WeddingTemplate";


export default function App() {
  return (
    <div>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<LandingPage />} />
          <Route path="login" element={<Login />} />
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="select-template" element={<SelectTemplate />} />
          <Route path="select-template/party" element={<PartyTemplate />} />
          <Route path="select-template/planning" element={<PlanningTemplate />} />
          <Route path="select-template/wedding" element={<WeddingTemplate />} />
        </Route>
      </Routes>
    </div>
  );
}

