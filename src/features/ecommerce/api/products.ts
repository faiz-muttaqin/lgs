import { apiClient, type ApiResponse } from '@/lib/api/client';
import type { Product, ProductDetail, ProductFilters, ProductListResponse, FilterWithOperator } from '../types';

/**
 * Helper to add filter with operator to params
 */
function addFilterParam(
	params: Record<string, string | number | boolean>,
	field: string,
	value: string | number | boolean | FilterWithOperator | undefined
) {
	if (value === undefined) return;

	if (typeof value === 'object' && 'operator' in value) {
		// Handle FilterWithOperator
		const filter = value as FilterWithOperator;
		if (filter.operator === 'between' && Array.isArray(filter.value)) {
			// For between, add two params: field[gte] and field[lte]
			params[`${field}[gte]`] = filter.value[0];
			params[`${field}[lte]`] = filter.value[1];
		} else if (filter.operator === 'in' && Array.isArray(filter.value)) {
			// For in, join array values with comma
			params[`${field}[in]`] = filter.value.join(',');
		} else {
			params[`${field}[${filter.operator}]`] = filter.value as string | number | boolean;
		}
	} else {
		// Default to 'eq' for direct values, except strings use 'contains'
		const operator = typeof value === 'string' ? 'contains' : 'eq';
		params[`${field}[${operator}]`] = value;
	}
}

/**
 * API service for product-related endpoints
 */
export const productsApi = {
	/**
	 * Fetch products with filters, search, and pagination
	 */
	async getProducts(filters?: ProductFilters): Promise<ApiResponse<ProductListResponse>> {
		const params: Record<string, string | number | boolean> = {};

		if (filters) {
			// Convert page/page_size to DataTables format (start/length)
			const page = filters.page || 1;
			const pageSize = filters.page_size || 20;
			params.start = (page - 1) * pageSize;
			params.length = pageSize;

			// Search (convert to name[contains])
			if (filters.search) {
				params['name[contains]'] = filters.search;
			}

			// Dynamic field filters with operator support
			addFilterParam(params, 'name', filters.name);
			addFilterParam(params, 'sku', filters.sku);
			addFilterParam(params, 'description', filters.description);
			addFilterParam(params, 'category_id', filters.category_id);
			addFilterParam(params, 'sub_category_id', filters.sub_category_id);
			addFilterParam(params, 'shop_id', filters.shop_id);
			addFilterParam(params, 'price', filters.price);
			addFilterParam(params, 'stock', filters.stock);
			addFilterParam(params, 'status', filters.status);
			addFilterParam(params, 'discount_pct', filters.discount_pct);
			addFilterParam(params, 'rating', filters.rating);

			// Convenience price range filters
			if (filters.min_price !== undefined) params['price[gte]'] = filters.min_price;
			if (filters.max_price !== undefined) params['price[lte]'] = filters.max_price;

			// Boolean filters
			if (filters.is_active !== undefined) params['is_active[eq]'] = filters.is_active;
			if (filters.is_featured !== undefined) params['is_featured[eq]'] = filters.is_featured;

			// Sorting (use filter DSL format: "-created_at" for desc, "created_at" for asc)
			if (filters.sort_by) {
				const sortOrder = filters.sort_order === 'asc' ? '' : '-';
				params.sort = `${sortOrder}${filters.sort_by}`;
			}
		}

		const response = await apiClient.get<Product[]>('/products', { params });
		
		// Transform DataTables response to our format
		if (response.success && response.data) {
			const start = params.start as number || 0;
			const length = params.length as number || 20;
			const page = Math.floor(start / length) + 1;
			const total = response.recordsFiltered || response.recordsTotal || 0;
			
			return {
				...response,
				data: {
					data: Array.isArray(response.data) ? response.data : [],
					meta: {
						page,
						page_size: length,
						total,
						total_page: Math.ceil(total / length),
					},
				},
			};
		}

		// Return empty response if no data
		return {
			...response,
			data: {
				data: [],
				meta: {
					page: 1,
					page_size: 20,
					total: 0,
					total_page: 0,
				},
			},
		};
	},

	/**
	 * Fetch a single product by ID with full details
	 */
	async getProductById(id: string | number): Promise<ApiResponse<ProductDetail>> {
		return apiClient.get<ProductDetail>(`/products/${id}`);
	},

	/**
	 * Create a new product (requires authentication)
	 */
	async createProduct(productData: Partial<Product>): Promise<ApiResponse<Product>> {
		return apiClient.post<Product>('/products', productData);
	},

	/**
	 * Update an existing product (requires authentication)
	 */
	async updateProduct(id: string | number, productData: Partial<Product>): Promise<ApiResponse<Product>> {
		return apiClient.put<Product>(`/products/${id}`, productData);
	},

	/**
	 * Delete a product (requires authentication)
	 */
	async deleteProduct(id: string | number): Promise<ApiResponse<void>> {
		return apiClient.delete(`/products/${id}`);
	},
};
