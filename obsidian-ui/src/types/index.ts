// Workspace and Table types
export interface Workspace {
  id: string;
  name: string;
  description?: string;
  createdAt: string;
  updatedAt: string;
}

export interface Table {
  id: string;
  workspaceId: string;
  name: string;
  description?: string;
  fields: Field[];
  createdAt: string;
  updatedAt: string;
}

// Field types
export type FieldType = 
  | 'text' 
  | 'number' 
  | 'select' 
  | 'multiSelect' 
  | 'date' 
  | 'checkbox'
  | 'file'
  | 'link'
  | 'user'
  | 'formula';

export interface Field {
  id: string;
  name: string;
  type: FieldType;
  options?: FieldOption[];
  isRequired?: boolean;
  isUnique?: boolean;
  defaultValue?: any;
}

export interface FieldOption {
  id: string;
  label: string;
  color?: string;
}

// Record types
export interface Record {
  id: string;
  tableId: string;
  values: { [fieldId: string]: any };
  createdAt: string;
  updatedAt: string;
}

// View types
export type ViewType = 'grid' | 'kanban' | 'calendar' | 'gallery';

export interface View {
  id: string;
  tableId: string;
  name: string;
  type: ViewType;
  filters?: Filter[];
  sorts?: Sort[];
  groupBy?: string; // Field ID
  fieldOrder?: string[]; // Field IDs in order
  hiddenFields?: string[]; // Field IDs that are hidden
}

// Filter types
export type FilterOperator = 
  | 'eq' | 'neq' 
  | 'contains' | 'notContains'
  | 'gt' | 'gte' | 'lt' | 'lte'
  | 'isEmpty' | 'isNotEmpty';

export interface Filter {
  id: string;
  fieldId: string;
  operator: FilterOperator;
  value: any;
}

// Sort types
export interface Sort {
  fieldId: string;
  direction: 'asc' | 'desc';
}

// Activity and Comment types
export interface Activity {
  id: string;
  tableId: string;
  recordId?: string;
  userId: string;
  action: 'create' | 'update' | 'delete';
  fieldId?: string;
  oldValue?: any;
  newValue?: any;
  timestamp: string;
}

export interface Comment {
  id: string;
  tableId: string;
  recordId: string;
  userId: string;
  text: string;
  createdAt: string;
  updatedAt?: string;
}

// User types
export interface User {
  id: string;
  name: string;
  email: string;
  avatarUrl?: string;
}