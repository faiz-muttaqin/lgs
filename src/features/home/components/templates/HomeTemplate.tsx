import { useEffect } from 'react';
import { Header } from '../organisms/Header';
import { BannerCarousel } from '../organisms/BannerCarousel';
import { ProductSection } from '../organisms/ProductSection';
import { Footer } from '../organisms/Footer';
import { mainBanners, sections } from '../../data/mockData';
import '../../styles/home.css';

export const HomeTemplate = () => {
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
          <BannerCarousel banners={mainBanners} />
        </section>

        {/* Product Sections */}
        {sections.map((section) => (
          <ProductSection key={section.id} section={section} />
        ))}
      </main>

      <Footer />
    </div>
  );
};
