import { useState } from 'react';
import { Link } from 'react-router-dom';
import { ChevronDown, ChevronRight, Plus, Layers, Folder, Users, FileText, Calendar, Settings } from 'lucide-react';
import { useWorkspace } from '../../context/WorkspaceContext';

const Sidebar = () => {
  const { workspaces, currentWorkspace, tables, setCurrentWorkspace, setCurrentTable } = useWorkspace();
  const [isWorkspacesOpen, setIsWorkspacesOpen] = useState(true);

  return (
    <aside className="w-64 h-full bg-secondary dark:bg-secondary border-r border-gray-200 dark:border-gray-700 flex-shrink-0 overflow-y-auto hidden lg:block">
      <div className="p-4">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-lg font-semibold">工作台</h2>
          <button className="p-1 rounded hover:bg-gray-200 dark:hover:bg-gray-700">
            <Plus size={18} />
          </button>
        </div>

        {/* Navigation */}
        <nav className="space-y-1 mb-6">
          <Link to="/" className="flex items-center px-2 py-2 text-sm rounded-md hover:bg-gray-200 dark:hover:bg-gray-700">
            <Layers size={18} className="mr-2" />
            所有表格
          </Link>
          <Link to="/calendar" className="flex items-center px-2 py-2 text-sm rounded-md hover:bg-gray-200 dark:hover:bg-gray-700">
            <Calendar size={18} className="mr-2" />
            日历视图
          </Link>
          <Link to="/documents" className="flex items-center px-2 py-2 text-sm rounded-md hover:bg-gray-200 dark:hover:bg-gray-700">
            <FileText size={18} className="mr-2" />
            文档
          </Link>
          <Link to="/team" className="flex items-center px-2 py-2 text-sm rounded-md hover:bg-gray-200 dark:hover:bg-gray-700">
            <Users size={18} className="mr-2" />
            团队
          </Link>
        </nav>

        {/* Workspaces */}
        <div className="mb-4">
          <div 
            className="flex items-center justify-between px-2 py-2 text-sm font-medium cursor-pointer"
            onClick={() => setIsWorkspacesOpen(!isWorkspacesOpen)}
          >
            <span>工作区</span>
            {isWorkspacesOpen ? <ChevronDown size={16} /> : <ChevronRight size={16} />}
          </div>
          
          {isWorkspacesOpen && (
            <div className="ml-2 mt-1 space-y-1">
              {workspaces.map(workspace => (
                <div 
                  key={workspace.id}
                  className={`flex items-center px-2 py-1.5 text-sm rounded-md cursor-pointer ${
                    currentWorkspace?.id === workspace.id ? 'bg-accent/10 text-accent' : 'hover:bg-gray-200 dark:hover:bg-gray-700'
                  }`}
                  onClick={() => setCurrentWorkspace(workspace)}
                >
                  <Folder size={16} className="mr-2" />
                  <span className="truncate">{workspace.name}</span>
                </div>
              ))}
              <div className="flex items-center px-2 py-1.5 text-sm text-text-secondary hover:bg-gray-200 dark:hover:bg-gray-700 rounded-md cursor-pointer">
                <Plus size={16} className="mr-2" />
                添加工作区
              </div>
            </div>
          )}
        </div>

        {/* Tables */}
        {currentWorkspace && (
          <div className="mb-4">
            <div className="flex items-center justify-between px-2 py-2 text-sm font-medium">
              <span>{currentWorkspace.name} - 表格</span>
              <button className="p-1 rounded hover:bg-gray-200 dark:hover:bg-gray-700">
                <Plus size={14} />
              </button>
            </div>
            
            <div className="ml-2 mt-1 space-y-1">
              {tables.map(table => (
                <Link 
                  key={table.id}
                  to={`/tables/${table.id}`}
                  className="flex items-center px-2 py-1.5 text-sm rounded-md hover:bg-gray-200 dark:hover:bg-gray-700"
                  onClick={() => setCurrentTable(table)}
                >
                  <Layers size={16} className="mr-2" />
                  <span className="truncate">{table.name}</span>
                </Link>
              ))}
            </div>
          </div>
        )}
      </div>

      {/* Footer */}
      <div className="mt-auto border-t border-gray-200 dark:border-gray-700 p-4">
        <Link to="/settings" className="flex items-center px-2 py-2 text-sm rounded-md hover:bg-gray-200 dark:hover:bg-gray-700">
          <Settings size={18} className="mr-2" />
          设置
        </Link>
      </div>
    </aside>
  );
};

export default Sidebar;