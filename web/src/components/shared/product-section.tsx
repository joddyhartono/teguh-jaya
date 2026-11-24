import { useAppSelector } from "@/store/hooks";
import { Button } from "../ui/button";
import useProducts from "@/hooks/use-products";
import ProductList from "./product-list";
import { ProductCardSkeleton } from "./product-card";

const ProductSection = () => {
  const token = useAppSelector((state) => state.auth.token);
  const { products, isLoading } = useProducts();

  return (
    <section className="flex flex-col gap-4 p-4">
      <div className="flex justify-between items-center">
        <h2 className="text-xl font-semibold">Produk</h2>
        {token && <Button>Buat Produk</Button>}
      </div>

      {isLoading ? (
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {Array.from({ length: 4 }).map((_, i) => (
            <ProductCardSkeleton key={i} />
          ))}
        </div>
      ) : (
        <ProductList products={products} />
      )}
    </section>
  );
};

export default ProductSection;
