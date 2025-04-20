import React, { useState, useEffect } from 'react';
import { useProject } from '../../context/ProjectContext';
import { cn } from '../../lib/utils';
import { ChevronLeft, ChevronRight } from 'lucide-react';
import { Task } from '../../types';

export function CalendarView() {
  const { currentProject } = useProject();
  const [currentDate, setCurrentDate] = useState(new Date());
  const [calendarDays, setCalendarDays] = useState<Date[]>([]);
  
  if (!currentProject) return null;
  
  const tasks = currentProject.tasks;

  useEffect(() => {
    generateCalendarDays(currentDate);
  }, [currentDate]);

  const generateCalendarDays = (date: Date) => {
    const year = date.getFullYear();
    const month = date.getMonth();
    
    // Get the first day of the month
    const firstDay = new Date(year, month, 1);
    // Get the last day of the month
    const lastDay = new Date(year, month + 1, 0);
    
    // Get the day of the week of the first day (0 = Sunday, 1 = Monday, etc.)
    const firstDayOfWeek = firstDay.getDay();
    
    // Calculate how many days from the previous month we need to show
    const daysFromPrevMonth = firstDayOfWeek === 0 ? 6 : firstDayOfWeek - 1;
    
    // Get the first day to show in the calendar
    const firstDayToShow = new Date(year, month, 1 - daysFromPrevMonth);
    
    // Generate an array of all days to show in the calendar (usually 6 weeks = 42 days)
    const days = [];
    for (let i = 0; i < 42; i++) {
      const day = new Date(firstDayToShow);
      day.setDate(firstDayToShow.getDate() + i);
      days.push(day);
    }
    
    setCalendarDays(days);
  };

  const getTasksForDay = (day: Date) => {
    return tasks.filter(task => {
      if (!task.dueDate) return false;
      
      const dueDate = new Date(task.dueDate);
      return (
        dueDate.getFullYear() === day.getFullYear() &&
        dueDate.getMonth() === day.getMonth() &&
        dueDate.getDate() === day.getDate()
      );
    });
  };

  const prevMonth = () => {
    const newDate = new Date(currentDate);
    newDate.setMonth(newDate.getMonth() - 1);
    setCurrentDate(newDate);
  };

  const nextMonth = () => {
    const newDate = new Date(currentDate);
    newDate.setMonth(newDate.getMonth() + 1);
    setCurrentDate(newDate);
  };

  const isToday = (day: Date) => {
    const today = new Date();
    return (
      day.getFullYear() === today.getFullYear() &&
      day.getMonth() === today.getMonth() &&
      day.getDate() === today.getDate()
    );
  };

  const isCurrentMonth = (day: Date) => {
    return day.getMonth() === currentDate.getMonth();
  };

  const renderTaskInCalendar = (task: Task) => {
    let bgColor = 'bg-accent-100 text-accent-800';
    
    if (task.priority === 'High' || task.priority === 'Urgent') {
      bgColor = 'bg-error-100 text-error-800';
    } else if (task.status === 'Completed') {
      bgColor = 'bg-success-100 text-success-800';
    }
    
    return (
      <div 
        key={task.id} 
        className={cn(
          "px-1 py-0.5 text-xs rounded truncate mb-1 cursor-pointer",
          bgColor
        )}
        title={task.title}
      >
        {task.title}
      </div>
    );
  };

  const dayNames = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
  const monthNames = [
    'January', 'February', 'March', 'April', 'May', 'June',
    'July', 'August', 'September', 'October', 'November', 'December'
  ];

  return (
    <div className="p-4">
      <div className="flex justify-between items-center mb-4">
        <h2 className="text-xl font-bold text-text dark:text-white">
          {monthNames[currentDate.getMonth()]} {currentDate.getFullYear()}
        </h2>
        
        <div className="flex space-x-2">
          <button
            onClick={prevMonth}
            className="p-1 rounded-md hover:bg-gray-100 dark:hover:bg-dark-hover"
          >
            <ChevronLeft size={20} className="text-text-secondary" />
          </button>
          
          <button
            onClick={nextMonth}
            className="p-1 rounded-md hover:bg-gray-100 dark:hover:bg-dark-hover"
          >
            <ChevronRight size={20} className="text-text-secondary" />
          </button>
        </div>
      </div>
      
      <div className="grid grid-cols-7 gap-px bg-gray-200 dark:bg-dark-border rounded-t-lg overflow-hidden">
        {dayNames.map(day => (
          <div 
            key={day} 
            className="bg-secondary dark:bg-dark-card p-2 text-center text-sm font-medium text-text-secondary dark:text-gray-400"
          >
            {day}
          </div>
        ))}
      </div>
      
      <div className="grid grid-cols-7 gap-px bg-gray-200 dark:bg-dark-border rounded-b-lg overflow-hidden">
        {calendarDays.map((day, index) => {
          const tasksForDay = getTasksForDay(day);
          
          return (
            <div 
              key={index}
              className={cn(
                "bg-white dark:bg-dark-bg min-h-24 p-1",
                !isCurrentMonth(day) && "opacity-40"
              )}
            >
              <div className="flex justify-between items-center mb-1">
                <span 
                  className={cn(
                    "text-sm font-medium p-1 flex items-center justify-center w-6 h-6",
                    isToday(day) && "bg-accent text-white rounded-full",
                    !isToday(day) && "text-text-secondary dark:text-gray-400"
                  )}
                >
                  {day.getDate()}
                </span>
              </div>
              
              <div className="overflow-y-auto max-h-20">
                {tasksForDay.slice(0, 3).map(task => renderTaskInCalendar(task))}
                
                {tasksForDay.length > 3 && (
                  <div className="text-xs text-text-secondary dark:text-gray-400 px-1">
                    + {tasksForDay.length - 3} more
                  </div>
                )}
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
}