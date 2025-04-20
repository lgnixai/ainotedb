import { Link } from 'react-router-dom';
import { Plus, FileText, Grid, Columns, Calendar, Star, Clock } from 'lucide-react';
import { useWorkspace } from '../context/WorkspaceContext';

const Dashboard = () => {
  const { workspaces, tables } = useWorkspace();

  return (
    <div className="container mx-auto p-6 max-w-6xl">
      <div className="flex items-center justify-between mb-8">
        <h1 className="text-2xl font-bold">仪表盘</h1>
        <button className="btn btn-primary">
          <Plus size={18} className="mr-1" />
          创建新表格
        </button>
      </div>

      {/* Recent tables */}
      <div className="mb-10">
        <h2 className="text-xl font-semibold mb-4">最近访问</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {tables.slice(0, 3).map(table => (
            <Link
              key={table.id}
              to={`/tables/${table.id}`}
              className="card hover:shadow-md transition-shadow flex items-start p-5"
            >
              <div className="mr-4 p-2 bg-accent/10 rounded-lg">
                <Grid size={24} className="text-accent" />
              </div>
              <div>
                <h3 className="font-medium mb-1">{table.name}</h3>
                <p className="text-sm text-text-secondary">
                  {table.description || '无描述'}
                </p>
                <div className="text-xs text-text-secondary mt-2 flex items-center">
                  <Clock size={12} className="mr-1" />
                  {new Date(table.updatedAt).toLocaleDateString('zh-CN')}
                </div>
              </div>
            </Link>
          ))}
        </div>
      </div>

      {/* Favorites */}
      <div className="mb-10">
        <h2 className="text-xl font-semibold mb-4">收藏</h2>
        <div className="card p-0 overflow-hidden">
          <div className="grid grid-cols-1 divide-y divide-gray-200 dark:divide-gray-700">
            {tables.slice(0, 2).map(table => (
              <Link
                key={table.id}
                to={`/tables/${table.id}`}
                className="flex items-center p-4 hover:bg-gray-50 dark:hover:bg-gray-800"
              >
                <div className="p-2 bg-accent/10 rounded-lg mr-4">
                  <Grid size={20} className="text-accent" />
                </div>
                <div className="flex-1">
                  <h3 className="font-medium">{table.name}</h3>
                  <p className="text-sm text-text-secondary">
                    {table.description || '无描述'}
                  </p>
                </div>
                <Star size={18} className="text-warning" fill="#F59E0B" />
              </Link>
            ))}
            {tables.length === 0 && (
              <div className="p-6 text-center text-text-secondary">
                你还没有收藏任何表格
              </div>
            )}
          </div>
        </div>
      </div>

      {/* Your workspaces */}
      <div>
        <h2 className="text-xl font-semibold mb-4">你的工作区</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {workspaces.map(workspace => (
            <div key={workspace.id} className="card hover:shadow-md transition-shadow">
              <h3 className="font-semibold mb-2">{workspace.name}</h3>
              <p className="text-sm text-text-secondary mb-4">
                {workspace.description || '无描述'}
              </p>
              <div className="flex flex-wrap gap-2">
                <div className="px-2 py-1 text-xs rounded-full bg-accent/10 text-accent">
                  {tables.filter(t => t.workspaceId === workspace.id).length} 个表格
                </div>
                <div className="px-2 py-1 text-xs rounded-full bg-gray-100 dark:bg-gray-800 text-text-secondary">
                  创建于 {new Date(workspace.createdAt).toLocaleDateString('zh-CN')}
                </div>
              </div>
            </div>
          ))}
          <div className="card border-2 border-dashed border-gray-200 dark:border-gray-700 flex items-center justify-center min-h-[160px] hover:border-accent cursor-pointer">
            <div className="text-center">
              <div className="w-12 h-12 rounded-full bg-accent/10 flex items-center justify-center mx-auto mb-2">
                <Plus size={24} className="text-accent" />
              </div>
              <p className="font-medium text-accent">创建新工作区</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;