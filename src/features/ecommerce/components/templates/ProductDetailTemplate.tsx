import { useState } from 'react';
import { useNavigate } from '@tanstack/react-router';
import { ShoppingCart, Heart, Share2, Star, MapPin, ChevronLeft, ChevronRight, Minus, Plus, MessageCircle } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Badge } from '@/components/ui/badge';
import { Separator } from '@/components/ui/separator';
import type { ProductDetail } from '../../types';
import { Header } from '../../components/organisms/Header';
import { Footer } from '../../components/organisms/Footer';

interface ProductDetailTemplateProps {
    product: ProductDetail;
}

export const ProductDetailTemplate = ({ product }: ProductDetailTemplateProps) => {
    const navigate = useNavigate();
    const [selectedImage, setSelectedImage] = useState(0);
    const [quantity, setQuantity] = useState(1);
    const [selectedVariant, setSelectedVariant] = useState<number | null>(
        product.variants?.length > 0 ? product.variants[0].id : null
    );

    const images = [
        product.image_url,
        ...(product.images?.map((img) => img.image_url) || []),
    ];

    const formatPrice = (price: number) => {
        return new Intl.NumberFormat('id-ID', {
            style: 'currency',
            currency: 'IDR',
            minimumFractionDigits: 0,
        }).format(price);
    };

    const handleQuantityChange = (delta: number) => {
        const newQuantity = quantity + delta;
        if (newQuantity >= 1 && newQuantity <= product.stock) {
            setQuantity(newQuantity);
        }
    };

    const currentPrice = selectedVariant
        ? product.variants?.find((v) => v.id === selectedVariant)?.price || product.price
        : product.price;

    const handleAddToCart = () => {
        // TODO: Implement add to cart functionality
    };

    const handleBuyNow = () => {
        // TODO: Implement buy now functionality
    };

    const handleContactSeller = () => {
        // TODO: Implement contact seller functionality
    };

    return (
        <div className="min-h-screen flex flex-col">
            <Header />

            <main className="flex-1 container mx-auto px-4 py-6">
                {/* Breadcrumb */}
                <div className="mb-4 text-sm text-muted-foreground flex items-center gap-2">
                    <button onClick={() => navigate({ to: '/' })} className="hover:underline">
                        Home
                    </button>
                    <span>/</span>
                    <button onClick={() => navigate({ to: '/products' })} className="hover:underline">
                        Products
                    </button>
                    {product.category && (
                        <>
                            <span>/</span>
                            <span>{product.category.name}</span>
                        </>
                    )}
                </div>

                <div className="grid md:grid-cols-2 gap-8 mb-8">
                    {/* Image Gallery */}
                    <div className="space-y-4">
                        <Card className="overflow-hidden">
                            <div className="relative aspect-square bg-muted">
                                <img
                                    src={images[selectedImage]}
                                    alt={product.name}
                                    className="w-full h-full object-cover"
                                    onError={(e) => {
                                        e.currentTarget.src = 'https://via.placeholder.com/600x600?text=Product';
                                    }}
                                />
                                {product.discount_pct && product.discount_pct > 0 && (
                                    <div className="absolute top-4 left-4 bg-red-500 text-white text-sm font-bold px-3 py-1 rounded">
                                        {product.discount_pct}% OFF
                                    </div>
                                )}
                                {images.length > 1 && (
                                    <>
                                        <Button
                                            variant="secondary"
                                            size="icon"
                                            className="absolute left-2 top-1/2 -translate-y-1/2"
                                            onClick={() =>
                                                setSelectedImage((prev) => (prev === 0 ? images.length - 1 : prev - 1))
                                            }
                                        >
                                            <ChevronLeft className="w-4 h-4" />
                                        </Button>
                                        <Button
                                            variant="secondary"
                                            size="icon"
                                            className="absolute right-2 top-1/2 -translate-y-1/2"
                                            onClick={() =>
                                                setSelectedImage((prev) => (prev === images.length - 1 ? 0 : prev + 1))
                                            }
                                        >
                                            <ChevronRight className="w-4 h-4" />
                                        </Button>
                                    </>
                                )}
                            </div>
                        </Card>

                        {/* Thumbnail Gallery */}
                        {images.length > 1 && (
                            <div className="grid grid-cols-5 gap-2">
                                {images.map((img, idx) => (
                                    <button
                                        key={idx}
                                        onClick={() => setSelectedImage(idx)}
                                        className={`aspect-square rounded border-2 overflow-hidden ${selectedImage === idx ? 'border-primary' : 'border-transparent'
                                            }`}
                                    >
                                        <img
                                            src={img}
                                            alt={`${product.name} ${idx + 1}`}
                                            className="w-full h-full object-cover"
                                        />
                                    </button>
                                ))}
                            </div>
                        )}
                    </div>

                    {/* Product Info */}
                    <div className="space-y-6">
                        <div>
                            <h1 className="text-3xl font-bold mb-2">{product.name}</h1>
                            {product.subtitle && (
                                <p className="text-muted-foreground">{product.subtitle}</p>
                            )}
                        </div>

                        <div className="flex items-center gap-4">
                            <div className="flex items-center gap-1">
                                <Star className="w-5 h-5 fill-yellow-400 text-yellow-400" />
                                <span className="font-semibold">{product.rating.toFixed(1)}</span>
                                <span className="text-muted-foreground">
                                    ({product.count_review} reviews)
                                </span>
                            </div>
                            <Separator orientation="vertical" className="h-6" />
                            <span className="text-muted-foreground">{product.count_sold} sold</span>
                        </div>

                        <Card>
                            <CardContent className="p-4">
                                <div className="flex items-baseline gap-2">
                                    <span className="text-3xl font-bold text-primary">
                                        {formatPrice(currentPrice)}
                                    </span>
                                    {product.slashed_price && product.slashed_price > product.price && (
                                        <span className="text-lg text-muted-foreground line-through">
                                            {formatPrice(product.slashed_price)}
                                        </span>
                                    )}
                                </div>
                            </CardContent>
                        </Card>

                        {/* Variants */}
                        {product.variants && product.variants.length > 0 && (
                            <div className="space-y-2">
                                <label className="text-sm font-medium">Select Variant</label>
                                <div className="flex flex-wrap gap-2">
                                    {product.variants.map((variant) => (
                                        <Button
                                            key={variant.id}
                                            variant={selectedVariant === variant.id ? 'default' : 'outline'}
                                            onClick={() => setSelectedVariant(variant.id)}
                                            disabled={!variant.is_available || variant.stock === 0}
                                        >
                                            {variant.name}
                                            {variant.price && variant.price !== product.price && (
                                                <span className="ml-2 text-xs">+{formatPrice(variant.price - product.price)}</span>
                                            )}
                                        </Button>
                                    ))}
                                </div>
                            </div>
                        )}

                        {/* Quantity */}
                        <div className="space-y-2">
                            <label className="text-sm font-medium">Quantity</label>
                            <div className="flex items-center gap-4">
                                <div className="flex items-center border rounded">
                                    <Button
                                        variant="ghost"
                                        size="icon"
                                        onClick={() => handleQuantityChange(-1)}
                                        disabled={quantity <= 1}
                                    >
                                        <Minus className="w-4 h-4" />
                                    </Button>
                                    <span className="px-4 py-2 min-w-[60px] text-center">{quantity}</span>
                                    <Button
                                        variant="ghost"
                                        size="icon"
                                        onClick={() => handleQuantityChange(1)}
                                        disabled={quantity >= product.stock}
                                    >
                                        <Plus className="w-4 h-4" />
                                    </Button>
                                </div>
                                <span className="text-sm text-muted-foreground">
                                    {product.stock} pieces available
                                </span>
                            </div>
                        </div>

                        {/* Actions */}
                        <div className="flex gap-2">
                            <Button onClick={handleAddToCart} variant="outline" className="flex-1">
                                <ShoppingCart className="w-4 h-4 mr-2" />
                                Add to Cart
                            </Button>
                            <Button onClick={handleBuyNow} className="flex-1">
                                Buy Now
                            </Button>
                        </div>

                        <div className="flex gap-2">
                            <Button variant="ghost" size="icon">
                                <Heart className="w-5 h-5" />
                            </Button>
                            <Button variant="ghost" size="icon">
                                <Share2 className="w-5 h-5" />
                            </Button>
                        </div>

                        <Separator />

                        {/* Shop Info */}
                        <Card>
                            <CardContent className="p-4">
                                <div className="flex items-center justify-between">
                                    <div className="flex items-center gap-3">
                                        <div className="w-12 h-12 rounded-full bg-muted flex items-center justify-center">
                                            {product.shop.image_url ? (
                                                <img
                                                    src={product.shop.image_url}
                                                    alt={product.shop.name}
                                                    className="w-full h-full rounded-full object-cover"
                                                />
                                            ) : (
                                                <span className="text-xl font-bold">
                                                    {product.shop.name.charAt(0).toUpperCase()}
                                                </span>
                                            )}
                                        </div>
                                        <div>
                                            <div className="flex items-center gap-2">
                                                <h3 className="font-semibold">{product.shop.name}</h3>
                                                {product.shop.is_official && (
                                                    <Badge variant="secondary" className="bg-green-100 text-green-800">
                                                        Official
                                                    </Badge>
                                                )}
                                            </div>
                                            {product.shop.city && (
                                                <p className="text-sm text-muted-foreground flex items-center gap-1">
                                                    <MapPin className="w-3 h-3" />
                                                    {product.shop.city}
                                                </p>
                                            )}
                                        </div>
                                    </div>
                                    <Button variant="outline" size="sm" onClick={handleContactSeller}>
                                        <MessageCircle className="w-4 h-4 mr-2" />
                                        Chat
                                    </Button>
                                </div>
                            </CardContent>
                        </Card>
                    </div>
                </div>

                {/* Product Details Tabs */}
                <Card>
                    <Tabs defaultValue="description" className="w-full">
                        <TabsList className="w-full justify-start rounded-none border-b">
                            <TabsTrigger value="description">Description</TabsTrigger>
                            <TabsTrigger value="specifications">Specifications</TabsTrigger>
                            <TabsTrigger value="reviews">Reviews ({product.count_review})</TabsTrigger>
                        </TabsList>
                        <TabsContent value="description" className="p-6">
                            <div className="prose max-w-none">
                                {product.description ? (
                                    <p className="whitespace-pre-wrap">{product.description}</p>
                                ) : (
                                    <p className="text-muted-foreground">No description available.</p>
                                )}
                            </div>
                        </TabsContent>
                        <TabsContent value="specifications" className="p-6">
                            <div className="space-y-2">
                                <div className="grid grid-cols-2 py-2 border-b">
                                    <span className="font-medium">SKU</span>
                                    <span>{product.sku || 'N/A'}</span>
                                </div>
                                <div className="grid grid-cols-2 py-2 border-b">
                                    <span className="font-medium">Weight</span>
                                    <span>{product.weight ? `${product.weight}g` : 'N/A'}</span>
                                </div>
                                <div className="grid grid-cols-2 py-2 border-b">
                                    <span className="font-medium">Stock</span>
                                    <span>{product.stock} pieces</span>
                                </div>
                                <div className="grid grid-cols-2 py-2 border-b">
                                    <span className="font-medium">Category</span>
                                    <span>{product.category?.name || 'N/A'}</span>
                                </div>
                                {product.sub_category && (
                                    <div className="grid grid-cols-2 py-2 border-b">
                                        <span className="font-medium">Sub Category</span>
                                        <span>{product.sub_category.name}</span>
                                    </div>
                                )}
                                <div className="grid grid-cols-2 py-2 border-b">
                                    <span className="font-medium">Status</span>
                                    <span className="capitalize">{product.status}</span>
                                </div>
                            </div>
                        </TabsContent>
                        <TabsContent value="reviews" className="p-6">
                            <div className="text-center py-8 text-muted-foreground">
                                Reviews section coming soon
                            </div>
                        </TabsContent>
                    </Tabs>
                </Card>
            </main>

            <Footer />
        </div>
    );
};
