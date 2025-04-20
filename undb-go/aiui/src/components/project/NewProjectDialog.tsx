import React, { useState } from 'react';
import { X } from 'lucide-react';
import { useProject } from '../../context/ProjectContext';
import { Button } from '../ui/Button';
import { Input } from '../ui/Input';
 
import { createSpace } from '../../lib/api'; // 路径根据实际情况调整

interface NewProjectDialogProps {
  isOpen: boolean;
  onClose: () => void;
}

export function NewProjectDialog({ isOpen, onClose }: NewProjectDialogProps) {
  const { addProject } = useProject();
  const [isLoading, setIsLoading] = useState(false);
  const [formData, setFormData] = useState({
    name: '',
    description: ''
  });

  if (!isOpen) return null;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
  
    try {
      const newSpace = await createSpace({
        name: formData.name,
        description: formData.description,
        visibility: 'public',
      });
      console.log('newSpace', newSpace); // 调试用
      addProject(newSpace); // 确认 newSpace 结构和 addProject 预期一致
      onClose();
    } catch (error: any) {
      alert('创建空间失败: ' + (error.message || '未知错误'));
      console.error('Failed to create project:', error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white dark:bg-dark-card rounded-lg shadow-lg w-full max-w-md">
        <div className="flex justify-between items-center p-4 border-b border-gray-200 dark:border-dark-border">
          <h2 className="text-lg font-semibold text-text dark:text-white">Create New Project</h2>
          <button
            onClick={onClose}
            className="text-text-secondary hover:text-text dark:text-gray-400 dark:hover:text-white"
          >
            <X size={20} />
          </button>
        </div>

        <form onSubmit={handleSubmit} className="p-4 space-y-4">
          <Input
            label="Project Name"
            placeholder="Enter project name"
            value={formData.name}
            onChange={(e) => setFormData({ ...formData, name: e.target.value })}
            required
          />

          <Input
            label="Description"
            placeholder="Enter project description"
            value={formData.description}
            onChange={(e) => setFormData({ ...formData, description: e.target.value })}
          />

          <div className="flex justify-end gap-2 mt-6">
            <Button
              type="button"
              variant="ghost"
              onClick={onClose}
              disabled={isLoading}
            >
              Cancel
            </Button>
            <Button
              type="submit"
              isLoading={isLoading}
            >
              Create Project
            </Button>
          </div>
        </form>
      </div>
    </div>
  );
}