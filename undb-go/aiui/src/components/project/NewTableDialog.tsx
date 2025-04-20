import React, { useState } from 'react';
import { X } from 'lucide-react';
import { Button } from '../ui/Button';
import { Input } from '../ui/Input';
import { createTable } from '../../lib/api';

interface NewTableDialogProps {
  isOpen: boolean;
  onClose: () => void;
  spaceId: string;
  onTableCreated: () => void;
}

export function NewTableDialog({ isOpen, onClose, spaceId, onTableCreated }: NewTableDialogProps) {
  const [isLoading, setIsLoading] = useState(false);
  const [name, setName] = useState('');

  if (!isOpen) return null;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    try {
      await createTable(spaceId, name);
      setName('');
      onTableCreated();
      onClose();
    } catch (error: any) {
      alert('创建表失败: ' + (error.message || '未知错误'));
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white dark:bg-dark-card rounded-lg shadow-lg w-full max-w-md">
        <div className="flex justify-between items-center p-4 border-b border-gray-200 dark:border-dark-border">
          <h2 className="text-lg font-semibold text-text dark:text-white">新建表</h2>
          <button
            onClick={onClose}
            className="text-text-secondary hover:text-text dark:text-gray-400 dark:hover:text-white"
          >
            <X size={20} />
          </button>
        </div>
        <form onSubmit={handleSubmit} className="p-4 space-y-4">
          <Input
            label="表名"
            placeholder="请输入表名"
            value={name}
            onChange={e => setName(e.target.value)}
            required
          />
          <div className="flex justify-end gap-2 mt-6">
            <Button type="button" variant="ghost" onClick={onClose} disabled={isLoading}>取消</Button>
            <Button type="submit" isLoading={isLoading}>创建表</Button>
          </div>
        </form>
      </div>
    </div>
  );
}
