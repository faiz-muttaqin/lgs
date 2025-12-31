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

  // Handle both old mock data and new API data
  const imageUrl = (product as any).image_url || (product as any).image;
  const discount = (product as any).discount_pct || (product as any).discount;
  const slashedPrice = (product as any).slashed_price || (product as any).originalPrice;
  const location = (product as any).shop?.city || (product as any).location;
  const reviews = (product as any).count_review || (product as any).reviews;
  const badge = (product as any).is_featured ? 'Featured' : (product as any).badge;

  return (
    <a href={`/products/${product.id}`} className="block">
      <Card className="overflow-hidden hover:shadow-lg transition-shadow cursor-pointer min-w-[180px] max-w-[220px]">
        <div className="relative aspect-square bg-muted">
          <img
            src={imageUrl}
            alt={product.name}
            className="w-full h-full object-cover"
            onError={(e) => {
              e.currentTarget.src = 'https://via.placeholder.com/300x300?text=Product';
            }}
          />
          {discount && <DiscountBadge discount={discount} />}
          {badge && (
            <div className="absolute top-2 right-2 bg-yellow-400 text-xs font-semibold px-2 py-1 rounded">
              {badge}
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
            {slashedPrice && slashedPrice > product.price && (
              <span className="text-xs text-muted-foreground line-through">
                {formatPrice(slashedPrice)}
              </span>
            )}
            {product.rating && (
              <div className="flex items-center gap-1 text-xs text-muted-foreground">
                <Star className="h-3 w-3 fill-yellow-400 text-yellow-400" />
                <span>{product.rating.toFixed(1)}</span>
                {reviews && (
                  <span className="text-muted-foreground/70">({reviews})</span>
                )}
              </div>
            )}
            {location && (
              <div className="flex items-center gap-1 text-xs text-muted-foreground">
                <MapPin className="h-3 w-3" />
                <span>{location}</span>
              </div>
            )}
          </div>
        </CardContent>
      </Card>
    </a>
  );
};
