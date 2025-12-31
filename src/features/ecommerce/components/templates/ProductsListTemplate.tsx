import { useState, useEffect } from 'react';
import { useNavigate } from '@tanstack/react-router';
import { SlidersHorizontal, Grid, List, X } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select';
import { Slider } from '@/components/ui/slider';
import { Sheet, SheetContent, SheetHeader, SheetTitle, SheetTrigger } from '@/components/ui/sheet';
import { Skeleton } from '@/components/ui/skeleton';
import type { Product, ProductFilters } from '../../types';
import { Header } from '../../components/organisms/Header';
import { Footer } from '../../components/organisms/Footer';
import { useSearchStore } from '@/stores/search-store';

interface ProductsListTemplateProps {
	products: Product[];
	meta: {
		page: number;
		page_size: number;
		total: number;
		total_page: number;
	};
	filters: ProductFilters;
	onFiltersChange: (filters: ProductFilters) => void;
	isLoading?: boolean;
}

const FilterSheet = ({
	priceRange,
	setPriceRange,
	formatPrice,
	handlePriceFilter,
	clearFilters,
}: {
	priceRange: [number, number];
	setPriceRange: (value: [number, number]) => void;
	formatPrice: (price: number) => string;
	handlePriceFilter: () => void;
	clearFilters: () => void;
}) => (
	<Sheet>
		<SheetTrigger asChild>
			<Button variant="outline" size="sm">
				<SlidersHorizontal className="w-4 h-4 mr-2" />
				Filters
			</Button>
		</SheetTrigger>
		<SheetContent>
			<SheetHeader>
				<SheetTitle>Filter Products</SheetTitle>
			</SheetHeader>
			<div className="py-6 space-y-6 px-3">
				<div className="space-y-2">
					<label className="text-sm font-medium">Price Range</label>
					<Slider
						min={0}
						max={10000000}
						step={100000}
						value={priceRange}
						onValueChange={(value) => setPriceRange(value as [number, number])}
						className="py-4"
					/>
					<div className="flex justify-between text-sm text-muted-foreground">
						<span>{formatPrice(priceRange[0])}</span>
						<span>{formatPrice(priceRange[1])}</span>
					</div>
				</div>

				<div className="flex gap-2">
					<Button onClick={handlePriceFilter} className="flex-1">
						Apply Filters
					</Button>
					<Button onClick={clearFilters} variant="outline">
						Clear
					</Button>
				</div>
			</div>
		</SheetContent>
	</Sheet>
);

export const ProductsListTemplate = ({
	products,
	meta,
	filters,
	onFiltersChange,
	isLoading = false,
}: ProductsListTemplateProps) => {
	const navigate = useNavigate();
	const { searchQuery, clearSearch } = useSearchStore();
	const [viewMode, setViewMode] = useState<'grid' | 'list'>('grid');
	const [priceRange, setPriceRange] = useState<[number, number]>([
		filters.min_price || 0,
		filters.max_price || 10000000,
	]);

	// Sync price range with filters when they change externally
	useEffect(() => {
		setPriceRange([filters.min_price || 0, filters.max_price || 10000000]);
	}, [filters.min_price, filters.max_price]);

	// Update filters when search query changes from the header
	useEffect(() => {
		if (searchQuery && searchQuery !== filters.search) {
			onFiltersChange({ ...filters, search: searchQuery, page: 1 });
		}
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [searchQuery]);


    const formatPrice = (price: number) => {
        return new Intl.NumberFormat('id-ID', {
            style: 'currency',
            currency: 'IDR',
            minimumFractionDigits: 0,
        }).format(price);
    };

    const handleSortChange = (value: string) => {
        const [sort_by, sort_order] = value.split('-') as [ProductFilters['sort_by'], ProductFilters['sort_order']];
        onFiltersChange({ ...filters, sort_by, sort_order, page: 1 });
    };

    const handlePriceFilter = () => {
        onFiltersChange({
            ...filters,
            min_price: priceRange[0],
            max_price: priceRange[1],
            page: 1,
        });
    };

    const clearFilters = () => {
        clearSearch();
        setPriceRange([0, 10000000]);
        onFiltersChange({
            page: 1,
            page_size: filters.page_size,
        });
    };

    const handleProductClick = (productId: number) => {
        navigate({ to: `/products/${productId}` });
    };

    const handlePageChange = (newPage: number) => {
        onFiltersChange({ ...filters, page: newPage });
        window.scrollTo({ top: 0, behavior: 'smooth' });
    };

    const ProductCard = ({ product }: { product: Product }) => (
        <Card
            className="overflow-hidden hover:shadow-lg transition-shadow cursor-pointer"
            onClick={() => handleProductClick(product.id)}
        >
            <div className="relative aspect-square bg-muted">
                <img
                    src={product.image_url}
                    alt={product.name}
                    className="w-full h-full object-cover"
                    onError={(e) => {
                        e.currentTarget.src = 'https://via.placeholder.com/300x300?text=Product';
                    }}
                />
                {product.discount_pct && product.discount_pct > 0 && (
                    <div className="absolute top-2 left-2 bg-red-500 text-white text-xs font-bold px-2 py-1 rounded">
                        {product.discount_pct}% OFF
                    </div>
                )}
                {product.is_featured && (
                    <div className="absolute top-2 right-2 bg-yellow-400 text-xs font-semibold px-2 py-1 rounded">
                        Featured
                    </div>
                )}
            </div>
            <CardContent className="p-4">
                <h3 className="font-medium text-sm text-foreground line-clamp-2 mb-2 h-10">
                    {product.name}
                </h3>
                <div className="flex flex-col gap-1">
                    <div className="flex items-center gap-2">
                        <span className="text-lg font-bold text-primary">
                            {formatPrice(product.price)}
                        </span>
                        {product.slashed_price && product.slashed_price > product.price && (
                            <span className="text-xs text-muted-foreground line-through">
                                {formatPrice(product.slashed_price)}
                            </span>
                        )}
                    </div>
                    <div className="flex items-center justify-between text-xs text-muted-foreground">
                        <div className="flex items-center gap-1">
                            <span>⭐ {product.rating.toFixed(1)}</span>
                            <span>({product.count_review})</span>
                        </div>
                        <span>{product.count_sold} sold</span>
                    </div>
                    {product.shop && (
                        <div className="text-xs text-muted-foreground mt-1 flex items-center gap-1">
                            <span className="line-clamp-1">{product.shop.name}</span>
                            {product.shop.is_official && (
                                <span className="text-green-600 font-semibold">✓</span>
                            )}
                        </div>
                    )}
                </div>
            </CardContent>
        </Card>
    );

    const ProductListItem = ({ product }: { product: Product }) => (
        <Card
            className="overflow-hidden hover:shadow-lg transition-shadow cursor-pointer"
            onClick={() => handleProductClick(product.id)}
        >
            <div className="flex gap-4 p-4">
                <div className="relative w-32 h-32 shrink-0 bg-muted rounded">
                    <img
                        src={product.image_url}
                        alt={product.name}
                        className="w-full h-full object-cover rounded"
                        onError={(e) => {
                            e.currentTarget.src = 'https://via.placeholder.com/300x300?text=Product';
                        }}
                    />
                    {product.discount_pct && product.discount_pct > 0 && (
                        <div className="absolute top-2 left-2 bg-red-500 text-white text-xs font-bold px-2 py-1 rounded">
                            {product.discount_pct}% OFF
                        </div>
                    )}
                </div>
                <div className="flex-1 flex flex-col justify-between">
                    <div>
                        <h3 className="font-medium text-base text-foreground line-clamp-2 mb-2">
                            {product.name}
                        </h3>
                        {product.subtitle && (
                            <p className="text-sm text-muted-foreground line-clamp-1 mb-2">
                                {product.subtitle}
                            </p>
                        )}
                    </div>
                    <div className="flex items-center justify-between">
                        <div className="flex flex-col gap-1">
                            <div className="flex items-center gap-2">
                                <span className="text-xl font-bold text-primary">
                                    {formatPrice(product.price)}
                                </span>
                                {product.slashed_price && product.slashed_price > product.price && (
                                    <span className="text-sm text-muted-foreground line-through">
                                        {formatPrice(product.slashed_price)}
                                    </span>
                                )}
                            </div>
                            <div className="flex items-center gap-3 text-sm text-muted-foreground">
                                <div className="flex items-center gap-1">
                                    <span>⭐ {product.rating.toFixed(1)}</span>
                                    <span>({product.count_review})</span>
                                </div>
                                <span>{product.count_sold} sold</span>
                            </div>
                        </div>
                        {product.shop && (
                            <div className="text-sm text-muted-foreground flex items-center gap-1">
                                <span>{product.shop.name}</span>
                                {product.shop.is_official && (
                                    <span className="text-green-600 font-semibold">✓</span>
                                )}
                            </div>
                        )}
                    </div>
                </div>
            </div>
        </Card>
    );

    return (
        <div className="min-h-screen flex flex-col">
            <Header />

            <main className="flex-1 container mx-auto px-4 py-6 mt-30">
                {/* Filters Bar */}
                <div className="mb-6 space-y-4">
                    <div className="flex items-center justify-between flex-wrap gap-4">
                        <div className="flex items-center gap-2">
                            <FilterSheet
								priceRange={priceRange}
								setPriceRange={setPriceRange}
								formatPrice={formatPrice}
								handlePriceFilter={handlePriceFilter}
								clearFilters={clearFilters}
							/>
                            {(filters.search || filters.min_price || filters.max_price || filters.category_id || filters.sub_category_id) && (
                                <Button
                                    variant="ghost"
                                    size="sm"
                                    onClick={clearFilters}
                                    className="text-muted-foreground"
                                >
                                    <X className="w-4 h-4 mr-2" />
                                    Clear filters
                                </Button>
                            )}
                            {filters.search && (
                                <div className="text-sm text-muted-foreground">
                                    Searching for: <span className="font-medium text-foreground">"{filters.search}"</span>
                                </div>
                            )}
                        </div>

                        <div className="flex items-center gap-2">
                            <Select
                                value={`${filters.sort_by || 'created_at'}-${filters.sort_order || 'desc'}`}
                                onValueChange={handleSortChange}
                            >
                                <SelectTrigger className="w-[180px]">
                                    <SelectValue placeholder="Sort by" />
                                </SelectTrigger>
                                <SelectContent>
                                    <SelectItem value="created_at-desc">Newest</SelectItem>
                                    <SelectItem value="created_at-asc">Oldest</SelectItem>
                                    <SelectItem value="price-asc">Price: Low to High</SelectItem>
                                    <SelectItem value="price-desc">Price: High to Low</SelectItem>
                                    <SelectItem value="rating-desc">Highest Rated</SelectItem>
                                    <SelectItem value="count_sold-desc">Most Sold</SelectItem>
                                    <SelectItem value="name-asc">Name: A to Z</SelectItem>
                                </SelectContent>
                            </Select>

                            <div className="flex gap-1 border rounded-md p-1">
                                <Button
                                    variant={viewMode === 'grid' ? 'secondary' : 'ghost'}
                                    size="sm"
                                    onClick={() => setViewMode('grid')}
                                >
                                    <Grid className="w-4 h-4" />
                                </Button>
                                <Button
                                    variant={viewMode === 'list' ? 'secondary' : 'ghost'}
                                    size="sm"
                                    onClick={() => setViewMode('list')}
                                >
                                    <List className="w-4 h-4" />
                                </Button>
                            </div>
                        </div>
                    </div>
                </div>

                {/* Results Info */}
                <div className="mb-4 text-sm text-muted-foreground">
                    Showing {products.length > 0 ? ((meta.page - 1) * meta.page_size) + 1 : 0} - {Math.min(meta.page * meta.page_size, meta.total)} of {meta.total} products
                </div>

                {/* Products Grid/List */}
                {isLoading ? (
                    <div className={viewMode === 'grid'
                        ? 'grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-4'
                        : 'space-y-4'
                    }>
                        {Array.from({ length: 10 }).map((_, i) => (
                            <Card key={i}>
                                <Skeleton className="aspect-square" />
                                <CardContent className="p-4">
                                    <Skeleton className="h-4 w-full mb-2" />
                                    <Skeleton className="h-4 w-2/3" />
                                </CardContent>
                            </Card>
                        ))}
                    </div>
                ) : products.length === 0 ? (
                    <div className="text-center py-12">
                        <p className="text-muted-foreground text-lg">No products found</p>
                        <Button onClick={clearFilters} variant="outline" className="mt-4">
                            Clear filters
                        </Button>
                    </div>
                ) : (
                    <>
                        {viewMode === 'grid' ? (
                            <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-4">
                                {products.map((product) => (
                                    <ProductCard key={product.id} product={product} />
                                ))}
                            </div>
                        ) : (
                            <div className="space-y-4">
                                {products.map((product) => (
                                    <ProductListItem key={product.id} product={product} />
                                ))}
                            </div>
                        )}

                        {/* Pagination */}
                        {meta.total_page > 1 && (
                            <div className="mt-8 flex justify-center items-center gap-2">
                                <Button
                                    variant="outline"
                                    disabled={meta.page === 1}
                                    onClick={() => handlePageChange(meta.page - 1)}
                                >
                                    Previous
                                </Button>

                                <div className="flex gap-1">
                                    {Array.from({ length: Math.min(5, meta.total_page) }, (_, i) => {
                                        let pageNum: number;
                                        if (meta.total_page <= 5) {
                                            pageNum = i + 1;
                                        } else if (meta.page <= 3) {
                                            pageNum = i + 1;
                                        } else if (meta.page >= meta.total_page - 2) {
                                            pageNum = meta.total_page - 4 + i;
                                        } else {
                                            pageNum = meta.page - 2 + i;
                                        }

                                        return (
                                            <Button
                                                key={pageNum}
                                                variant={pageNum === meta.page ? 'default' : 'outline'}
                                                onClick={() => handlePageChange(pageNum)}
                                                size="sm"
                                            >
                                                {pageNum}
                                            </Button>
                                        );
                                    })}
                                </div>

                                <Button
                                    variant="outline"
                                    disabled={meta.page === meta.total_page}
                                    onClick={() => handlePageChange(meta.page + 1)}
                                >
                                    Next
                                </Button>
                            </div>
                        )}
                    </>
                )}
            </main>

            <Footer />
        </div>
    );
};
