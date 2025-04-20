import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { login } from '../../lib/api';
import { Input } from '../ui/Input';
import { Button } from '../ui/Button';

export default function Login() {
  const [form, setForm] = useState({ email: '', password: '' });
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    try {
      const res = await login(form.email, form.password);
      localStorage.setItem('token', res.token);
      navigate('/');
    } catch (err: any) {
      setError(err.message || '登录失败');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <form className="w-full max-w-sm p-8 bg-white rounded shadow" onSubmit={handleSubmit}>
        <h2 className="text-2xl font-bold mb-6 text-center">登录</h2>
        <Input name="email" placeholder="邮箱" value={form.email} onChange={handleChange} required />
        <Input name="password" type="password" placeholder="密码" value={form.password} onChange={handleChange} required className="mt-4" />
        {error && <p className="text-red-500 mt-2">{error}</p>}
        <Button type="submit" className="w-full mt-6" isLoading={loading}>登录</Button>
        <div className="text-center mt-4">
          没有账号？ <a className="text-blue-600 underline" href="/register">注册</a>
        </div>
      </form>
    </div>
  );
}
