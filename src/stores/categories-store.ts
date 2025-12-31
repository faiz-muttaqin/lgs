import { create } from 'zustand';
import type { Category } from '@/features/ecommerce/types';

interface CategoriesState {
	categories: Category[];
	isLoading: boolean;
	error: string | null;
	setCategories: (categories: Category[]) => void;
	setLoading: (isLoading: boolean) => void;
	setError: (error: string | null) => void;
	clearCategories: () => void;
}

export const useCategoriesStore = create<CategoriesState>((set) => ({
	categories: [],
	isLoading: false,
	error: null,
	setCategories: (categories) => set({ categories, error: null }),
	setLoading: (isLoading) => set({ isLoading }),
	setError: (error) => set({ error, isLoading: false }),
	clearCategories: () => set({ categories: [], error: null }),
}));
