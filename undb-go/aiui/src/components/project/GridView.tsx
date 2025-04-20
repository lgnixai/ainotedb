import React, { useState, useMemo } from 'react';
import {
  useReactTable,
  getCoreRowModel,
  flexRender,
  createColumnHelper,
  ColumnDef,
  ColumnResizeMode,
} from '@tanstack/react-table';
import { useProject } from '../../context/ProjectContext';
import { Task } from '../../types';
import { Badge } from '../ui/Badge';
import { Input } from '../ui/Input';
import { Select } from '../ui/Select';
import { cn, formatDate } from '../../lib/utils';
import {
  Plus,
  MoreHorizontal,
  Check,
  AlertCircle,
  Clock,
  GripVertical,
} from 'lucide-react';
import {
  DndContext,
  DragEndEvent,
  MouseSensor,
  TouchSensor,
  useSensor,
  useSensors,
} from '@dnd-kit/core';
import {
  SortableContext,
  arrayMove,
  useSortable,
  verticalListSortingStrategy,
} from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';

const columnHelper = createColumnHelper<Task>();

const getPriorityBadge = (priority: string) => {
  switch (priority) {
    case 'Low':
      return <Badge variant="success">Low</Badge>;
    case 'Medium':
      return <Badge variant="warning">Medium</Badge>;
    case 'High':
      return <Badge variant="primary">High</Badge>;
    case 'Urgent':
      return <Badge variant="danger">Urgent</Badge>;
    default:
      return <Badge>{priority}</Badge>;
  }
};

const getStatusIcon = (status: string) => {
  switch (status) {
    case 'Completed':
      return <Check size={14} className="text-success-500 mr-1" />;
    case 'In Progress':
      return <Clock size={14} className="text-warning-500 mr-1" />;
    case 'On Hold':
      return <AlertCircle size={14} className="text-error-500 mr-1" />;
    default:
      return null;
  }
};

interface DraggableRowProps {
  row: any;
  children: React.ReactNode;
}

function DraggableRow({ row, children }: DraggableRowProps) {
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition,
    isDragging,
  } = useSortable({
    id: row.original.id,
  });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    opacity: isDragging ? 0.5 : 1,
  };

  return (
    <tr
      ref={setNodeRef}
      style={style}
      className={cn(
        'hover:bg-gray-50 dark:hover:bg-dark-hover cursor-pointer border-b border-gray-200 dark:border-dark-border',
        isDragging && 'bg-gray-50 dark:bg-dark-hover'
      )}
      {...attributes}
    >
      <td className="w-8 px-2">
        <div
          className="cursor-grab hover:bg-gray-100 dark:hover:bg-dark-card p-1 rounded"
          {...listeners}
        >
          <GripVertical size={16} className="text-text-secondary" />
        </div>
      </td>
      {children}
    </tr>
  );
}

export function GridView() {
  const { currentProject, currentView, updateTask } = useProject();
  const [columnResizeMode] = useState<ColumnResizeMode>('onChange');
  const [rowOrder, setRowOrder] = useState<string[]>([]);
  const [editingCell, setEditingCell] = useState<{
    rowId: string;
    columnId: string;
  } | null>(null);

  const sensors = useSensors(
    useSensor(MouseSensor, {
      activationConstraint: {
        distance: 8,
      },
    }),
    useSensor(TouchSensor, {
      activationConstraint: {
        delay: 200,
        tolerance: 8,
      },
    })
  );

  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event;
    if (!over || active.id === over.id) return;

    const oldIndex = rowOrder.indexOf(active.id as string);
    const newIndex = rowOrder.indexOf(over.id as string);

    setRowOrder(arrayMove(rowOrder, oldIndex, newIndex));
  };

  const columns = useMemo(() => {
    if (!currentView) return [];

    return [
      columnHelper.display({
        id: 'drag',
        size: 40,
        cell: () => null,
      }),
      ...currentView.columns.map((col) => {
        return columnHelper.accessor(col.accessor as keyof Task, {
          id: col.id,
          header: col.title,
          size: col.width,
          cell: ({ row, column, getValue }) => {
            const value = getValue();
            const isEditing =
              editingCell?.rowId === row.original.id &&
              editingCell?.columnId === column.id;

            if (isEditing) {
              if (col.type === 'select') {
                return (
                  <Select
                    options={
                      col.options?.map((opt) => ({
                        value: opt,
                        label: opt,
                      })) || []
                    }
                    value={value as string}
                    onChange={(e) => {
                      updateTask(row.original.id, {
                        [col.accessor]: e.target.value,
                      });
                      setEditingCell(null);
                    }}
                    autoFocus
                    onBlur={() => setEditingCell(null)}
                  />
                );
              }

              return (
                <Input
                  value={value as string}
                  onChange={(e) => {
                    updateTask(row.original.id, {
                      [col.accessor]: e.target.value,
                    });
                  }}
                  autoFocus
                  onBlur={() => setEditingCell(null)}
                  onKeyDown={(e) => {
                    if (e.key === 'Enter') {
                      setEditingCell(null);
                    }
                  }}
                />
              );
            }

            if (col.accessor === 'status') {
              return (
                <div className="flex items-center">
                  {getStatusIcon(value as string)}
                  <span>{value as string}</span>
                </div>
              );
            }

            if (col.accessor === 'priority') {
              return getPriorityBadge(value as string);
            }

            if (col.accessor === 'dueDate' && value) {
              return formatDate(value as Date);
            }

            if (col.accessor === 'assignee') {
              return (
                <div className="flex items-center">
                  <div className="w-6 h-6 rounded-full bg-accent-100 text-accent-700 flex items-center justify-center text-xs font-medium mr-2">
                    {(value as string)?.charAt(0)}
                  </div>
                  <span>{value as string}</span>
                </div>
              );
            }

            return (
              <div
                className="py-2 px-1 -mx-1 rounded hover:bg-gray-100 dark:hover:bg-dark-card"
                onClick={() =>
                  setEditingCell({
                    rowId: row.original.id,
                    columnId: column.id,
                  })
                }
              >
                {value as string}
              </div>
            );
          },
        });
      }),
      columnHelper.display({
        id: 'actions',
        size: 40,
        cell: () => (
          <button className="p-1 rounded-md hover:bg-gray-100 dark:hover:bg-dark-card">
            <MoreHorizontal size={16} className="text-text-secondary" />
          </button>
        ),
      }),
    ];
  }, [currentView, editingCell, updateTask]);

  const data = useMemo(() => {
    if (!currentProject) return [];
    return currentProject.tasks;
  }, [currentProject]);

  const table = useReactTable({
    data,
    columns,
    columnResizeMode,
    getCoreRowModel: getCoreRowModel(),
  });

  if (!currentProject || !currentView) return null;

  return (
    <div className="overflow-x-auto">
      <DndContext sensors={sensors} onDragEnd={handleDragEnd}>
        <table className="w-full border-collapse">
          <thead>
            {table.getHeaderGroups().map((headerGroup) => (
              <tr key={headerGroup.id} className="bg-secondary dark:bg-dark-card">
                {headerGroup.headers.map((header) => (
                  <th
                    key={header.id}
                    className="text-left px-4 py-3 text-sm font-medium text-text-secondary dark:text-gray-400 border-b border-gray-200 dark:border-dark-border relative"
                    style={{ width: header.getSize() }}
                  >
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                    <div
                      onMouseDown={header.getResizeHandler()}
                      onTouchStart={header.getResizeHandler()}
                      className={`absolute right-0 top-0 h-full w-1 cursor-col-resize select-none touch-none ${
                        header.column.getIsResizing()
                          ? 'bg-accent'
                          : 'bg-gray-200 dark:bg-dark-border'
                      }`}
                    />
                  </th>
                ))}
              </tr>
            ))}
          </thead>
          <tbody>
            <SortableContext
              items={data.map((item) => item.id)}
              strategy={verticalListSortingStrategy}
            >
              {table.getRowModel().rows.map((row) => (
                <DraggableRow key={row.id} row={row}>
                  {row.getVisibleCells().map((cell) => (
                    <td
                      key={cell.id}
                      className="px-4 py-3 text-sm text-text-secondary dark:text-gray-300"
                    >
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </td>
                  ))}
                </DraggableRow>
              ))}
            </SortableContext>
            <tr className="hover:bg-gray-50 dark:hover:bg-dark-hover cursor-pointer border-b border-gray-200 dark:border-dark-border">
              <td colSpan={table.getAllColumns().length} className="px-4 py-3">
                <button className="text-sm text-text-secondary dark:text-gray-400 flex items-center hover:text-accent">
                  <Plus size={16} className="mr-1" />
                  Add new task
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </DndContext>
    </div>
  );
}