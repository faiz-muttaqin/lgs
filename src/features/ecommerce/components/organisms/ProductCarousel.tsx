import { useRef, useState, useEffect } from 'react';
import { ProductItem } from '../molecules/ProductItem';
import { ScrollButton } from '../atoms/ScrollButton';
import type { Product } from '../../types';

interface ProductCarouselProps {
  products: Product[];
}

export const ProductCarousel = ({ products }: ProductCarouselProps) => {
  const scrollRef = useRef<HTMLDivElement>(null);
  const [canScrollLeft, setCanScrollLeft] = useState(false);
  const [canScrollRight, setCanScrollRight] = useState(true);

  const checkScroll = () => {
    if (scrollRef.current) {
      const { scrollLeft, scrollWidth, clientWidth } = scrollRef.current;
      setCanScrollLeft(scrollLeft > 0);
      setCanScrollRight(scrollLeft < scrollWidth - clientWidth - 10);
    }
  };

  useEffect(() => {
    checkScroll();
    const scrollElement = scrollRef.current;
    scrollElement?.addEventListener('scroll', checkScroll);
    return () => scrollElement?.removeEventListener('scroll', checkScroll);
  }, []);

  const scroll = (direction: 'left' | 'right') => {
    if (scrollRef.current) {
      const scrollAmount = 400;
      scrollRef.current.scrollBy({
        left: direction === 'left' ? -scrollAmount : scrollAmount,
        behavior: 'smooth',
      });
    }
  };

  return (
    <div className="relative group">
      <div
        ref={scrollRef}
        className="flex gap-4 overflow-x-auto scrollbar-hide scroll-smooth pb-2"
      >
        {products.map((product) => (
          <ProductItem key={product.id} product={product} />
        ))}
      </div>

      {/* Scroll Buttons */}
      {canScrollLeft && (
        <div className="absolute left-0 top-1/2 -translate-y-1/2 -translate-x-2 opacity-0 group-hover:opacity-100 transition-opacity hidden md:block">
          <ScrollButton direction="left" onClick={() => scroll('left')} />
        </div>
      )}
      {canScrollRight && (
        <div className="absolute right-0 top-1/2 -translate-y-1/2 translate-x-2 opacity-0 group-hover:opacity-100 transition-opacity hidden md:block">
          <ScrollButton direction="right" onClick={() => scroll('right')} />
        </div>
      )}
    </div>
  );
};
