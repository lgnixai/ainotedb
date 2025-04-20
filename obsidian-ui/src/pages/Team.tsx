import { Users, UserPlus, Mail, Phone, CheckCircle, Clock, Search } from 'lucide-react';

const Team = () => {
  // Mock team members data
  const members = [
    { 
      id: 1, 
      name: '张三', 
      role: '项目经理', 
      email: 'zhangsan@example.com', 
      phone: '13888881111',
      avatar: 'https://i.pravatar.cc/150?img=11',
      status: 'active'
    },
    { 
      id: 2, 
      name: '李四', 
      role: '开发工程师', 
      email: 'lisi@example.com', 
      phone: '13888882222',
      avatar: 'https://i.pravatar.cc/150?img=12',
      status: 'active'
    },
    { 
      id: 3, 
      name: '王五', 
      role: '设计师', 
      email: 'wangwu@example.com', 
      phone: '13888883333',
      avatar: 'https://i.pravatar.cc/150?img=13',
      status: 'active'
    },
    { 
      id: 4, 
      name: '赵六', 
      role: '产品经理', 
      email: 'zhaoliu@example.com', 
      phone: '13888884444',
      avatar: 'https://i.pravatar.cc/150?img=14',
      status: 'away'
    },
    { 
      id: 5, 
      name: '孙七', 
      role: '测试工程师', 
      email: 'sunqi@example.com', 
      phone: '13888885555',
      avatar: 'https://i.pravatar.cc/150?img=15',
      status: 'inactive'
    },
  ];

  // Mock pending invitations
  const pendingInvitations = [
    { email: 'new1@example.com', role: '开发工程师', sentAt: '2023-06-10T14:45:00Z' },
    { email: 'new2@example.com', role: '内容编辑', sentAt: '2023-06-11T09:30:00Z' },
  ];

  return (
    <div className="container mx-auto p-6 max-w-6xl">
      <div className="flex items-center justify-between mb-8">
        <h1 className="text-2xl font-bold">团队</h1>
        <div className="flex space-x-3">
          <div className="relative max-w-md">
            <input
              type="text"
              placeholder="搜索成员..."
              className="input py-1.5 pl-10"
            />
            <Search size={18} className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500" />
          </div>
          <button className="btn btn-primary">
            <UserPlus size={18} className="mr-1" />
            邀请成员
          </button>
        </div>
      </div>

      {/* Active members */}
      <div className="mb-10">
        <h2 className="text-xl font-semibold mb-4">团队成员</h2>
        <div className="card p-0 overflow-hidden">
          <div className="overflow-x-auto">
            <table className="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
              <thead className="bg-secondary dark:bg-secondary/50">
                <tr>
                  <th className="px-6 py-3 text-left text-xs font-medium text-text-secondary uppercase tracking-wider">
                    成员
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-text-secondary uppercase tracking-wider">
                    联系方式
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-text-secondary uppercase tracking-wider">
                    状态
                  </th>
                  <th className="px-6 py-3 text-left text-xs font-medium text-text-secondary uppercase tracking-wider">
                    操作
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white dark:bg-primary divide-y divide-gray-200 dark:divide-gray-700">
                {members.map(member => (
                  <tr key={member.id} className="hover:bg-gray-50 dark:hover:bg-gray-800">
                    <td className="px-6 py-4 whitespace-nowrap">
                      <div className="flex items-center">
                        <div className="h-10 w-10 rounded-full overflow-hidden mr-3">
                          <img src={member.avatar} alt={member.name} className="h-full w-full object-cover" />
                        </div>
                        <div>
                          <div className="font-medium">{member.name}</div>
                          <div className="text-sm text-text-secondary">{member.role}</div>
                        </div>
                      </div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      <div className="flex flex-col">
                        <div className="flex items-center text-sm text-text-secondary">
                          <Mail size={14} className="mr-1" />
                          {member.email}
                        </div>
                        <div className="flex items-center text-sm text-text-secondary mt-1">
                          <Phone size={14} className="mr-1" />
                          {member.phone}
                        </div>
                      </div>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap">
                      <span className={`px-2 py-1 text-xs rounded-full ${
                        member.status === 'active' 
                          ? 'bg-success/10 text-success' 
                          : member.status === 'away' 
                            ? 'bg-warning/10 text-warning' 
                            : 'bg-gray-100 dark:bg-gray-700 text-text-secondary'
                      }`}>
                        {member.status === 'active' ? '在线' : member.status === 'away' ? '离开' : '离线'}
                      </span>
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm">
                      <button className="text-accent hover:text-accent/80 mr-3">
                        编辑
                      </button>
                      <button className="text-error hover:text-error/80">
                        移除
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>

      {/* Pending invitations */}
      <div>
        <h2 className="text-xl font-semibold mb-4">待处理邀请</h2>
        {pendingInvitations.length > 0 ? (
          <div className="card p-0 overflow-hidden">
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                <thead className="bg-secondary dark:bg-secondary/50">
                  <tr>
                    <th className="px-6 py-3 text-left text-xs font-medium text-text-secondary uppercase tracking-wider">
                      邮箱
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-text-secondary uppercase tracking-wider">
                      职位
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-text-secondary uppercase tracking-wider">
                      发送时间
                    </th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-text-secondary uppercase tracking-wider">
                      操作
                    </th>
                  </tr>
                </thead>
                <tbody className="bg-white dark:bg-primary divide-y divide-gray-200 dark:divide-gray-700">
                  {pendingInvitations.map((invitation, index) => (
                    <tr key={index} className="hover:bg-gray-50 dark:hover:bg-gray-800">
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="flex items-center">
                          <Mail size={16} className="mr-2 text-text-secondary" />
                          <span>{invitation.email}</span>
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        {invitation.role}
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap">
                        <div className="flex items-center text-sm text-text-secondary">
                          <Clock size={14} className="mr-1" />
                          {new Date(invitation.sentAt).toLocaleString('zh-CN')}
                        </div>
                      </td>
                      <td className="px-6 py-4 whitespace-nowrap text-sm">
                        <button className="text-accent hover:text-accent/80 mr-3">
                          重发邀请
                        </button>
                        <button className="text-error hover:text-error/80">
                          取消
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        ) : (
          <div className="card p-6 text-center text-text-secondary">
            目前没有待处理的邀请
          </div>
        )}
      </div>
    </div>
  );
};

export default Team;