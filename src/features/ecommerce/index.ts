export { HomeTemplate } from './components/templates/HomeTemplate';
export { ProductsListTemplate } from './components/templates/ProductsListTemplate';
export { ProductDetailTemplate } from './components/templates/ProductDetailTemplate';
export { Header } from './components/organisms/Header';
export { Footer } from './components/organisms/Footer';
export { BannerCarousel } from './components/organisms/BannerCarousel';
export { ProductSection } from './components/organisms/ProductSection';
export { ProductCarousel } from './components/organisms/ProductCarousel';

// Types
export type { Product, ProductDetail, ProductFilters, ProductListResponse, Category, SubCategory, Shop, Section, Banner } from './types';

// API
export { productsApi } from './api/products';
export { categoriesApi } from './api/categories';

// Data
export { mockProducts, categories, sections, mainBanners } from './data/mockData';
