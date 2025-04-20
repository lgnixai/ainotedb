import { Routes, Route, Navigate } from 'react-router-dom';
import TableView from '../components/table/TableView';
import Dashboard from '../pages/Dashboard';
import Settings from '../pages/Settings';
import CalendarView from '../pages/CalendarView';
import Documents from '../pages/Documents';
import Team from '../pages/Team';

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />} />
      <Route path="/tables/:tableId" element={<TableView />} />
      <Route path="/calendar" element={<CalendarView />} />
      <Route path="/documents" element={<Documents />} />
      <Route path="/team" element={<Team />} />
      <Route path="/settings" element={<Settings />} />
      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  );
};

export default AppRoutes;