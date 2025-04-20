import { useState } from 'react';
import { Menu, X, Search, Moon, Sun, Settings, Bell, HelpCircle, User } from 'lucide-react';
import { useTheme } from '../../context/ThemeContext';
import { useWorkspace } from '../../context/WorkspaceContext';

const Header = () => {
  const { theme, toggleTheme } = useTheme();
  const { currentWorkspace } = useWorkspace();
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const [isSearchOpen, setIsSearchOpen] = useState(false);

  return (
    <header className="bg-white dark:bg-primary border-b border-gray-200 dark:border-gray-700 h-14 flex items-center px-4 justify-between shadow-sm z-10">
      {/* Left section */}
      <div className="flex items-center">
        <button 
          className="lg:hidden mr-2 p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-800"
          onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
        >
          {isMobileMenuOpen ? <X size={20} /> : <Menu size={20} />}
        </button>
        
        <div className="flex items-center">
          <h1 className="text-lg font-semibold truncate max-w-[150px] md:max-w-xs">
            {currentWorkspace?.name || '一个类似 Airtable 的前端项目'}
          </h1>
        </div>
      </div>

      {/* Middle section - Search (hidden on mobile unless active) */}
      <div className={`${isSearchOpen ? 'absolute left-0 right-0 px-4 bg-white dark:bg-primary z-20' : 'hidden md:flex'} items-center mx-4 flex-1 max-w-xl`}>
        <div className="relative w-full">
          <input
            type="text"
            placeholder="搜索..."
            className="w-full py-1.5 pl-10 pr-4 rounded-md bg-gray-100 dark:bg-gray-800 border border-transparent focus:border-accent focus:bg-white dark:focus:bg-gray-700 focus:outline-none"
          />
          <Search size={18} className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500" />
          {isSearchOpen && (
            <button 
              className="absolute right-2 top-1/2 transform -translate-y-1/2 md:hidden"
              onClick={() => setIsSearchOpen(false)}
            >
              <X size={18} className="text-gray-500" />
            </button>
          )}
        </div>
      </div>

      {/* Right section */}
      <div className="flex items-center">
        <button 
          className="md:hidden mr-2 p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-800"
          onClick={() => setIsSearchOpen(true)}
        >
          <Search size={20} />
        </button>
        
        <button 
          className="p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-800"
          onClick={toggleTheme}
          aria-label={theme === 'dark' ? '切换到亮色模式' : '切换到暗色模式'}
        >
          {theme === 'dark' ? <Sun size={20} /> : <Moon size={20} />}
        </button>
        
        <button className="p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-800 relative">
          <Bell size={20} />
          <span className="absolute top-1 right-1 w-2 h-2 bg-accent rounded-full"></span>
        </button>
        
        <button className="p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-800">
          <HelpCircle size={20} />
        </button>
        
        <button className="p-2 rounded-full hover:bg-gray-100 dark:hover:bg-gray-800">
          <Settings size={20} />
        </button>
        
        <div className="ml-2 pl-2 border-l border-gray-200 dark:border-gray-700">
          <button className="flex items-center space-x-2 p-1 rounded-full hover:bg-gray-100 dark:hover:bg-gray-800">
            <div className="w-8 h-8 rounded-full bg-accent/20 flex items-center justify-center">
              <User size={16} className="text-accent" />
            </div>
          </button>
        </div>
      </div>
    </header>
  );
};

export default Header;