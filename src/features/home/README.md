# Home Feature - E-commerce Front Page

This feature implements a complete e-commerce front page using **Atomic Design** principles with **shadcn/ui** components.

## 📁 Folder Structure

```
src/features/home/
├── components/
│   ├── atoms/              # Basic building blocks
│   │   ├── Logo.tsx
│   │   ├── CategoryBadge.tsx
│   │   ├── SearchInput.tsx
│   │   ├── CartIcon.tsx
│   │   ├── ScrollButton.tsx
│   │   └── DiscountBadge.tsx
│   ├── molecules/          # Combinations of atoms
│   │   ├── SearchBar.tsx
│   │   ├── CartButton.tsx
│   │   ├── AuthButtons.tsx
│   │   ├── ProductItem.tsx
│   │   ├── SectionHeader.tsx
│   │   └── CategoryList.tsx
│   ├── organisms/          # Complex UI sections
│   │   ├── Header.tsx
│   │   ├── BannerCarousel.tsx
│   │   ├── ProductCarousel.tsx
│   │   ├── ProductSection.tsx
│   │   └── Footer.tsx
│   └── templates/          # Page templates
│       └── HomeTemplate.tsx
├── data/
│   └── mockData.ts         # Mock data for products, categories, etc.
├── types/
│   └── index.ts            # TypeScript type definitions
├── styles/
│   └── home.css            # Custom styles for the home page
└── index.ts                # Public exports
```

## 🏗️ Atomic Design Architecture

### Atoms
Smallest, reusable UI components:
- **Logo**: Brand logo with icon and text
- **CategoryBadge**: Category pill/badge button
- **SearchInput**: Search text input with icon
- **CartIcon**: Shopping cart icon with item count badge
- **ScrollButton**: Left/right navigation button for carousels
- **DiscountBadge**: Discount percentage badge

### Molecules
Combinations of atoms forming meaningful UI units:
- **SearchBar**: Search input with button
- **CartButton**: Cart icon in a button with click handler
- **AuthButtons**: Login and Register buttons
- **ProductItem**: Product card with image, name, price, discount
- **SectionHeader**: Section title with "View All" link
- **CategoryList**: Horizontal scrollable category list

### Organisms
Complex, self-contained sections:
- **Header**: Fixed top navigation with logo, search, cart, auth buttons, and categories
- **BannerCarousel**: Auto-rotating promotional banner with navigation
- **ProductCarousel**: Horizontal scrollable product list
- **ProductSection**: Complete section with header and product carousel
- **Footer**: Full footer with links, payment methods, social media

### Templates
Page-level layouts:
- **HomeTemplate**: Complete homepage layout combining all organisms

## 🎨 Features

### Header
- ✅ Fixed top position
- ✅ Logo with brand identity
- ✅ Search bar (responsive)
- ✅ Cart with item count
- ✅ Login/Register buttons
- ✅ Horizontal scrollable category list
- ✅ Promotional top banner

### Banner
- ✅ Auto-rotating carousel (5s interval)
- ✅ Manual navigation (prev/next buttons)
- ✅ Dot indicators
- ✅ Responsive aspect ratio

### Product Sections
- ✅ Multiple product sections (Flash Sale, Top Brands, etc.)
- ✅ Horizontal scroll with lazy loading support
- ✅ Hover-to-show scroll buttons
- ✅ Product cards with:
  - Product image
  - Name (2-line clamp)
  - Price with discount
  - Rating and reviews
  - Badges (discount, best seller)

### Footer
- ✅ Brand information
- ✅ Social media links
- ✅ Customer service links
- ✅ About company links
- ✅ Payment method icons
- ✅ Shipping partner icons
- ✅ Copyright and legal links

## 🎯 Design Matches

Based on the provided design screenshot:
- ✅ Green (#16a34a) primary color theme
- ✅ Fixed header with promotional banner
- ✅ Logo | Categories | Search | Cart | Login/Register layout
- ✅ Main promotional banner carousel
- ✅ Multiple product sections
- ✅ Horizontal scrollable product carousels
- ✅ Discount badges and pricing
- ✅ Comprehensive footer with payment/shipping info

## 🔧 Usage

### Import and Use
```tsx
import { HomeTemplate } from '@/features/home';

function HomePage() {
  return <HomeTemplate />;
}
```

### Individual Components
```tsx
import { 
  Header, 
  Footer, 
  ProductSection,
  BannerCarousel 
} from '@/features/home';
```

## 📝 Mock Data

Mock data is provided in `data/mockData.ts`:
- `categories` - Product categories
- `mainBanners` - Hero banners
- `mockProducts` - Sample products
- `sections` - Product sections

## 🎨 Styling

- Uses **Tailwind CSS** for styling
- Uses **shadcn/ui** components (Button, Card, Badge, etc.)
- Custom CSS in `styles/home.css` for:
  - Horizontal scroll hiding
  - Smooth scrolling
  - Image optimization

## 📱 Responsive Design

- Mobile-first approach
- Breakpoints:
  - Mobile: Default
  - Tablet: `md:` (768px+)
  - Desktop: Larger screens

## 🚀 Performance Features

- ✅ Lazy horizontal scrolling
- ✅ Optimized images with error fallbacks
- ✅ Vertical page lazy scroll ready
- ✅ Component code splitting ready
- ✅ Smooth scroll behavior

## 🔄 Future Enhancements

- [ ] Infinite scroll for product sections
- [ ] Virtual scrolling for large product lists
- [ ] Image lazy loading with placeholder
- [ ] Add to cart functionality
- [ ] Product quick view
- [ ] Search autocomplete
- [ ] Category filters
- [ ] Wishlist functionality

## 📦 Dependencies

- React
- TanStack Router
- shadcn/ui components
- lucide-react (icons)
- Tailwind CSS

---

Built with ❤️ using Atomic Design principles
