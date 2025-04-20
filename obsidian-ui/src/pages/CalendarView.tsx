import { useState } from 'react';
import { ChevronLeft, ChevronRight } from 'lucide-react';

const CalendarView = () => {
  const [currentDate, setCurrentDate] = useState(new Date());
  
  const daysInMonth = new Date(
    currentDate.getFullYear(), 
    currentDate.getMonth() + 1, 
    0
  ).getDate();
  
  const firstDayOfMonth = new Date(
    currentDate.getFullYear(), 
    currentDate.getMonth(), 
    1
  ).getDay();

  const monthNames = [
    '一月', '二月', '三月', '四月', '五月', '六月',
    '七月', '八月', '九月', '十月', '十一月', '十二月'
  ];

  const dayNames = ['周日', '周一', '周二', '周三', '周四', '周五', '周六'];

  const previousMonth = () => {
    setCurrentDate(new Date(currentDate.getFullYear(), currentDate.getMonth() - 1, 1));
  };

  const nextMonth = () => {
    setCurrentDate(new Date(currentDate.getFullYear(), currentDate.getMonth() + 1, 1));
  };

  // Generate calendar days
  const calendarDays = [];
  
  // Add empty cells for days before the first day of the month
  for (let i = 0; i < firstDayOfMonth; i++) {
    calendarDays.push(null);
  }
  
  // Add days of the month
  for (let day = 1; day <= daysInMonth; day++) {
    calendarDays.push(day);
  }

  // Mock events (in a real app, these would come from your data)
  const events = [
    { day: 5, title: '团队会议', color: '#3B82F6' },
    { day: 12, title: '项目截止日期', color: '#F97316' },
    { day: 15, title: '产品发布', color: '#22C55E' },
    { day: 20, title: '客户演示', color: '#A855F7' },
    { day: 25, title: '月度回顾', color: '#3B82F6' },
  ];

  return (
    <div className="container mx-auto p-6 max-w-6xl">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">日历视图</h1>
        <div className="flex items-center space-x-2">
          <button
            className="p-2 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800"
            onClick={previousMonth}
          >
            <ChevronLeft size={20} />
          </button>
          <span className="text-lg font-medium">
            {monthNames[currentDate.getMonth()]} {currentDate.getFullYear()}
          </span>
          <button
            className="p-2 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800"
            onClick={nextMonth}
          >
            <ChevronRight size={20} />
          </button>
        </div>
      </div>

      <div className="card overflow-hidden p-0">
        {/* Calendar header */}
        <div className="grid grid-cols-7 border-b border-gray-200 dark:border-gray-700 bg-secondary dark:bg-secondary/50">
          {dayNames.map((day, index) => (
            <div 
              key={index} 
              className="py-2 text-center font-medium text-text-secondary text-sm"
            >
              {day}
            </div>
          ))}
        </div>

        {/* Calendar grid */}
        <div className="grid grid-cols-7 auto-rows-fr">
          {calendarDays.map((day, index) => {
            const dayEvents = day ? events.filter(event => event.day === day) : [];
            
            return (
              <div 
                key={index} 
                className={`min-h-[100px] p-2 border-b border-r border-gray-200 dark:border-gray-700 ${
                  day === null ? 'bg-gray-50 dark:bg-gray-800/30' : 'hover:bg-gray-50 dark:hover:bg-gray-800/10'
                }`}
              >
                {day !== null && (
                  <>
                    <div className={`text-right mb-1 ${
                      day === new Date().getDate() && 
                      currentDate.getMonth() === new Date().getMonth() && 
                      currentDate.getFullYear() === new Date().getFullYear()
                        ? 'bg-accent text-white w-6 h-6 rounded-full flex items-center justify-center ml-auto'
                        : ''
                    }`}>
                      {day}
                    </div>
                    <div className="space-y-1">
                      {dayEvents.map((event, eventIndex) => (
                        <div 
                          key={eventIndex}
                          className="text-xs p-1 rounded truncate cursor-pointer"
                          style={{ backgroundColor: `${event.color}20`, color: event.color }}
                        >
                          {event.title}
                        </div>
                      ))}
                    </div>
                  </>
                )}
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default CalendarView;