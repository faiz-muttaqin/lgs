import { useEffect } from 'react';
import { useNavigate, useSearch } from '@tanstack/react-router';
import { Logo } from '../atoms/Logo';
import { SearchBar } from '../molecules/SearchBar';
import { CartButton } from '../molecules/CartButton';
import { ProfileDropdown } from '@/components/ProfileDropdown';
import { CategoryList } from '../molecules/CategoryList';
import { useCategoriesStore } from '@/stores/categories-store';
import { categoriesApi } from '../../api/categories';
import ThemePresetSelect from '@/features/editor-theme/components/theme-preset-select';

export const Header = () => {
  const navigate = useNavigate();
  const { categories, isLoading, setCategories, setLoading, setError } = useCategoriesStore();
  
  // Get current category from URL if on products page
  const searchParams = useSearch({ strict: false }) as { category_id?: string };
  const activeCategoryId = searchParams?.category_id;

  useEffect(() => {
    // Fetch categories on mount if not already loaded
    if (categories.length === 0 && !isLoading) {
      const fetchCategories = async () => {
        setLoading(true);
        try {
          const response = await categoriesApi.getCategories();
          if (response.success && response.data) {
            setCategories(response.data);
          } else {
            setError('Failed to load categories');
          }
        } catch (error) {
          console.error('Error fetching categories:', error);
          setError('Failed to load categories');
        } finally {
          setLoading(false);
        }
      };

      fetchCategories();
    }
  }, [categories.length, isLoading, setCategories, setLoading, setError]);

  const handleCategoryClick = (categoryId: number) => {
    // Navigate to products page with category filter
    navigate({
      to: '/products',
      search: {
        category_id: String(categoryId),
        page: '1',
      },
    });
  };

  return (
    <header className="fixed top-0 left-0 right-0 z-50 bg-background border-b shadow-sm">
      {/* Top Bar */}
      {/* <div className="bg-primary text-primary-foreground text-xs py-1">
        <div className="container mx-auto px-4 flex justify-between items-center">
          <span>Promo 12.12 - Diskon hingga 90%</span>
          <span>Gratis Ongkir Min. Belanja 0Rp</span>
        </div>
      </div> */}

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
        {isLoading ? (
          <div className="flex items-center justify-center py-2">
            <span className="text-sm text-muted-foreground">Loading categories...</span>
          </div>
        ) : (
          <CategoryList 
            categories={categories} 
            activeCategory={activeCategoryId}
            onCategoryClick={handleCategoryClick}
          />
        )}
      </div>
    </header>
  );
};
