import { useAppSelector } from "@/store/hooks";
import CategoryList from "./category-list";
import { Button } from "../ui/button";
import { CategoryCardSkeleton } from "./category-card";
import useCategories from "@/hooks/use-categories";

const CategorySection = () => {
  const { categories } = useCategories();
  const status = useAppSelector((state) => state.category.status);

  return (
    <section className="flex flex-col gap-4 p-4">
      <div className="flex justify-between items-center">
        <h2 className="text-xl font-semibold">Kategori</h2>
        <Button>Buat Kategori</Button>
      </div>

      {status === "loading" && (
        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
          {Array.from({ length: 4 }).map((_, i) => (
            <CategoryCardSkeleton key={i} />
          ))}
        </div>
      )}
      {status === "succeeded" && <CategoryList categories={categories} />}
    </section>
  );
};

export default CategorySection;
