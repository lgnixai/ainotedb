import { FileText, Folder, Search, Plus, Star, Clock, MoreHorizontal } from 'lucide-react';

const Documents = () => {
  // Mock documents data
  const documents = [
    { id: 1, title: '项目说明文档', starred: true, modified: '2023-06-15T10:30:00Z' },
    { id: 2, title: '团队协作指南', starred: false, modified: '2023-06-10T14:45:00Z' },
    { id: 3, title: '产品需求文档', starred: true, modified: '2023-06-05T09:20:00Z' },
    { id: 4, title: '营销策略', starred: false, modified: '2023-05-28T16:15:00Z' },
    { id: 5, title: '用户反馈汇总', starred: false, modified: '2023-05-20T11:00:00Z' },
  ];

  // Mock folders
  const folders = [
    { id: 1, name: '项目文档', count: 12 },
    { id: 2, name: '团队资源', count: 8 },
    { id: 3, name: '客户资料', count: 5 },
  ];

  return (
    <div className="container mx-auto p-6 max-w-6xl">
      <div className="flex items-center justify-between mb-8">
        <h1 className="text-2xl font-bold">文档</h1>
        <div className="flex space-x-3">
          <div className="relative max-w-md">
            <input
              type="text"
              placeholder="搜索文档..."
              className="input py-1.5 pl-10"
            />
            <Search size={18} className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500" />
          </div>
          <button className="btn btn-primary">
            <Plus size={18} className="mr-1" />
            新建文档
          </button>
        </div>
      </div>

      {/* Recent documents */}
      <div className="mb-10">
        <h2 className="text-xl font-semibold mb-4">最近文档</h2>
        <div className="card p-0 overflow-hidden">
          <div className="grid grid-cols-1 divide-y divide-gray-200 dark:divide-gray-700">
            {documents.map(doc => (
              <div 
                key={doc.id}
                className="flex items-center p-4 hover:bg-gray-50 dark:hover:bg-gray-800 cursor-pointer"
              >
                <div className="p-2 bg-accent/10 rounded-lg mr-4">
                  <FileText size={20} className="text-accent" />
                </div>
                <div className="flex-1">
                  <h3 className="font-medium">{doc.title}</h3>
                  <div className="text-xs text-text-secondary mt-1 flex items-center">
                    <Clock size={12} className="mr-1" />
                    修改于 {new Date(doc.modified).toLocaleDateString('zh-CN')}
                  </div>
                </div>
                {doc.starred && (
                  <Star size={18} className="text-warning mr-2" fill="#F59E0B" />
                )}
                <button className="p-1 rounded-full hover:bg-gray-200 dark:hover:bg-gray-700">
                  <MoreHorizontal size={18} />
                </button>
              </div>
            ))}
          </div>
        </div>
      </div>

      {/* Folders */}
      <div>
        <h2 className="text-xl font-semibold mb-4">文件夹</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          {folders.map(folder => (
            <div 
              key={folder.id}
              className="card hover:shadow-md transition-shadow cursor-pointer"
            >
              <div className="flex items-start">
                <div className="p-2 bg-secondary dark:bg-gray-700 rounded-lg mr-3">
                  <Folder size={24} className="text-accent" />
                </div>
                <div>
                  <h3 className="font-medium">{folder.name}</h3>
                  <p className="text-sm text-text-secondary mt-1">
                    {folder.count} 个文档
                  </p>
                </div>
              </div>
            </div>
          ))}
          <div className="card border-2 border-dashed border-gray-200 dark:border-gray-700 flex items-center justify-center min-h-[100px] hover:border-accent cursor-pointer">
            <div className="flex items-center">
              <Plus size={20} className="text-accent mr-2" />
              <span className="font-medium text-accent">新建文件夹</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Documents;