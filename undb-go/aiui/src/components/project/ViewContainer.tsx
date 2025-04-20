import React from 'react';
import { useProject } from '../../context/ProjectContext';
import { GridView } from './GridView';
import { KanbanView } from './KanbanView';
import { CalendarView } from './CalendarView';

export function ViewContainer() {
  const { currentView } = useProject();
  
  if (!currentView) return null;

  const renderView = () => {
    switch (currentView.type) {
      case 'grid':
        return <GridView />;
      case 'kanban':
        return <KanbanView />;
      case 'calendar':
        return <CalendarView />;
      default:
        return <GridView />;
    }
  };

  return (
    <div className="flex-1 overflow-hidden flex flex-col">
      {renderView()}
    </div>
  );
}