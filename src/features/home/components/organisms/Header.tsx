import { Logo } from '../atoms/Logo';
import { SearchBar } from '../molecules/SearchBar';
import { CartButton } from '../molecules/CartButton';
import { ProfileDropdown } from '@/components/ProfileDropdown';
import { CategoryList } from '../molecules/CategoryList';
import { categories } from '../../data/mockData';
import ThemePresetSelect from '@/features/editor-theme/components/theme-preset-select';

export const Header = () => {
  return (
    <header className="fixed top-0 left-0 right-0 z-50 bg-background border-b shadow-sm">
      {/* Top Bar */}
      <div className="bg-primary text-primary-foreground text-xs py-1">
        <div className="container mx-auto px-4 flex justify-between items-center">
          <span>Promo 12.12 - Diskon hingga 90%</span>
          <span>Gratis Ongkir Min. Belanja 0Rp</span>
        </div>
      </div>

      {/* Main Header */}
      <div className="container mx-auto px-4 py-3">
        <div className="flex items-center justify-between gap-4">
          <Logo />
          <div className="flex-1 max-w-2xl hidden md:block">
            <SearchBar />
          </div>
          <div className="flex items-center gap-2">
            <ThemePresetSelect withCycleThemes={false} className="w-10" />
            <CartButton count={3} />
            {/* <AuthButtons /> */}
            <ProfileDropdown />
          </div>
        </div>

        {/* Mobile Search */}
        <div className="mt-3 md:hidden">
          <SearchBar />
        </div>
      </div>

      {/* Categories */}
      <div className="container mx-auto px-4 py-3 border-t">
        <CategoryList categories={categories} />
      </div>
    </header>
  );
};
