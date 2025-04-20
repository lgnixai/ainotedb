import React from 'react';
import { cn } from '../../lib/utils';

type BadgeVariant = 'default' | 'primary' | 'secondary' | 'success' | 'warning' | 'danger' | 'outline';

interface BadgeProps {
  variant?: BadgeVariant;
  className?: string;
  children: React.ReactNode;
}

const badgeVariants = {
  default: "bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-300",
  primary: "bg-accent-100 text-accent-800 dark:bg-accent-900 dark:text-accent-300",
  secondary: "bg-secondary text-text-secondary dark:bg-gray-700 dark:text-gray-300",
  success: "bg-success-100 text-success-800 dark:bg-success-900 dark:text-success-300",
  warning: "bg-warning-100 text-warning-800 dark:bg-warning-900 dark:text-warning-300",
  danger: "bg-error-100 text-error-800 dark:bg-error-900 dark:text-error-300",
  outline: "bg-transparent border border-gray-300 text-gray-700 dark:border-gray-600 dark:text-gray-300",
};

export function Badge({ variant = 'default', className, children }: BadgeProps) {
  return (
    <span
      className={cn(
        "inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium transition-colors",
        badgeVariants[variant],
        className
      )}
    >
      {children}
    </span>
  );
}