import React from 'react';
import { ThemeProvider } from './context/ThemeContext';
import { ProjectProvider } from './context/ProjectContext';
import { Header } from './components/layout/Header';
import { Sidebar } from './components/layout/Sidebar';
import { ProjectHeader } from './components/project/ProjectHeader';
import { ViewContainer } from './components/project/ViewContainer';

function App() {
  return (
    <ThemeProvider>
      <ProjectProvider>
        <div className="min-h-screen flex flex-col bg-gray-50 dark:bg-dark-bg text-text dark:text-gray-200 font-open-sans">
          <Header />
          
          <div className="flex flex-1 overflow-hidden">
            <Sidebar />
            
            <main className="flex-1 flex flex-col overflow-hidden">
              <ProjectHeader />
              <ViewContainer />
            </main>
          </div>
        </div>
      </ProjectProvider>
    </ThemeProvider>
  );
}

export default App;