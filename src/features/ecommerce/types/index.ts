// Backend API Product Types
export interface Product {
  id: number;
  external_id?: string;
  sku?: string;
  name: string;
  slug: string;
  subtitle?: string;
  description?: string;
  image_url: string;
  price: number;
  slashed_price?: number;
  discount_pct?: number;
  stock: number;
  rating: number;
  count_review: number;
  count_sold: number;
  weight?: number;
  is_active: boolean;
  is_featured: boolean;
  status: string;
  category_id?: number;
  sub_category_id?: number;
  shop_id: number;
  created_at: string;
  updated_at: string;
  
  // Relations (optional, loaded with preload)
  category?: Category;
  sub_category?: SubCategory;
  shop?: Shop;
  images?: ProductImage[];
  labels?: ProductLabel[];
  badges?: ProductBadge[];
  variants?: ProductVariant[];
}

export interface ProductDetail extends Product {
  // Always includes relations for detail view
  category: Category;
  sub_category?: SubCategory;
  shop: Shop;
  images: ProductImage[];
  labels: ProductLabel[];
  badges: ProductBadge[];
  variants: ProductVariant[];
}

export interface ProductImage {
  id: number;
  product_id: number;
  image_url: string;
  position: number;
  is_main: boolean;
  created_at: string;
  updated_at: string;
}

export interface ProductLabel {
  id: number;
  product_id: number;
  title: string;
  type?: string;
  position?: string;
  url?: string;
  created_at: string;
  updated_at: string;
}

export interface ProductBadge {
  id: number;
  product_id: number;
  title: string;
  image_url?: string;
  created_at: string;
  updated_at: string;
}

export interface ProductVariant {
  id: number;
  product_id: number;
  sku?: string;
  name: string;
  price?: number;
  stock: number;
  weight?: number;
  image_url?: string;
  is_available: boolean;
  created_at: string;
  updated_at: string;
}

export interface Category {
  id: number;
  name: string;
  slug: string;
  description?: string;
  image_url?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
  sub_categories?: SubCategory[];
}

export interface SubCategory {
  id: number;
  category_id: number;
  name: string;
  slug: string;
  description?: string;
  image_url?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
  category?: Category;
}

export interface Shop {
  id: number;
  user_id: number;
  name: string;
  slug: string;
  domain?: string;
  city?: string;
  image_url?: string;
  reputation?: string;
  is_official: boolean;
  is_active: boolean;
  description?: string;
  created_at: string;
  updated_at: string;
}

// Filter and Query Types
// Operator types based on API schema
export type FilterOperator = 
  | 'eq' | 'ne' | 'gt' | 'gte' | 'lt' | 'lte' 
  | 'contains' | 'starts_with' | 'ends_with' 
  | 'in' | 'between' | 'is_null';

// Generic filter value type
export type FilterValue = string | number | boolean | string[] | number[] | [number, number];

// Filter with operator support
export interface FilterWithOperator {
  operator: FilterOperator;
  value: FilterValue;
}

export interface ProductFilters {
  // Simple search (will be converted to name[contains])
  search?: string;
  
  // Direct field filters with operators
  name?: string | FilterWithOperator;
  sku?: string | FilterWithOperator;
  description?: string | FilterWithOperator;
  category_id?: number | FilterWithOperator;
  sub_category_id?: number | FilterWithOperator;
  shop_id?: number | FilterWithOperator;
  price?: number | FilterWithOperator;
  min_price?: number; // Convenience for price[gte]
  max_price?: number; // Convenience for price[lte]
  stock?: number | FilterWithOperator;
  status?: string | FilterWithOperator;
  is_active?: boolean;
  is_featured?: boolean;
  discount_pct?: number | FilterWithOperator;
  rating?: number | FilterWithOperator;
  
  // Sorting
  sort_by?: 'created_at' | 'price' | 'rating' | 'count_sold' | 'name' | 'stock' | 'status' | 'id';
  sort_order?: 'asc' | 'desc';
  
  // Pagination
  page?: number;
  page_size?: number;
}

export interface ProductListResponse {
  data: Product[];
  meta: {
    page: number;
    page_size: number;
    total: number;
    total_page: number;
  };
}

// Legacy types for backward compatibility with mock data
export interface Section {
  id: string;
  title: string;
  type: 'products' | 'banner' | 'flash-sale' | 'voucher';
  products?: Product[];
  bannerImage?: string;
  link?: string;
}

export interface Banner {
  id: string;
  image: string;
  title?: string;
  subtitle?: string;
  link?: string;
}
