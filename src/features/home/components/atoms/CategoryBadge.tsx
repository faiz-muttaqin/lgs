import { cn } from '@/lib/utils';

interface CategoryBadgeProps {
  name: string;
  isActive?: boolean;
  onClick?: () => void;
}

export const CategoryBadge = ({ name, isActive, onClick }: CategoryBadgeProps) => {
  return (
    <button
      onClick={onClick}
      className={cn(
        'px-4 py-2 rounded-full text-sm font-medium whitespace-nowrap transition-colors',
        isActive
          ? 'bg-primary text-primary-foreground'
          : 'bg-secondary text-secondary-foreground hover:bg-secondary/80'
      )}
    >
      {name}
    </button>
  );
};
