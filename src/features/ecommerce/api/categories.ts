import { apiClient, type ApiResponse } from '@/lib/api/client';
import type { Category, SubCategory } from '../types';

/**
 * API service for category-related endpoints
 */
export const categoriesApi = {
	/**
	 * Fetch all categories
	 */
	async getCategories(): Promise<ApiResponse<Category[]>> {
		const response = await apiClient.get<Category[]>('/categories', {
			params: {
				length: 100,
				start: 0,
			},
		});

		// Transform DataTables response
		if (response.success && response.data) {
			return {
				...response,
				data: Array.isArray(response.data) ? response.data : [],
			};
		}

		return response;
	},

	/**
	 * Fetch categories with subcategories
	 */
	async getCategoriesWithSubCategories(): Promise<ApiResponse<Category[]>> {
		const response = await apiClient.get<Category[]>('/categories/sub-categories', {
			params: {
				length: 100,
				start: 0,
			},
		});

		// Transform DataTables response
		if (response.success && response.data) {
			return {
				...response,
				data: Array.isArray(response.data) ? response.data : [],
			};
		}

		return response;
	},

	/**
	 * Fetch all subcategories
	 */
	async getSubCategories(): Promise<ApiResponse<SubCategory[]>> {
		const response = await apiClient.get<SubCategory[]>('/sub-categories', {
			params: {
				length: 100,
				start: 0,
			},
		});

		// Transform DataTables response
		if (response.success && response.data) {
			return {
				...response,
				data: Array.isArray(response.data) ? response.data : [],
			};
		}

		return response;
	},
};
