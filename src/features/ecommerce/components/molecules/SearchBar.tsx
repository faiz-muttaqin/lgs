import { useState } from 'react';
import { useNavigate } from '@tanstack/react-router';
import { SearchInput } from '../atoms/SearchInput';
import { Button } from '@/components/ui/button';
import { Search } from 'lucide-react';
import { useSearchStore } from '@/stores/search-store';

export const SearchBar = () => {
	const navigate = useNavigate();
	const { setSearchQuery } = useSearchStore();
	const [localSearch, setLocalSearch] = useState('');

	const handleSearch = () => {
		if (localSearch.trim()) {
			setSearchQuery(localSearch.trim());
			navigate({
				to: '/products',
				search: { search: localSearch.trim(), page: "1" },
			});
		}
	};

	const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
		if (e.key === 'Enter') {
			handleSearch();
		}
	};

	return (
		<div className="flex gap-2 w-full max-w-2xl">
			<SearchInput
				placeholder="Cari di LGS..."
				value={localSearch}
				onChange={(e) => setLocalSearch(e.target.value)}
				onKeyDown={handleKeyDown}
			/>
			<Button onClick={handleSearch}>
				<Search className="h-4 w-4" />
			</Button>
		</div>
	);
};
