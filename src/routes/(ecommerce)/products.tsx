import { createFileRoute } from '@tanstack/react-router';
import { useState, useEffect } from 'react';
import { ProductsListTemplate } from '@/features/ecommerce/components/templates/ProductsListTemplate';
import { productsApi } from '@/features/ecommerce/api/products';
import type { ProductFilters, Product } from '@/features/ecommerce/types';

interface ProductsSearch {
    search?: string;
    category_id?: string;
    sub_category_id?: string;
    min_price?: string;
    max_price?: string;
    sort_by?: string;
    sort_order?: string;
    page?: string;
    page_size?: string;
}

export const Route = createFileRoute('/(ecommerce)/products')({
    validateSearch: (search: Record<string, unknown>): ProductsSearch => {
        return {
            search: (search.search as string) || undefined,
            category_id: search.category_id ? String(search.category_id) : undefined,
            sub_category_id: search.sub_category_id ? String(search.sub_category_id) : undefined,
            min_price: search.min_price ? String(search.min_price) : undefined,
            max_price: search.max_price ? String(search.max_price) : undefined,
            sort_by: (search.sort_by as string) || undefined,
            sort_order: (search.sort_order as string) || undefined,
            page: search.page ? String(search.page) : undefined,
            page_size: search.page_size ? String(search.page_size) : undefined,
        };
    },
    loaderDeps: ({ search }) => ({ search }),
    loader: async ({ deps: { search } }) => {
        const filters: ProductFilters = {
            search: search.search,
            category_id: search.category_id ? Number(search.category_id) : undefined,
            sub_category_id: search.sub_category_id ? Number(search.sub_category_id) : undefined,
            min_price: search.min_price ? Number(search.min_price) : undefined,
            max_price: search.max_price ? Number(search.max_price) : undefined,
            sort_by: search.sort_by as ProductFilters['sort_by'],
            sort_order: search.sort_order as ProductFilters['sort_order'],
            page: search.page ? Number(search.page) : 1,
            page_size: search.page_size ? Number(search.page_size) : 20,
        };

        const response = await productsApi.getProducts(filters);
        return {
            products: response.data?.data || [],
            meta: response.data?.meta || { page: 1, page_size: 20, total: 0, total_page: 0 },
            initialFilters: filters,
        };
    },
    component: RouteComponent,
});

function RouteComponent() {
    const loaderData = Route.useLoaderData();
    const navigate = Route.useNavigate();
    const [products, setProducts] = useState<Product[]>(loaderData.products);
    const [meta, setMeta] = useState(loaderData.meta);
    const [filters, setFilters] = useState<ProductFilters>(loaderData.initialFilters);
    const [isLoading, setIsLoading] = useState(false);

    // Sync state when loader data changes (e.g., when navigating from header category click)
    useEffect(() => {
        setProducts(loaderData.products);
        setMeta(loaderData.meta);
        setFilters(loaderData.initialFilters);
    }, [loaderData]);

    const handleFiltersChange = async (newFilters: ProductFilters) => {
        setIsLoading(true);
        setFilters(newFilters);

        // Extract simple values for URL (convert FilterWithOperator to simple values)
        const getCategoryId = (value: typeof newFilters.category_id): number | undefined => {
            if (!value) return undefined;
            return typeof value === 'object' && 'value' in value 
                ? Number(value.value) 
                : Number(value);
        };

        const getSubCategoryId = (value: typeof newFilters.sub_category_id): number | undefined => {
            if (!value) return undefined;
            return typeof value === 'object' && 'value' in value 
                ? Number(value.value) 
                : Number(value);
        };

        // Update URL with new filters
        await navigate({
            search: {
                search: newFilters.search,
                category_id: getCategoryId(newFilters.category_id)?.toString(),
                sub_category_id: getSubCategoryId(newFilters.sub_category_id)?.toString(),
                min_price: newFilters.min_price?.toString(),
                max_price: newFilters.max_price?.toString(),
                sort_by: newFilters.sort_by,
                sort_order: newFilters.sort_order,
                page: newFilters.page?.toString(),
                page_size: newFilters.page_size?.toString(),
            },
        });

        // Fetch new data
        try {
            const response = await productsApi.getProducts(newFilters);
            setProducts(response.data?.data || []);
            setMeta(response.data?.meta || { page: 1, page_size: 20, total: 0, total_page: 0 });
        } catch (error) {
            console.error('Failed to fetch products:', error);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <ProductsListTemplate
            products={products}
            meta={meta}
            filters={filters}
            onFiltersChange={handleFiltersChange}
            isLoading={isLoading}
        />
    );
}
