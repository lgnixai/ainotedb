import { Workspace, Table, Record, View, User, Activity, Comment, Field, FieldType } from '../types';

// Mock data generation utilities
const generateId = () => Math.random().toString(36).substr(2, 9);

const generateTimestamp = (daysAgo = 0) => {
  const date = new Date();
  date.setDate(date.getDate() - daysAgo);
  return date.toISOString();
};

// Sample users
const users: User[] = [
  { 
    id: 'user1', 
    name: '张三', 
    email: 'zhangsan@example.com',
    avatarUrl: 'https://i.pravatar.cc/150?img=1'
  },
  { 
    id: 'user2', 
    name: '李四', 
    email: 'lisi@example.com',
    avatarUrl: 'https://i.pravatar.cc/150?img=2'
  },
  { 
    id: 'user3', 
    name: '王五', 
    email: 'wangwu@example.com',
    avatarUrl: 'https://i.pravatar.cc/150?img=3'
  },
];

// Sample workspaces
const workspaces: Workspace[] = [
  {
    id: 'ws1',
    name: '项目管理',
    description: '团队项目跟踪和任务管理',
    createdAt: generateTimestamp(30),
    updatedAt: generateTimestamp(2),
  },
  {
    id: 'ws2',
    name: '客户关系管理',
    description: '客户信息和销售线索跟踪',
    createdAt: generateTimestamp(20),
    updatedAt: generateTimestamp(5),
  },
  {
    id: 'ws3',
    name: '产品开发',
    description: '产品路线图和功能规划',
    createdAt: generateTimestamp(15),
    updatedAt: generateTimestamp(1),
  },
];

// Generate fields for a table
const generateFields = (tableType: string): Field[] => {
  const commonFields: Field[] = [
    {
      id: 'name',
      name: '名称',
      type: 'text',
      isRequired: true,
    },
    {
      id: 'notes',
      name: '备注',
      type: 'text',
    },
  ];

  if (tableType === 'projects') {
    return [
      ...commonFields,
      {
        id: 'status',
        name: '状态',
        type: 'select',
        options: [
          { id: 's1', label: '未开始', color: '#CBD5E1' },
          { id: 's2', label: '进行中', color: '#3B82F6' },
          { id: 's3', label: '已完成', color: '#22C55E' },
          { id: 's4', label: '已搁置', color: '#F97316' },
        ],
      },
      {
        id: 'priority',
        name: '优先级',
        type: 'select',
        options: [
          { id: 'p1', label: '低', color: '#CBD5E1' },
          { id: 'p2', label: '中', color: '#3B82F6' },
          { id: 'p3', label: '高', color: '#F97316' },
          { id: 'p4', label: '紧急', color: '#EF4444' },
        ],
      },
      {
        id: 'dueDate',
        name: '截止日期',
        type: 'date',
      },
      {
        id: 'assignee',
        name: '负责人',
        type: 'user',
      },
    ];
  } else if (tableType === 'tasks') {
    return [
      ...commonFields,
      {
        id: 'status',
        name: '状态',
        type: 'select',
        options: [
          { id: 's1', label: '待办', color: '#CBD5E1' },
          { id: 's2', label: '进行中', color: '#3B82F6' },
          { id: 's3', label: '审核中', color: '#F97316' },
          { id: 's4', label: '已完成', color: '#22C55E' },
        ],
      },
      {
        id: 'completed',
        name: '已完成',
        type: 'checkbox',
      },
      {
        id: 'dueDate',
        name: '截止日期',
        type: 'date',
      },
      {
        id: 'assignee',
        name: '负责人',
        type: 'user',
      },
      {
        id: 'estimatedHours',
        name: '预估时间（小时）',
        type: 'number',
      },
    ];
  } else if (tableType === 'clients') {
    return [
      ...commonFields,
      {
        id: 'company',
        name: '公司',
        type: 'text',
      },
      {
        id: 'contactName',
        name: '联系人',
        type: 'text',
      },
      {
        id: 'email',
        name: '邮箱',
        type: 'text',
      },
      {
        id: 'phone',
        name: '电话',
        type: 'text',
      },
      {
        id: 'status',
        name: '状态',
        type: 'select',
        options: [
          { id: 's1', label: '潜在客户', color: '#CBD5E1' },
          { id: 's2', label: '活跃客户', color: '#22C55E' },
          { id: 's3', label: '非活跃客户', color: '#F97316' },
        ],
      },
    ];
  }

  return commonFields;
};

// Generate tables for each workspace
const tables: Table[] = [
  {
    id: 'table1',
    workspaceId: 'ws1',
    name: '项目',
    description: '所有项目及其状态',
    fields: generateFields('projects'),
    createdAt: generateTimestamp(28),
    updatedAt: generateTimestamp(3),
  },
  {
    id: 'table2',
    workspaceId: 'ws1',
    name: '任务',
    description: '项目任务和进度跟踪',
    fields: generateFields('tasks'),
    createdAt: generateTimestamp(27),
    updatedAt: generateTimestamp(1),
  },
  {
    id: 'table3',
    workspaceId: 'ws2',
    name: '客户',
    description: '客户信息和联系方式',
    fields: generateFields('clients'),
    createdAt: generateTimestamp(19),
    updatedAt: generateTimestamp(4),
  },
];

// Generate records for each table
const generateRecords = (tableId: string, count = 10): Record[] => {
  const table = tables.find(t => t.id === tableId);
  if (!table) return [];

  const records: Record[] = [];
  
  for (let i = 0; i < count; i++) {
    const record: Record = {
      id: generateId(),
      tableId,
      values: {},
      createdAt: generateTimestamp(Math.floor(Math.random() * 20)),
      updatedAt: generateTimestamp(Math.floor(Math.random() * 5)),
    };

    // Fill with some sample data based on field type
    table.fields.forEach(field => {
      switch (field.type) {
        case 'text':
          if (field.id === 'name') {
            record.values[field.id] = `示例 ${i + 1}`;
          } else if (field.id === 'company') {
            record.values[field.id] = `公司 ${i + 1}`;
          } else if (field.id === 'contactName') {
            record.values[field.id] = users[i % users.length].name;
          } else if (field.id === 'email') {
            record.values[field.id] = `example${i}@example.com`;
          } else if (field.id === 'phone') {
            record.values[field.id] = `1388888${1000 + i}`;
          } else {
            record.values[field.id] = Math.random() > 0.3 ? `备注内容 ${i + 1}` : '';
          }
          break;
        case 'number':
          record.values[field.id] = Math.floor(Math.random() * 40) + 1;
          break;
        case 'select':
          if (field.options) {
            record.values[field.id] = field.options[i % field.options.length].id;
          }
          break;
        case 'date':
          const date = new Date();
          date.setDate(date.getDate() + Math.floor(Math.random() * 30));
          record.values[field.id] = date.toISOString().split('T')[0];
          break;
        case 'checkbox':
          record.values[field.id] = Math.random() > 0.5;
          break;
        case 'user':
          record.values[field.id] = users[i % users.length].id;
          break;
        default:
          break;
      }
    });

    records.push(record);
  }

  return records;
};

// Generate views for each table
const generateViews = (tableId: string): View[] => {
  const table = tables.find(t => t.id === tableId);
  if (!table) return [];

  return [
    {
      id: `view-grid-${tableId}`,
      tableId,
      name: '表格视图',
      type: 'grid',
      fieldOrder: table.fields.map(field => field.id),
    },
    {
      id: `view-kanban-${tableId}`,
      tableId,
      name: '看板视图',
      type: 'kanban',
      fieldOrder: table.fields.map(field => field.id),
      groupBy: 'status',
    },
  ];
};

// Create sample records and views
let records: Record[] = [];
let views: View[] = [];

tables.forEach(table => {
  records = [...records, ...generateRecords(table.id, 15)];
  views = [...views, ...generateViews(table.id)];
});

// Generate mock activities
const activities: Activity[] = records.slice(0, 20).map((record, index) => ({
  id: generateId(),
  tableId: record.tableId,
  recordId: record.id,
  userId: users[index % users.length].id,
  action: index % 3 === 0 ? 'create' : (index % 3 === 1 ? 'update' : 'delete'),
  fieldId: index % 3 === 1 ? Object.keys(record.values)[0] : undefined,
  oldValue: index % 3 === 1 ? '旧值' : undefined,
  newValue: index % 3 === 1 ? '新值' : undefined,
  timestamp: generateTimestamp(Math.floor(Math.random() * 10)),
}));

// Generate mock comments
const comments: Comment[] = records.slice(0, 10).map((record, index) => ({
  id: generateId(),
  tableId: record.tableId,
  recordId: record.id,
  userId: users[index % users.length].id,
  text: `这是一条示例评论。评论编号 ${index + 1}`,
  createdAt: generateTimestamp(Math.floor(Math.random() * 10)),
}));

// API delay simulation
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

// Mock API functions
export const fetchWorkspaces = async (): Promise<Workspace[]> => {
  await delay(300);
  return [...workspaces];
};

export const fetchTables = async (workspaceId: string): Promise<Table[]> => {
  await delay(300);
  return tables.filter(table => table.workspaceId === workspaceId);
};

export const fetchRecords = async (tableId: string): Promise<Record[]> => {
  await delay(300);
  return records.filter(record => record.tableId === tableId);
};

export const fetchRecord = async (recordId: string): Promise<Record | null> => {
  await delay(200);
  return records.find(record => record.id === recordId) || null;
};

export const fetchViews = async (tableId: string): Promise<View[]> => {
  await delay(200);
  return views.filter(view => view.tableId === tableId);
};

export const fetchActivities = async (tableId: string, recordId?: string): Promise<Activity[]> => {
  await delay(200);
  return activities.filter(activity => {
    if (recordId) {
      return activity.tableId === tableId && activity.recordId === recordId;
    }
    return activity.tableId === tableId;
  });
};

export const fetchComments = async (recordId: string): Promise<Comment[]> => {
  await delay(200);
  return comments.filter(comment => comment.recordId === recordId);
};

export const fetchUsers = async (): Promise<User[]> => {
  await delay(200);
  return [...users];
};

export const createRecord = async (tableId: string, values: { [key: string]: any }): Promise<Record> => {
  await delay(400);
  const newRecord: Record = {
    id: generateId(),
    tableId,
    values,
    createdAt: new Date().toISOString(),
    updatedAt: new Date().toISOString(),
  };
  records.push(newRecord);
  return newRecord;
};

export const updateRecord = async (recordId: string, values: { [key: string]: any }): Promise<Record> => {
  await delay(400);
  const recordIndex = records.findIndex(r => r.id === recordId);
  if (recordIndex === -1) {
    throw new Error('Record not found');
  }
  
  const updatedRecord = {
    ...records[recordIndex],
    values: { ...records[recordIndex].values, ...values },
    updatedAt: new Date().toISOString(),
  };
  
  records[recordIndex] = updatedRecord;
  return updatedRecord;
};

export const deleteRecord = async (recordId: string): Promise<boolean> => {
  await delay(400);
  const initialLength = records.length;
  records = records.filter(r => r.id !== recordId);
  return records.length < initialLength;
};

export const addComment = async (recordId: string, userId: string, text: string): Promise<Comment> => {
  await delay(300);
  const record = records.find(r => r.id === recordId);
  if (!record) {
    throw new Error('Record not found');
  }
  
  const newComment: Comment = {
    id: generateId(),
    tableId: record.tableId,
    recordId,
    userId,
    text,
    createdAt: new Date().toISOString(),
  };
  
  comments.push(newComment);
  return newComment;
};