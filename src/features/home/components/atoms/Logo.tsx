import { Link } from '@tanstack/react-router';
import { ShoppingBag } from 'lucide-react';

export const Logo = () => {
  return (
    <Link to="/" className="flex items-center gap-2 hover:opacity-80 transition-opacity">
      <ShoppingBag className="h-8 w-8 text-primary" />
      <div className="flex flex-col leading-none">
        <span className="text-xl font-bold text-foreground">LGS</span>
        <span className="text-xs text-muted-foreground">Lucky Good Store</span>
      </div>
    </Link>
  );
};
