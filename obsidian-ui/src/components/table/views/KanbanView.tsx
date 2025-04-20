import { useState } from 'react';
import { Plus, MoreHorizontal } from 'lucide-react';
import { Field, Record, FieldOption } from '../../../types';
import toast from 'react-hot-toast';
import { updateRecord } from '../../../api/mockApi';

interface KanbanViewProps {
  records: Record[];
  fields: Field[];
  groupByField: string;
  onRecordSelect: (record: Record) => void;
  selectedRecordId?: string;
}

const KanbanView: React.FC<KanbanViewProps> = ({ 
  records, 
  fields, 
  groupByField, 
  onRecordSelect,
  selectedRecordId
}) => {
  const [draggingRecord, setDraggingRecord] = useState<string | null>(null);

  // Find the field we're grouping by
  const groupField = fields.find(f => f.id === groupByField);
  
  if (!groupField || !groupField.options) {
    return (
      <div className="p-6 text-center text-text-secondary">
        无法使用看板视图，因为没有选择一个有选项的字段进行分组。
      </div>
    );
  }

  // Group records by the selected field
  const groupedRecords: Record<string, Record[]> = {};
  
  // Initialize each group to ensure we display empty columns
  groupField.options.forEach(option => {
    groupedRecords[option.id] = [];
  });
  
  // Add records to appropriate groups
  records.forEach(record => {
    const value = record.values[groupByField];
    if (value) {
      if (!groupedRecords[value]) {
        groupedRecords[value] = [];
      }
      groupedRecords[value].push(record);
    }
  });

  const handleDragStart = (recordId: string) => {
    setDraggingRecord(recordId);
  };

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault(); // Allow drop
  };

  const handleDrop = async (e: React.DragEvent, groupId: string) => {
    e.preventDefault();
    if (!draggingRecord) return;

    try {
      // Find the record
      const record = records.find(r => r.id === draggingRecord);
      if (!record) return;
      
      // Update the record's group value
      await updateRecord(draggingRecord, {
        ...record.values,
        [groupByField]: groupId
      });
      
      toast.success('记录已更新');
    } catch (error) {
      console.error('Failed to update record:', error);
      toast.error('更新记录失败');
    } finally {
      setDraggingRecord(null);
    }
  };

  return (
    <div className="h-full overflow-auto p-4">
      <div className="flex space-x-4 h-full">
        {groupField.options?.map(option => {
          const records = groupedRecords[option.id] || [];
          
          return (
            <div 
              key={option.id}
              className="flex-shrink-0 w-72 flex flex-col h-full"
              onDragOver={handleDragOver}
              onDrop={(e) => handleDrop(e, option.id)}
            >
              {/* Column header */}
              <div 
                className="bg-white dark:bg-secondary rounded-t-lg p-3 border border-gray-200 dark:border-gray-700 flex items-center justify-between"
                style={{ 
                  borderLeftWidth: '4px',
                  borderLeftColor: option.color || '#CBD5E1',
                }}
              >
                <h3 className="font-medium">{option.label} ({records.length})</h3>
                <button className="p-1 rounded hover:bg-gray-100 dark:hover:bg-gray-700">
                  <MoreHorizontal size={16} />
                </button>
              </div>
              
              {/* Column content */}
              <div className="flex-1 bg-gray-50 dark:bg-secondary/50 p-2 rounded-b-lg border border-t-0 border-gray-200 dark:border-gray-700 overflow-y-auto">
                {/* Add button */}
                <button className="w-full text-left py-2 px-3 mb-2 rounded bg-white dark:bg-secondary border border-dashed border-gray-300 dark:border-gray-600 hover:border-accent text-text-secondary hover:text-accent">
                  <Plus size={16} className="inline mr-1" />
                  添加记录
                </button>
                
                {/* Cards */}
                {records.map(record => (
                  <div 
                    key={record.id}
                    className={`mb-2 p-3 bg-white dark:bg-secondary rounded border border-gray-200 dark:border-gray-700 cursor-pointer hover:shadow-md transition-shadow ${
                      selectedRecordId === record.id ? 'ring-2 ring-accent' : ''
                    }`}
                    onClick={() => onRecordSelect(record)}
                    draggable
                    onDragStart={() => handleDragStart(record.id)}
                  >
                    {/* Record preview - show first few fields */}
                    <div className="font-medium mb-1 truncate">
                      {record.values['name'] || '无标题'}
                    </div>
                    
                    {fields.slice(0, 3).map(field => {
                      if (field.id === 'name' || field.id === groupByField) return null;
                      return (
                        <div key={field.id} className="text-sm text-text-secondary truncate">
                          <span className="font-medium mr-1">{field.name}:</span>
                          {record.values[field.id] !== undefined ? String(record.values[field.id]) : '—'}
                        </div>
                      );
                    })}
                  </div>
                ))}
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default KanbanView;