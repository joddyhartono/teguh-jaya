import type { Category } from "@/types/category";
import { Button } from "../ui/button";
import { SquarePen, Trash } from "lucide-react";
import { toImageSrc } from "@/lib/utils";

type CategoryCardProps = {
  category: Category;
};

const CategoryCard = ({ category }: CategoryCardProps) => {
  return (
    <div className="flex flex-col gap-2 justify-between items-center border rounded-lg bg-card p-4 shadow-sm">
      <h3 className="font-semibold">{category.name}</h3>
      <img
        src={toImageSrc(category.image)}
        alt={category.name}
        className="h-56 w-auto"
      />
      <div className="flex flex-end flex-1 justify-between items-center ">
        <Button>
          <SquarePen />
        </Button>
        <Button>
          <Trash />
        </Button>
      </div>
    </div>
  );
};

export default CategoryCard;
