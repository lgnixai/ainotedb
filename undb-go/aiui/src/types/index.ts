export type Status = 'Not Started' | 'In Progress' | 'Completed' | 'On Hold';

export type Priority = 'Low' | 'Medium' | 'High' | 'Urgent';

export interface Task {
  id: string;
  title: string;
  description?: string;
  status: Status;
  priority: Priority;
  assignee?: string;
  dueDate?: Date;
  createdAt: Date;
  updatedAt: Date;
  tags: string[];
}

export interface Project {
  id: string;
  name: string;
  description?: string;
  createdAt: Date;
  updatedAt: Date;
  tasks: Task[];
  team: string[];
}

export interface User {
  id: string;
  name: string;
  email: string;
  avatar?: string;
  role: string;
}

export interface Column {
  id: string;
  title: string;
  accessor: string;
  type: 'text' | 'date' | 'select' | 'multiselect' | 'number' | 'user' | 'checkbox';
  width?: number;
  options?: string[];
}

export interface View {
  id: string;
  name: string;
  type: 'grid' | 'kanban' | 'calendar' | 'list';
  columns: Column[];
  filters?: any[];
  sorts?: any[];
}

export type Theme = 'light' | 'dark';