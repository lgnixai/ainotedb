import { useEffect } from 'react';
import { useLocation, useNavigate, Outlet } from 'react-router-dom';

export default function RequireAuth() {
  const token = localStorage.getItem('token');
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    if (!token) {
      navigate('/login', { replace: true, state: { from: location } });
    }
  }, [token, navigate, location]);

  if (!token) return null;
  return <Outlet />;
}

