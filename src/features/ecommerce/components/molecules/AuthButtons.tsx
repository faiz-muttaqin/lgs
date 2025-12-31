import { Button } from '@/components/ui/button';
import { Link } from '@tanstack/react-router';

export const AuthButtons = () => {
  return (
    <div className="flex items-center gap-2">
      <a href="/login">
        <Button variant="ghost">
          Masuk
        </Button>
      </a>
      <Link to="/register" search={{}}>
        <Button>
          Daftar
        </Button>
      </Link>
    </div>
  );
};
