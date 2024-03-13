import { BrowserRouter, Routes, Route } from "react-router-dom";
import { RegisterPage, CheckUsersPage } from "./pages";
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<RegisterPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/check-user" element={<CheckUsersPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
