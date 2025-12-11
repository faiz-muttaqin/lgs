export interface Product {
  id: string;
  name: string;
  price: number;
  originalPrice?: number;
  discount?: number;
  image: string;
  rating?: number;
  reviews?: number;
  badge?: string;
  location?: string;
}

export interface Category {
  id: string;
  name: string;
  icon?: string;
  image?: string;
}

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
