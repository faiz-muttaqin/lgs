import { Badge } from '@/components/ui/badge';

interface DiscountBadgeProps {
  discount: number;
}

export const DiscountBadge = ({ discount }: DiscountBadgeProps) => {
  return (
    <Badge variant="destructive" className="absolute top-2 left-2 font-semibold">
      {discount}%
    </Badge>
  );
};
