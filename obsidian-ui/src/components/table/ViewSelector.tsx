import { useState } from 'react';
import { ChevronDown, Grid, Columns, Calendar, Image } from 'lucide-react';
import { View, ViewType } from '../../types';

interface ViewSelectorProps {
  views: View[];
  currentViewId: string;
  onViewChange: (viewId: string) => void;
}

const ViewSelector: React.FC<ViewSelectorProps> = ({ views, currentViewId, onViewChange }) => {
  const [isOpen, setIsOpen] = useState(false);
  
  const currentView = views.find(view => view.id === currentViewId);
  
  const getViewIcon = (type: ViewType) => {
    switch (type) {
      case 'grid':
        return <Grid size={16} />;
      case 'kanban':
        return <Columns size={16} />;
      case 'calendar':
        return <Calendar size={16} />;
      case 'gallery':
        return <Image size={16} />;
      default:
        return <Grid size={16} />;
    }
  };

  return (
    <div className="relative">
      <button
        className="flex items-center space-x-2 px-3 py-1.5 rounded-md bg-white dark:bg-secondary border border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-700"
        onClick={() => setIsOpen(!isOpen)}
      >
        {currentView && (
          <>
            <span className="flex items-center">
              {getViewIcon(currentView.type)}
            </span>
            <span>{currentView.name}</span>
          </>
        )}
        <ChevronDown size={16} />
      </button>

      {isOpen && (
        <div className="absolute right-0 mt-1 bg-white dark:bg-secondary border border-gray-200 dark:border-gray-700 rounded-md shadow-lg z-10 w-48 py-1 animate-fade-in">
          {views.map(view => (
            <button
              key={view.id}
              className={`w-full text-left px-4 py-2 flex items-center space-x-2 hover:bg-gray-100 dark:hover:bg-gray-700 ${
                view.id === currentViewId ? 'bg-gray-100 dark:bg-gray-700' : ''
              }`}
              onClick={() => {
                onViewChange(view.id);
                setIsOpen(false);
              }}
            >
              <span className="flex items-center">
                {getViewIcon(view.type)}
              </span>
              <span>{view.name}</span>
            </button>
          ))}
          <div className="border-t border-gray-200 dark:border-gray-700 my-1"></div>
          <button
            className="w-full text-left px-4 py-2 flex items-center space-x-2 text-accent hover:bg-gray-100 dark:hover:bg-gray-700"
            onClick={() => {
              // This would open a dialog to create a new view
              setIsOpen(false);
            }}
          >
            <span>+ 创建新视图</span>
          </button>
        </div>
      )}
    </div>
  );
};

export default ViewSelector;