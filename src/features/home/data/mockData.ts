import type { Banner, Category, Product, Section } from '../types';

export const categories: Category[] = [
  { id: '1', name: 'Fashion Pria' },
  { id: '2', name: 'Fashion Wanita' },
  { id: '3', name: 'Elektronik' },
  { id: '4', name: 'Handphone & Aksesoris' },
  { id: '5', name: 'Komputer & Laptop' },
  { id: '6', name: 'Rumah Tangga' },
  { id: '7', name: 'Olahraga' },
  { id: '8', name: 'Kecantikan' },
];

export const mainBanners: Banner[] = [
  {
    id: '1',
    image: '/images/banner-1.jpg',
    title: 'Promo 12.12',
    subtitle: 'Diskon s.d 1,2 juta',
    link: '/promo',
  },
];

export const mockProducts: Product[] = [
  {
    id: '1',
    name: 'Sepatu Sneakers Pria',
    price: 299000,
    originalPrice: 599000,
    discount: 50,
    image: '/images/product-1.jpg',
    rating: 4.5,
    reviews: 120,
    badge: 'Best Seller',
  },
  {
    id: '2',
    name: 'Tas Ransel Premium',
    price: 199000,
    originalPrice: 399000,
    discount: 50,
    image: '/images/product-2.jpg',
    rating: 4.8,
    reviews: 89,
  },
  {
    id: '3',
    name: 'Jam Tangan Digital',
    price: 150000,
    originalPrice: 300000,
    discount: 50,
    image: '/images/product-3.jpg',
    rating: 4.3,
    reviews: 56,
  },
  {
    id: '4',
    name: 'Kaos Polos Premium',
    price: 89000,
    originalPrice: 150000,
    discount: 40,
    image: '/images/product-4.jpg',
    rating: 4.6,
    reviews: 200,
  },
  {
    id: '5',
    name: 'Celana Jeans Slim Fit',
    price: 249000,
    originalPrice: 450000,
    discount: 45,
    image: '/images/product-5.jpg',
    rating: 4.7,
    reviews: 150,
  },
  {
    id: '6',
    name: 'Jaket Hoodie',
    price: 179000,
    originalPrice: 350000,
    discount: 49,
    image: '/images/product-6.jpg',
    rating: 4.5,
    reviews: 95,
  },
];

export const sections: Section[] = [
  {
    id: '1',
    title: 'Flash Sale Hari Ini',
    type: 'flash-sale',
    products: mockProducts,
  },
  {
    id: '2',
    title: 'Pilihan Brand Top',
    type: 'products',
    products: mockProducts,
  },
  {
    id: '3',
    title: 'Diskon Elektronik',
    type: 'products',
    products: mockProducts,
  },
  {
    id: '4',
    title: 'Fashion Terlaris',
    type: 'products',
    products: mockProducts,
  },
  {
    id: '5',
    title: 'Peralatan Rumah',
    type: 'products',
    products: mockProducts,
  },
];
