import React, { createContext, useContext, useState, ReactNode, useEffect } from 'react';
import { Project, Task, View } from '../types';
import { generateId } from '../lib/utils';
import { getSpaces } from '../lib/api'; // 引入获取spaces的API


interface ProjectContextType {
  projects: Project[];
  currentProject: Project | null;
  currentView: View;
  views: View[];
  setCurrentProject: (project: Project | null) => void;
  addProject: (project: Project) => void;
  updateProject: (id: string, project: Partial<Project>) => void;
  deleteProject: (id: string) => void;
  addTask: (task: Omit<Task, 'id' | 'createdAt' | 'updatedAt'>) => void;
  updateTask: (id: string, task: Partial<Task>) => void;
  deleteTask: (id: string) => void;
  setCurrentView: (view: View) => void;
  addView: (view: Omit<View, 'id'>) => void;
  updateView: (id: string, view: Partial<View>) => void;
  deleteView: (id: string) => void;
}

const defaultColumns = [
  { id: 'name', title: 'Name', accessor: 'title', type: 'text' as const, width: 200 },
  { id: 'status', title: 'Status', accessor: 'status', type: 'select' as const, width: 150, options: ['Not Started', 'In Progress', 'Completed', 'On Hold'] },
  { id: 'priority', title: 'Priority', accessor: 'priority', type: 'select' as const, width: 120, options: ['Low', 'Medium', 'High', 'Urgent'] },
  { id: 'assignee', title: 'Assignee', accessor: 'assignee', type: 'user' as const, width: 150 },
  { id: 'dueDate', title: 'Due Date', accessor: 'dueDate', type: 'date' as const, width: 120 }
];

const defaultViews = [
  { id: 'grid', name: 'Grid', type: 'grid' as const, columns: defaultColumns },
  { id: 'kanban', name: 'Kanban', type: 'kanban' as const, columns: defaultColumns },
  { id: 'calendar', name: 'Calendar', type: 'calendar' as const, columns: defaultColumns }
];

const defaultProject: Project = {
  id: 'default-project',
  name: 'My First Project',
  description: 'A sample project to get you started',
  createdAt: new Date(),
  updatedAt: new Date(),
  tasks: [
    {
      id: 'task-1',
      title: 'Design project dashboard',
      description: 'Create wireframes for the main dashboard',
      status: 'In Progress',
      priority: 'High',
      assignee: 'Jane Smith',
      dueDate: new Date(Date.now() + 86400000 * 2), // 2 days from now
      createdAt: new Date(),
      updatedAt: new Date(),
      tags: ['design', 'dashboard']
    },
    {
      id: 'task-2',
      title: 'Setup API endpoints',
      description: 'Configure backend API endpoints for project data',
      status: 'Not Started',
      priority: 'Medium',
      assignee: 'John Doe',
      dueDate: new Date(Date.now() + 86400000 * 5), // 5 days from now
      createdAt: new Date(),
      updatedAt: new Date(),
      tags: ['backend', 'api']
    },
    {
      id: 'task-3',
      title: 'User testing session',
      description: 'Conduct user testing for the new features',
      status: 'Completed',
      priority: 'Low',
      assignee: 'Alex Johnson',
      dueDate: new Date(Date.now() - 86400000 * 1), // 1 day ago
      createdAt: new Date(),
      updatedAt: new Date(),
      tags: ['testing', 'ux']
    }
  ],
  team: ['Jane Smith', 'John Doe', 'Alex Johnson']
};

const ProjectContext = createContext<ProjectContextType | undefined>(undefined);

export function ProjectProvider({ children }: { children: ReactNode }) {
  const [projects, setProjects] = useState<Project[]>([]);
  const [currentProject, setCurrentProject] = useState<Project | null>(null);
  const [views, setViews] = useState<View[]>(defaultViews);
  const [currentView, setCurrentView] = useState<View>(defaultViews[0]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    async function fetchProjects() {
      setLoading(true);
      try {
        const data = await getSpaces();
        const spaces = Array.isArray(data) ? data : (data.spaces || []);
        // 将 spaces 映射为 Project 类型，补齐 tasks/team 字段
        const mappedProjects: Project[] = spaces.map((space: any) => ({
          id: String(space.id),
          name: space.name,
          description: space.description || '',
          createdAt: new Date(space.created_at),
          updatedAt: new Date(space.updated_at),
          tasks: [], // 可后续扩展为真实数据
          team: [],  // 可后续扩展为真实数据
        }));
        setProjects(mappedProjects);
        setCurrentProject(mappedProjects[0] || null);
        setError('');
      } catch (err: any) {
        setError(err.message || '获取空间列表失败');
      } finally {
        setLoading(false);
      }
    }
    fetchProjects();
  }, []);


  const addProject = (project: Project) => {
    setProjects(prev => [...prev, project]);
    setCurrentProject(project);
  };

  const updateProject = (id: string, updatedFields: Partial<Project>) => {
    setProjects(prevProjects => 
      prevProjects.map(project => 
        project.id === id 
          ? { ...project, ...updatedFields, updatedAt: new Date() } 
          : project
      )
    );

    if (currentProject?.id === id) {
      setCurrentProject(prev => prev ? { ...prev, ...updatedFields, updatedAt: new Date() } : null);
    }
  };

  const deleteProject = (id: string) => {
    setProjects(prevProjects => prevProjects.filter(project => project.id !== id));
    
    if (currentProject?.id === id) {
      setCurrentProject(projects.length > 1 ? projects.find(p => p.id !== id) || null : null);
    }
  };

  const addTask = (task: Omit<Task, 'id' | 'createdAt' | 'updatedAt'>) => {
    if (!currentProject) return;

    const newTask: Task = {
      ...task,
      id: generateId(),
      createdAt: new Date(),
      updatedAt: new Date()
    };

    const updatedTasks = [...currentProject.tasks, newTask];
    
    updateProject(currentProject.id, { tasks: updatedTasks });
  };

  const updateTask = (id: string, updatedFields: Partial<Task>) => {
    if (!currentProject) return;

    const updatedTasks = currentProject.tasks.map(task => 
      task.id === id 
        ? { ...task, ...updatedFields, updatedAt: new Date() } 
        : task
    );

    updateProject(currentProject.id, { tasks: updatedTasks });
  };

  const deleteTask = (id: string) => {
    if (!currentProject) return;

    const updatedTasks = currentProject.tasks.filter(task => task.id !== id);
    updateProject(currentProject.id, { tasks: updatedTasks });
  };

  const addView = (view: Omit<View, 'id'>) => {
    const newView: View = {
      ...view,
      id: generateId()
    };
    
    setViews([...views, newView]);
    setCurrentView(newView);
  };

  const updateView = (id: string, updatedFields: Partial<View>) => {
    setViews(prevViews => 
      prevViews.map(view => 
        view.id === id 
          ? { ...view, ...updatedFields } 
          : view
      )
    );

    if (currentView.id === id) {
      setCurrentView(prev => ({ ...prev, ...updatedFields }));
    }
  };

  const deleteView = (id: string) => {
    if (views.length <= 1) return; // Don't delete the last view
    
    setViews(prevViews => prevViews.filter(view => view.id !== id));
    
    if (currentView.id === id) {
      setCurrentView(views.find(v => v.id !== id) || views[0]);
    }
  };

  return (
    <ProjectContext.Provider 
      value={{ 
        projects, 
        currentProject, 
        currentView,
        views,
        setCurrentProject,
        addProject,
        updateProject,
        deleteProject,
        addTask,
        updateTask,
        deleteTask,
        setCurrentView,
        addView,
        updateView,
        deleteView
      }}
    >
      {children}
    </ProjectContext.Provider>
  );
}

export function useProject() {
  const context = useContext(ProjectContext);
  if (context === undefined) {
    throw new Error('useProject must be used within a ProjectProvider');
  }
  return context;
}