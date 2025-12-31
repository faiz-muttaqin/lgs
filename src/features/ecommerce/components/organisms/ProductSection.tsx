import { SectionHeader } from '../molecules/SectionHeader';
import { ProductCarousel } from './ProductCarousel';
import type { Section } from '../../types';

interface ProductSectionProps {
  section: Section;
}

export const ProductSection = ({ section }: ProductSectionProps) => {
  if (!section.products || section.products.length === 0) {
    return null;
  }

  return (
    <section className="py-6">
      <SectionHeader
        title={section.title}
        link={`/category/${section.id}`}
        showViewAll={true}
      />
      <ProductCarousel products={section.products} />
    </section>
  );
};
