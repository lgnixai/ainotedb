import { useState, useEffect } from 'react';
import { X, Send } from 'lucide-react';
import { fetchActivities, fetchComments, addComment, fetchUsers } from '../../api/mockApi';
import { Activity, Comment, User } from '../../types';
import toast from 'react-hot-toast';

interface ActivityPanelProps {
  recordId: string;
  tableId: string;
  onClose: () => void;
}

const ActivityPanel: React.FC<ActivityPanelProps> = ({ recordId, tableId, onClose }) => {
  const [activities, setActivities] = useState<Activity[]>([]);
  const [comments, setComments] = useState<Comment[]>([]);
  const [users, setUsers] = useState<User[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [commentText, setCommentText] = useState('');
  const [activeTab, setActiveTab] = useState<'all' | 'comments' | 'activity'>('all');

  useEffect(() => {
    const loadData = async () => {
      setIsLoading(true);
      try {
        const [activitiesData, commentsData, usersData] = await Promise.all([
          fetchActivities(tableId, recordId),
          fetchComments(recordId),
          fetchUsers()
        ]);
        
        setActivities(activitiesData);
        setComments(commentsData);
        setUsers(usersData);
      } catch (error) {
        console.error('Failed to load activity data:', error);
        toast.error('加载活动记录失败');
      } finally {
        setIsLoading(false);
      }
    };

    loadData();
  }, [recordId, tableId]);

  const handleAddComment = async () => {
    if (!commentText.trim()) return;
    
    try {
      // Use the first user for mock purposes - in a real app, this would be the current user
      const userId = users[0]?.id;
      
      if (!userId) {
        toast.error('无法添加评论：用户信息不可用');
        return;
      }
      
      await addComment(recordId, userId, commentText);
      setCommentText('');
      
      // Refresh comments
      const newComments = await fetchComments(recordId);
      setComments(newComments);
      
      toast.success('评论已添加');
    } catch (error) {
      console.error('Failed to add comment:', error);
      toast.error('添加评论失败');
    }
  };

  const getUserById = (userId: string): User | undefined => {
    return users.find(user => user.id === userId);
  };

  const renderActivityItem = (activity: Activity) => {
    const user = getUserById(activity.userId);
    const date = new Date(activity.timestamp).toLocaleString('zh-CN');
    
    let actionText = '';
    switch (activity.action) {
      case 'create':
        actionText = '创建了记录';
        break;
      case 'update':
        actionText = '更新了记录';
        break;
      case 'delete':
        actionText = '删除了记录';
        break;
    }

    return (
      <div className="flex items-start space-x-3 p-3 hover:bg-gray-50 dark:hover:bg-gray-800/50 rounded-md">
        <div className="w-8 h-8 rounded-full bg-gray-200 dark:bg-gray-700 flex-shrink-0 overflow-hidden">
          {user?.avatarUrl ? (
            <img src={user.avatarUrl} alt={user.name} className="w-full h-full object-cover" />
          ) : (
            <div className="w-full h-full flex items-center justify-center text-text-secondary">
              {user?.name.charAt(0) || '?'}
            </div>
          )}
        </div>
        <div className="flex-1">
          <div className="flex items-baseline justify-between">
            <div className="font-medium">{user?.name || '未知用户'}</div>
            <div className="text-xs text-text-secondary">{date}</div>
          </div>
          <div className="text-sm mt-1">{actionText}</div>
          {activity.action === 'update' && activity.fieldId && (
            <div className="mt-1 text-sm bg-gray-50 dark:bg-gray-800 p-2 rounded">
              <div className="text-text-secondary">字段：{activity.fieldId}</div>
              <div className="line-through text-text-secondary">旧值：{activity.oldValue}</div>
              <div className="text-accent">新值：{activity.newValue}</div>
            </div>
          )}
        </div>
      </div>
    );
  };

  const renderCommentItem = (comment: Comment) => {
    const user = getUserById(comment.userId);
    const date = new Date(comment.createdAt).toLocaleString('zh-CN');

    return (
      <div className="flex items-start space-x-3 p-3 hover:bg-gray-50 dark:hover:bg-gray-800/50 rounded-md">
        <div className="w-8 h-8 rounded-full bg-gray-200 dark:bg-gray-700 flex-shrink-0 overflow-hidden">
          {user?.avatarUrl ? (
            <img src={user.avatarUrl} alt={user.name} className="w-full h-full object-cover" />
          ) : (
            <div className="w-full h-full flex items-center justify-center text-text-secondary">
              {user?.name.charAt(0) || '?'}
            </div>
          )}
        </div>
        <div className="flex-1">
          <div className="flex items-baseline justify-between">
            <div className="font-medium">{user?.name || '未知用户'}</div>
            <div className="text-xs text-text-secondary">{date}</div>
          </div>
          <div className="text-sm mt-1 whitespace-pre-wrap">{comment.text}</div>
        </div>
      </div>
    );
  };

  // Combine and sort activities and comments chronologically
  const allItems = [...activities, ...comments].sort((a, b) => {
    const dateA = new Date(a.timestamp || a.createdAt);
    const dateB = new Date(b.timestamp || b.createdAt);
    return dateB.getTime() - dateA.getTime();
  });

  const filteredItems = activeTab === 'all' 
    ? allItems 
    : activeTab === 'comments' 
      ? comments 
      : activities;

  return (
    <div className="h-full flex flex-col bg-white dark:bg-primary">
      {/* Header */}
      <div className="p-4 border-b border-gray-200 dark:border-gray-700 flex items-center justify-between">
        <h3 className="text-lg font-medium">活动记录</h3>
        <button 
          className="p-1.5 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800"
          onClick={onClose}
        >
          <X size={18} />
        </button>
      </div>
      
      {/* Tabs */}
      <div className="flex border-b border-gray-200 dark:border-gray-700">
        <button 
          className={`flex-1 py-2 text-sm font-medium ${
            activeTab === 'all' 
              ? 'text-accent border-b-2 border-accent' 
              : 'text-text-secondary hover:text-text'
          }`}
          onClick={() => setActiveTab('all')}
        >
          全部
        </button>
        <button 
          className={`flex-1 py-2 text-sm font-medium ${
            activeTab === 'comments' 
              ? 'text-accent border-b-2 border-accent' 
              : 'text-text-secondary hover:text-text'
          }`}
          onClick={() => setActiveTab('comments')}
        >
          评论
        </button>
        <button 
          className={`flex-1 py-2 text-sm font-medium ${
            activeTab === 'activity' 
              ? 'text-accent border-b-2 border-accent' 
              : 'text-text-secondary hover:text-text'
          }`}
          onClick={() => setActiveTab('activity')}
        >
          历史
        </button>
      </div>

      {/* Content */}
      <div className="flex-1 overflow-y-auto p-2">
        {isLoading ? (
          <div className="flex items-center justify-center h-32">
            <div className="w-8 h-8 border-4 border-accent border-t-transparent rounded-full animate-spin"></div>
          </div>
        ) : filteredItems.length === 0 ? (
          <div className="text-center text-text-secondary p-6">
            {activeTab === 'all' 
              ? '暂无活动记录'
              : activeTab === 'comments'
                ? '暂无评论'
                : '暂无历史记录'}
          </div>
        ) : (
          <div className="space-y-1">
            {filteredItems.map(item => (
              'text' in item 
                ? renderCommentItem(item as Comment) 
                : renderActivityItem(item as Activity)
            ))}
          </div>
        )}
      </div>

      {/* Comment form */}
      <div className="p-3 border-t border-gray-200 dark:border-gray-700">
        <div className="flex items-center space-x-2">
          <input
            type="text"
            placeholder="添加评论..."
            className="input py-1.5"
            value={commentText}
            onChange={(e) => setCommentText(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === 'Enter' && !e.shiftKey) {
                e.preventDefault();
                handleAddComment();
              }
            }}
          />
          <button 
            className={`p-2 rounded-md ${
              commentText.trim() 
                ? 'bg-accent text-white hover:bg-accent/90' 
                : 'bg-gray-100 text-text-secondary dark:bg-gray-800'
            }`}
            onClick={handleAddComment}
            disabled={!commentText.trim()}
          >
            <Send size={18} />
          </button>
        </div>
      </div>
    </div>
  );
};

export default ActivityPanel;