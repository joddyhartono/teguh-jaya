import { toImageSrc } from "@/lib/utils";
import type { Product } from "@/types/product";
import { Button } from "../ui/button";
import { SquarePen, Trash } from "lucide-react";
import { Skeleton } from "../ui/skeleton";

type ProductCardProps = {
  product: Product;
};

const ProductCard = ({ product }: ProductCardProps) => {
  return (
    <div className="flex flex-col gap-2 justify-between items-center border rounded-lg bg-card p-4 shadow-sm">
      <h3 className="font-semibold">{product.name}</h3>
      <img
        src={toImageSrc(product.image)}
        alt={product.name}
        className="h-56 w-auto"
      />
      <div className="flex flex-1 justify-between items-center ">
        <Button className="cursor-pointer">
          <SquarePen />
        </Button>
        <Button className="cursor-pointer">
          <Trash />
        </Button>
      </div>
    </div>
  );
};

export const ProductCardSkeleton = () => {
  return (
    <div className="flex flex-col gap-2 justify-between items-center border rounded-lg bg-card p-4 shadow-sm">
      <Skeleton className="h-6 w-24" />
      <Skeleton className="h-56 w-40" />
      <div className="flex flex-1 w-full justify-between items-center">
        <Skeleton className="h-9 w-10" />
        <Skeleton className="h-9 w-10" />
      </div>
    </div>
  );
};

export default ProductCard;
