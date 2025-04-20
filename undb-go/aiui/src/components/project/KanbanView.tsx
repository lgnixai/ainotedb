import React from 'react';
import { useProject } from '../../context/ProjectContext';
import { formatDate, cn } from '../../lib/utils';
import { Badge } from '../ui/Badge';
import { Status, Task } from '../../types';
import { Plus, MoreHorizontal } from 'lucide-react';

const statusOptions: Status[] = ['Not Started', 'In Progress', 'Completed', 'On Hold'];

const statusColors = {
  'Not Started': 'bg-gray-100 dark:bg-dark-card',
  'In Progress': 'bg-accent-50 dark:bg-accent-900/30',
  'Completed': 'bg-success-50 dark:bg-success-900/30',
  'On Hold': 'bg-warning-50 dark:bg-warning-900/30'
};

const getPriorityBadge = (priority: string) => {
  switch (priority) {
    case 'Low':
      return <Badge variant="success">Low</Badge>;
    case 'Medium':
      return <Badge variant="warning">Medium</Badge>;
    case 'High':
      return <Badge variant="primary">High</Badge>;
    case 'Urgent':
      return <Badge variant="danger">Urgent</Badge>;
    default:
      return <Badge>{priority}</Badge>;
  }
};

export function KanbanView() {
  const { currentProject, updateTask } = useProject();
  
  if (!currentProject) return null;
  
  const tasks = currentProject.tasks;

  const getTasksByStatus = (status: Status) => {
    return tasks.filter(task => task.status === status);
  };

  const handleDragStart = (e: React.DragEvent, task: Task) => {
    e.dataTransfer.setData('taskId', task.id);
  };

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault();
  };

  const handleDrop = (e: React.DragEvent, status: Status) => {
    e.preventDefault();
    const taskId = e.dataTransfer.getData('taskId');
    updateTask(taskId, { status });
  };

  return (
    <div className="flex overflow-x-auto p-4 h-[calc(100vh-13.5rem)]">
      {statusOptions.map(status => (
        <div 
          key={status}
          className={cn(
            "flex-shrink-0 w-72 mr-4 rounded-lg overflow-hidden",
            statusColors[status]
          )}
          onDragOver={handleDragOver}
          onDrop={(e) => handleDrop(e, status)}
        >
          <div className="p-3 border-b border-gray-200 dark:border-dark-border">
            <div className="flex justify-between items-center">
              <h3 className="font-medium text-text dark:text-white">{status}</h3>
              <span className="text-sm text-text-secondary dark:text-gray-400 bg-white dark:bg-dark-hover px-2 py-0.5 rounded-full">
                {getTasksByStatus(status).length}
              </span>
            </div>
          </div>
          
          <div className="p-2 h-full overflow-y-auto">
            {getTasksByStatus(status).map(task => (
              <div
                key={task.id}
                className="mb-2 p-3 bg-white dark:bg-dark-card rounded-md shadow-sm border border-gray-200 dark:border-dark-border cursor-grab"
                draggable
                onDragStart={(e) => handleDragStart(e, task)}
              >
                <div className="flex justify-between items-start mb-2">
                  <h4 className="font-medium text-text dark:text-white">{task.title}</h4>
                  <button className="p-1 rounded-full hover:bg-gray-100 dark:hover:bg-dark-hover">
                    <MoreHorizontal size={14} className="text-text-secondary" />
                  </button>
                </div>
                
                {task.description && (
                  <p className="text-xs text-text-secondary dark:text-gray-400 mb-2 line-clamp-2">
                    {task.description}
                  </p>
                )}
                
                <div className="flex flex-wrap gap-2 mt-2">
                  {getPriorityBadge(task.priority)}
                  
                  {task.dueDate && (
                    <Badge variant="outline" className="text-xs">
                      Due {formatDate(task.dueDate)}
                    </Badge>
                  )}
                </div>
                
                {task.assignee && (
                  <div className="mt-3 pt-2 border-t border-gray-100 dark:border-dark-border flex items-center">
                    <div className="w-5 h-5 rounded-full bg-accent-100 text-accent-700 flex items-center justify-center text-xs font-medium">
                      {task.assignee.charAt(0)}
                    </div>
                    <span className="ml-1 text-xs text-text-secondary dark:text-gray-400">
                      {task.assignee}
                    </span>
                  </div>
                )}
              </div>
            ))}
            
            <button className="w-full p-2 rounded-md border border-dashed border-gray-300 dark:border-dark-border text-sm text-text-secondary dark:text-gray-400 hover:bg-gray-50 dark:hover:bg-dark-hover mt-1">
              <span className="flex items-center justify-center">
                <Plus size={14} className="mr-1" />
                Add task
              </span>
            </button>
          </div>
        </div>
      ))}
    </div>
  );
}