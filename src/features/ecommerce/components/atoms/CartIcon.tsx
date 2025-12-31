import { ShoppingCart } from 'lucide-react';
import { Badge } from '@/components/ui/badge';

interface CartIconProps {
  count?: number;
}

export const CartIcon = ({ count = 0 }: CartIconProps) => {
  return (
    <div className="relative">
      <ShoppingCart className="h-6 w-6 text-foreground" />
      {count > 0 && (
        <Badge
          variant="destructive"
          className="absolute -top-2 -right-2 h-5 w-5 flex items-center justify-center p-0 text-xs"
        >
          {count > 9 ? '9+' : count}
        </Badge>
      )}
    </div>
  );
};
