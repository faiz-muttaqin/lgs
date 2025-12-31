import { ChevronLeft, ChevronRight } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { cn } from '@/lib/utils';

interface ScrollButtonProps {
  direction: 'left' | 'right';
  onClick: () => void;
  className?: string;
  disabled?: boolean;
}

export const ScrollButton = ({ direction, onClick, className, disabled }: ScrollButtonProps) => {
  const Icon = direction === 'left' ? ChevronLeft : ChevronRight;

  return (
    <Button
      variant="outline"
      size="icon"
      onClick={onClick}
      disabled={disabled}
      className={cn(
        'h-10 w-10 rounded-full shadow-md',
        'disabled:opacity-50 disabled:cursor-not-allowed',
        className
      )}
    >
      <Icon className="h-5 w-5" />
    </Button>
  );
};
