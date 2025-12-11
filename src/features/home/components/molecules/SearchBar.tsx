import { SearchInput } from '../atoms/SearchInput';
import { Button } from '@/components/ui/button';
import { Search } from 'lucide-react';

export const SearchBar = () => {
  const handleSearch = () => {
    // Search logic here
  };

  return (
    <div className="flex gap-2 w-full max-w-2xl">
      <SearchInput placeholder="Cari di LGS..." />
      <Button onClick={handleSearch}>
        <Search className="h-4 w-4" />
      </Button>
    </div>
  );
};
