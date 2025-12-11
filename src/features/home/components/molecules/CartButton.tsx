import { Button } from '@/components/ui/button';
import { CartIcon } from '../atoms/CartIcon';

interface CartButtonProps {
  count?: number;
  onClick?: () => void;
}

export const CartButton = ({ count = 0, onClick }: CartButtonProps) => {
  return (
    <Button
      variant="ghost"
      size="icon"
      onClick={onClick}
      className="relative"
    >
      <CartIcon count={count} />
    </Button>
  );
};
