import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import App from './App';
import Login from './components/auth/Login';
import Register from './components/auth/Register';
import RequireAuth from './components/auth/RequireAuth';
import './index.css';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route element={<RequireAuth />}>
          <Route path="/*" element={<App />} />
        </Route>
      </Routes>
    </BrowserRouter>
  </StrictMode>
);
