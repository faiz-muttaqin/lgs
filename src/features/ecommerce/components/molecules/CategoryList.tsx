import { CategoryBadge } from '../atoms/CategoryBadge';
import type { Category } from '../../types';
import { useRef } from 'react';
import { ScrollButton } from '../atoms/ScrollButton';

interface CategoryListProps {
  categories: Category[];
  activeCategory?: string | number;
  onCategoryClick?: (categoryId: number) => void;
}

export const CategoryList = ({ categories, activeCategory, onCategoryClick }: CategoryListProps) => {
  const scrollRef = useRef<HTMLDivElement>(null);

  const scroll = (direction: 'left' | 'right') => {
    if (scrollRef.current) {
      const scrollAmount = 300;
      scrollRef.current.scrollBy({
        left: direction === 'left' ? -scrollAmount : scrollAmount,
        behavior: 'smooth',
      });
    }
  };

  return (
    <div className="relative flex items-center gap-2">
      <ScrollButton
        direction="left"
        onClick={() => scroll('left')}
        className="hidden md:flex"
      />
      <div
        ref={scrollRef}
        className="flex gap-2 overflow-x-auto scrollbar-hide scroll-smooth flex-1"
      >
        {categories.map((category) => (
          <CategoryBadge
            key={category.id}
            name={category.name}
            isActive={activeCategory === category.id || activeCategory === String(category.id)}
            onClick={() => onCategoryClick?.(category.id)}
          />
        ))}
      </div>
      <ScrollButton
        direction="right"
        onClick={() => scroll('right')}
        className="hidden md:flex"
      />
    </div>
  );
};
