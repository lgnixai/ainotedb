import React, { useState } from 'react';
import { useTheme } from '../../context/ThemeContext';
import { useProject } from '../../context/ProjectContext';
import { Button } from '../ui/Button';
import { Input } from '../ui/Input';
import { 
  Moon, 
  Sun, 
  Search, 
  Menu, 
  Plus, 
  Bell, 
  User, 
  HelpCircle,
  X
} from 'lucide-react';

export function Header() {
  const { theme, toggleTheme } = useTheme();
  const { currentProject } = useProject();
  const [searchOpen, setSearchOpen] = useState(false);
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  return (
    <header className="sticky top-0 z-40 w-full border-b border-gray-200 bg-primary dark:bg-dark-bg dark:border-dark-border">
      <div className="container flex h-16 items-center justify-between px-4 md:px-6">
        <div className="flex items-center gap-4">
          <button 
            className="block md:hidden rounded-lg p-1.5 text-text-secondary hover:bg-gray-100 dark:hover:bg-dark-hover"
            onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
          >
            {mobileMenuOpen ? <X size={24} /> : <Menu size={24} />}
          </button>
          
          <div className="flex items-center">
            <span className="text-xl font-bold text-text dark:text-white">一个类似 airtable 的前端项目</span>
          </div>
        </div>

        {!searchOpen ? (
          <div className="hidden md:flex items-center gap-2">
            <Button 
              variant="ghost" 
              size="sm" 
              onClick={() => setSearchOpen(true)}
              className="text-text-secondary"
              leftIcon={<Search size={18} />}
            >
              Search...
            </Button>
          </div>
        ) : (
          <div className="hidden md:flex items-center flex-1 max-w-md mx-6">
            <Input
              placeholder="Search tasks, projects..."
              leftIcon={<Search size={18} />}
              rightIcon={
                <button onClick={() => setSearchOpen(false)}>
                  <X size={18} />
                </button>
              }
              className="w-full"
              autoFocus
              onBlur={() => setSearchOpen(false)}
              onKeyDown={(e) => e.key === 'Escape' && setSearchOpen(false)}
            />
          </div>
        )}

        <div className="flex items-center gap-2">
          <Button 
            variant="primary" 
            size="sm" 
            className="hidden md:flex" 
            leftIcon={<Plus size={16} />}
          >
            New Task
          </Button>
          
          <div className="border-l h-6 mx-2 border-gray-200 dark:border-dark-border"></div>
          
          <button className="p-1.5 rounded-lg text-text-secondary hover:bg-gray-100 dark:hover:bg-dark-hover">
            <HelpCircle size={20} />
          </button>
          
          <button className="p-1.5 rounded-lg text-text-secondary hover:bg-gray-100 dark:hover:bg-dark-hover">
            <Bell size={20} />
          </button>
          
          <button 
            className="p-1.5 rounded-lg text-text-secondary hover:bg-gray-100 dark:hover:bg-dark-hover"
            onClick={toggleTheme}
          >
            {theme === 'light' ? <Moon size={20} /> : <Sun size={20} />}
          </button>
          
          <button className="rounded-full overflow-hidden border-2 border-accent">
            <div className="w-8 h-8 flex items-center justify-center bg-gray-200 text-text-secondary dark:bg-dark-hover dark:text-gray-300">
              <User size={18} />
            </div>
          </button>
        </div>
      </div>
      
      {/* Mobile menu */}
      {mobileMenuOpen && (
        <div className="md:hidden border-t border-gray-200 bg-primary dark:bg-dark-bg dark:border-dark-border py-4 px-4 animate-slide-down">
          <div className="flex mb-4">
            <Input
              placeholder="Search tasks, projects..."
              leftIcon={<Search size={18} />}
              className="w-full"
            />
          </div>
          
          <div className="space-y-2">
            <Button 
              variant="primary" 
              className="w-full justify-center" 
              leftIcon={<Plus size={16} />}
            >
              New Task
            </Button>
            
            <div className="pt-2 border-t border-gray-200 dark:border-dark-border">
              <p className="text-sm font-medium text-text-secondary mb-2">Current Project:</p>
              <p className="text-text dark:text-white font-medium">
                {currentProject?.name || 'No project selected'}
              </p>
            </div>
          </div>
        </div>
      )}
    </header>
  );
}