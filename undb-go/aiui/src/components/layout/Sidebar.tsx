import React, { useState } from 'react';
import { useProject } from '../../context/ProjectContext';
import { cn, getRandomColor } from '../../lib/utils';
import { Button } from '../ui/Button';
import { NewProjectDialog } from '../project/NewProjectDialog';
import { 
  Plus, 
  FolderPlus, 
  Layout, 
  Calendar, 
  Kanban, 
  Table, 
  List, 
  Users, 
  Settings, 
  HelpCircle, 
  ChevronDown 
} from 'lucide-react';

export function Sidebar() {
  const { 
    projects, 
    currentProject, 
    setCurrentProject, 
    views,
    currentView,
    setCurrentView
  } = useProject();
  
  const [projectsOpen, setProjectsOpen] = React.useState(true);
  const [viewsOpen, setViewsOpen] = React.useState(true);
  const [isNewProjectDialogOpen, setIsNewProjectDialogOpen] = useState(false);

  const handleProjectClick = (project: any) => {
    setCurrentProject(project);
  };

  const handleViewClick = (view: any) => {
    setCurrentView(view);
  };

  const getViewIcon = (type: string) => {
    switch (type) {
      case 'grid':
        return <Table size={18} />;
      case 'kanban':
        return <Kanban size={18} />;
      case 'calendar':
        return <Calendar size={18} />;
      case 'list':
        return <List size={18} />;
      default:
        return <Layout size={18} />;
    }
  };

  return (
    <>
      <aside className="hidden md:flex flex-col w-64 border-r border-gray-200 bg-primary dark:bg-dark-bg dark:border-dark-border">
        <div className="flex flex-col h-full overflow-y-auto">
          <div className="p-4">
            <Button 
              variant="primary" 
              className="w-full justify-center"
              leftIcon={<Plus size={16} />}
              onClick={() => setIsNewProjectDialogOpen(true)}
            >
              New Project
            </Button>
          </div>
          
          <div className="px-3 py-2">
            <button
              className="flex items-center justify-between w-full px-3 py-2 text-sm font-medium text-text dark:text-gray-200 rounded-md hover:bg-gray-100 dark:hover:bg-dark-hover"
              onClick={() => setProjectsOpen(!projectsOpen)}
            >
              <span className="flex items-center">
                <FolderPlus size={18} className="mr-2 text-text-secondary" />
                Projects
              </span>
              <ChevronDown
                size={16}
                className={cn(
                  "transition-transform duration-200 text-text-secondary",
                  projectsOpen ? "transform rotate-180" : ""
                )}
              />
            </button>
            
            {projectsOpen && (
              <div className="pl-9 mt-1 space-y-1">
                {projects.map((project) => (
                  <button
                    key={project.id}
                    className={cn(
                      "flex items-center w-full px-3 py-2 text-sm rounded-md hover:bg-gray-100 dark:hover:bg-dark-hover",
                      currentProject?.id === project.id
                        ? "bg-gray-100 text-accent font-medium dark:bg-dark-hover"
                        : "text-text-secondary dark:text-gray-400"
                    )}
                    onClick={() => handleProjectClick(project)}
                  >
                    <div className={cn("w-2 h-2 rounded-full mr-2", getRandomColor())}></div>
                    <span className="truncate">{project.name}</span>
                  </button>
                ))}
                
                <button
                  className="flex items-center w-full px-3 py-2 text-sm text-text-secondary rounded-md hover:bg-gray-100 dark:hover:bg-dark-hover dark:text-gray-400"
                  onClick={() => setIsNewProjectDialogOpen(true)}
                >
                  <Plus size={16} className="mr-2" />
                  Add Project
                </button>
              </div>
            )}
          </div>
          
          {currentProject && (
            <div className="px-3 py-2">
              <button
                className="flex items-center justify-between w-full px-3 py-2 text-sm font-medium text-text dark:text-gray-200 rounded-md hover:bg-gray-100 dark:hover:bg-dark-hover"
                onClick={() => setViewsOpen(!viewsOpen)}
              >
                <span className="flex items-center">
                  <Layout size={18} className="mr-2 text-text-secondary" />
                  Views
                </span>
                <ChevronDown
                  size={16}
                  className={cn(
                    "transition-transform duration-200 text-text-secondary",
                    viewsOpen ? "transform rotate-180" : ""
                  )}
                />
              </button>
              
              {viewsOpen && (
                <div className="pl-9 mt-1 space-y-1">
                  {views.map((view) => (
                    <button
                      key={view.id}
                      className={cn(
                        "flex items-center w-full px-3 py-2 text-sm rounded-md hover:bg-gray-100 dark:hover:bg-dark-hover",
                        currentView?.id === view.id
                          ? "bg-gray-100 text-accent font-medium dark:bg-dark-hover"
                          : "text-text-secondary dark:text-gray-400"
                      )}
                      onClick={() => handleViewClick(view)}
                    >
                      {getViewIcon(view.type)}
                      <span className="ml-2 truncate">{view.name}</span>
                    </button>
                  ))}
                  
                  <button
                    className="flex items-center w-full px-3 py-2 text-sm text-text-secondary rounded-md hover:bg-gray-100 dark:hover:bg-dark-hover dark:text-gray-400"
                  >
                    <Plus size={16} className="mr-2" />
                    Add View
                  </button>
                </div>
              )}
            </div>
          )}

          <div className="px-3 py-2">
            <button className="flex items-center w-full px-3 py-2 text-sm font-medium text-text dark:text-gray-200 rounded-md hover:bg-gray-100 dark:hover:bg-dark-hover">
              <Users size={18} className="mr-2 text-text-secondary" />
              Team
            </button>
          </div>
          
          <div className="mt-auto px-3 py-4 border-t border-gray-200 dark:border-dark-border">
            <button className="flex items-center w-full px-3 py-2 text-sm font-medium text-text dark:text-gray-200 rounded-md hover:bg-gray-100 dark:hover:bg-dark-hover">
              <Settings size={18} className="mr-2 text-text-secondary" />
              Settings
            </button>
            <button className="flex items-center w-full px-3 py-2 text-sm font-medium text-text dark:text-gray-200 rounded-md hover:bg-gray-100 dark:hover:bg-dark-hover">
              <HelpCircle size={18} className="mr-2 text-text-secondary" />
              Help & Support
            </button>
          </div>
        </div>
      </aside>

      <NewProjectDialog 
        isOpen={isNewProjectDialogOpen}
        onClose={() => setIsNewProjectDialogOpen(false)}
      />
    </>
  );
}