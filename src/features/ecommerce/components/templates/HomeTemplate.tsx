import { useEffect } from 'react';
import { Header } from '../organisms/Header';
// import { BannerCarousel } from '../organisms/BannerCarousel';
// import { ProductSection } from '../organisms/ProductSection';
import { Footer } from '../organisms/Footer';
// import { mainBanners, sections as mockSections } from '../../data/mockData';
import type { Section, Category } from '../../types';
import '../../styles/home.css';

interface HomeTemplateProps {
	sections?: Section[];
	categories?: Category[];
	isLoading?: boolean;
	onRefresh?: () => void;
}

export const HomeTemplate = ({ 
	// sections = mockSections, 
	isLoading = false,
	onRefresh: _onRefresh
}: HomeTemplateProps) => {
	useEffect(() => {
		// Scroll to top on mount
		window.scrollTo(0, 0);
	}, []);

	return (
		<div className="min-h-screen bg-background">
			<Header />

			{/* Add spacing for fixed header - approximate height */}
			<div className="h-[180px] md:h-40" />

			<main className="container mx-auto px-4">
				{/* Main Banner */}
				<section className="py-6">
					{/* <BannerCarousel banners={mainBanners} /> */}
				</section>

				{/* Product Sections */}
				{isLoading ? (
					<div className="text-center py-12">
						<p className="text-muted-foreground">Loading products...</p>
					</div>
				): (
					<div className="text-center py-12">
						<p className="text-muted-foreground">No products available.</p>
					</div>
				)}
			</main>

			<Footer />
		</div>
	);
};
