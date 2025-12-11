import { Card, CardContent } from '@/components/ui/card';
import { DiscountBadge } from '../atoms/DiscountBadge';
import type { Product } from '../../types';
import { Star, MapPin } from 'lucide-react';

interface ProductItemProps {
  product: Product;
}

export const ProductItem = ({ product }: ProductItemProps) => {
  const formatPrice = (price: number) => {
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0,
    }).format(price);
  };

  return (
    <a href={`/product/${product.id}`} className="block">
      <Card className="overflow-hidden hover:shadow-lg transition-shadow cursor-pointer min-w-[180px] max-w-[220px]">
        <div className="relative aspect-square bg-muted">
          <img
            src={product.image}
            alt={product.name}
            className="w-full h-full object-cover"
            onError={(e) => {
              e.currentTarget.src = 'https://via.placeholder.com/300x300?text=Product';
            }}
          />
          {product.discount && <DiscountBadge discount={product.discount} />}
          {product.badge && (
            <div className="absolute top-2 right-2 bg-yellow-400 text-xs font-semibold px-2 py-1 rounded">
              {product.badge}
            </div>
          )}
        </div>
        <CardContent className="p-3">
          <h3 className="font-medium text-sm text-foreground line-clamp-2 mb-1 h-10">
            {product.name}
          </h3>
          <div className="flex flex-col gap-1">
            <div className="flex items-center gap-1">
              <span className="text-lg font-bold text-primary">
                {formatPrice(product.price)}
              </span>
            </div>
            {product.originalPrice && (
              <span className="text-xs text-muted-foreground line-through">
                {formatPrice(product.originalPrice)}
              </span>
            )}
            {product.rating && (
              <div className="flex items-center gap-1 text-xs text-muted-foreground">
                <Star className="h-3 w-3 fill-yellow-400 text-yellow-400" />
                <span>{product.rating}</span>
                {product.reviews && (
                  <span className="text-muted-foreground/70">({product.reviews})</span>
                )}
              </div>
            )}
            {product.location && (
              <div className="flex items-center gap-1 text-xs text-muted-foreground">
                <MapPin className="h-3 w-3" />
                <span>{product.location}</span>
              </div>
            )}
          </div>
        </CardContent>
      </Card>
    </a>
  );
};
