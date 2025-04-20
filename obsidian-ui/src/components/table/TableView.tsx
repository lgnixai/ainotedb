import { useState, useEffect } from 'react';
import { fetchRecords, fetchViews } from '../../api/mockApi';
import { useWorkspace } from '../../context/WorkspaceContext';
import { Record, View, ViewType } from '../../types';
import GridView from './views/GridView';
import KanbanView from './views/KanbanView';
import ViewSelector from './ViewSelector';
import RecordDetail from '../record/RecordDetail';
import ActivityPanel from '../activity/ActivityPanel';

const TableView = () => {
  const { currentTable } = useWorkspace();
  const [records, setRecords] = useState<Record[]>([]);
  const [views, setViews] = useState<View[]>([]);
  const [currentView, setCurrentView] = useState<View | null>(null);
  const [selectedRecord, setSelectedRecord] = useState<Record | null>(null);
  const [isActivityPanelOpen, setIsActivityPanelOpen] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    if (!currentTable) return;

    const loadData = async () => {
      setIsLoading(true);
      try {
        const [recordsData, viewsData] = await Promise.all([
          fetchRecords(currentTable.id),
          fetchViews(currentTable.id)
        ]);
        
        setRecords(recordsData);
        setViews(viewsData);
        
        if (viewsData.length > 0) {
          setCurrentView(viewsData[0]);
        }
      } catch (error) {
        console.error('Failed to load table data:', error);
      } finally {
        setIsLoading(false);
      }
    };

    loadData();
    
    // Reset selected record when table changes
    setSelectedRecord(null);
    setIsActivityPanelOpen(false);
  }, [currentTable]);

  const handleViewChange = (viewId: string) => {
    const view = views.find(v => v.id === viewId);
    if (view) {
      setCurrentView(view);
    }
  };

  const handleRecordSelect = (record: Record) => {
    setSelectedRecord(record);
  };

  const toggleActivityPanel = () => {
    setIsActivityPanelOpen(!isActivityPanelOpen);
  };

  if (!currentTable) {
    return (
      <div className="h-full flex items-center justify-center text-text-secondary">
        请选择一个表格
      </div>
    );
  }

  if (isLoading) {
    return (
      <div className="h-full flex items-center justify-center">
        <div className="w-10 h-10 border-4 border-accent border-t-transparent rounded-full animate-spin"></div>
      </div>
    );
  }

  const renderView = () => {
    if (!currentView) return null;

    switch (currentView.type) {
      case 'grid':
        return (
          <GridView 
            records={records} 
            fields={currentTable.fields}
            onRecordSelect={handleRecordSelect}
            selectedRecordId={selectedRecord?.id}
          />
        );
      case 'kanban':
        return (
          <KanbanView 
            records={records} 
            fields={currentTable.fields}
            groupByField={currentView.groupBy || 'status'}
            onRecordSelect={handleRecordSelect}
            selectedRecordId={selectedRecord?.id}
          />
        );
      default:
        return (
          <div className="flex items-center justify-center h-64 text-text-secondary">
            该视图类型尚未实现
          </div>
        );
    }
  };

  return (
    <div className="h-full flex flex-col">
      {/* Table header */}
      <div className="bg-white dark:bg-primary border-b border-gray-200 dark:border-gray-700 p-4">
        <div className="flex items-center justify-between">
          <h2 className="text-xl font-semibold">{currentTable.name}</h2>
          <ViewSelector 
            views={views} 
            currentViewId={currentView?.id || ''} 
            onViewChange={handleViewChange} 
          />
        </div>
      </div>

      {/* Table content */}
      <div className="flex-1 flex overflow-hidden">
        {/* Main view area */}
        <div className={`flex-1 ${selectedRecord ? 'hidden md:block' : ''}`}>
          {renderView()}
        </div>

        {/* Record detail panel */}
        {selectedRecord && (
          <div className={`${isActivityPanelOpen ? 'w-1/3' : 'w-2/5'} border-l border-gray-200 dark:border-gray-700 h-full ${isActivityPanelOpen ? 'hidden lg:block' : ''}`}>
            <RecordDetail 
              record={selectedRecord} 
              fields={currentTable.fields}
              onClose={() => setSelectedRecord(null)}
              onToggleActivity={toggleActivityPanel}
              isActivityPanelOpen={isActivityPanelOpen}
            />
          </div>
        )}

        {/* Activity panel */}
        {selectedRecord && isActivityPanelOpen && (
          <div className="w-1/4 border-l border-gray-200 dark:border-gray-700 h-full">
            <ActivityPanel 
              recordId={selectedRecord.id} 
              tableId={currentTable.id}
              onClose={() => setIsActivityPanelOpen(false)}
            />
          </div>
        )}
      </div>
    </div>
  );
};

export default TableView;