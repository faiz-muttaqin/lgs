import { ShoppingBag, Facebook, Instagram, Twitter, Youtube } from 'lucide-react';
import { Separator } from '@/components/ui/separator';

export const Footer = () => {
  return (
    <footer className="bg-muted/50 border-t mt-12">
      <div className="container mx-auto px-4 py-12">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-8">
          {/* Brand */}
          <div>
            <div className="flex items-center gap-2 mb-4">
              <ShoppingBag className="h-8 w-8 text-primary" />
              <div className="flex flex-col leading-none">
                <span className="text-xl font-bold text-foreground">LGS</span>
                <span className="text-xs text-muted-foreground">Lucky Good Store</span>
              </div>
            </div>
            <p className="text-sm text-muted-foreground mb-4">
              Platform belanja online terpercaya dengan jutaan produk berkualitas dan harga terbaik.
            </p>
            <div className="flex gap-3">
              <a href="#" className="text-muted-foreground hover:text-primary transition-colors">
                <Facebook className="h-5 w-5" />
              </a>
              <a href="#" className="text-muted-foreground hover:text-primary transition-colors">
                <Instagram className="h-5 w-5" />
              </a>
              <a href="#" className="text-muted-foreground hover:text-primary transition-colors">
                <Twitter className="h-5 w-5" />
              </a>
              <a href="#" className="text-muted-foreground hover:text-primary transition-colors">
                <Youtube className="h-5 w-5" />
              </a>
            </div>
          </div>

          {/* Customer Service */}
          <div>
            <h3 className="font-semibold text-foreground mb-4">Layanan Pelanggan</h3>
            <ul className="space-y-2 text-sm text-muted-foreground">
              <li>
                <a href="/help" className="hover:text-primary transition-colors">
                  Pusat Bantuan
                </a>
              </li>
              <li>
                <a href="/payment" className="hover:text-primary transition-colors">
                  Cara Pembelian
                </a>
              </li>
              <li>
                <a href="/shipping" className="hover:text-primary transition-colors">
                  Pengiriman
                </a>
              </li>
              <li>
                <a href="/returns" className="hover:text-primary transition-colors">
                  Pengembalian Barang
                </a>
              </li>
              <li>
                <a href="/contact" className="hover:text-primary transition-colors">
                  Hubungi Kami
                </a>
              </li>
            </ul>
          </div>

          {/* About LGS */}
          <div>
            <h3 className="font-semibold text-foreground mb-4">Tentang LGS</h3>
            <ul className="space-y-2 text-sm text-muted-foreground">
              <li>
                <a href="/about" className="hover:text-primary transition-colors">
                  Tentang Kami
                </a>
              </li>
              <li>
                <a href="/careers" className="hover:text-primary transition-colors">
                  Karir
                </a>
              </li>
              <li>
                <a href="/blog" className="hover:text-primary transition-colors">
                  Blog
                </a>
              </li>
              <li>
                <a href="/seller" className="hover:text-primary transition-colors">
                  Jual di LGS
                </a>
              </li>
              <li>
                <a href="/affiliate" className="hover:text-primary transition-colors">
                  Program Afiliasi
                </a>
              </li>
            </ul>
          </div>

          {/* Payment & Shipping */}
          <div>
            <h3 className="font-semibold text-foreground mb-4">Metode Pembayaran</h3>
            <div className="grid grid-cols-3 gap-2 mb-4">
              {['BCA', 'BNI', 'Mandiri', 'BRI', 'Gopay', 'OVO'].map((payment) => (
                <div
                  key={payment}
                  className="bg-background border rounded p-2 text-xs text-center font-medium text-foreground"
                >
                  {payment}
                </div>
              ))}
            </div>
            <h3 className="font-semibold text-foreground mb-4">Jasa Pengiriman</h3>
            <div className="grid grid-cols-3 gap-2">
              {['JNE', 'J&T', 'SiCepat', 'Anteraja', 'Ninja'].map((shipping) => (
                <div
                  key={shipping}
                  className="bg-background border rounded p-2 text-xs text-center font-medium text-foreground"
                >
                  {shipping}
                </div>
              ))}
            </div>
          </div>
        </div>

        <Separator className="my-8" />

        <div className="flex flex-col md:flex-row justify-between items-center gap-4 text-sm text-muted-foreground">
          <p>© 2024 Lucky Good Store. All rights reserved.</p>
          <div className="flex gap-6">
            <a href="/privacy" className="hover:text-primary transition-colors">
              Kebijakan Privasi
            </a>
            <a href="/terms" className="hover:text-primary transition-colors">
              Syarat & Ketentuan
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
};
