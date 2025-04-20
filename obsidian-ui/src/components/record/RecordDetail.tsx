import { useState } from 'react';
import { X, MessageSquare, Activity, Clock, Calendar, User, Edit3, Trash2, Save } from 'lucide-react';
import { Record, Field } from '../../types';
import toast from 'react-hot-toast';
import { updateRecord, deleteRecord } from '../../api/mockApi';

interface RecordDetailProps {
  record: Record;
  fields: Field[];
  onClose: () => void;
  onToggleActivity: () => void;
  isActivityPanelOpen: boolean;
}

const RecordDetail: React.FC<RecordDetailProps> = ({ 
  record, 
  fields, 
  onClose, 
  onToggleActivity,
  isActivityPanelOpen
}) => {
  const [isEditing, setIsEditing] = useState(false);
  const [editedValues, setEditedValues] = useState<{ [key: string]: any }>(record.values);

  const handleEdit = () => {
    setIsEditing(true);
  };

  const handleCancel = () => {
    setIsEditing(false);
    setEditedValues(record.values);
  };

  const handleSave = async () => {
    try {
      await updateRecord(record.id, editedValues);
      setIsEditing(false);
      toast.success('记录已更新');
    } catch (error) {
      console.error('Failed to update record:', error);
      toast.error('更新记录失败');
    }
  };

  const handleDelete = async () => {
    if (window.confirm('确定要删除这条记录吗？此操作不可撤销。')) {
      try {
        await deleteRecord(record.id);
        onClose();
        toast.success('记录已删除');
      } catch (error) {
        console.error('Failed to delete record:', error);
        toast.error('删除记录失败');
      }
    }
  };

  const handleFieldChange = (fieldId: string, value: any) => {
    setEditedValues(prev => ({
      ...prev,
      [fieldId]: value
    }));
  };

  return (
    <div className="h-full flex flex-col bg-white dark:bg-primary">
      {/* Header */}
      <div className="p-4 border-b border-gray-200 dark:border-gray-700 flex items-center justify-between">
        <h3 className="text-lg font-medium">
          {isEditing ? '编辑记录' : '记录详情'}
        </h3>
        <div className="flex items-center space-x-1">
          <button 
            className={`p-1.5 rounded-md ${isActivityPanelOpen ? 'bg-accent/10 text-accent' : 'hover:bg-gray-100 dark:hover:bg-gray-800'}`}
            onClick={onToggleActivity}
          >
            <Activity size={18} />
          </button>
          <button 
            className="p-1.5 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800"
            onClick={onClose}
          >
            <X size={18} />
          </button>
        </div>
      </div>

      {/* Content */}
      <div className="flex-1 overflow-y-auto p-4">
        {isEditing ? (
          // Edit form
          <div className="space-y-4">
            {fields.map(field => (
              <div key={field.id} className="space-y-1">
                <label className="block text-sm font-medium">
                  {field.name}
                </label>
                <input
                  type="text"
                  className="input"
                  value={editedValues[field.id] || ''}
                  onChange={(e) => handleFieldChange(field.id, e.target.value)}
                />
              </div>
            ))}
          </div>
        ) : (
          // View details
          <div className="space-y-4">
            {fields.map(field => (
              <div key={field.id} className="border-b border-gray-200 dark:border-gray-700 pb-3 last:border-b-0">
                <div className="text-sm text-text-secondary mb-1">{field.name}</div>
                <div className="font-medium">
                  {record.values[field.id] !== undefined ? String(record.values[field.id]) : '—'}
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Footer */}
      <div className="p-4 border-t border-gray-200 dark:border-gray-700">
        {isEditing ? (
          <div className="flex items-center justify-between">
            <button 
              className="text-text-secondary hover:text-text"
              onClick={handleCancel}
            >
              取消
            </button>
            <button 
              className="btn btn-primary"
              onClick={handleSave}
            >
              <Save size={16} className="mr-1" />
              保存
            </button>
          </div>
        ) : (
          <div className="flex items-center justify-between">
            <div className="text-sm text-text-secondary flex items-center">
              <Clock size={14} className="mr-1" />
              {new Date(record.updatedAt).toLocaleString('zh-CN')}
            </div>
            <div className="flex space-x-2">
              <button 
                className="p-1.5 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800 text-text-secondary hover:text-accent"
                onClick={handleEdit}
              >
                <Edit3 size={18} />
              </button>
              <button 
                className="p-1.5 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800 text-text-secondary hover:text-error"
                onClick={handleDelete}
              >
                <Trash2 size={18} />
              </button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default RecordDetail;