import { useEffect, useState } from 'react';
import { BrowserRouter } from 'react-router-dom';
import { Toaster } from 'react-hot-toast';
import Layout from './components/layout/Layout';
import ThemeProvider from './context/ThemeContext';
import WorkspaceProvider from './context/WorkspaceContext';
import AppRoutes from './routes/AppRoutes';

function App() {
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsLoading(false);
    }, 800);
    return () => clearTimeout(timer);
  }, []);

  if (isLoading) {
    return (
      <div className="fixed inset-0 flex items-center justify-center bg-white dark:bg-gray-900">
        <div className="w-12 h-12 rounded-full border-4 border-accent border-t-transparent animate-spin"></div>
      </div>
    );
  }

  return (
    <BrowserRouter>
      <ThemeProvider>
        <WorkspaceProvider>
          <Layout>
            <AppRoutes />
          </Layout>
          <Toaster
            position="bottom-right"
            toastOptions={{
              className: 'dark:bg-gray-800 dark:text-white',
              duration: 3000,
            }}
          />
        </WorkspaceProvider>
      </ThemeProvider>
    </BrowserRouter>
  );
}

export default App;