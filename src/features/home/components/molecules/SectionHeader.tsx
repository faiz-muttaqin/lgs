import { ChevronRight } from 'lucide-react';
import { Button } from '@/components/ui/button';

interface SectionHeaderProps {
  title: string;
  link?: string;
  showViewAll?: boolean;
}

export const SectionHeader = ({ title, link, showViewAll = true }: SectionHeaderProps) => {
  return (
    <div className="flex items-center justify-between mb-4">
      <h2 className="text-xl md:text-2xl font-bold text-foreground">{title}</h2>
      {showViewAll && link && (
        <a href={link}>
          <Button variant="ghost" className="text-primary">
            Lihat Semua
            <ChevronRight className="h-4 w-4 ml-1" />
          </Button>
        </a>
      )}
    </div>
  );
};
