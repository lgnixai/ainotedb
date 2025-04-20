import React, { createContext, useContext, useState, useEffect } from 'react';
import { fetchWorkspaces, fetchTables } from '../api/mockApi';
import { Workspace, Table } from '../types';

interface WorkspaceContextType {
  workspaces: Workspace[];
  currentWorkspace: Workspace | null;
  tables: Table[];
  currentTable: Table | null;
  isLoading: boolean;
  setCurrentWorkspace: (workspace: Workspace) => void;
  setCurrentTable: (table: Table) => void;
}

const WorkspaceContext = createContext<WorkspaceContextType | undefined>(undefined);

export function useWorkspace() {
  const context = useContext(WorkspaceContext);
  if (context === undefined) {
    throw new Error('useWorkspace must be used within a WorkspaceProvider');
  }
  return context;
}

interface WorkspaceProviderProps {
  children: React.ReactNode;
}

export default function WorkspaceProvider({ children }: WorkspaceProviderProps) {
  const [workspaces, setWorkspaces] = useState<Workspace[]>([]);
  const [currentWorkspace, setCurrentWorkspace] = useState<Workspace | null>(null);
  const [tables, setTables] = useState<Table[]>([]);
  const [currentTable, setCurrentTable] = useState<Table | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  // Fetch workspaces on initial load
  useEffect(() => {
    const loadWorkspaces = async () => {
      try {
        const data = await fetchWorkspaces();
        setWorkspaces(data);
        if (data.length > 0) {
          setCurrentWorkspace(data[0]);
        }
      } catch (error) {
        console.error('Failed to fetch workspaces:', error);
      } finally {
        setIsLoading(false);
      }
    };

    loadWorkspaces();
  }, []);

  // Fetch tables when workspace changes
  useEffect(() => {
    const loadTables = async () => {
      if (!currentWorkspace) return;
      
      setIsLoading(true);
      try {
        const data = await fetchTables(currentWorkspace.id);
        setTables(data);
        if (data.length > 0) {
          setCurrentTable(data[0]);
        } else {
          setCurrentTable(null);
        }
      } catch (error) {
        console.error('Failed to fetch tables:', error);
      } finally {
        setIsLoading(false);
      }
    };

    loadTables();
  }, [currentWorkspace]);

  return (
    <WorkspaceContext.Provider
      value={{
        workspaces,
        currentWorkspace,
        tables,
        currentTable,
        isLoading,
        setCurrentWorkspace,
        setCurrentTable,
      }}
    >
      {children}
    </WorkspaceContext.Provider>
  );
}