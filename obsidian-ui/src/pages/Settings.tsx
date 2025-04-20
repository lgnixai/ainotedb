import { useState } from 'react';
import { Save, User, Lock, Bell, Globe, Database, Shield, Monitor } from 'lucide-react';

const Settings = () => {
  const [activeTab, setActiveTab] = useState('profile');

  const renderTabContent = () => {
    switch (activeTab) {
      case 'profile':
        return (
          <div className="space-y-6">
            <h2 className="text-xl font-semibold mb-4">个人资料</h2>
            
            <div className="flex flex-col md:flex-row md:space-x-4">
              <div className="mb-4 md:mb-0 md:w-1/3">
                <div className="flex flex-col items-center p-6 bg-secondary dark:bg-secondary/50 rounded-lg">
                  <div className="w-24 h-24 rounded-full bg-accent/20 flex items-center justify-center mb-4">
                    <User size={40} className="text-accent" />
                  </div>
                  <button className="btn btn-secondary w-full mb-2">更换头像</button>
                  <button className="text-error text-sm">删除头像</button>
                </div>
              </div>
              
              <div className="md:w-2/3">
                <div className="space-y-4">
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div>
                      <label className="block text-sm font-medium mb-1">姓名</label>
                      <input type="text" className="input" defaultValue="张三" />
                    </div>
                    <div>
                      <label className="block text-sm font-medium mb-1">显示名称</label>
                      <input type="text" className="input" defaultValue="zhangsan" />
                    </div>
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium mb-1">邮箱</label>
                    <input type="email" className="input" defaultValue="zhangsan@example.com" />
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium mb-1">职位</label>
                    <input type="text" className="input" defaultValue="项目经理" />
                  </div>
                  
                  <div>
                    <label className="block text-sm font-medium mb-1">个人简介</label>
                    <textarea 
                      className="input min-h-[100px]" 
                      defaultValue="负责项目管理和团队协调，具有5年项目管理经验。"
                    />
                  </div>
                  
                  <div className="flex justify-end">
                    <button className="btn btn-primary">
                      <Save size={16} className="mr-1" />
                      保存更改
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        );
      
      case 'account':
        return (
          <div className="space-y-6">
            <h2 className="text-xl font-semibold mb-4">账户设置</h2>
            
            <div className="card p-6">
              <h3 className="text-lg font-medium mb-4">更改密码</h3>
              <div className="space-y-4">
                <div>
                  <label className="block text-sm font-medium mb-1">当前密码</label>
                  <input type="password" className="input" />
                </div>
                <div>
                  <label className="block text-sm font-medium mb-1">新密码</label>
                  <input type="password" className="input" />
                </div>
                <div>
                  <label className="block text-sm font-medium mb-1">确认新密码</label>
                  <input type="password" className="input" />
                </div>
                <div className="flex justify-end">
                  <button className="btn btn-primary">更新密码</button>
                </div>
              </div>
            </div>
            
            <div className="card p-6">
              <h3 className="text-lg font-medium mb-4">账户安全</h3>
              <div className="space-y-4">
                <div className="flex items-center justify-between p-4 border border-gray-200 dark:border-gray-700 rounded-md">
                  <div>
                    <div className="font-medium">双因素认证</div>
                    <div className="text-sm text-text-secondary">使用手机应用进行双重验证</div>
                  </div>
                  <div>
                    <button className="btn btn-secondary">设置</button>
                  </div>
                </div>
                
                <div className="flex items-center justify-between p-4 border border-gray-200 dark:border-gray-700 rounded-md">
                  <div>
                    <div className="font-medium">登录历史</div>
                    <div className="text-sm text-text-secondary">查看账户的登录历史记录</div>
                  </div>
                  <div>
                    <button className="btn btn-secondary">查看</button>
                  </div>
                </div>
                
                <div className="flex items-center justify-between p-4 border border-error/20 bg-error/5 rounded-md">
                  <div>
                    <div className="font-medium text-error">删除账户</div>
                    <div className="text-sm text-text-secondary">永久删除您的账户和所有数据</div>
                  </div>
                  <div>
                    <button className="btn btn-danger">删除账户</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        );
        
      case 'appearance':
        return (
          <div className="space-y-6">
            <h2 className="text-xl font-semibold mb-4">外观设置</h2>
            
            <div className="card p-6">
              <h3 className="text-lg font-medium mb-4">主题</h3>
              <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
                <div className="border border-accent rounded-md p-4 flex flex-col items-center">
                  <div className="bg-white w-full h-20 mb-3 rounded flex items-center justify-center">
                    <Sun size={24} className="text-text" />
                  </div>
                  <div className="font-medium">浅色模式</div>
                </div>
                
                <div className="border border-gray-200 dark:border-gray-700 rounded-md p-4 flex flex-col items-center">
                  <div className="bg-gray-900 w-full h-20 mb-3 rounded flex items-center justify-center">
                    <Moon size={24} className="text-white" />
                  </div>
                  <div className="font-medium">深色模式</div>
                </div>
                
                <div className="border border-gray-200 dark:border-gray-700 rounded-md p-4 flex flex-col items-center">
                  <div className="bg-gradient-to-b from-white to-gray-900 w-full h-20 mb-3 rounded flex items-center justify-center">
                    <Monitor size={24} className="text-gray-700" />
                  </div>
                  <div className="font-medium">跟随系统</div>
                </div>
              </div>
            </div>
            
            <div className="card p-6">
              <h3 className="text-lg font-medium mb-4">其他设置</h3>
              <div className="space-y-4">
                <div className="flex items-center justify-between">
                  <div>
                    <div className="font-medium">紧凑模式</div>
                    <div className="text-sm text-text-secondary">减少元素间距，显示更多内容</div>
                  </div>
                  <label className="relative inline-flex items-center cursor-pointer">
                    <input type="checkbox" className="sr-only peer" />
                    <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-accent/20 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                  </label>
                </div>
                
                <div className="flex items-center justify-between">
                  <div>
                    <div className="font-medium">动画效果</div>
                    <div className="text-sm text-text-secondary">启用界面过渡动画</div>
                  </div>
                  <label className="relative inline-flex items-center cursor-pointer">
                    <input type="checkbox" className="sr-only peer" defaultChecked />
                    <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-accent/20 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                  </label>
                </div>
              </div>
            </div>
          </div>
        );
        
      case 'preferences':
        return (
          <div className="space-y-6">
            <h2 className="text-xl font-semibold mb-4">偏好设置</h2>
            
            <div className="card p-6">
              <h3 className="text-lg font-medium mb-4">通知</h3>
              <div className="space-y-4">
                <div className="flex items-center justify-between">
                  <div>
                    <div className="font-medium">电子邮件通知</div>
                    <div className="text-sm text-text-secondary">接收有关项目更新的电子邮件</div>
                  </div>
                  <label className="relative inline-flex items-center cursor-pointer">
                    <input type="checkbox" className="sr-only peer" defaultChecked />
                    <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-accent/20 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                  </label>
                </div>
                
                <div className="flex items-center justify-between">
                  <div>
                    <div className="font-medium">应用内通知</div>
                    <div className="text-sm text-text-secondary">显示应用内通知</div>
                  </div>
                  <label className="relative inline-flex items-center cursor-pointer">
                    <input type="checkbox" className="sr-only peer" defaultChecked />
                    <div className="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-accent/20 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-accent"></div>
                  </label>
                </div>
              </div>
            </div>
            
            <div className="card p-6">
              <h3 className="text-lg font-medium mb-4">语言和区域</h3>
              <div className="space-y-4">
                <div>
                  <label className="block text-sm font-medium mb-1">语言</label>
                  <select className="input">
                    <option value="zh-CN">简体中文</option>
                    <option value="en-US">English (US)</option>
                    <option value="ja-JP">日本語</option>
                  </select>
                </div>
                
                <div>
                  <label className="block text-sm font-medium mb-1">时区</label>
                  <select className="input">
                    <option value="Asia/Shanghai">中国标准时间 (GMT+8)</option>
                    <option value="America/New_York">Eastern Time (GMT-5)</option>
                    <option value="Europe/London">Greenwich Mean Time (GMT+0)</option>
                  </select>
                </div>
                
                <div>
                  <label className="block text-sm font-medium mb-1">日期格式</label>
                  <select className="input">
                    <option value="yyyy-MM-dd">YYYY-MM-DD</option>
                    <option value="MM/dd/yyyy">MM/DD/YYYY</option>
                    <option value="dd/MM/yyyy">DD/MM/YYYY</option>
                  </select>
                </div>
              </div>
            </div>
          </div>
        );
        
      default:
        return null;
    }
  };

  return (
    <div className="container mx-auto p-6 max-w-6xl">
      <h1 className="text-2xl font-bold mb-8">设置</h1>
      
      <div className="flex flex-col md:flex-row md:space-x-6">
        {/* Sidebar */}
        <div className="w-full md:w-64 mb-6 md:mb-0">
          <div className="bg-white dark:bg-primary rounded-lg shadow-sm">
            <nav className="p-2">
              <button
                className={`w-full flex items-center px-4 py-2 rounded-md text-left mb-1 ${
                  activeTab === 'profile' 
                    ? 'bg-accent/10 text-accent' 
                    : 'hover:bg-gray-100 dark:hover:bg-gray-800'
                }`}
                onClick={() => setActiveTab('profile')}
              >
                <User size={18} className="mr-2" />
                个人资料
              </button>
              
              <button
                className={`w-full flex items-center px-4 py-2 rounded-md text-left mb-1 ${
                  activeTab === 'account' 
                    ? 'bg-accent/10 text-accent' 
                    : 'hover:bg-gray-100 dark:hover:bg-gray-800'
                }`}
                onClick={() => setActiveTab('account')}
              >
                <Lock size={18} className="mr-2" />
                账户安全
              </button>
              
              <button
                className={`w-full flex items-center px-4 py-2 rounded-md text-left mb-1 ${
                  activeTab === 'appearance' 
                    ? 'bg-accent/10 text-accent' 
                    : 'hover:bg-gray-100 dark:hover:bg-gray-800'
                }`}
                onClick={() => setActiveTab('appearance')}
              >
                <Monitor size={18} className="mr-2" />
                外观设置
              </button>
              
              <button
                className={`w-full flex items-center px-4 py-2 rounded-md text-left mb-1 ${
                  activeTab === 'preferences' 
                    ? 'bg-accent/10 text-accent' 
                    : 'hover:bg-gray-100 dark:hover:bg-gray-800'
                }`}
                onClick={() => setActiveTab('preferences')}
              >
                <Bell size={18} className="mr-2" />
                偏好设置
              </button>
            </nav>
          </div>
        </div>
        
        {/* Main content */}
        <div className="flex-1">
          <div className="bg-white dark:bg-primary rounded-lg shadow-sm p-6">
            {renderTabContent()}
          </div>
        </div>
      </div>
    </div>
  );
};

// Mocked SVG icons for themes
const Sun = ({ size, className }) => (
  <svg xmlns="http://www.w3.org/2000/svg" width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className={className}>
    <circle cx="12" cy="12" r="5"></circle>
    <line x1="12" y1="1" x2="12" y2="3"></line>
    <line x1="12" y1="21" x2="12" y2="23"></line>
    <line x1="4.22" y1="4.22" x2="5.64" y2="5.64"></line>
    <line x1="18.36" y1="18.36" x2="19.78" y2="19.78"></line>
    <line x1="1" y1="12" x2="3" y2="12"></line>
    <line x1="21" y1="12" x2="23" y2="12"></line>
    <line x1="4.22" y1="19.78" x2="5.64" y2="18.36"></line>
    <line x1="18.36" y1="5.64" x2="19.78" y2="4.22"></line>
  </svg>
);

const Moon = ({ size, className }) => (
  <svg xmlns="http://www.w3.org/2000/svg" width={size} height={size} viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className={className}>
    <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
  </svg>
);

export default Settings;