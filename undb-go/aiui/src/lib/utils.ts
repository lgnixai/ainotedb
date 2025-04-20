import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function formatDate(date: Date): string {
  return new Intl.DateTimeFormat('en-US', {
    month: 'short',
    day: 'numeric',
    year: 'numeric'
  }).format(date);
}

export function getRandomColor(): string {
  const colors = [
    'bg-accent-100',
    'bg-success-100',
    'bg-warning-100',
    'bg-error-100',
    'bg-purple-100',
    'bg-pink-100',
    'bg-indigo-100'
  ];
  return colors[Math.floor(Math.random() * colors.length)];
}

export function getRandomColorText(): string {
  const colors = [
    'text-accent-700',
    'text-success-700',
    'text-warning-700',
    'text-error-700',
    'text-purple-700',
    'text-pink-700',
    'text-indigo-700'
  ];
  return colors[Math.floor(Math.random() * colors.length)];
}

export function truncateText(text: string, maxLength: number): string {
  if (text.length <= maxLength) return text;
  return text.slice(0, maxLength) + '...';
}

export function generateId(): string {
  return Math.random().toString(36).substring(2, 9);
}