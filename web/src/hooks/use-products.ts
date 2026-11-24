import { getProducts } from "@/services/product-service";
import type { Product } from "@/types/product";
import { useEffect, useState } from "react";
import { toast } from "sonner";

const useProducts = () => {
  const [products, setProducts] = useState<Product[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        setIsLoading(true);
        const data = await getProducts();
        if (data) {
          setProducts(data);
        }
      } catch (error) {
        if (typeof error === "object" && error !== null && "message" in error) {
          toast.error(String(error.message));
        }
      } finally {
        setIsLoading(false);
      }
    };

    fetchProducts();
  }, []);

  return {
    products,
    isLoading,
  };
};

export default useProducts;
