import React, { useEffect, useState } from 'react';
import { useProject } from '../../context/ProjectContext';
import { getSpaces } from '../../lib/api'; // 引入获取spaces的API
import { getTablesBySpaceId } from '../../lib/api';
import { NewTableDialog } from './NewTableDialog';
import { Select } from '../ui/Select';

import { Button } from '../ui/Button';
import {
  Grid,
  CalendarDays,
  Layout,
  Filter,
  ArrowDownUp,
  Users,
  Share,
  MoreHorizontal,
  Plus
} from 'lucide-react';
import { formatDate } from '../../lib/utils';
import { Badge } from '../ui/Badge';

export function ProjectHeader() {
  const { projects, currentProject, setCurrentProject, currentView, views, setCurrentView } = useProject();
  const [spaces, setSpaces] = useState<any[]>([]);
  const [loadingSpaces, setLoadingSpaces] = useState(true);
  const [error, setError] = useState('');

  // 新增表相关
  const [tables, setTables] = useState<any[]>([]);
  const [loadingTables, setLoadingTables] = useState(false);
  const [selectedTable, setSelectedTable] = useState<string>('');
  const [showNewTableDialog, setShowNewTableDialog] = useState(false);

  // 拉取表列表
  const fetchTables = async (spaceId: string) => {
    setLoadingTables(true);
    try {
      const data = await getTablesBySpaceId(spaceId);
      setTables(Array.isArray(data) ? data : (data.tables || []));
      if (Array.isArray(data) && data.length > 0) setSelectedTable(String(data[0].id));
      else if (data.tables && data.tables.length > 0) setSelectedTable(String(data.tables[0].id));
      else setSelectedTable('');
    } catch (e) {
      setTables([]);
      setSelectedTable('');
    } finally {
      setLoadingTables(false);
    }
  };

  // 当前空间变化时刷新表列表
  useEffect(() => {
    if (currentProject && currentProject.id) {
      fetchTables(currentProject.id);
    }
  }, [currentProject]);

  useEffect(() => {
    async function fetchSpaces() {
      setLoadingSpaces(true);
      try {
        const data = await getSpaces();
        setSpaces(Array.isArray(data) ? data : (data.spaces || []));
        setError('');
      } catch (err: any) {
        setError(err.message || '获取空间列表失败');
      } finally {
        setLoadingSpaces(false);
      }
    }
    fetchSpaces();
  }, []);

  // 切换空间
  const handleSpaceChange = (space: any) => {
    setCurrentProject(space);
  };

  if (!currentProject) return null;

  return (
    <div className="border-b border-gray-200 dark:border-dark-border bg-primary dark:bg-dark-bg px-4 py-4">
      <div className="flex flex-col space-y-3">
        <div className="flex justify-between items-center">
          {/* Space 列表展示与切换 */}
          <div className="flex items-center space-x-3">
            <span className="font-bold">空间：</span>
            {loadingSpaces ? (
              <span className="text-gray-400">加载中...</span>
            ) : error ? (
              <span className="text-red-500">{error}</span>
            ) : (
              <select
                className="border rounded px-2 py-1 text-base bg-white dark:bg-dark-bg text-text dark:text-white"
                value={currentProject?.id || ''}
                onChange={e => {
                  const selected = spaces.find(s => String(s.id) === e.target.value);
                  if (selected) handleSpaceChange(selected);
                }}
              >
                {spaces.map(space => (
                  <option key={space.id} value={space.id}>{space.name}</option>
                ))}
              </select>
            )}

            {/* 表下拉框和新建表按钮 */}
            <div className="flex items-center space-x-2">
              {loadingTables ? (
                <span className="text-gray-400">表加载中...</span>
              ) : (
                <Select
                  className="min-w-[120px]"
                  value={selectedTable}
                  onChange={e => setSelectedTable(e.target.value)}
                  options={tables.map(table => ({ value: String(table.id), label: table.name }))}
                  placeholder="选择表"
                />
              )}
              <Button
                variant="outline"
                size="sm"
                className="border-accent text-accent"
                onClick={() => setShowNewTableDialog(true)}
              >
                新建表
              </Button>
              <NewTableDialog
                isOpen={showNewTableDialog}
                onClose={() => setShowNewTableDialog(false)}
                spaceId={currentProject ? Number(currentProject.id) : 0}
                onTableCreated={() => fetchTables(currentProject.id)}
              />
            </div>
          </div>

          <div className="flex items-center space-x-2">
            <Button
              variant="ghost"
              size="sm"
              leftIcon={<Users size={16} />}
            >
              Team
            </Button>

            <Button
              variant="ghost"
              size="sm"
              leftIcon={<Share size={16} />}
            >
              Share
            </Button>

            <Button
              variant="ghost"
              size="sm"
            >
              <MoreHorizontal size={16} />
            </Button>
          </div>
        </div>

        {currentProject.description && (
          <p className="text-text-secondary dark:text-gray-400 text-sm">{currentProject.description}</p>
        )}

        <div className="flex flex-col sm:flex-row sm:justify-between sm:items-center space-y-3 sm:space-y-0">
          <div className="flex space-x-2 overflow-x-auto pb-1">
            {views.map(view => (
              <button
                key={view.id}
                className={`inline-flex items-center px-3 py-1.5 rounded-md text-sm font-medium ${
                  currentView?.id === view.id
                    ? 'bg-accent text-white'
                    : 'text-text-secondary hover:bg-gray-100 dark:hover:bg-dark-hover dark:text-gray-300'
                }`}
                onClick={() => setCurrentView(view)}
              >
                {view.type === 'grid' && <Grid size={16} className="mr-1.5" />}
                {view.type === 'calendar' && <CalendarDays size={16} className="mr-1.5" />}
                {view.type === 'kanban' && <Layout size={16} className="mr-1.5" />}
                {view.name}
              </button>
            ))}

            <button className="inline-flex items-center px-3 py-1.5 rounded-md text-sm font-medium text-text-secondary hover:bg-gray-100 dark:hover:bg-dark-hover dark:text-gray-300">
              <Plus size={16} className="mr-1.5" />
              New View
            </button>
          </div>

          <div className="flex items-center space-x-2">
            <Button
              variant="ghost"
              size="sm"
              leftIcon={<Filter size={16} />}
            >
              Filter
            </Button>

            <Button
              variant="ghost"
              size="sm"
              leftIcon={<ArrowDownUp size={16} />}
            >
              Sort
            </Button>

            <Badge variant="outline" className="text-xs">
              Created {formatDate(currentProject.createdAt)}
            </Badge>
          </div>
        </div>
      </div>
    </div>
  );
}
