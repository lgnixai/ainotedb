import { useState } from 'react';
import { Plus, Filter, SortAsc, Download, Upload, MoreHorizontal } from 'lucide-react';
import { Record, Field } from '../../../types';
import { createRecord } from '../../../api/mockApi';
import toast from 'react-hot-toast';

interface GridViewProps {
  records: Record[];
  fields: Field[];
  onRecordSelect: (record: Record) => void;
  selectedRecordId?: string;
}

const GridView: React.FC<GridViewProps> = ({ records, fields, onRecordSelect, selectedRecordId }) => {
  const [isAddingRecord, setIsAddingRecord] = useState(false);
  const [newRecordValues, setNewRecordValues] = useState<{ [key: string]: any }>({});

  const handleAddRecord = async () => {
    try {
      if (isAddingRecord) {
        // Save the new record
        await createRecord(records[0]?.tableId || '', newRecordValues);
        setIsAddingRecord(false);
        setNewRecordValues({});
        toast.success('记录已添加');
      } else {
        // Start adding a new record
        setIsAddingRecord(true);
      }
    } catch (error) {
      console.error('Failed to add record:', error);
      toast.error('添加记录失败');
    }
  };

  const handleFieldChange = (fieldId: string, value: any) => {
    setNewRecordValues(prev => ({
      ...prev,
      [fieldId]: value
    }));
  };

  return (
    <div className="h-full flex flex-col">
      {/* Toolbar */}
      <div className="bg-white dark:bg-primary border-b border-gray-200 dark:border-gray-700 p-3 flex items-center justify-between">
        <div className="flex items-center space-x-2">
          <button 
            className="btn btn-primary py-1 px-3 text-sm"
            onClick={handleAddRecord}
          >
            <Plus size={16} className="mr-1" />
            {isAddingRecord ? '保存' : '添加记录'}
          </button>
          
          <button className="btn btn-secondary py-1 px-3 text-sm">
            <Filter size={16} className="mr-1" />
            筛选
          </button>
          
          <button className="btn btn-secondary py-1 px-3 text-sm">
            <SortAsc size={16} className="mr-1" />
            排序
          </button>
        </div>
        
        <div className="flex items-center space-x-2">
          <button className="btn btn-secondary py-1 px-3 text-sm">
            <Upload size={16} className="mr-1" />
            导入
          </button>
          
          <button className="btn btn-secondary py-1 px-3 text-sm">
            <Download size={16} className="mr-1" />
            导出
          </button>
          
          <button className="p-1 rounded hover:bg-gray-100 dark:hover:bg-gray-700">
            <MoreHorizontal size={18} />
          </button>
        </div>
      </div>

      {/* Table */}
      <div className="flex-1 overflow-auto">
        <table className="min-w-full border-collapse">
          <thead className="sticky top-0 bg-secondary dark:bg-secondary/80 backdrop-blur-sm z-10">
            <tr>
              {fields.map(field => (
                <th 
                  key={field.id}
                  className="table-header whitespace-nowrap text-left"
                  style={{ minWidth: field.type === 'text' ? '200px' : '150px' }}
                >
                  {field.name}
                </th>
              ))}
            </tr>
          </thead>
          <tbody>
            {isAddingRecord && (
              <tr className="bg-accent/5">
                {fields.map(field => (
                  <td key={field.id} className="table-cell">
                    <input
                      type="text"
                      className="w-full p-1 border border-accent/30 rounded"
                      placeholder={`输入${field.name}`}
                      value={newRecordValues[field.id] || ''}
                      onChange={(e) => handleFieldChange(field.id, e.target.value)}
                    />
                  </td>
                ))}
              </tr>
            )}
            
            {records.map(record => (
              <tr 
                key={record.id}
                className={`hover:bg-gray-50 dark:hover:bg-gray-800/50 cursor-pointer ${
                  selectedRecordId === record.id ? 'bg-accent/10' : ''
                }`}
                onClick={() => onRecordSelect(record)}
              >
                {fields.map(field => (
                  <td key={field.id} className="table-cell">
                    {record.values[field.id] !== undefined ? String(record.values[field.id]) : '—'}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default GridView;