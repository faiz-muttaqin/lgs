import { Search } from 'lucide-react';
import { Input } from '@/components/ui/input';
import { forwardRef } from 'react';

interface SearchInputProps {
  placeholder?: string;
  value?: string;
  onChange?: (value: string) => void;
}

export const SearchInput = forwardRef<HTMLInputElement, SearchInputProps>(
  ({ placeholder = 'Cari produk...', value, onChange }, ref) => {
    return (
      <div className="relative flex-1">
        <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
        <Input
          ref={ref}
          type="text"
          placeholder={placeholder}
          value={value}
          onChange={(e) => onChange?.(e.target.value)}
          className="pl-10 pr-4 h-10"
        />
      </div>
    );
  }
);

SearchInput.displayName = 'SearchInput';
