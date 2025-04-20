import React, { useState } from 'react';
import { ChevronLeft, ChevronRight, Menu, FileText, Settings, Search, Command } from 'lucide-react';
import { useWorkspace } from '../../context/WorkspaceContext';

interface LayoutProps {
  children: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  const { isLoading } = useWorkspace();
  const [isToolbarExpanded, setIsToolbarExpanded] = useState(true);
  const [isLeftSidebarOpen, setIsLeftSidebarOpen] = useState(true);
  const [isRightSidebarOpen, setIsRightSidebarOpen] = useState(true);

  const toggleToolbar = () => setIsToolbarExpanded(!isToolbarExpanded);
  const toggleLeftSidebar = () => setIsLeftSidebarOpen(!isLeftSidebarOpen);
  const toggleRightSidebar = () => setIsRightSidebarOpen(!isRightSidebarOpen);

  return (
    <div className="h-screen flex bg-primary text-text">
      {/* Left Toolbar */}
      <div className={`flex flex-col border-r border-gray-200 dark:border-gray-700 bg-secondary dark:bg-secondary transition-all ${
        isToolbarExpanded ? 'w-14' : 'w-10'
      }`}>
        <div className="flex-1 py-2">
          <button 
            className="w-full p-2 hover:bg-gray-200 dark:hover:bg-gray-700 flex items-center justify-center"
            onClick={toggleToolbar}
          >
            <Menu size={18} />
          </button>
          <button className="w-full p-2 hover:bg-gray-200 dark:hover:bg-gray-700 flex items-center justify-center">
            <Search size={18} />
          </button>
          <button className="w-full p-2 hover:bg-gray-200 dark:hover:bg-gray-700 flex items-center justify-center">
            <Command size={18} />
          </button>
          <button className="w-full p-2 hover:bg-gray-200 dark:hover:bg-gray-700 flex items-center justify-center">
            <FileText size={18} />
          </button>
        </div>
        <div className="py-2 border-t border-gray-200 dark:border-gray-700">
          <button className="w-full p-2 hover:bg-gray-200 dark:hover:bg-gray-700 flex items-center justify-center">
            <Settings size={18} />
          </button>
        </div>
      </div>

      {/* Left Sidebar */}
      <div className="relative flex">
        <div className={`border-r border-gray-200 dark:border-gray-700 bg-secondary dark:bg-secondary transition-all ${
          isLeftSidebarOpen ? 'w-64' : 'w-0 overflow-hidden'
        }`}>
          <div className="p-4">
            <div className="flex items-center justify-between mb-4">
              <h2 className="text-lg font-semibold">文件</h2>
              <button 
                className="p-1 hover:bg-gray-200 dark:hover:bg-gray-700 rounded"
                onClick={toggleLeftSidebar}
              >
                <ChevronLeft size={18} />
              </button>
            </div>
            <div className="space-y-2">
              {/* Sidebar content */}
            </div>
          </div>
        </div>
        {/* Left sidebar toggle button when closed */}
        {!isLeftSidebarOpen && (
          <button
            className="absolute -right-6 top-1/2 -translate-y-1/2 bg-secondary dark:bg-secondary border border-gray-200 dark:border-gray-700 rounded-r p-1 hover:bg-gray-200 dark:hover:bg-gray-700"
            onClick={toggleLeftSidebar}
          >
            <ChevronRight size={18} />
          </button>
        )}
      </div>

      {/* Main Content */}
      <div className="flex-1 flex flex-col overflow-hidden">
        {isLoading ? (
          <div className="h-full flex items-center justify-center">
            <div className="w-10 h-10 border-4 border-accent border-t-transparent rounded-full animate-spin"></div>
          </div>
        ) : (
          <div className="flex-1 overflow-auto">
            {children}
          </div>
        )}
      </div>

      {/* Right Sidebar */}
      <div className="relative flex">
        {/* Right sidebar toggle button when closed */}
        {!isRightSidebarOpen && (
          <button
            className="absolute -left-6 top-1/2 -translate-y-1/2 bg-secondary dark:bg-secondary border border-gray-200 dark:border-gray-700 rounded-l p-1 hover:bg-gray-200 dark:hover:bg-gray-700"
            onClick={toggleRightSidebar}
          >
            <ChevronLeft size={18} />
          </button>
        )}
        <div className={`border-l border-gray-200 dark:border-gray-700 bg-secondary dark:bg-secondary transition-all ${
          isRightSidebarOpen ? 'w-64' : 'w-0 overflow-hidden'
        }`}>
          <div className="p-4">
            <div className="flex items-center justify-between mb-4">
              <button 
                className="p-1 hover:bg-gray-200 dark:hover:bg-gray-700 rounded"
                onClick={toggleRightSidebar}
              >
                <ChevronRight size={18} />
              </button>
              <h2 className="text-lg font-semibold">属性</h2>
            </div>
            <div className="space-y-2">
              {/* Sidebar content */}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Layout;